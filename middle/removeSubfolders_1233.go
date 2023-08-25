package middle

import (
	"fmt"
	"sort"
	"strings"
)

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	res := make([]string, 0)
	res = append(res, folder[0])
	pre := folder[0]
	if !strings.HasSuffix(pre, "/") {
		pre = pre + "/"
	}
	for i := 1; i < len(folder); i++ {
		cur := folder[i]
		if !strings.HasSuffix(cur, "/") {
			cur = cur + "/"
		}
		if !strings.HasPrefix(cur, pre) {
			res = append(res, folder[i])
			pre = cur
		}
	}
	return res
}

func RemoveSubfolders_tree(folder []string) []string {
	type folderDic struct {
		nexts map[string]*folderDic
		path  string
		isEnd bool
	}
	root := &folderDic{path: "", isEnd: false, nexts: map[string]*folderDic{}}
out:
	for _, path := range folder {
		splits := strings.Split(strings.TrimSuffix(path, "/"), "/")
		pre := root
		for i := 1; i < len(splits); i++ {
			cur := splits[i]
			if pre.nexts[cur] == nil {
				pre.nexts[cur] = &folderDic{path: cur, isEnd: false, nexts: map[string]*folderDic{}}
			} else if pre.nexts[cur].isEnd {
				continue out
			}
			pre = pre.nexts[cur]
		}
		if pre != root {
			pre.isEnd = true
		}
	}
	res := make([]string, 0)
	paths := make([]string, 0)
	var dfs func(root *folderDic, paths []string)
	dfs = func(root *folderDic, paths []string) {
		if root.isEnd {
			res = append(res, "/"+strings.Join(paths, "/"))
			return
		}
		for _, v := range root.nexts {
			paths = append(paths, v.path)
			dfs(v, paths)
			paths = paths[:len(paths)-1]
		}
	}
	//todo:dcl
	dfs(root, paths)

	res1 := make([]string, 0)

	type bfsInfo struct {
		folder *folderDic
		path   string
	}

	bfs := func(root *folderDic) {
		work := make([]bfsInfo, 0)
		work = append(work, bfsInfo{folder: root, path: ""})
		for len(work) > 0 {
			cur := work[0]
			work = work[1:]
			if cur.folder.isEnd {
				res1 = append(res1, cur.path)
				continue
			} else {
				for _, v := range cur.folder.nexts {
					work = append(work, bfsInfo{folder: v, path: cur.path + "/" + v.path})
				}
			}
		}
	}
	bfs(root)
	fmt.Println(res, res1)
	return res
}
