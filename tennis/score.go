package score

type ScoreBoard struct {
	aPoint int
	bPoint int
}

func NewScoreBoard() ScoreBoard {
	return ScoreBoard{0, 0}
}

func (s *ScoreBoard) AGetPoint() string {
	var scoreText string
	if s.aPoint == 0 {
		s.aPoint = 15
		scoreText = "15 - 0"
	} else {
		s.aPoint = 30
		scoreText = "30 - 0"
	}
	return scoreText
}
