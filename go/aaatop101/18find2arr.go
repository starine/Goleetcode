package main

func Find(target int, array [][]int) bool {
	n := len(array)
	m := len(array[0])
	if n == 0 || m == 0 {
		return false
	}
	i := n - 1
	j := 0
	for i >= 0 && j < m {
		if target < array[i][j] {
			i--
		} else if target > array[i][j] {
			j++
		} else {
			return true
		}
	}
	return false
}
func main() {

}
