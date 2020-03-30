package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSortBoolChannel(arr []int) []int{
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr)/2

	var lArr,rArr []int
	lChan := make(chan bool)
	rChan := make(chan bool)

	go func() {
		lArr = mergeSortBoolChannel(arr[:mid])
		lChan <- true
	}()

	go func() {
		rArr = mergeSortBoolChannel(arr[mid:])
		rChan <- true
	}()
	<-lChan
	<-rChan

	return mergeTheTwoArrayBoolChannel(lArr, rArr)
}

func mergeTheTwoArrayBoolChannel(lArr, rArr []int) []int {
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
	// arr := []int{2,0,34,12,13,1,4,98,67,45}
	arr := GetRandomArray()
	t1 := time.Now()
	sorted := mergeSortBoolChannel(arr)
	t2 := time.Since(t1)
	fmt.Println("Time Taken :", t2.Milliseconds())
	fmt.Println(len(sorted))
	//fmt.Println("Sorted", sortedArr)
}

func GetRandomArray() []int{
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	arr := make([]int, 10000000)
	for i:=0;i<10000000;i++ {
		arr[i] = r.Intn(1000000000)
	}
	return arr
}
