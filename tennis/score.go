package score

type ScoreBoard struct{}

func NewScoreBoard() ScoreBoard {
	return ScoreBoard{}
}

func (s ScoreBoard) AGetPoint() string {
	return "15 - 0"
}
