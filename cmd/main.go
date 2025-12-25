package main

import (
	"bufio"
	"connect4/internal/app/game"
	"connect4/internal/app/model"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("тут приветствие и правила игры")

	player1 := model.NewPlayer(readInput("введите имя первого игрока"), '⊗')
	player2 := model.NewPlayer(readInput("введите имя второго игрока"), '●')
	ctrl := game.NewGame(player1, player2)

	fmt.Println(ctrl.GetUserField())

	for {
		if err := ctrl.MakeDrop(readInput("Ход игрока " + ctrl.CurrentPlayer().Name + ": ")); err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(ctrl.GetUserField())
	}

	// TODO

	//if ctrl.IsGameOver() {
	//fmt.Println("объявление победителя")
	//break

}

// приветствия и объяснение правил игры перед первым ходом
// можно запросить правила по кодовому слову
func readInput(p string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(p)

	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}
