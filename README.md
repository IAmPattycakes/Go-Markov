# Go-Markov

Go-Markov is a Markov chain text generator implemented in Go, with a simple interface.
Currently the only context that is taken into account is the current word, but a configurable context depth is planned for the future where the model can take more than one word of context into account when it comes to text generation.

## Usage

Make a graph, add stuff to the graph with graph.LoadPhrase("this is a sentence that will be added to the text generator") and get out a phrase using graph.GenerateMarkovString() 
