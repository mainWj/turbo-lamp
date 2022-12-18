package back_tracking
// 77 组合问题，利用回溯解决，注意重复问题
func combine(n int, k int) [][]int {
	var res [][]int
	visit := map[int]bool{}
	var backTrack func(path []int, j int)
	backTrack = func(path []int, j int){
		if len(path) == k{
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i:=j; i <= n; i++{
			if visit[i]{
				continue
			}
			visit[i] = true
			path = append(path, i)
			backTrack(path, i)
			path = path[:len(path)-1]
			visit[i] = false
		}
	}
	backTrack([]int{}, 1)
	return res
}


// 78 组合 回溯法的变种
func subsets(nums []int) [][]int {
	var res [][]int
	visit := map[int]bool{}
	var traceBake func([]int,int)
	// res = append(res, []int{})
	traceBake = func(path []int, j int){
		// if true{
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		//return
		// }
		for i := j;i < len(nums); i ++{
			if visit[nums[i]]{
				continue
			}
			visit[nums[i]] = true
			path = append(path, nums[i])
			traceBake(path, i)
			visit[nums[i]] = false
			path = path[:len(path)-1]
		}
	}
	traceBake([]int{}, 0)
	return res
}


// 131 分割回文串，分割问题，也可以当作组合问题来处理，例如，先拿一个字符类比于切割出也给字符
func partition(s string) [][]string {
	var res [][]string
	var path []string
	var traceBake func(start int, path []string)
	traceBake = func(start int, path []string){
		if start >= len(s){
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i:=start; i < len(s); i++{
			if isPalindrome(s, start, i){
				path = append(path, s[start:i+1])
			} else {
				continue
			}
			traceBake(i + 1, path)
			path = path[:len(path)-1]
		}
	}
	traceBake(0,path)
	return res
}
func isPalindrome(s string, start int, end int) bool{
	i:=start
	j:=end
	for ; i < j;  {
		if s[i] != s[j] {
			return false
		}
		i ++
		j--
	}
	return true
}
