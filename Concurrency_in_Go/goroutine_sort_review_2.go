package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*

Task:
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is
sorted by a different goroutine. Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should
print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

Submission: Upload your source code for the program.


Review Criteria:

The goal of this activity is to explore the use of threads by creating a program for sorting integers that uses four
goroutines to create four sub-arrays and then merge the arrays into a single array.
Review criteria less Students will receive five points if the program sorts the integers
and five additional points if there are four goroutines that each prints out a set of array elements
that it is storing.
*/

func mergeArrays(arr1, arr2 []int) []int {
	arr1Len := len(arr1)
	arr2Len := len(arr2)
	n := arr1Len + arr2Len
	merge := make([]int, n, n)
	i := 0
	j := 0
	for idx := 0; idx < len(merge); idx++ {

		if i >= arr1Len {
			merge[idx] = arr2[j]
			j++
			continue
		}

		if j >= arr2Len {
			merge[idx] = arr1[i]
			i++
			continue
		}

		if arr1[i] < arr2[j] {
			merge[idx] = arr1[i]
			i++
		} else {
			merge[idx] = arr2[j]
			j++
		}
	}
	return merge
}

func sortIntegers(sli []int, c chan []int, goroutineId int) {
	fmt.Printf("Subarray to be sorted by goroutine with ID [%d]: %v\n", goroutineId, sli)
	sort.Ints(sli)
	c <- sli
}

func ReadNumbers() []int {
	fmt.Println("Enter a sequence of integers separated by space. Example: 291 6 11 134 1 3 10 29 29 16 12 233 2")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputLine := scanner.Text()
		tokens := strings.Fields(inputLine)

		ints := make([]int, len(tokens), len(tokens))
		for i, t := range tokens {
			val, err := strconv.Atoi(t)
			if err != nil {
				panic("failed to convert input to integer" + err.Error())
			}
			ints[i] = val
		}
		return ints
	}
	return nil
}

func partitionArray(arr []int, numPartitions int) [][]int {

	partitions := make([][]int, numPartitions, numPartitions)
	for i, val := range arr {
		partitions[i%numPartitions] = append(partitions[i%numPartitions], val)
	}
	return partitions
}

func main() {

	ints := ReadNumbers()

	const N = 4.

	subArrays := make([]chan []int, N, N)
	partitions := partitionArray(ints, N)

	for i := 0; i < N; i++ {
		arr := make(chan []int)
		subArrays[i] = arr
		go sortIntegers(partitions[i], subArrays[i], i)
	}

	a1 := <-subArrays[0]
	a2 := <-subArrays[1]
	a3 := <-subArrays[2]
	a4 := <-subArrays[3]
	arr1 := mergeArrays(a1, a2)
	arr2 := mergeArrays(a3, a4)
	sortedArray := mergeArrays(arr1, arr2)
	fmt.Println(">>> RESULT: Final array", sortedArray)
}
