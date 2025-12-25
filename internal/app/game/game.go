package game

import (
	"connect4/internal/app/model"
	"errors"
	"fmt"
	"strconv"
)

// 7 columns × 6 rows
// знает правила игры:
// чей ход
// когда игра окончена (выигрыш или закончились фишки/пустые поля)

// показывает анимацию падения фишки
// поле
// сколько фишек у текущего игрока

const fieldSizeColum = 7
const fieldSizeRow = 6
const cellEmpty = '◌'

type Game struct {
	field         [][]rune
	players       map[int]model.Player
	currentPlayer int
}

func NewGame(players ...model.Player) *Game {
	field := make([][]rune, 0, fieldSizeRow)
	for i := 0; i < fieldSizeRow; i++ {
		r := make([]rune, 0, fieldSizeColum)
		for j := 0; j < fieldSizeColum; j++ {
			r = append(r, cellEmpty)
		}
		field = append(field, r)
	}

	playersMap := make(map[int]model.Player)
	for n, pl := range players {
		playersMap[n] = pl
	}

	return &Game{
		players:       playersMap,
		currentPlayer: 0,
		field:         field,
	}
}

func (c *Game) GetUserField() string {
	s := " 1234567  \n"
	s += "┌" + repeat('─', fieldSizeColum) + "┐\n"
	for _, row := range c.field {
		str := ""
		for _, cell := range row {
			str += string(cell)
		}
		s += fmt.Sprintf("│%s│\n", str)
	}

	s += "└" + repeat('─', fieldSizeColum) + "┘  "
	return s
}

func (c *Game) MakeDrop(s string) error { // p - текущий игрока
	y, err := c.checkDrop(s)
	if err != nil {
		return err
	}

	for x := fieldSizeRow - 1; x >= 0; x-- {
		if c.field[x][y-1] == cellEmpty {
			c.field[x][y-1] = c.CurrentPlayer().Sign
			c.changePlayer()
			return nil
		}
	}

	return errors.New("колонка заполнена")
}

func (c Game) CurrentPlayer() model.Player {
	return c.players[c.currentPlayer]
}

func (c *Game) SlowMoDrop() {
	//todo
}

func (c *Game) IsGameOver() bool {
	//todo
	//если закончились ходы - проиграл
	return true
}

func (c *Game) Win() bool {
	//todo
	return true
}

func (c *Game) changePlayer() {
	c.currentPlayer++
	if c.currentPlayer > len(c.players)-1 {
		c.currentPlayer = 0
	}
}

func (c *Game) checkDrop(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("Введите номер колонки")
	}
	if i < 1 || i > 7 {
		return 0, errors.New("Введите номер колонки от 1 до 7")
	}
	return i, nil
}

func repeat(s rune, n int) string {
	str := ""
	for i := 0; i < n; i++ {
		str += string(s)
	}
	return str
}
