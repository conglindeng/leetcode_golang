package middle

func beautySum(s string) int {
	ans := 0
	for i := range s {
		cnt := [26]int{}
		maxFre := 0
		for _,c := range s[i:] {
			minFre:=len(s)
			cnt[c-'a']++
			maxFre=max(maxFre,cnt[c-'a'])
			for _,c:=range cnt{
				if c>0{
					minFre=min(minFre,c)
				}
			}
			ans+=maxFre-minFre
		}
	}
	return ans
}
