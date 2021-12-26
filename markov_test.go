package markov

import (
	"fmt"
	"testing"
)

func TestStringGeneration(t *testing.T) {
	g := NewGraph(1)
	strlist := []string{"1 2 3 4 5", "1 3 5 4 2", "1 3 4 5 2", "2 3 4 5 1", "1 5 3 2 4", "5 4 3 2 1", "4 3 1 2 5"}
	for _, v := range strlist {
		g.LoadPhrase(v)
	}
	arr := make([]string, 100)
	notSeenBefore := false
	for i := 0; i < len(arr); i++ {
		arr[i] = g.GenerateMarkovString()
	}
	for _, v := range arr {
		if !notSeenBefore { //Checking if this is a new string, not just a copy from the given data
			inThisList := false
			for _, st := range strlist {
				if v == st {
					inThisList = true
				}
			}
			if !inThisList {
				notSeenBefore = true
			}
		}
		fmt.Println(v)
	}
	if !notSeenBefore {
		t.Error("No new strings were ever generated")
	}
}

func safeval(list []*string, num int) string {
	if num >= len(list) {
		return ""
	}
	return *list[num]
}
