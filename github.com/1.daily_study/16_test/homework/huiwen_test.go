package homework

import (
	"fmt"
	"testing"
)

func TestJudgePalindrome(t *testing.T){
	tests := [][]interface{}{
		{"中文回文", "杨凯是凯杨", true},
		{"中文非回文", "杨凯是杨凯", false},
		{"英文非回文", "are you ok?", false},
		{"英文回文", "MadamImadam", true},
		{"空", "", true},
	}

	for _, line := range tests{
		t.Run(line[0].(string), func(t *testing.T){
			s := line[1].(string)
			fmt.Println(s)
			want := line[2].(bool)
			got := JudgePalindrome(s)
			if got != want {
				t.Errorf("want %t, got %t\n", want, got)
			}
		})
	}
}
