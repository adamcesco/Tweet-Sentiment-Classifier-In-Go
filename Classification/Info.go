package Classification

type ConfusionMatrix struct {
	ConditionPos int
	ConditionNeg int
	TrueNeg      int
	TruePos      int
	FalseNeg     int
	FalsePos     int
}

func (cm ConfusionMatrix) Accuracy() float32 {
	return float32(cm.TruePos+cm.TrueNeg) / float32(cm.TruePos+cm.TrueNeg+cm.FalsePos+cm.FalseNeg)
}

type FeatureData struct {
	NegCount int
	PosCount int
	CM       ConfusionMatrix
	Flag     bool
}
