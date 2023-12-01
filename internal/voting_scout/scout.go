package voting_scout

import (
	"errors"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	cache_factory "github.com/matDobek/gov--attendance-check/internal/cache/factory"
	httpclient "github.com/matDobek/gov--attendance-check/internal/http_client"
	"github.com/matDobek/gov--attendance-check/internal/parsers/html"
)

type Voting struct {
	num        int
	officeTerm int
	session    int
	name       string
	votes      []Vote
	votes_no   int
}

type Vote struct {
	party string
	name  string
	val   string
}

//
// TODO
// link does change: "https://www.sejm.gov.pl/Sejm10.nsf/agent.xsp?symbol=glosowania&NrKadencji=10&NrPosiedzenia=1&NrGlosowania=23"
// previous request ( and still missing functionality to get paginated ) is nn [or maybe theres no pagination, dunno]
// -> Yeah, Statue Title is way better
//		ScourStatues(...)
// -> also I can get additional info like Date and Time
//
// TODO rate limiter
// time.Sleep(2 * time.Second)
//
// TODO: 1 resultat, wymaga osobnego parsera: "Wybór Marszałka Sejmu RP"
// -> we can delete for now
//
// ToDO: filter stuff like: "propozycji wysluchania oswiadczen", "wniosek o przerwe", "przejscie do porzadku dziennego"
//
// TODO: log url / status / err when doing http request, or if it was takan from cache
// Extend Client to return ErrNotFound, and other status codes? [guess Im only expecting 200]
//

type Discovered struct {
	termNo    int
	sessionNo int
	votingNo  int
	symbol    string
	url       string
}

// FIXME clean me, this is a mess
func Discovery() ([]string, error) {
	cache, err := cache_factory.FileCache()
	if err != nil {
		return nil, err
	}
	client := httpclient.New(cache)

	termNo := 10
	sessionNo := 1
	votingNo := 1

	symbols := []string{"glosowania", "glosowaniaL"}
	currSymbol := 0

	retry := 0
	votings := []string{}
	keepGoing := true
	for keepGoing {
		time.Sleep(500 * time.Millisecond) // FIXME dont sleep on cache hit

		url := fmt.Sprintf("https://www.sejm.gov.pl/Sejm10.nsf/agent.xsp?symbol=%v&NrKadencji=%v&NrPosiedzenia=%v&NrGlosowania=%v", symbols[currSymbol], termNo, sessionNo, votingNo)
		response, err := client.CachedGet(url)

		switch {
		case errors.Is(err, httpclient.ErrStatusClientError): // FIXME provide more granual status errors like 404
			if currSymbol == 0 {
				currSymbol++
			} else if votingNo == 1 {
				keepGoing = false
			} else {
				currSymbol = 0
				votingNo = 1
				sessionNo++
			}
		case errors.Is(err, httpclient.ConnectionError{}):
			if retry > 2 {
				return nil, err
			}
			retry++
			time.Sleep(5 * time.Second)
		case err != nil:
			return nil, err
		default:
			votings = append(votings, string(response))
			votingNo++
			retry = 0
			currSymbol = 0
		}
	}

	return votings, nil
}

// Fixme: skip votings with symbol "glosowaniaL"  due to different format
func Run() ([]Voting, error) {
	result := []Voting{}

	cache, err := cache_factory.FileCache()
	if err != nil {
		return result, err
	}
	client := httpclient.New(cache)

	doc, err := trafficControl(client, "https://www.sejm.gov.pl/Sejm10.nsf/agent.xsp?symbol=GLOSNAPOS&NrPos=1&NrKadencji=10")
	if err != nil {
		return result, err
	}

	statues, err := extractAndZip(doc,
		[][]string{
			{"body #title_content table tbody tr td:4 a"},
			{"body #title_content table tbody tr td:1 a", "href"},
		},
	)

	if err != nil {
		return result, err
	}

	for _, statue := range statues {
		name := statue[0]
		url := "https://www.sejm.gov.pl/Sejm10.nsf/" + statue[1]

		doc, err := trafficControl(client, url)
		if err != nil {
			return []Voting{}, err
		}

		//
		// Results will look like this:
		// {
		//   {"PiS", "part/of/url/1"}
		//   {"KO",  "part/of/url/2"}
		//   ...
		// }
		//
		clubs, err := extractAndZip(doc, [][]string{
			{"body #title_content table tbody tr td:1 a"},         // club name
			{"body #title_content table tbody tr td:1 a", "href"}, // url - club voting results
		})

		if err != nil {
			return result, err
		}

		votes := []Vote{}
		for _, club := range clubs {
			clubName := club[0]
			url = "https://www.sejm.gov.pl/Sejm10.nsf/" + club[1]

			doc, err := trafficControl(client, url)
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

			v := Vote{party: clubName}
			for i, vr := range votingResults {
				switch i % 3 {
				case 0:
				case 1:
					v.name = vr[0]
				case 2:
					v.val = vr[0]
					votes = append(votes, v)
				}
			}
		}

		ving := Voting{
			name:       name,
			num:        -1, // TODO
			session:    -1,
			officeTerm: -1,
			votes:      votes,
			votes_no:   len(votes),
		}

		result = append(result, ving)
	}

	spew.Dump(result)

	return result, err
}

func trafficControl(client *httpclient.HttpClient, url string) (string, error) {
	response, err := client.CachedGet(url)

	switch {
	case errors.Is(err, httpclient.ConnectionError{}):
		// TODO: retry
		return "", err
	case err != nil:
		return "", err
	}

	return string(response), nil
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
