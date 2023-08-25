package middle

type AuthNode struct {
	token     string
	t         int
	pre, next *AuthNode
}

type AuthenticationManager_L struct {
	head, tail *AuthNode
	ttl        int
}

func Constructor_L(timeToLive int) AuthenticationManager_L {
	auth := AuthenticationManager_L{}
	auth.ttl = timeToLive
	head := &AuthNode{t: -1}
	tail := &AuthNode{t: -1}
	head.next = tail
	tail.pre = head
	auth.head = head
	auth.tail = tail
	return auth
}

func (this *AuthenticationManager_L) Generate_L(tokenId string, currentTime int) {
	node := AuthNode{token: tokenId, t: currentTime}
	pre := this.tail.pre
	pre.next = &node
	node.pre = pre
	this.tail.pre = &node
	node.next = this.tail
}

func (this *AuthenticationManager_L) Renew_L(tokenId string, currentTime int) {
	cur := this.head.next
	for cur.t != -1 {
		if cur.token == tokenId {
			if cur.t+this.ttl > currentTime {
				cur.pre.next = cur.next
				cur.next.pre = cur.pre
				break
			} else {
				return
			}
		}
		cur = cur.next
	}
	if cur.t == -1 {
		return
	}
	pre := this.tail.pre
	pre.next = cur
	cur.pre = pre
	this.tail.pre = cur
	cur.next = this.tail
	cur.t = currentTime
}

func (this *AuthenticationManager_L) CountUnexpiredTokens_L(currentTime int) int {
	res := 0
	cur := this.head.next
	for cur.t != -1 {
		if cur.t+this.ttl <= currentTime {
			next := cur.next
			cur.pre.next = next
			next.pre = cur.pre
			cur = next
		} else {
			res++
			cur = cur.next
		}
	}
	return res
}
