package model

type Player struct {
	Name  string
	Chips int
	Sign  int32
}

func NewPlayer(name string, sign int32) Player {
	return Player{
		Name:  name,
		Chips: 21,
		Sign:  sign,
	}
}
