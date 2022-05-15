package main

import (
	"TSC.go/Classification"
	"fmt"
	"strings"
)

func main() {
	tweet := Classification.NewTrainTweet(strings.Split("4,2014351367,Wed Jun 03 01:14:33 PDT 2009,NO_QUERY,cathyleehart,@just_tam21 hahah its kind of the same as adelaides except a bit bigger http://example.com", ","))
	for _, tok := range tweet.Content.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
		// Go NNP B-GPE
		// is VBZ O
		// an DT O
		// ...
	}

	// Iterate over the doc's named-entities:
	fmt.Println("========================================")
	for _, ent := range tweet.Content.Entities() {
		fmt.Println(ent.Text, ent.Label)
		// Go GPE
		// Google GPE
	}

	// Iterate over the doc's sentences:
	fmt.Println("========================================")
	for _, sent := range tweet.Content.Sentences() {
		fmt.Println(sent.Text)
		// Go is an open-source programming language created at Google.
	}
}
