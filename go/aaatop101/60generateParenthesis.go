package main

func generateParenthesis(n int) (ans []string) {
	var trace func(k int, cnt int, tmp []byte)
	trace = func(k int, cnt int, tmp []byte) {
		if k == 2*n {
			// 右括号必需占一半
			if cnt != n {
				return
			}
			ans = append(ans, string(tmp))
			return
		}
		// 右括号大于左括号数量，剪枝
		// 左括号数量已经大于n，剪枝
		if cnt > len(tmp)-cnt || len(tmp)-cnt > n {
			return
		}
		// 先放 (
		tmp = append(tmp, '(')
		trace(k+1, cnt, tmp)
		tmp = tmp[:len(tmp)-1]
		tmp = append(tmp, ')')
		trace(k+1, cnt+1, tmp)
	}

	trace(0, 0, []byte{})
	return
}
