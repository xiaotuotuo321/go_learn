package split

import "strings"

//func Split(s, sep string) (result []string) {
//	i := strings.Index(s, sep)
//
//	for i > -1 {
//		result = append(result, s[:i])
//
//		s = s[i+len(sep):]
//		i = strings.Index(s, sep)
//	}
//
//	result = append(result, s)
//	return
//}

// 优化使用make函数将result出事化为一个容量足够大的切片，而不再像之前一样通过调用append函数来追加。
func Split(s, sep string) (result []string){
	result = make([]string, 0, strings.Count(s, sep) + 1)
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

