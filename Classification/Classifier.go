package Classification

type Classifier struct {
	TrainingData TrainingHandler
	Tweets       []Tweet
}

func (cl *Classifier) Classify() {
	for i, tweet := range cl.Tweets {
		grade := float32(0.0)
		empty := true
		for _, tok := range tweet.Tokens {
			data, contained := cl.TrainingData.WordMap[tok.Text]
			if !contained {
				delete(cl.TrainingData.WordMap, tok.Text)
				continue
			}

			if data.PosCount > data.NegCount {
				empty = false
				grade += data.CM.Accuracy()
			} else if data.PosCount < data.NegCount {
				empty = false
				grade -= data.CM.Accuracy()
			}
		}

		cl.Tweets[i].Senti = Guess
		if !empty {
			gradeInt := int(grade * 100)

			if gradeInt < 0 {
				cl.Tweets[i].Senti = Negative
			} else if gradeInt > 0 {
				cl.Tweets[i].Senti = Positive
			}
		}
	}
}

func CalcConfusionMatrix(classifier Classifier, sentimentMap map[string]Sentiment) ConfusionMatrix {
	var confusionMatrix ConfusionMatrix
	for _, tweet := range classifier.Tweets {
		actualSentiment := sentimentMap[tweet.ID]

		if actualSentiment == Positive {
			confusionMatrix.ConditionPos++
		} else if actualSentiment == Negative {
			confusionMatrix.ConditionNeg++
		}

		if tweet.Senti == Guess {
			if actualSentiment == Positive {
				confusionMatrix.TruePos++
			} else {
				confusionMatrix.FalsePos++
			}
			continue
		}

		if actualSentiment == Positive {
			if tweet.Senti == actualSentiment {
				confusionMatrix.TruePos++
			} else {
				confusionMatrix.FalsePos++
			}
		} else if actualSentiment == Negative {
			if tweet.Senti == actualSentiment {
				confusionMatrix.TrueNeg++
			} else {
				confusionMatrix.FalseNeg++
			}
		}
	}
	return confusionMatrix
}
