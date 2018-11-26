package score

import "fmt"

type ScoreBoard struct {
	aPoint int
	bPoint int
}

func NewScoreBoard() ScoreBoard {
	return ScoreBoard{0, 0}
}

func (s *ScoreBoard) AGetPoint() string {
	if s.bPoint == 4 {
		s.bPoint--
	} else {
		s.aPoint++
	}
	aPointText := pointCalc(s.aPoint)
	bPointText := pointCalc(s.bPoint)
	scoreText := fmt.Sprintf("%s - %s", aPointText, bPointText)
	return scoreText
}

func (s *ScoreBoard) BGetPoint() string {
	if s.aPoint == 4 {
		s.aPoint--

	} else {
		s.bPoint++
	}
	aPointText := pointCalc(s.aPoint)
	bPointText := pointCalc(s.bPoint)
	scoreText := fmt.Sprintf("%s - %s", aPointText, bPointText)
	return scoreText
}

func pointCalc(point int) string {
	score := [5]string{"0", "15", "30", "40", "setpoint"}
	return score[point]
}
