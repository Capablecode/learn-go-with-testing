package struc

type student struct{
	Name	string
	Age		int
	Score	[]float64
}

func (s student) AverageScore() float64{
	if len(s.Score) == 0{
		return 0
	}
	var total float64

	for _, score := range s.Score{
		total += score
	}
	return total / float64(len(s.Score))
}

func (s student) hasPassed() bool{
	currentAvaerage := s.AverageScore()
	if currentAvaerage >= 50{
		return true
	}else{
		return false
	}
}