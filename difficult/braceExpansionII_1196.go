package difficult

func braceExpansionII(expression string) []string {
	stack := make([]byte, 0)
	for _, c := range expression {
		if c == '{' {

		} else if c == '}' {

		} else {
			stack = append(stack, byte(c))
		}
	}
	return nil
}
