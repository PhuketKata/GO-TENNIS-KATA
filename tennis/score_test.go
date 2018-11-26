package score

import (
	"testing"
)

func TestAGot15AndBGot0(t *testing.T) {
	scoreBoard := NewScoreBoard()
	scoreText := scoreBoard.AGetPoint()
	if scoreText != "15 - 0" {
		t.Errorf("score value should be 15 - 0 but %s", scoreText)
	}
}

func TestAGot30AndBGot0(t *testing.T) {
	ScoreBoard := NewScoreBoard()
	ScoreBoard.AGetPoint()
	scoreText := ScoreBoard.AGetPoint()
	if scoreText != "30 - 0" {
		t.Errorf("score value should be 30 - 0 but %s", scoreText)
	}
}

func TestAGot40AndBGot0(t *testing.T) {
	ScoreBoard := NewScoreBoard()
	ScoreBoard.AGetPoint()
	ScoreBoard.AGetPoint()
	scoreText := ScoreBoard.AGetPoint()
	if scoreText != "40 - 0" {
		t.Errorf("score value should be 40 - 0 but %s", scoreText)
	}
}

func TestAGot40AndBGot15(t *testing.T) {
	ScoreBoard := NewScoreBoard()
	ScoreBoard.AGetPoint()
	ScoreBoard.AGetPoint()
	ScoreBoard.AGetPoint()
	scoreText := ScoreBoard.BGetPoint()
	if scoreText != "40 - 15" {
		t.Errorf("score value should be 40 - 15 but %s", scoreText)
	}
}

func TestAGot40AndBGot30(t *testing.T) {
	ScoreBoard := NewScoreBoard()
	ScoreBoard.AGetPoint()
	ScoreBoard.AGetPoint()
	ScoreBoard.AGetPoint()
	ScoreBoard.BGetPoint()
	scoreText := ScoreBoard.BGetPoint()
	if scoreText != "40 - 30" {
		t.Errorf("score value should be 40 - 30 but %s", scoreText)
	}
}

func TestAGot40AndBGot40(t *testing.T) {
	ScoreBoard := NewScoreBoard()
	ScoreBoard.AGetPoint()
	ScoreBoard.AGetPoint()
	ScoreBoard.AGetPoint()
	ScoreBoard.BGetPoint()
	ScoreBoard.BGetPoint()
	scoreText := ScoreBoard.BGetPoint()
	if scoreText != "40 - 40" {
		t.Errorf("score value should be 40 - 40 but %s", scoreText)
	}
}
