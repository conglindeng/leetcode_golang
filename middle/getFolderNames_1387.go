package middle

import "strconv"

func GetFolderNames(names []string) []string {
	count := map[string]int{}
	res := make([]string, 0)
	for _, name := range names {
		if _, exsit := count[name]; !exsit {
			res = append(res, name)
			count[name] = 1
		} else {
			newName := name + "(" + strconv.Itoa(count[name]) + ")"
			for {
				if _, ex := count[newName]; !ex {
					break
				}
				count[name]++
				newName = name + "(" + strconv.Itoa(count[name]) + ")"
			}
			res = append(res, newName)
			count[newName] = 1
		}
	}
	return res
}
