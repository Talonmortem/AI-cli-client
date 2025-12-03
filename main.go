package main

import (
	"bufio"
	"flag"
	"fmt"

	"os"
	"sync"

	"github.com/Talonmortem/AI-cli-client/AI-cli-client/client"
	"github.com/Talonmortem/AI-cli-client/AI-cli-client/host"
	"github.com/Talonmortem/AI-cli-client/AI-cli-client/logger"
	"github.com/charmbracelet/glamour"

	pm "github.com/Talonmortem/AI-cli-client/AI-cli-client/promptmanager"
	"github.com/Talonmortem/AI-cli-client/AI-cli-client/requests"
)

// Input represents the parsed input from CLI
// Mode: "flag" | "interactive"
// Query: user question

type Input struct {
	Mode  string
	Query string
}

var wg sync.WaitGroup
var log = logger.For("main")

func ParseInput() Input {
	// Define flags
	question := flag.String("f", "", "Вопрос для помощника")
	flag.Parse()

	// Case 1: user provided -f "..."
	if *question != "" {
		return Input{
			Mode:  "flag",
			Query: *question,
		}
	}

	// Case 2: arguments without flag (./aihelper "как сжать файл")
	// flag.Args() returns non-flag arguments
	args := flag.Args()
	if len(args) > 0 {
		return Input{
			Mode:  "flag",
			Query: args[0],
		}
	}

	// Case 3: fallback — interactive mode
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return Input{
			Mode:  "interactive",
			Query: scanner.Text(),
		}
	}

	return Input{Mode: "interactive", Query: ""}
}

func main() {
	debugMode := false // можно брать из флага
	isGlamour := true  // красивый рендер в терминале
	err := logger.Init("app.log", debugMode)
	if err != nil {
		log.Errorf("Failed to initialize logger: %v", err)
		return
	}

	ws := initWSClient()
	host, _ := host.NewHost()

	wg.Add(1)
	ws.StartReader(isGlamour)

	in := ParseInput()

	prompt := pm.GeneratePrompt(host)

	//prepare request
	req := requests.Request{
		Prompt:   prompt,
		Text:     in.Query,
		Model:    "qwen2.5-coder:14b-instruct",
		Role:     "assistant",
		Provider: "ollama",
		APIKey:   "",
	}

	//send request to server
	fmt.Println()
	resp, err := requests.SendRequest(req)
	if err != nil {
		log.Printf("Error sending request: %v. Response: %v", err, resp)
	} else {
		log.Println(resp)
	}
	wg.Wait()
}

func initWSClient() *client.WSClient {
	ws, err := client.NewWSClient("http://localhost:8082/ws")
	if err != nil {
		panic(err)
	}

	ws.TokenHandler = func(tok string) {
		fmt.Print(tok)
	}

	ws.FullResponseHandler = func(response string) {
		log.Printf("len: %v \n", len(response))

		out, err := glamour.Render(response, "dark")
		if err != nil {
			fmt.Printf("Ошибка при рендеринге Markdown: %v\n", err)
			fmt.Println(response) // fallback
			return
		}
		fmt.Print(out)
	}

	ws.DoneHandler = func() {
		wg.Done()
	}

	return ws
}
