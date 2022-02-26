# Go-Markov

Go-Markov is a Markov chain text generator implemented in Go, with a simple interface.
Currently the only context that is taken into account is the current word, but a configurable context depth is planned for the future where the model can take more than one word of context into account when it comes to text generation.

## Usage
It's meant to be simple, there's only 3 functions you need to think about right now. 

`NewGraph(int) *Graph` is going to be your starting point for making a graph that handles your markov chain. The input is the context depth you want. More about "context depth" is below. 

`Graph.LoadPhrase(string)` is how you input phrases into the graph. These phrases are expected to be words separated by spaces. 

`Graph.GenerateMarkovString() string` is how to get things out of the graph. It generates a string randomly based off of the relationships between words in the graph already. 

A very basic usage of it can be seen in the markov_test.go that shows things getting loaded in, then strings being generated. 

## How it all works
Wikipedia has a good writeup on markov chains as a whole, but I'll go into how it pertains to this library narrowly. 

When a phrase is loaded into the model it's broken up word by word on space character boundaries. The first word is added to a list of "starting words" that will get picked from to start the sentence with. 

After the starting word is picked randomly from the starter list, the graph will pick from a list of words that have followed that word before. 

Once that word is picked from the list, the graph will start picking the next word based off of the word that was just generated, as well as however many words before as dictated by the context depth. 

### Context depth
This is a term I kinda made up on my own, possibly with inspiration from other things I read. It basically just means how far back the model looks to get context on the next word to be said. If the model has already generated the phrase `"I am a"` it is currently at the "a" in order to generate the next word. At a context depth of 0, it looks back 0 words and thus the only information it has to pick from is all previously loaded phrases that had "a" as a word. At a context depth of 1, it will look back 1 word, and thus generate the next word based off of the previous phrases that had "am a" in them. 

This is, in essence, a randomization dampener. It limits the choices by forcing them to make sense in the context they're picked in. 

A **VERY** importent thing to note, is that **a lack of context is context** and that means when you're at a context level of 2, and the graph is generating the third word, it's looking at the second word. It can see the first word as expected, but it also knows that there is no word before the first word. This forces it to not generate any words that haven't been the third word in a sentence before.

#### Where this matters
Lots of context can give you much more proper sentences out. Zero context is often very terrible, and while it was good for generating laughs, it wasn't anywhere close to believable. However, with higher context levels you can get very severe overfitting, where your sentence has no possible branches to pick from after it's first few words. Due to how the context works. the first `contextLevel + 1` words will have always been said together in a phrase in your dataset. You will never get a pair of 2 words at the front of your sentence that haven't ever been at the front of a sentence that was entered if you have a context depth of 1, the first 3 words will have always been said before at the start of a sentence for a context depth of 2, and so on. Since you will always have things that match a sentence at the start, if you don't have enough phrases that start with that string of words and then diverge, you will never get anything that wasn't said before. And that kinda defeats the purpose of a text gnerator if you don't get anything out of it that is generated and not already said.

#### How to tune properly
Basically, too much and you no longer have a text generator. Too little, and you just have a terrible text generator. That's why it's reconfigurable. Metrics/benchmarks will be coming in a future update, to show the randomness possible in the sentence generation, to help with the tuning more than just what feels right to the human eye. But right now, it's a manual process. I'd be surprised if this can be used past a context depth of 3 right now, because every level of context is a theoretical factorial amount of data it can pull from. That's not really how languages work, but increasing context depth is definitely more than an exponential increase if you have enough data to use the increase.  
