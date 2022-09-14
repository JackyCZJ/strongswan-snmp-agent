package snmp

import (
	"github.com/posteo/go-agentx"
)

const DefaultSocketPath = "/var/agentx/master"

func NewSession(network, address string) (*agentx.Session, error) {
	agent, err := agentx.Dial(network, address)
	if err != nil {
		return nil, err
	}
	agentSession, err := agent.Session()
	if err != nil {
		return nil, err
	}
	return agentSession, nil
}
