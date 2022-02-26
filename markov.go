package markov

import (
	"math/rand"
	"strings"
)

//PUBLIC
type Graph interface {
	LoadPhrase(string)
	GenerateMarkovString() string
}

//Graph is the generic holder of all of the data for a certain graph of words needed to generate a markov chain.
type RamGraph struct {
	starters     list
	allWords     list
	strings      []string
	contextDepth int
	nodes        map[string]*link
}

func NewGraph(depth int) *RamGraph {
	var g RamGraph
	g.nodes = make(map[string]*link)
	g.contextDepth = depth
	return &g
}

//LoadPhrase loads a whole phrase (be it a sentence, single word that someone said, or paragraph) into the graph.
//Splits the phrase along spaces in order to load each word into the chain.
func (graph *RamGraph) LoadPhrase(phrase string) {
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
		prevVal = graph.loadWord(words[i], prevVal, (i == 0), context)
	}
}

//GenerateMarkovString generates a random string based off of the current model in the graph.
func (graph *RamGraph) GenerateMarkovString() string {
	var ret string
	var currLink *link = graph.starters.links[rand.Intn(len(graph.starters.links))]
	ret += *currLink.value
	currLink = currLink.links[rand.Intn(len(currLink.links))]
	for currLink != nil {
		ret += " " + *currLink.value
		currLink = currLink.links[rand.Intn(len(currLink.links))]
	}
	return ret
}

//private, implementation details.

type link struct {
	value   *string
	links   []*link
	context []*string
}

type list struct {
	links []*link
}

//loadWord loads a word into the respective graph. It creates a link if one doesn't exist,
// and links the current word to both previous (if higher context is on) and next words.
func (graph *RamGraph) loadWord(val string, nextval *link, starter bool, context []*string) *link {
	mapKey := buildMapKey(val, context)
	l := graph.nodes[mapKey]
	if l == nil {
		ctx := make([]*string, len(context))
		for i, v := range context {
			ctx[i] = graph.findString(*v)
		}
		l = &link{
			value:   graph.findString(val),
			links:   []*link{nextval},
			context: ctx,
		}
		graph.allWords.links = append(graph.allWords.links, l)
		graph.nodes[mapKey] = l
	} else {
		l.links = append(l.links, nextval)
	}

	if starter {
		graph.starters.links = append(graph.starters.links, l)
	}
	return l
}

func (graph *RamGraph) findInGraph(val string, context []*string) *link {
	key := ""
	for _, ctx := range context {
		key += *ctx + " "
	}
	key += val
	return graph.nodes[key]

	// for i, v := range graph.allWords.links {
	// 	if *v.value == val {
	// 		contextMatching := true
	// 		if len(v.context) != len(context) {
	// 			continue
	// 		}
	// 		for j, ctx := range v.context {
	// 			if *ctx != *context[j] {
	// 				contextMatching = false
	// 			}
	// 		}
	// 		if contextMatching {
	// 			return graph.allWords.links[i]
	// 		}
	// 	}
	// }
	// return nil
}

func (graph *RamGraph) findString(val string) *string {

	for i, v := range graph.strings {
		if v == val {
			return &graph.strings[i]
		}
	}
	graph.strings = append(graph.strings, val)
	return &graph.strings[len(graph.strings)-1]
}

func lesser(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func buildMapKey(val string, context []*string) string {
	key := ""
	for _, ctx := range context {
		key += *ctx + " "
	}
	key += val
	return key
}
