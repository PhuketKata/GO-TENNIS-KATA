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
	if s.aPoint == 40 {
		scoreText = fmt.Sprintf("setpoint - %d", s.bPoint)
	} else {
		nextPoint := pointCalc(s.aPoint)
		s.aPoint = nextPoint
		scoreText = fmt.Sprintf("%d - %d", s.aPoint, s.bPoint)
	}
	return scoreText
}

func (s *ScoreBoard) BGetPoint() string {
	var scoreText string
	if s.bPoint == 40 {
		scoreText = fmt.Sprintf("%d - setpoint", s.aPoint)
	} else {
		nextPoint := pointCalc(s.bPoint)
		s.bPoint = nextPoint
		scoreText = fmt.Sprintf("%d - %d", s.aPoint, s.bPoint)
	}
	return scoreText
}

func pointCalc(point int) int {
	var nextPoint int
	if point == 0 {
		nextPoint = 15
	} else if point == 15 {
		nextPoint = 30
	} else {
		nextPoint = 40
	}
	return nextPoint
}
