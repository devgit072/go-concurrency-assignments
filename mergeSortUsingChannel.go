package main

import "fmt"

func mergeSortCon(arr []int, ch chan []int) {
	if len(arr) == 0 {
		return
	}
	if len(arr) == 1 {
		ch <- arr
		return
	}
	lChan := make(chan []int)
	rChan := make(chan []int)
	mid := len(arr)/2

	go mergeSortCon(arr[:mid], lChan)
	go mergeSortCon(arr[mid:], rChan)
	lSortedArray := <- lChan
	rSortedArray := <- rChan
	close(lChan)
	close(rChan)
	sortedArr := mergeTheTwoArrayCon(lSortedArray, rSortedArray)
	ch <- sortedArr
}

func mergeTheTwoArrayCon(lArr, rArr []int) []int {
	totalLen := len(lArr) + len(rArr)
	newArr := make([]int, totalLen)
	i := 0
	l := 0
	r := 0
	for(l<len(lArr) && r <len(rArr)) {
		if lArr[l] < rArr[r] {
			newArr[i] = lArr[l]
			i++
			l++
		} else if lArr[l] > rArr[r] {
			newArr[i] = rArr[r]
			i++
			r++
		} else {
			newArr[i] = lArr[l]
			i++
			l++
			newArr[i] = rArr[r]
			i++
			r++
		}
	}

	for (l<len(lArr)) {
		newArr[i] = lArr[l]
		l++
		i++
	}
	for (r<len(rArr)) {
		newArr[i] = rArr[r]
		r++
		i++
	}
	return newArr
}

func main() {
	arr := []int{2,0,34,12,13,1,4,98,67,45}
	ch := make(chan []int)
	go mergeSortCon(arr, ch)
	sortedArr := <- ch
	fmt.Println("Sorted", sortedArr)
	close(ch)
}
