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
	if s.aPoint == 3 && s.bPoint < 3 {
		return fmt.Sprintf("A WIN")
	} else if s.aPoint == 4 {
		return fmt.Sprintf("A WIN")
	}
	aPoint, bPoint := pointCalc(s.aPoint, s.bPoint, "A")
	s.aPoint = aPoint
	s.bPoint = bPoint
	aPointText := pointToString(s.aPoint)
	bPointText := pointToString(s.bPoint)
	scoreText := fmt.Sprintf("%s - %s", aPointText, bPointText)
	return scoreText
}

func (s *ScoreBoard) BGetPoint() string {
	aPoint, bPoint := pointCalc(s.aPoint, s.bPoint, "B")
	s.aPoint = aPoint
	s.bPoint = bPoint
	aPointText := pointToString(s.aPoint)
	bPointText := pointToString(s.bPoint)
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
