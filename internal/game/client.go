package game

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"strconv"
)

var (
	defStyle        = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	yellowCharStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorYellow)
	redCharStyle    = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorRed)
	greenCharStyle  = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorGreen)
	scoreStyle      = tcell.StyleDefault.Background(tcell.Color20).Foreground(tcell.ColorWhite)
)

func showText(
	s tcell.Screen,
	x, y int,
	currentCharHighlight, prevTextStyle, defaultStyle tcell.Style,
	str string,
	letterIndex int,
) {
	var xs int
	for index, char := range str {
		if char == '\n' {
			xs = 0
			y += 1
			continue
		} else if index < letterIndex {
			s.SetContent(xs+x, y, char, nil, prevTextStyle)
		} else if index == letterIndex {
			s.SetContent(xs+x, y, char, nil, currentCharHighlight)
		} else {
			s.SetContent(xs+x, y, char, nil, defaultStyle)
		}
		xs += 1
	}
}

func showEndGame(s tcell.Screen, wpm, score, totalPress, typosCount int) {
	const (
		hor  = '-'
		vert = '|'
	)
	deffaultStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorWhite)
	accuracyInfoStyle := tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorWhite)
	typosInfoStyle := tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tcell.ColorWhite)
	styleEsc := tcell.StyleDefault.Background(tcell.Color20).Foreground(tcell.ColorWhite)
	// draw box
	h := 10
	w := 35
	for i := 1; i < w; i++ {
		s.SetContent(i, 0, hor, nil, deffaultStyle)
		s.SetContent(i, h, hor, nil, deffaultStyle)
	}
	for i := 0; i < h+1; i++ {
		s.SetContent(0, i, vert, nil, deffaultStyle)
		s.SetContent(w, i, vert, nil, deffaultStyle)
	}
	offset := 2
	resultString := "Your word per minute is: " + strconv.Itoa(wpm)
	for i, char := range resultString {
		s.SetContent(i+offset, 2, char, nil, deffaultStyle)
	}
	accuracy := 100 * score / totalPress
	if accuracy < 0 {
		accuracy = 0
	}
	accuracyInfo := "Your accuracy is: " + strconv.Itoa(accuracy) + "%"
	for i, char := range accuracyInfo {
		s.SetContent(i+offset, 4, char, nil, accuracyInfoStyle)
	}
	typosInfo := "You've done " + strconv.Itoa(typosCount) + " typos"
	for i, char := range typosInfo {
		s.SetContent(i+offset, 5, char, nil, typosInfoStyle)
	}
	espString := "Press Esc to exit"
	for i, char := range espString {
		s.SetContent(i+offset, 8, char, nil, styleEsc)
	}
}

func showScore(s tcell.Screen, x, y, wpm int, style tcell.Style) {
	// clear
	s.SetContent(x, y, ' ', nil, tcell.StyleDefault.Background(tcell.ColorReset))
	s.SetContent(x+1, y, ' ', nil, tcell.StyleDefault.Background(tcell.ColorReset))
	s.SetContent(x+2, y, ' ', nil, tcell.StyleDefault.Background(tcell.ColorReset))
	// show wpm
	str := strconv.Itoa(wpm)
	for i, char := range str {
		s.SetContent(x+i, y, char, nil, style)
	}
}

func SoloTyper() {
	var textX = 0
	var textY = 0
	var scoreX = 123
	var scoreY = 6

	// Initialize screen
	s, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	quit := func() {
		// You have to catch panics in defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	tg := TyperGame{}
	err = tg.StartGame()
	showText(s, textX, textY, defStyle, defStyle, defStyle, string(tg.Text), 0)
	//drawBox(s, boxX, boxY, redCharStyle)
	for {
		showScore(s, scoreX, scoreY, tg.PlayerScore, scoreStyle)
		s.Sync()
		if tg.GameState == false {
			s.Clear()
			showEndGame(s, tg.WPM, tg.PlayerScore, tg.PressCounter, tg.TyposCount)
			s.Sync()
			for {
				ev := s.PollEvent()
				switch ev := ev.(type) {
				case *tcell.EventKey:
					if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
						return
					}
				}
			}
		}
		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else {
				index, correct := tg.HandleKey(ev.Rune())
				if correct {
					showText(s, textX, textY, yellowCharStyle, greenCharStyle, defStyle, string(tg.Text), index)
				} else {
					showText(s, textX, textY, redCharStyle, greenCharStyle, defStyle, string(tg.Text), index)
				}
			}
		}

	}
}