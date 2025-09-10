package helpers

func GetRTP(sequence []float64, multiplicator []float64) (rtp float64) {
	var sum float64

	for i := 0; i < len(sequence); i++ {
		if sequence[i] >= multiplicator[i] {
			sequence[i] = 0
		}
		sum += sequence[i]
	}
	return sum / float64(len(sequence))
}
