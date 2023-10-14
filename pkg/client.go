package afc

import (
	"fmt"

	bridge "bitbucket.org/newyuinc/platform-mobile-bridge"
	iafc "bitbucket.org/newyuinc/sandbox-klee-pm-nfl-afc/service"
)

type Client struct {
	svc *iafc.Service
}

func NewClient(b *bridge.Bridge) (*Client, error) {
	svc, ok := b.ServiceInstance("nfl-afc").(*iafc.Service)
	if !ok || svc == nil {
		return nil, fmt.Errorf("could not locate AFC service")
	}

	return &Client{
		svc: svc,
	}, nil
}

func (c *Client) Teams() []string {
	return c.svc.Teams()
}

func (c *Client) ValidateTeam(team string) bool {
	return c.svc.ValidateTeam(team)
}

func (c *Client) PredictScore(team string) int {
	return c.svc.PredictScore(team)
}
