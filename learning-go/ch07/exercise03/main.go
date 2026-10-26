package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(firstTeam string, scoreFirstTeam int, secondTeam string, scoreSecondTeam int) {
	if _, ok := l.Teams[firstTeam]; !ok {
		return
	}
	if _, ok := l.Teams[secondTeam]; !ok {
		return
	}
	if scoreFirstTeam == scoreSecondTeam {
		return
	}
	if scoreFirstTeam > scoreSecondTeam {
		l.Wins[firstTeam]++
	} else {
		l.Wins[secondTeam]++
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	return names
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(ranker Ranker, writer io.Writer) {
	result := ranker.Ranking()
	for _, r := range result {
		io.WriteString(writer, r)
		writer.Write([]byte("\n"))
	}
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Germany": {
				Name:    "Germany",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Spain": {
				Name:    "Spain",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 100, "France", 70)
	l.MatchResult("Germany", 85, "Spain", 80)
	l.MatchResult("Italy", 100, "Germany", 55)
	l.MatchResult("France", 100, "Spain", 110)
	l.MatchResult("Italy", 100, "Spain", 70)
	l.MatchResult("France", 95, "Germany", 80)
	RankPrinter(l, os.Stdout)
}
