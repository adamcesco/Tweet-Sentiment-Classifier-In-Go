package main

import (
	"TSC.go/Classification"
	"TSC.go/Utilities"
	"fmt"
)

func main() {
	var trainingTweets []Classification.Tweet
	{
		rows, err := Utilities.ReadCSVFile("./data/train_dataset_20k.csv")
		if err != nil {
			panic(err)
		}

		trainingTweets = make([]Classification.Tweet, len(rows)-1)
		for i, row := range rows {
			if i == 0 {
				continue
			}
			tweet := Classification.NewTrainTweet(row)
			tweet.Clean()
			trainingTweets[i-1] = tweet
		}
	}

	trainingHandler := Classification.TrainingHandler{
		WordMap: make(map[string]Classification.FeatureData),
	}

	trainingHandler.CountWordSentFrom(trainingTweets)
	trainingHandler.CalcWordSentiAcc(trainingTweets)
	trainingHandler.CleanOutliers()

	//---------------------------------------------

	var testTweets []Classification.Tweet
	{
		rows, err := Utilities.ReadCSVFile("./data/test_dataset_10k.csv")
		if err != nil {
			panic(err)
		}

		testTweets = make([]Classification.Tweet, len(rows)-1)
		for i, row := range rows {
			if i == 0 {
				continue
			}
			tweet := Classification.NewTestTweet(row)
			tweet.Clean()
			testTweets[i-1] = tweet
		}
	}

	classifier := Classification.Classifier{
		TrainingData: trainingHandler,
		Tweets:       testTweets,
	}

	classifier.Classify()

	//---------------------------------------------

	var answerKey map[string]Classification.Sentiment
	{
		rows, err := Utilities.ReadCSVFile("./data/test_dataset_sentiment_10k.csv")
		if err != nil {
			panic(err)
		}

		answerKey = make(map[string]Classification.Sentiment)
		for i, row := range rows {
			if i == 0 {
				continue
			}

			var senti Classification.Sentiment
			if row[0][0] == '4' {
				senti = Classification.Positive
			} else if row[0][0] == '0' {
				senti = Classification.Negative
			}

			answerKey[row[1]] = senti
		}
	}

	confusionMatrix := Classification.CalcConfusionMatrix(classifier, answerKey)
	fmt.Println("Classifier Statistics:")
	fmt.Println("\tClassifier Accuracy: ", confusionMatrix.Accuracy())
	fmt.Println("\tClassified correctly: ", confusionMatrix.TruePos+confusionMatrix.TrueNeg)
	fmt.Println("\tClassified incorrectly: ", confusionMatrix.FalseNeg+confusionMatrix.FalsePos)
	fmt.Println("\tTotal amount of tweets: ", confusionMatrix.ConditionPos+confusionMatrix.ConditionNeg)
}
