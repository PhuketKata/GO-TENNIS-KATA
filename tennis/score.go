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
	s.aPoint++
	aPointText := pointCalc(s.aPoint)
	bPointText := pointCalc(s.bPoint)
	scoreText := fmt.Sprintf("%s - %s", aPointText, bPointText)
	return scoreText
}

func (s *ScoreBoard) BGetPoint() string {
	s.bPoint++
	aPointText := pointCalc(s.aPoint)
	bPointText := pointCalc(s.bPoint)
	scoreText := fmt.Sprintf("%s - %s", aPointText, bPointText)
	return scoreText
}

func pointCalc(point int) string {
	score := [5]string{"0", "15", "30", "40", "setpoint"}
	return score[point]
}
