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
