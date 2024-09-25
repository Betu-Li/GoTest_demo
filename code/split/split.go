package split

import "strings"

// Split 将字符串按照分隔符进行分割，返回分割后的字符串切片
/* s: 要分割的字符串, sep: 分隔符*/
func Split(s, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
