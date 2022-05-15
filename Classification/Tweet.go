package Classification

import "strings"

type Sentiment int8

const (
	Guess Sentiment = iota
	Positive
	Negative
)

type Tweet struct {
	Content []string
	ID      string
	Senti   Sentiment
}

func (tw *Tweet) ReadTestTweet(cells []string) {
	tw.Senti = Guess
	tw.ID = cells[0]
	tw.Content = strings.Split(cells[4], " ")
}

func (tw *Tweet) ReadTrainTweet(cells []string) {
	if cells[0] == "4" {
		tw.Senti = Positive
	} else {
		tw.Senti = Negative
	}
	tw.ID = cells[1]
	tw.Content = strings.Split(cells[5], " ")
}

func (tw *Tweet) Clean() {

}
