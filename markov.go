package markov

import (
	"math/rand"
	"strings"
)

//PUBLIC

//Graph is the generic holder of all of the data for a certain graph of words needed to generate a markov chain.
type Graph struct {
	starters     list
	allWords     list
	strings      []string
	contextDepth int
}

//LoadPhrase loads a whole phrase (be it a sentence, single word that someone said, or paragraph) into the graph.
//Splits the phrase along spaces in order to load each word into the chain.
func (graph *Graph) LoadPhrase(phrase string) {
	words := strings.Split(phrase, " ")
	// Get the words in reverse order, just to make it easier to input the next word in the chain.
	var prevVal *link = nil //Start prevVal as nil, as the first word to appear will be the last one in the phrase.
	for i := len(words) - 1; i >= 0; i-- {
		depthToGo := lesser(i, graph.contextDepth)
		context := make([]*string, depthToGo)
		for j := depthToGo; j > 0; j-- {
			indexToGrab := i - j
			context[j-1] = &words[indexToGrab]
		}
		prevVal = graph.loadWord(words[i], prevVal, (i == 0), context) //Last arg, if i is 0 then it is the start of the sentence, despite being at the end of the loop.
	}
}

//GenerateMarkovString generates a random string based off of the current model in the graph.
func (graph *Graph) GenerateMarkovString() string {
	var ret string
	var currLink *link = graph.starters.links[rand.Intn(len(graph.starters.links))]
	for currLink != nil {
		ret += currLink.value + " "
		currLink = currLink.links[rand.Intn(len(currLink.links))]
	}
	return ret
}

//private, implementation details.

type link struct {
	value   string
	links   []*link
	context []*string
}

type list struct {
	links []*link
}

//loadWord loads a word into the respective graph. It creates a link if one doesn't exist,
// and links the current word to both previous (if higher context is on) and next words.
func (graph *Graph) loadWord(val string, nextval *link, starter bool, context []*string) *link {
	l := graph.findInGraph(val)
	if l == nil {
		l = &link{
			value: val,
			links: []*link{nextval},
		}
		graph.allWords.links = append(graph.allWords.links, l)
	} else {
		l.links = append(l.links, nextval)
	}

	if starter {
		graph.starters.links = append(graph.starters.links, l)
	}
	return l
}

func (graph *Graph) findInGraph(val string) *link {
	for i, v := range graph.allWords.links {
		if v.value == val {
			return graph.allWords.links[i]
		}
	}
	return nil
}

func lesser(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
