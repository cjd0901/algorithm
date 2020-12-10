package golang

func LemonadeChange(bills []int) (B bool) {
	faceKey := []int{5, 10, 20}
	faceValue := make(map[int]int)
	if bills[0] > 5 {
		return
	}
	var greedy func(rest, index int) bool
	greedy = func(rest, index int) bool {
		if rest == 0 {
			return true
		}
		if index < 0 {
			return false
		}
		need := rest / faceKey[index]
		stock := faceValue[faceKey[index]]
		if rest < faceKey[index] {
			return greedy(rest, index-1)
		}else {
			if stock >= need {
				faceValue[faceKey[index]] -= need
				return greedy(rest - faceKey[index] * need, index-1)
			}else {
				faceValue[faceKey[index]] = 0
				return greedy(rest - stock * need, index-1)
			}
		}
	}
	for _, bill := range bills {
		faceValue[bill] += 1
		B = greedy(bill-5, len(faceKey) - 1)
		if !B {
			return
		}
	}
	return
}