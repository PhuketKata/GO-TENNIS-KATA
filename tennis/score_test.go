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
