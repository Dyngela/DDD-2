package CLI

import (
	"awesomeProject/pkg/PFC"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ICli interface {
	Init(historic PFC.IHistoric)
}

type cli struct{}

func NewCli() ICli {
	return &cli{}
}

func (c cli) Init(historic PFC.IHistoric) {

	fmt.Println("Simple Shi Fu Mi")
	fmt.Println("---------------------")
	fmt.Println("Type score to show the current score")
	fmt.Println("Type exit to stop the game")
	fmt.Println("Type 0 to play Rock")
	fmt.Println("Type 1 to play Paper")
	fmt.Println("Type 2 to play Cisor")
	handleInput(historic)
}

func handleInput(historic PFC.IHistoric) {
	game := PFC.NewGame(historic)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if strings.Compare("score", text) == 0 {
			game.GetHistoric()
			handleInput(historic)
		} else if strings.Compare("0", text) == 0 {
			game.Play(PFC.Rock)
			handleInput(historic)
		} else if "1" == text {
			game.Play(PFC.Paper)
			handleInput(historic)
		} else if strings.Compare("2", text) == 0 {
			game.Play(PFC.Cisor)
			handleInput(historic)
		} else if strings.Compare("exit", text) == 0 {
			fmt.Println("final score :")
			game.GetHistoric()
			fmt.Println("-----------------------")
			game.End()
			os.Exit(1)
		} else {
			fmt.Println("input: " + text + " is unknown command please try again")
		}
	}
}
