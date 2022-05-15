package Classification

type Feature struct {
	Data     string
	NegCount int
	PosCount int
	CM       ConfusionMatrix
}
