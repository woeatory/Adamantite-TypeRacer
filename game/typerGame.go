package game

import (
	"os"
	"time"
)

const GAME_TIME = time.Second * 3

type TyperGame struct {
	Text          []rune
	PlayerScore   int
	currentRune   rune
	index         int
	GameState     bool
	WordsCount    int
	WordsComplete int
	WPM           int
}

func (tg *TyperGame) initGame(text string) {
	tg.Text = []rune(text)
	tg.currentRune = tg.Text[0]
}

func (tg *TyperGame) HandleKey(pressed rune) (int, bool) {
	var correct bool
	if tg.currentRune != pressed {
		tg.PlayerScore--
		correct = false
	} else {
		tg.PlayerScore++
		tg.index++
		for true {
			if tg.Text[tg.index] == '\n' || tg.Text[tg.index] == '\r' {
				tg.index++
			} else {
				break
			}
		}
		tg.currentRune = tg.Text[tg.index]
		correct = true
	}
	if tg.index > len(tg.Text) {
		tg.GameState = false
	}
	if tg.Text[tg.index] == '\n' || tg.Text[tg.index] == ' ' {
		tg.WordsComplete++
		tg.WPM = tg.WordsComplete
	}
	return tg.index, correct
}

func (tg *TyperGame) StartGame() error {
	text, err := os.ReadFile("game/texts/text.txt")
	if err != nil {
		return err
	}
	tg.Text = []rune(string(text))
	for _, ch := range tg.Text {
		if ch == ' ' || ch == '\n' {
			tg.WordsCount++
		}
	}
	tg.currentRune = tg.Text[0]
	tg.GameState = true
	go func() {
		time.Sleep(GAME_TIME)
		tg.GameState = false
	}()
	return nil
}
