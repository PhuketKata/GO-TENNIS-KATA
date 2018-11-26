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
	var scoreText string
	if s.aPoint == 3 {
		bPointText := pointCalc(s.bPoint)
		scoreText = fmt.Sprintf("setpoint - %s", bPointText)
	} else {
		s.aPoint++
		aPointText := pointCalc(s.aPoint)
		bPointText := pointCalc(s.bPoint)
		scoreText = fmt.Sprintf("%s - %s", aPointText, bPointText)
	}
	return scoreText
}

func (s *ScoreBoard) BGetPoint() string {
	var scoreText string
	if s.bPoint == 3 {
		aPointText := pointCalc(s.aPoint)
		scoreText = fmt.Sprintf("%s - setpoint", aPointText)
	} else {
		s.bPoint++
		aPointText := pointCalc(s.aPoint)
		bPointText := pointCalc(s.bPoint)
		scoreText = fmt.Sprintf("%s - %s", aPointText, bPointText)
	}
	return scoreText
}

func pointCalc(point int) string {
	score := [4]string{"0", "15", "30", "40"}
	return score[point]
}
