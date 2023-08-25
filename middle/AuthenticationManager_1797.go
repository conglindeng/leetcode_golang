package middle

type AuthenticationManager struct {
	timeout int
	tokens  map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{timeout: timeToLive, tokens: map[string]int{}}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.tokens[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if v, exist := this.tokens[tokenId]; exist {
		if v+this.timeout > currentTime {
			this.tokens[tokenId] = currentTime
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	res := 0
	for _, v := range this.tokens {
		if v+this.timeout > currentTime {
			res++
		}
	}
	return res
}