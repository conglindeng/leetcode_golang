package easy

func KItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	res := 0
	if k >= numOnes {
		res += numOnes
		k -= numOnes
	} else {
		return k
	}
	if k < numZeros {
		return res
	}
	k -= numZeros
	if k > numNegOnes {
		res += -numNegOnes
	} else {
		res += -k
	}
	return res
}
