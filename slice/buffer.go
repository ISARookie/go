package main

import "bytes"

func main() {
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok {
			buffer.WriteString(s)
		} else {
			break
		}
	}
}

func getNextString() (string, bool) {
	// 这里是获取下一个字符串的逻辑
	// 假设字符串获取成功
	return "NextString", true
}
