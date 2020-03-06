package sortUtil

import (
	"fmt"
	"math/rand"
)

//3,9,2,8,1,7,4,6,5,10
//3  9,2,8,1,7,4,6,5,10
//  2,1    3,     9,2,8,,7,4,6,5,10

// 9,2,8,,7,4,6,5,10
// ,8,7,4,6,5   9  10
//4 5 6

func QuickSort(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		splitdata := arr[0]          //以第一个为基准
		low := make([]int, 0, 0)     //存储比我小的
		high := make([]int, 0, 0)    //存储比我大的
		mid := make([]int, 0, 0)     //存储与我相等
		mid = append(mid, splitdata) //加入第一个相等

		for i := 1; i < length; i++ {
			if arr[i] < splitdata {       //从小到大，从大到小，只要把 <改成 >
				low = append(low, arr[i])
			} else if arr[i] > splitdata {  //从小到大，从大到小，只要把 <改成 >
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort(low), QuickSort(high) //切割递归处理
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}

func QuickSort2(arr []int) []int {
	length := len(arr) //数组长度
	if length <= 1 {
		return arr //一个元素的数组，直接返回
	} else {
		//n:=length-1//n>=0  && n<=length-1
		n := rand.Int() % length
		splitdata := arr[n]          //以第一个为基准
		low := make([]int, 0, 0)     //存储比我小的
		high := make([]int, 0, 0)    //存储比我大的
		mid := make([]int, 0, 0)     //存储与我相等
		//mid = append(mid, splitdata) //加入第一个相等

		for i := 0; i < length; i++ {
			//if i == n {
			//	continue
			//}
			if arr[i] < splitdata {
				low = append(low, arr[i])
			} else if arr[i] > splitdata {
				high = append(high, arr[i])
			} else {
				mid = append(mid, arr[i])
			}
		}
		low, high = QuickSort2(low), QuickSort2(high) //切割递归处理
		myarr := append(append(low, mid...), high...)
		return myarr
	}
}

func QuickSortTest() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	fmt.Println(QuickSort2(arr))
}
