package main

import (
	"fmt"
	"sort"

	"github.com/matDobek/gov--attendance-check/internal/discovery"
	"github.com/matDobek/gov--attendance-check/internal/logger"
)

func main() {
	statues, err := discovery.Run()

	if err != nil {
		fmt.Println(err)
	}

	// responses := []string{"Nieobecny", "Przeciw", "Wstrzymał się", "Za"}

	truthTeller(statues, "Nieobecny", "Najwiekszy obibok [nieobecny]")
	truthTeller(statues, "Wstrzymał się", "Najbardziej wstrzemięźliwy [wstrzymał się]")
	truthTeller(statues, "Za", "yes men [za]")
	truthTeller(statues, "Przeciw", "no men [przeciw]")
}

func truthTeller(statues []discovery.Statue, response string, msg string) {
	logger.Info("========  %v  ========", msg)
	votes := make(map[string]int)
	for _, s := range statues {
		for _, v := range s.Votes {
			if v.Response == response {
				votes[v.Name] += 1
			}
		}
	}

	keys := make([]string, 0, len(votes))
	for k := range votes {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return votes[keys[i]] > votes[keys[j]]
	})

	l := len(keys)
	if l > 5 {
		l = 5
	}
	for _, k := range keys[:l] {
		logger.Info("%v/%v - %v", votes[k], len(statues), k)
	}
}
