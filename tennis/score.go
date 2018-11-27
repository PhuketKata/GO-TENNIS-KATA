package score

import "fmt"

type ScoreBoard struct {
	a int
	b int
}

func NewScoreBoard() ScoreBoard {
	return ScoreBoard{0, 0}
}

func (s *ScoreBoard) AGetPoint() string {
	if s.a == 3 && s.b < 3 {
		return fmt.Sprintf("A WIN")
	} else if s.a == 4 {
		return fmt.Sprintf("A WIN")
	}
	aPoint, bPoint := pointCalc(s.a, s.b, "A")
	s.a = aPoint
	s.b = bPoint
	aPointText := pointToString(s.a)
	bPointText := pointToString(s.b)
	scoreText := fmt.Sprintf("%s - %s", aPointText, bPointText)
	return scoreText
}

func (s *ScoreBoard) BGetPoint() string {
	if s.b == 3 && s.a < 3 {
		return fmt.Sprintf("B WIN")
	} else if s.b == 4 {
		return fmt.Sprintf("B WIN")
	}
	aPoint, bPoint := pointCalc(s.a, s.b, "B")
	s.a = aPoint
	s.b = bPoint
	aPointText := pointToString(s.a)
	bPointText := pointToString(s.b)
	scoreText := fmt.Sprintf("%s - %s", aPointText, bPointText)
	return scoreText
}

func pointCalc(aPoint int, bPoint int, lastPoint string) (int, int) {
	if aPoint == 4 && lastPoint == "B" {
		nextApoint := aPoint - 1
		return nextApoint, bPoint
	} else if bPoint == 4 && lastPoint == "A" {
		nextBpoint := bPoint - 1
		return aPoint, nextBpoint
	} else if lastPoint == "A" {
		nextApoint := aPoint + 1
		return nextApoint, bPoint
	} else {
		nextBpoint := bPoint + 1
		return aPoint, nextBpoint
	}
}

func pointToString(point int) string {
	score := [5]string{"0", "15", "30", "40", "setpoint"}
	return score[point]
}
