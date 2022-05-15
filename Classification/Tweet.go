package Classification

import (
	"TSC.go/Utilities"
	"github.com/jdkato/prose/v2"
	"strings"
)

type Sentiment int8

const (
	Guess Sentiment = iota
	Positive
	Negative
)

type Tweet struct {
	OriginalContent prose.Document
	Tokens          []prose.Token
	Entities        []prose.Entity
	ID              string
	Senti           Sentiment
}

func NewTestTweet(cells []string) Tweet {
	var tw Tweet
	tw.Senti = Guess
	tw.ID = cells[0]
	docCpy, err := prose.NewDocument(cells[4])
	if err != nil {
		panic(err)
	}
	tw.OriginalContent = *docCpy
	tw.Tokens = docCpy.Tokens()
	tw.Entities = docCpy.Entities()
	return tw
}

func NewTrainTweet(cells []string) Tweet {
	var tw Tweet
	if cells[0] == "4" {
		tw.Senti = Positive
	} else {
		tw.Senti = Negative
	}
	tw.ID = cells[1]
	docCpy, err := prose.NewDocument(cells[5])
	if err != nil {
		panic(err)
	}
	tw.OriginalContent = *docCpy
	tw.Tokens = docCpy.Tokens()
	tw.Entities = docCpy.Entities()
	return tw
}

func (tw *Tweet) Clean() {
	stopWordsMap := Utilities.StopWords()
	index := 0
	cleanTokens := make([]prose.Token, len(tw.Tokens))
	for _, tok := range tw.Tokens {
		tok.Text = strings.ToLower(tok.Text)
		_, isSW := stopWordsMap[tok.Text]
		if !isSW && !Utilities.IsUrl(tok.Text) && !Utilities.IsUsername(tok.Text) {
			tok.Text = Utilities.Stem(tok.Text)
			cleanTokens[index] = tok
			index++
		}
	}
	tw.Tokens = cleanTokens
}
