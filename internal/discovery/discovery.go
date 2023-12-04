package discovery

import (
	"errors"
	"fmt"

	"github.com/matDobek/gov--attendance-check/internal/cache"
	"github.com/matDobek/gov--attendance-check/internal/cache/factory"
	"github.com/matDobek/gov--attendance-check/internal/db"
	httpclient "github.com/matDobek/gov--attendance-check/internal/http_client"
	"github.com/matDobek/gov--attendance-check/internal/parsers/html"
)

type StatueToParse struct {
	q   StatueQuery
	str string
}

type StatueQuery struct {
	termNo    int
	sessionNo int
	votingNo  int
	symbol    string
}

func (e *StatueQuery) url() string {
	return fmt.Sprintf(
		"https://www.sejm.gov.pl/Sejm10.nsf/agent.xsp?symbol=%v&NrKadencji=%v&NrPosiedzenia=%v&NrGlosowania=%v",
		e.symbol,
		e.termNo,
		e.sessionNo,
		e.votingNo,
	)
}

var request = initRequest()

func initRequest() func(string) (string, error) {
	var init bool
	var c *httpclient.HttpClient
	var err error

	return func(url string) (string, error) {
		if !init {
			var cache cache.Cache
			cache, err = factory.FileCache()
			c = httpclient.New(cache)
		}

		if err != nil {
			return "", err
		}

		resp, err := c.CachedGet(url)
		if err != nil {
			return "", err
		}

		return string(resp), nil
	}
}

func Run() ([]db.Statue, error) {
	result := []db.Statue{}

	statuesToParse, err := getStatuesToParse()
	if err != nil {
		return nil, err
	}

	for _, statueToParse := range statuesToParse {
		if statueToParse.q.symbol == "glosowaniaL" { // fixme, different parser required
			continue
		}

		//
		// {
		//		{Pkt. 25 Zmiany w składach osobowych komisji sejmowych (druk no 70) },
		//		{głosowanie nad przyjęciem wniosku z druku.},
		// }
		//
		titles, err := extractAndZip(statueToParse.str, [][]string{
			{"body #title_content p.subbig"}, // title
		})
		var title string
		for _, i := range titles {
			for _, j := range i {
				title = title + j
			}
		}

		//
		// Results will look like this:
		// {
		//   {"PiS", "part/of/url/1"}
		//   {"KO",  "part/of/url/2"}
		//   ...
		// }
		//
		clubs, err := extractAndZip(statueToParse.str, [][]string{
			{"body #title_content table tbody tr td:1 a"},         // club name
			{"body #title_content table tbody tr td:1 a", "href"}, // url - club voting results
		})

		if err != nil {
			return []db.Statue{}, err
		}

		votes := []db.Vote{}
		for _, club := range clubs {
			clubName := club[0]
			url := "https://www.sejm.gov.pl/Sejm10.nsf/" + club[1]

			doc, err := request(url)
			if err != nil {
				return result, err
			}

			//
			// Note: results are not grouped, but will follow a pattern:
			// {
			//		{"1."}, {"Jan Kowalski"}, {"za"},
			//		{"2."}, {"Jan Nowak"},		{"za"},
			//		...
			// }
			//
			votingResults, err := extractAndZip(doc, [][]string{
				{"body #title_content table tbody tr td"},
			})

			if err != nil {
				return result, err
			}

			v := db.Vote{Party: clubName}
			for i, vr := range votingResults {
				switch i % 3 {
				case 0:
				case 1:
					v.Name = vr[0]
				case 2:
					v.Response = vr[0]
					votes = append(votes, v)
				}
			}
		}

		ving := db.Statue{
			Title:     title,
			VotingNo:  statueToParse.q.votingNo,
			SessionNo: statueToParse.q.sessionNo,
			TermNo:    statueToParse.q.termNo,
			Votes:     votes,
		}

		result = append(result, ving)
	}

	return result, err
}

func getStatuesToParse() ([]StatueToParse, error) {
	statuesToParse := []StatueToParse{}
	statueQuery := StatueQuery{
		termNo:    10,
		sessionNo: 1,
		votingNo:  0,
	}

	for found := true; found; {
		for _, sq := range nextStatues(statueQuery) {
			response, err := request(sq.url())

			if err != nil {
				found = false
				continue
			}

			found = true
			statueQuery = sq
			statuesToParse = append(statuesToParse, StatueToParse{sq, string(response)})
			break
		}
	}

	return statuesToParse, nil
}

func nextStatues(e StatueQuery) []StatueQuery {
	var dx []StatueQuery

	nextVoting := StatueQuery{
		termNo:    e.termNo,
		sessionNo: e.sessionNo,
		votingNo:  e.votingNo + 1,
	}

	nextSession := StatueQuery{
		termNo:    e.termNo,
		sessionNo: e.sessionNo + 1,
		votingNo:  1,
	}

	symbols := []string{"glosowania", "glosowaniaL"}

	//
	//	1/ Possible outputs for current session in front
	//	2/ Possible outputs for next session later
	//

	for _, symbol := range symbols {
		nextVoting.symbol = symbol

		dx = append(dx, nextVoting)
	}

	for _, symbol := range symbols {
		nextSession.symbol = symbol

		dx = append(dx, nextSession)
	}

	return dx
}

func extractAndZip(doc string, extractArgs [][]string) ([][]string, error) {
	var out [][]string

	var extracted [][]string
	for _, args := range extractArgs {
		var e []string
		var err error

		if len(args) == 1 {
			e, err = html.Extract(doc, args[0])
		} else {
			e, err = html.ExtractAttr(doc, args[0], args[1])
		}

		if err != nil {
			return out, err
		}

		extracted = append(extracted, e)
	}

	out, ok := zip(extracted...)
	if !ok {
		err := errors.New("Number of statues, does not mach number of links to statues")
		return out, err
	}

	return out, nil
}

func zip[T any](xs ...[]T) ([][]T, bool) {
	result := [][]T{}
	ok := true

	var lengths []int
	for _, x := range xs {
		lengths = append(lengths, len(x))

		if len(x) != len(xs[0]) {
			ok = false
		}
	}

	for i := 0; i < min(lengths...); i++ {
		elem := []T{}

		for _, x := range xs {
			elem = append(elem, x[i])
		}

		result = append(result, elem)
	}

	return result, ok
}

func min(xs ...int) int {
	m := xs[0]

	for _, x := range xs {
		if x < m {
			x = m
		}
	}

	return m
}
