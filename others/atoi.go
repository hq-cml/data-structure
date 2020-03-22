package main

import (
	"errors"
	"fmt"
)

func Atoi(str string) (int, error) {
	if str == "" {
		return 0, errors.New("Invalid input")
	}

	//过滤空格字符
	beg := 0
	for i:=0; i<len(str); i++ {
		if string(str[i]) == " " {
			beg ++
		}else {
			break
		}
	}
	str = string(str[beg:])

	//正负号
	sign := 1
	if string(str[0]) == "+" {
		str = string(str[1:])
	} else if string(str[0]) == "-" {
		str = string(str[1:])
		sign = -1
	}

	if len(str) == 0 {
		//仅有正负号
		return 0, errors.New("Invalid input")
	}

	ret := 0
	addFlag := false //小数进位标记
	for i:=0; i<len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			ret = ret * 10 + int(str[i] - '0')
		} else if string(str[i]) == " " {
			//do nothing
		} else if string(str[i]) == "." {
			if i+1 <= len(str) - 1 && str[i+1] >= '0' && str[i+1] <= '9'{
				if str[i+1] <= '4' {
					addFlag = false
				} else {
					addFlag = true
				}
				break
			}else {
				return 0, errors.New("Invalid input")
			}
		} else {
			return 0, errors.New("Invalid input")
		}
	}

	if addFlag {
		return sign * (ret + 1), nil
	}

	return sign * ret, nil
}

func main() {
	fmt.Println(Atoi("108"))
	fmt.Println(Atoi("-"))
	fmt.Println(Atoi(" 10 8 "))
	fmt.Println(Atoi(" 10x8 "))
	fmt.Println(Atoi(" -10 8 "))
	fmt.Println(Atoi("108.1"))
	fmt.Println(Atoi("108.5"))
}