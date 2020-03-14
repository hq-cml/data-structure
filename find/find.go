package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type IntArr []int

func (s IntArr) Len() int {
	return len(s)
}
func (s IntArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s IntArr) Less(i, j int) bool {
	return s[i] < s[j]
}

//给一个随机数组，然后求出和为给定数字的所有成员的组合
//数组元素全部>0
func FindNum(arr IntArr, sum int) ([][]int, bool) {
	//异常输入
	if sum < 0 {
		return nil, false
	}

	length := len(arr)
	//排序
	sort.Sort(arr)

	//滑动窗口
	i := 0
	j := 0
	curr := arr[i]
	flag := false
	ret := [][]int{}
	uniq := map[string]struct{}{}
	for i < length && j < length {
		if curr == sum {
			if i <= j {
				tmp := []int{}
				builder := strings.Builder{}
				for k := i; k <= j; k++ {
					builder.WriteString(strconv.Itoa(arr[k]))
					tmp = append(tmp, arr[k])
				}
				_, ok := uniq[builder.String()]
				if !ok {
					ret = append(ret, tmp)
					uniq[builder.String()] = struct{}{}
				}
			}

			if arr[i] == 0 {
				i++
				if i > (length-1) { //边界
					break
				}
			} else {
				j++
				if j > (length-1) { //边界
					break
				}
				curr = curr + arr[j]
			}

			flag = true
		}

		//右边界前进
		if curr < sum {
			j++
			if j > 9 {
				break
			}
			curr = curr + arr[j]
		}

		//左边界
		if curr > sum {
			curr = curr - arr[i]
			i++
		}
	}

	return ret, flag
}

func main() {
	//异常输入
	a := IntArr{9, 0, 8, 2, 6, 5, 4, 3, 7, 1}
	fmt.Println(FindNum(a, -1))

	//常规
	a = IntArr{9, 0, 8, 2, 6, 5, 4, 3, 7, 1}
	fmt.Println(FindNum(a, 7))

	//恰好第一个元素
	a = IntArr{9, 0, 8, 2, 6, 5, 4, 3, 7, 1}
	fmt.Println(FindNum(a, 0))

	//恰好最后一元素
	a = IntArr{9, 0, 8, 2, 6, 5, 4, 3, 7, 1}
	fmt.Println(FindNum(a, 9))

	//找不到
	a = IntArr{9, 0, 8, 2, 6, 5, 4, 3, 7, 1}
	fmt.Println(FindNum(a, 100))

	//元素全部相同
	a = IntArr{7, 7, 7, 7, 7, 7, 7, 7, 7, 7}
	fmt.Println(FindNum(a, 7))

	//元素全部相同，且符合多个元素相加
	a = IntArr{7, 7, 7, 7, 7, 7, 7, 7, 7, 7}
	fmt.Println(FindNum(a, 14))

	//找不到
	a = IntArr{9, 0, 8, 2, 6, 5, 4, 3, 7, 1}
	fmt.Println(FindNum(a, 1))
}
