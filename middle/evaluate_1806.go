package middle

func evaluate(s string, knowledge [][]string) string {
	keyValue := map[string]string{}
	for _, v := range knowledge {
		keyValue[v[0]] = v[1]
	}
	res := make([]byte, 0)
	i := 0
	for i < len(s) {
		//find a left token
		if s[i] == '(' {
			j := i + 1
			for j < len(s) {
				if s[j] == ')' {
					key := string(s[i+1 : j])
					if v, ok := keyValue[key]; ok {
						res = append(res, []byte(v)...)
					} else {
						res = append(res, '?')
					}
					break
				} else {
					j++
				}
			}
			i = j + 1
		} else {
			res = append(res, s[i])
			i++
		}
	}

	return string(res)
}
