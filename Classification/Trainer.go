package Classification

type Trainer struct {
	WordMap map[string]FeatureData
}

func (tr *Trainer) CountWordSentFrom(tweetArr []Tweet) {
	for i := range tweetArr {
		for _, tok := range tweetArr[i].Tokens {
			if tweetArr[i].Senti == Positive {
				tr.WordMap[tok.Text] = FeatureData{
					PosCount: tr.WordMap[tok.Text].PosCount + 1,
					NegCount: tr.WordMap[tok.Text].NegCount,
					CM:       tr.WordMap[tok.Text].CM,
					Flag:     tr.WordMap[tok.Text].Flag,
				}
			} else if tweetArr[i].Senti == Negative {
				tr.WordMap[tok.Text] = FeatureData{
					PosCount: tr.WordMap[tok.Text].PosCount,
					NegCount: tr.WordMap[tok.Text].NegCount + 1,
					CM:       tr.WordMap[tok.Text].CM,
					Flag:     tr.WordMap[tok.Text].Flag,
				}
			}
		}
	}
}

func (tr *Trainer) CalcWordSentiAcc(tweetArr []Tweet) {

}

func (tr *Trainer) CleanOutliers() {

}
