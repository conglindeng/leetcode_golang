package easy

func distMoney(money int, children int) int {
	if money <= children {
		return -1
	}
	money -= children
	cnt := min(children, money/7)
	money -= cnt * 7
	children -= cnt
	if (children == 0 && money > 0) || (children == 1 && money == 3) {
		cnt--
	}
	return cnt
}
