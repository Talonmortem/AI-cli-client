package prompt

import (
	"fmt"

	"github.com/Talonmortem/AI-cli-client/AI-cli-client/logger"

	"github.com/Talonmortem/AI-cli-client/AI-cli-client/host"
)

var log = logger.For("promptmanager")

func GeneratePrompt(host *host.Host) string {
	instruction := fmt.Sprintln("Instruction: Dont use markdown. Generate a response to the following query. Use best practise. Answer single line command if possible. Give a few options if needed.")
	os := fmt.Sprintln("OS:", host.HostInfo.Platform)
	osv := fmt.Sprintln("OS Version:", host.HostInfo.PlatformVersion)
	hm := fmt.Sprintln("Hostname:", host.HostInfo.Hostname)
	prompt := fmt.Sprintf("%s\n%s\n%s\n%s\n", instruction, os, osv, hm)
	log.Printf("Generated Prompt: %s", prompt)
	return prompt
}
