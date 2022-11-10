//go:build !windows

package rig

import (
	"fmt"
	"net"
	"os"

	"github.com/k0sproject/rig/log"
	"golang.org/x/crypto/ssh/agent"
)

func agentClient() (agent.Agent, error) {
	sshAgentSock := os.Getenv("SSH_AUTH_SOCK")
	if sshAgentSock == "" {
		return nil, fmt.Errorf("SSH_AUTH_SOCK is empty")
	}
	log.Debugf("using SSH_AUTH_SOCK=%s", sshAgentSock)
	sshAgent, err := net.Dial("unix", sshAgentSock)
	if err != nil {
		return nil, fmt.Errorf("can't connect to SSH agent: %w", err)
	}
	return agent.NewClient(sshAgent), nil
}