package voting

import (
	"fmt"
	"testing"
)

func TestBallotBoxCount(t *testing.T) {
	box := NewBallotBox()
	if box.Count() != 0 {
		t.Errorf("Newly created Ballot box should be empty")
	}

	box.Vote(0, 1)
	box.Vote(0, 1)
	box.Vote(0, 1)

	count := box.Count()

	if count != 3 {
		t.Errorf("Count resulted in %v instead of %v", count, 3)
	}
}

func TestBallotBoxResult(t *testing.T) {
	testCases := []struct {
		votes []int
		want  int
	}{
		{[]int{1, 1, 2}, 1},
		{[]int{1, 2, 2}, 2},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v -> %v", tc.votes, tc.want), func(t *testing.T) {
			box := NewBallotBox()
			for i := 0; i < len(tc.votes); i++ {
				box.Vote(tc.votes[i], 1)
			}
			vote := box.Result()
			if vote != tc.want {
				t.Errorf("got %v; want %v", vote, tc.want)
			}

		})
	}
}
