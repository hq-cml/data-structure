package main

import "fmt"

//给定一个字符串，寻找不含交集的字符子串
// 比如：输入arrab，返回[arra，b]；输入arraba，则仅返回：[arraba]
func FindSubStrSet(str string) []string {
	length := len(str)

	if length == 0 {
		return nil
	}

	flags := map[byte]int{}
	for i:= length-1; i >=0; i -- {
		c := str[i]
		_, ok := flags[c]
		if !ok {
			flags[c] = i
		}
	}

	ret := []string{}
	for i:=0; i < length; i++ {
		c := str[i]
		end, ok := flags[c]
		if ok {
			ret = append(ret, str[i: end+1])
			i = end
		} else {
			ret = append(ret, string(str[i]))
		}
	}

	return ret
}

func main() {
	fmt.Println(FindSubStrSet(""))
	fmt.Println(FindSubStrSet("arraba"))
	fmt.Println(FindSubStrSet("arrab"))
	fmt.Println(FindSubStrSet("abcdef"))
	fmt.Println(FindSubStrSet("aaaaaa"))
}