package game_test

import (
	"github.com/devashishRaj/goTests/game"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestListItems(t *testing.T) {
	t.Parallel()
	type testcase struct {
		input []string
		want  string
	}
	testcases := []testcase{
		// multiple input
		{
			input: []string{
				"a battery",
				"a key",
				"a tourist map",
			},
			want: "You can see here a battery, a key, and a tourist map.",
		},
		// dual input
		{
			input: []string{
				"a battery",
				"a key",
			},
			want: "You can see here a battery and a key.",
		},
		// single input
		{
			input: []string{
				"a battery",
			},
			want: "You can see a battery here.",
		},
		// empty input
		{
			input: []string{},
			want:  "",
		},
	}
	for _, tc := range testcases {
		got := game.ListItems(tc.input)
		if !cmp.Equal(tc.want, got) {
			t.Error(cmp.Diff(tc.want, got))
		}
	}
}