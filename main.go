package main
import (
"fmt"
"bufio"
"os"
"errors"
"context"
"github.com/openai/openai-go/v3"
"github.com/openai/openai-go/v3/option"
)

func main() {

	fmt.Println("Welcome to GO AI Agent\n")

	var userName string = ""

	fmt.Print("Please enter your name:")
	fmt.Scan(&userName)

	fmt.Printf("\nWelcome %s \n", userName)
	fmt.Println("(type 'exit', 'quit' or 'q' to quit)")

	repl()
}

func repl() {

	for {

		userInput, err := prompt("\n > ")

		if err != nil {
			fmt.Println("Error:", err)
		}


		exitCommands := map[string]bool{"exit": true, "q": true, "quit": true}
		if exitCommands[userInput] { break }

		response, err := llm_request(userInput)

		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Printf("\nAI: %s \n", response)
	}
}

func prompt(msg string) (string, error) {
	fmt.Print(msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if scanner.Text() == "" {
		return "", errors.New("Please type something")
	}

	return scanner.Text(), nil
}

func llm_request(userInput string) (string, error) {

	client := openai.NewClient(
		option.WithBaseURL("http://192.168.1.202:8080"),
		option.WithAPIKey("not-needed"),
	)

	ctx := context.Background()

	completion, err := client.Chat.Completions.New(
		ctx,
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(userInput),
			},
			Model: "gpt-oss-120b",
		})

	if err != nil {
		return "", err
	}

	return completion.Choices[0].Message.Content, nil
}
