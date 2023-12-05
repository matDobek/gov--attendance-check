package db

const (
	VoteResponseAbsent = "Nieobecny"
	VoteResponseYes    = "Za"
	VoteResponseNo     = "Przeciw"
	VoteResponseMaybe  = "Wstrzymał się"
)

type GovStore struct {
	statues     []Statue
	politicians []Politician
	votes       []Vote
}

type Statue struct {
	Id        int    `json:"id"`
	VotingNo  int    `json:"votingNumber"`
	SessionNo int    `json:"sessionNumber"`
	TermNo    int    `json:"termNumber""`
	Title     string `json:"title"`
}

type Vote struct {
	Id           int    `json:"id"`
	PoliticianId int    `json:"politicianId"`
	StatueId     int    `json:"statueId"`
	Response     string `json:"response"`
}

type Politician struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Party string `json:"party"`
}

func NewGovStore(statues []Statue, politicians []Politician, votes []Vote) *GovStore {
	return &GovStore{
		statues:     statues,
		politicians: politicians,
		votes:       votes,
	}
}

func (s *GovStore) GetStatues() []Statue {
	return s.statues
}

func (s *GovStore) GetVotes() []Vote {
	return s.votes
}

func (s *GovStore) GetPoliticians() []Politician {
	return s.politicians
}

func (s *GovStore) GetTopPoliticians() []Politician {
	return s.politicians
}

// truthTeller(statues, "Nieobecny", "Najwiekszy obibok [nieobecny]")
// truthTeller(statues, "Wstrzymał się", "Najbardziej wstrzemięźliwy [wstrzymał się]")
// truthTeller(statues, "Za", "yes men [za]")
// truthTeller(statues, "Przeciw", "no men [przeciw]")
// func truthTeller(statues []Statue, response string, msg string) {
// 	logger.Info("========  %v  ========", msg)
// 	votes := make(map[string]int)
// 	for _, s := range statues {
// 		for _, v := range s.Votes {
// 			if v.Response == response {
// 				votes[v.Name] += 1
// 			}
// 		}
// 	}
//
// 	keys := make([]string, 0, len(votes))
// 	for k := range votes {
// 		keys = append(keys, k)
// 	}
//
// 	sort.Slice(keys, func(i, j int) bool {
// 		return votes[keys[i]] > votes[keys[j]]
// 	})
//
// 	l := len(keys)
// 	if l > 5 {
// 		l = 5
// 	}
// 	for _, k := range keys[:l] {
// 		logger.Info("%v/%v - %v", votes[k], len(statues), k)
// 	}
//
// }
