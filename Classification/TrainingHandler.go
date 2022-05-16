package Classification

type TrainingHandler struct {
	WordMap map[string]FeatureData
}

func (tr *TrainingHandler) CountWordSentFrom(tweetArr []Tweet) {
	for _, tweet := range tweetArr {
		for _, tok := range tweet.Tokens {
			data := tr.WordMap[tok.Text]
			if tweet.Senti == Positive {
				data.PosCount++
			} else if tweet.Senti == Negative {
				data.NegCount++
			}
			tr.WordMap[tok.Text] = data
		}
	}
}

func (tr *TrainingHandler) CalcWordSentiAcc(tweetArr []Tweet) {
	for _, tweet := range tweetArr {
		for _, tok := range tweet.Tokens {
			data := tr.WordMap[tok.Text]
			if tweet.Senti == Positive {
				data.CM.ConditionPos++

				if data.PosCount > data.NegCount {
					data.CM.TruePos++
				} else if data.PosCount < data.NegCount {
					data.CM.FalseNeg++
				}
			} else if tweet.Senti == Negative {
				data.CM.ConditionNeg++

				if data.PosCount < data.NegCount {
					data.CM.TrueNeg++
				} else if data.PosCount > data.NegCount {
					data.CM.FalsePos++
				}
			}
			tr.WordMap[tok.Text] = data
		}
	}
}

func (tr *TrainingHandler) CleanOutliers() {
	for word, data := range tr.WordMap {
		if (data.PosCount+data.NegCount) < 3 || int(data.CM.Accuracy()*100) < 54 {
			delete(tr.WordMap, word)
		}
	}
}
