package fid

// Fid 计算斐波那契数列的第n项
func Fid(n int) int {
	if n < 2 {
		return n
	}
	return Fid(n-1) + Fid(n-2)
}
