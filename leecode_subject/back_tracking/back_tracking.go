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