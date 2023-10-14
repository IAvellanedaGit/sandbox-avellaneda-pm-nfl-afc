package afc

import (
	"math/rand"
	"strings"

	bridge "github.com/IAvellanedaGit/nuevo-platform-mobile-bridge"
)

type Service struct {
	b      *bridge.Bridge
	logger bridge.Logger
}

func NewService(b *bridge.Bridge, logger bridge.Logger) *Service {
	svc := &Service{
		b:      b,
		logger: logger,
	}

	b.RegisterService("nfl-afc", svc)

	return svc
}

func (s *Service) Teams() []string {
	return []string{
		"chiefs",
		"raiders",
		"broncos",
		"chargers",
		"steelers",
		"browns",
		"bengals",
		"ravens",
		"bills",
		"jets",
		"dolphins",
		"patriots",
		"texans",
		"colts",
		"jaguars",
		"titans",
	}
}

func (s *Service) ValidateTeam(team string) bool {

	t := strings.ToLower(team)
	for _, n := range s.Teams() {
		if t == n {
			return true
		}
	}

	return false
}

func (s *Service) PredictScore(team string) int {

	score := rand.Intn(63)
	if strings.ToLower(team) == "chiefs" {
		score = rand.Intn(80)
	}

	return score
}
