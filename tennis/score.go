package score

import "fmt"

type ScoreBoard struct {
	a     int
	b     int
	isEnd bool
}

func NewScoreBoard() ScoreBoard {
	return ScoreBoard{0, 0, false}
}

func (s *ScoreBoard) AGetPoint() string {
	if s.isEnd {
		return "error"
	}
	if isPlayerMatchPoint(s.a, s.b) {
		s.isEnd = true
		return "A WIN"
	}
	if isPlayerSetPoint(s.b, s.a) {
		s.b = 3
	} else {
		s.a++
	}
	return fmt.Sprintf("%s - %s", pointToString(s.a), pointToString(s.b))
}

func (s *ScoreBoard) BGetPoint() string {
	if s.isEnd {
		return "error"
	}
	if isPlayerMatchPoint(s.b, s.a) {
		s.isEnd = true
		return "B WIN"
	}
	if isPlayerSetPoint(s.a, s.b) {
		s.a = 3
	} else {
		s.b++
	}
	return fmt.Sprintf("%s - %s", pointToString(s.a), pointToString(s.b))
}

func pointToString(point int) string {
	score := [5]string{"0", "15", "30", "40", "setpoint"}
	return score[point]
}

func isPlayerMatchPoint(playerScore1 int, playerScore2 int) bool {
	return (playerScore1 == 3 && playerScore2 < 3) || (playerScore1 == 4 && playerScore2 == 3)
}

func isPlayerSetPoint(playerScore1 int, playerScore2 int) bool {
	return playerScore1 == 4 && playerScore2 == 3
}
