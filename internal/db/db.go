package db

type GovStore struct {
	statues []Statue
}

type Statue struct {
	VotingNo  int    `json:"id"`
	SessionNo int    `json:"sessionNumber"`
	TermNo    int    `json:"termNumber""`
	Title     string `json:"title"`
	Votes     []Vote `json:"votes"`
}

type Vote struct {
	Name     string `json:"name"`
	Party    string `json:"party"`
	Response string `json:"response"`
}

func NewGovStore(statues []Statue) *GovStore {
	return &GovStore{
		statues: statues,
	}
}

func (s *GovStore) GetStatues() []Statue {
	return s.statues
}
