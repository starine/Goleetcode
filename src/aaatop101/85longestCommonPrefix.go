package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 { //空字符串数组
		return ""
	}
	for i := 0; i < len(strs[0]); i++ { //遍历第⼀个字符串的⻓度
		for j := 1; j < len(strs); j++ { //遍历后续的字符串
			if i == len(strs[j]) || strs[j][i] != strs[0][i] { //⽐较每个字符串该位置是否和第⼀个相同
				return strs[0][:i] //不相同则结束
			}
		}
	}
	return strs[0]
} //63ms 16360kb 时间复杂度O(mn)空间复杂度为O(1)

////分治法
//func longestCommonPrefix(strs []string) string {
//	if len(strs) == 0 {
//		return ""
//	}
//	var lcp func(int, int) string
//	lcp = func(start, end int) string {
//		if start == end {
//			return strs[start]
//		}
//		mid := (start + end) / 2
//		lcpLeft, lcpRight := lcp(start, mid), lcp(mid + 1, end)
//		minLength := min(len(lcpLeft), len(lcpRight))
//		for i := 0; i < minLength; i++ {
//			if lcpLeft[i] != lcpRight[i] {
//				return lcpLeft[:i]
//			}
//		}
//		return lcpLeft[:minLength]
//	}
//	return lcp(0, len(strs)-1)
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}//89ms,16340kb 时间复杂度O(mn)空间复杂度为O(mlogn)
