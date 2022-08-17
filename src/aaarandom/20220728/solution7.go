/*
* @Author: starine
* @Date:   2022/8/7 11:11
 */

package main

func reachableNodes(n int, edges [][]int, restricted []int) int {
	re := make(map[int]struct{})
	for _, v := range restricted {
		re[v] = struct{}{}
	}
	mp := make(map[int][]int)
	for i, _ := range edges {
		if _,ok1 := re[edges[i][0]]; !ok1 {
			continue
		}
		mp[edges[i][0]] = append(mp[edges[i][0]],edges[i][1])
		mp[edges[i][1]] = append(mp[edges[i][1]],edges[i][0])
	}
	res := 0
	vis := make(map[int]struct{})
	var dfs func(root int)
	dfs = func(root int) {
		tmp := mp[root]
		for _, v := range tmp {
			if _, ok := vis[v]; !ok {
				res++
				vis[v] = struct{}{}
				dfs(v)
			}
		}
	}
	if _,ok := mp[0]; ok {
		res++
		dfs(0)
		vis[0] = struct{}{}
	}
	return res
}

