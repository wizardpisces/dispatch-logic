package generateParenthesis

// 高度 2*n 的树，左子树是 "(" ,右子树是 ")"，递归版实际是先序遍历

func fn(numOfParenthesis int) (ans []string) {
	var dfs func(str string, open int, close int)
	dfs = func(str string, open int, close int) {
		if len(str) == 2*numOfParenthesis {
			ans = append(ans, str)
			return
		}
		if open < close { // 剪枝
			return
		}
		if open < numOfParenthesis {
			dfs(str+"(", open+1, close)
		}
		if close < numOfParenthesis {
			dfs(str+")", open, close+1)
		}

	}
	dfs("", 0, 0)
	return ans
}

func fn2(n int) (ans []string) {
	path := make([]byte, 2*n)
	var dfs func(open int, position int)
	dfs = func(open int, position int) {
		if position == 2*n {
			ans = append(ans, string(path))
			return
		}
		if open < position-open { // open < close  剪枝
			return
		}
		if open < n {
			path[position] = '('
			dfs(open+1, position+1)
		}

		if position-open < n {
			path[position] = ')'
			dfs(open, position+1)
		}
	}
	dfs(0, 0)
	return ans
}

type Stack struct {
	left  int
	right int
	res   string
}

func nonRecursive(n int) (ans []string) {
	stackList := []Stack{{left: n, right: n, res: ""}}

	if n == 0 {
		return ans
	}

	for len(stackList) != 0 {
		stack := stackList[0]
		stackList = stackList[1:]
		left, right, str := stack.left, stack.right, stack.res
		if left == 0 && right == 0 {
			ans = append(ans, str)
			continue
		}
		if left > right {
			continue
		}
		if left > 0 { //优先添加左括号
			stackList = append(stackList, Stack{res: str + "(", left: left - 1, right: right})
		}
		if right > 0 {
			stackList = append(stackList, Stack{res: str + ")", left: left, right: right - 1})
		}
	}

	return ans
}

// 动态规划版本

func generateParenthesis(n int) []string {
	return dp(n)[n]
}

func dp(n int) map[int][]string {
	if n == 0 {
		return map[int][]string{0: {""}}
	}
	if n == 1 {
		return map[int][]string{0: {""}, 1: {"()"}}
	}
	lastMap := dp(n - 1)
	var oneRes []string
	for i := 0; i < n; i++ {
		inners := lastMap[i]
		outers := lastMap[n-1-i]
		for _, inner := range inners {
			for _, outer := range outers {
				oneRes = append(oneRes, "("+inner+")"+outer)
			}
		}
	}
	lastMap[n] = oneRes
	return lastMap
}
