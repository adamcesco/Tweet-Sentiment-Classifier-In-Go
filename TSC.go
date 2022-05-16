package main

import (
	"TSC.go/Classification"
	"fmt"
	"strings"
)

func main() {
	tweet := Classification.NewTrainTweet(strings.Split("4,2014351367,Wed Jun 03 01:14:33 PDT 2009,NO_QUERY,cathyleehart,@just_tam21 hahah its kind of the same as adelaides except a bit bigger http://example.com", ","))
	for _, tok := range tweet.OriginalContent.Tokens() {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
	}
	fmt.Println("==========================")

	for _, tok := range tweet.Tokens {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
	}
	fmt.Println("==========================")

	tweet.Clean()
	for _, tok := range tweet.Tokens {
		fmt.Println(tok.Text, tok.Tag, tok.Label)
	}
}
