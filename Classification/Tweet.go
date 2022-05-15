package Classification

import (
	"github.com/jdkato/prose/v2"
)

type Sentiment int8

const (
	Guess Sentiment = iota
	Positive
	Negative
)

type Tweet struct {
	Content prose.Document
	ID      string
	Senti   Sentiment
}

func NewTestTweet(cells []string) Tweet {
	var tw Tweet
	tw.Senti = Guess
	tw.ID = cells[0]
	docCpy, err := prose.NewDocument(cells[4])
	if err != nil {
		panic(err)
	}
	tw.Content = *docCpy
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
	tw.Content = *docCpy
	return tw
}
