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
	if s.aPoint == 0 {
		s.aPoint = 15
	} else if s.aPoint == 15 {
		s.aPoint = 30
	} else {
		s.aPoint = 40
	}
	scoreText := fmt.Sprintf("%d - %d", s.aPoint, s.bPoint)
	return scoreText
}

func (s *ScoreBoard) BGetPoint() string {
	if s.bPoint == 0 {
		s.bPoint = 15
	} else {
		s.bPoint = 30
	}
	scoreText := fmt.Sprintf("%d - %d", s.aPoint, s.bPoint)
	return scoreText
}
