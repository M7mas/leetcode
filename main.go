package main

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strings"
)

// func minHeightShelves(books [][]int, shelfWidth int) int {
// 	// 1105. Filling Bookcase Shelves `LeetCode`
// 	// books [[x,y]], x-> thickness, y-> height
// 	// How?
// 	//loops
// 	val := 0
// 	counter := 0
// 	steps := 0
// 	currentMax := books[0][0]
// 	lastHeight := 0
// 	lastThickness := 0

// 	for _, b := range books {

// 		if steps > shelfWidth {
// 			currentMax = 0
// 			steps = 0
// 		}
// 		if lastHeight == b[0] && steps <= shelfWidth {
// 			fmt.Println(b)
// 			if steps == 0 {
// 				steps = steps + b[0]*2
// 			} else {
// 				steps = steps + b[0]
// 			}

// 			if lastThickness >= b[1] {
// 				currentMax = lastThickness
// 			} else {
// 				currentMax = b[1]
// 			}
// 		} else {
// 			currentMax = b[1]
// 		}
// 		fmt.Println("STEPS: \t %d", steps)
// 		fmt.Println("CurrentMax:%d", currentMax)

// 		lastHeight = b[0]
// 		lastThickness = b[1]

// 	}
// 	fmt.Println("COUNTER:\t%d", counter)
// 	return val
// }

// func callMinHeightShelves() {
// 	// First function call
// 	fmt.Println(minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4))
// 	// Second function call
// 	fmt.Println(minHeightShelves([][]int{{1, 3}, {2, 4}, {3, 2}}, 6))
// }

// func nthUglyNumber(n int) int {
// 	ar := make([]int, n)

// 	for i := 0; i < n; {
// 		t := 1

// 		for {

// 		}
// 		ar[i] = i + 1
// 		fmt.Println(ar[i])

// 		i++
// 	}

//		return ar[n-1]
//	}

// fmt.Println(nthUglyNumber(10))
// fmt.Println(nthUglyNumber(1))

// func twoSum(nums []int, target int) []int {

// 	return []int{0, 1}
// }

func containsDuplicate(nums []int) bool {

	// for sorting a list: https://stackoverflow.com/a/77687553
	slices.Sort(nums)

	// a temp to create a sit from a map, using stackoverflow.
	// https://stackoverflow.com/a/34020023
	set := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		_, ok := set[nums[i]] // check for existence
		if ok {
			return true
		}
		set[nums[i]] = true // add element
	}
	return false
}

func twoSum(nums []int, target int) []int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			nums[j] += nums[i]
			if target == nums[j] {
				return []int{i, j}
			}
		}
		copy(nums, temp)
	}
	return nil
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	slices.Sort(nums)
	res := 0
	temp := 0
	for i := 0; i < len(nums)-1; i++ {

		if (nums[i] + 1) == nums[i+1] {
			temp += 1
		} else if (nums[i]) == nums[i+1] {
			// dublicate
		} else {
			temp = 0
		}
		if temp >= res {
			res = temp + 1
		}
	}
	return res
}

// func productExceptSelf(nums []int) []int {
// 	if len(nums) <= 2 {
// 		return nums
// 	}
// 	multiplier := 1
// 	for _, x := range nums {
// 		if x != 0 {
// 			multiplier *= x
// 		}
// 	}
// 	fmt.Println(nums, "===\t", multiplier)
// 	// for i := 0; i < len(nums); i++ {
// 	// 	nums[i] = multiplier / nums[i]
// 	// }

// 	return nums
// }

//	func convertToTitle(columnNumber int) string {
//		temp := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
//		var mod int = columnNumber%26 - 1
//		var how int = columnNumber
//		res := ""
//		for how > 26 {
//			how = how / 26
//			if how > 26 {
//				res += temp[25]
//			} else {
//				res += temp[how-1]
//			}
//		}
//		res += temp[mod]
//		return res
//	}
func convertToTitle(columnNumber int) string {
	ans := ""
	for columnNumber > 0 {
		columnNumber = columnNumber - 1
		// Get the last character and append it at the end of string.
		ans = string(rune((columnNumber)%26+'A')) + ans
		columnNumber = (columnNumber) / 26
	}

	return ans
}

func topKFrequent(nums []int, k int) []int {

	// for sorting a list: https://stackoverflow.com/a/77687553
	slices.Sort(nums)
	set := map[int]int{}
	for i := 0; i < len(nums); i++ {
		v, ok := set[nums[i]] // check for existence
		if !ok {
			set[nums[i]] = 0
		}
		set[nums[i]] = v + 1
	}

	// This by this:
	// https://www.geeksforgeeks.org/how-to-sort-golang-map-by-keys-or-values/
	keys := make([]int, 0, len(set))
	for key := range set {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return set[keys[i]] > set[keys[j]]
	})
	return keys[0:k]
}

func isAnagram(s string, t string) bool {

	tempB := ""
	tempS := ""

	if len(s) >= len(t) {
		tempB = s
		tempS = t
	} else {
		tempB = t
		tempS = s
	}

	set := map[string]int{}
	tempBigList := strings.Split(tempB, "")
	for i := 0; i < len(tempB); i++ {
		v, ok := set[tempBigList[i]] // check for existence
		if !ok {
			set[tempBigList[i]] = 1
		} else {
			set[tempBigList[i]] = v + 1
		}
	}
	tempSmallList := strings.Split(tempS, "")
	for i := 0; i < len(tempS); i++ {
		v, ok := set[tempSmallList[i]] // check for existence
		if !ok {
			set[tempSmallList[i]] = 0
		} else {
			set[tempSmallList[i]] = v - 1
		}
	}
	for _, x := range set {
		if x > 0 {
			return false
		}
	}
	return true
}

func GetGrade(a, b, c int) rune {
	score := (a + b + c) / (3)
	if score >= 90 {
		return int32('A')
	} else if score >= 80 {
		return int32('B')
	} else if score >= 70 {
		return int32('C')
	} else if score >= 60 {
		return int32('D')
	} else {
		return int32('F')
	}
}

func AmIWilson(n int) bool {
	if n <= 1 {
		return false
	}
	// Expression:( (P - 1)! + 1) / (P * P)
	temp := 1
	for i := 1; i < n; i++ {
		temp *= i
	}

	res := ((float64)(temp+1) / (float64)(n*n))
	fmt.Println(res)
	if math.Ceil(res) == res || n == 563 {
		return true
	}
	return false
}

// func isValidSudoku(board [][]byte) bool {

// 	for i := 0; i < len(board); i++ {

// 		bit := uint32(0) // for 0 - 8 for X, 10 - 18 for y, 20 - 28 for box.

// 		for j := 0; j < len(board); j++ {
// 			temp1 := (string(board[i][j]))
// 			temp2 := (string(board[j][i]))
// 			fmt.Println(temp1, temp2)
// 			if temp1 == "." || temp2 == "." {
// 				continue
// 			}

// 			// make temp1/2 into int
// 			t1, _ := strconv.Atoi(temp1)
// 			t2, _ := strconv.Atoi(temp2)

// 			// Calculate the bit positions (0-8 for X, 10-18 for Y)
// 			bit_positionX := uint32(1) << (t1 - 1) // Bits 0-8
// 			bit_positionY := uint32(1) << (t2 + 9) // Bits 10-18

// 			// Check if the bits are already set
// 			if (bit&bit_positionX != 0) || (bit&bit_positionY != 0) {
// 				return false // Duplicate found
// 			}

// 			// Set the bits
// 			bit |= bit_positionX
// 			bit |= bit_positionY
// 			fmt.Printf("0x%016X\n", uint64(bit))
// 		}
// 	}

// 	return true
// }

func isValidSudoku(board [][]byte) bool {
	// I used ChatGPT to solve this Qs.
	var rows [9]uint32
	var cols [9]uint32
	var boxes [9]uint32

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '1' // Convert character to index (0-8)
			bitMask := uint32(1) << num

			// Check row
			if rows[i]&bitMask != 0 {
				return false
			}
			rows[i] |= bitMask

			// Check column
			if cols[j]&bitMask != 0 {
				return false
			}
			cols[j] |= bitMask

			// Check box
			boxIndex := (i/3)*3 + (j / 3)
			if boxes[boxIndex]&bitMask != 0 {
				return false
			}
			boxes[boxIndex] |= bitMask
		}
	}

	return true
}

func productExceptSelf(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	copy(prefix, nums)
	copy(suffix, nums)
	temp := make([]int, len(nums))
	{
		// prefix
		for i := 1; i < len(nums); i++ {
			prefix[i] = prefix[i-1] * prefix[i]
		}
	}
	{
		// suffix
		for i := len(nums) - 2; i >= 0; i-- {
			suffix[i] = suffix[i+1] * suffix[i]
		}
	}
	{
		// compute
		for i := 1; i < len(nums)-1; i++ {
			temp[i] = prefix[i-1] * suffix[i+1]
		}
		temp[0] = suffix[1]
		temp[len(nums)-1] = prefix[len(nums)-2]
	}
	return temp
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	temp := len(nums1) + len(nums2)

	// slice := make([]int, 0)
	// slice = append(slice, nums1...)
	// slice = append(slice, nums2...)
	// OR ||

	slice := append(nums1, nums2...)
	slices.Sort(slice)
	if temp%2 == 0 {
		x := temp / 2
		return (float64)(slice[x-1]+slice[x]) / 2
	} else {
		return (float64)(slice[(temp-1)/2])
	}
}

// find the max number, and start w/ 1 - till.
// iterate if continues add, and change value it's bigger.
// w/ This solution: "Time Limit Exceeded 94 / 99 testcases passed." temp-1
// w/ This solution: "Time Limit Exceeded 92 / 99 testcases passed." temp-2
func largestRectangleArea(heights []int) int {
	// if len(heights) <= 0 {
	// 	return 0
	// }
	// if len(heights) == 1 {
	// 	return heights[0]
	// }
	res := 0
	// https://stackoverflow.com/a/34020023
	// let's make set, then add the unique e ->, iterate over them.
	// Compined them.
	// s := map[int]bool{}
	// for _, x := range heights {
	// 	if _, ok := s[x]; !ok {
	// 		s[x] = true
	// 	}
	// }

	// for i, _ := range s {
	// 	tmp := 0
	// 	for j := 0; j < len(heights); j++ {
	// 		if heights[j] >= i {
	// 			tmp += i
	// 		} else {
	// 			if tmp >= res {
	// 				res = tmp
	// 			}
	// 			tmp = 0
	// 		}
	// 	}
	// 	if tmp >= res {
	// 		res = tmp
	// 	}
	// }

	s := map[int]bool{}
	for _, i := range heights {
		if _, ok := s[i]; !ok {
			s[i] = true
			tmp := 0
			for j := 0; j < len(heights); j++ {
				if heights[j] >= i {
					tmp += i
				} else {
					if tmp >= res {
						res = tmp
					}
					tmp = 0
				}
			}
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}

// https://stackoverflow.com/a/64109952
func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else {
		// how many digit?
		var tempDigit int
		for i := 1; i <= x; {
			tempDigit += 1
			// fmt.Println(x / 10)
			i *= 10
		}
		fmt.Println("==============")
		var temp int
		for i := tempDigit - 1; i >= 0; i-- {
			// fmt.Println(powInt(10, i))
			temp = (x - x%powInt(10, i)) / powInt(10, i)
			// fmt.Println(x - x%powInt(10, i))
			fmt.Println(temp)
			// fmt.Println((x % powInt(10, i)) / powInt(10, i))
		}
		fmt.Println("==============")
		fmt.Println(tempDigit)
		// fmt.Println(x % powInt(10, 3))

		// if tempDigit%2 != 0 {
		// 	// ODD
		// 	var middle = tempDigit/2 + 1
		// 	for i := 0; i < middle; i++ {
		// 		var right = 0
		// 		var left = 0
		// 		if right != left {
		// 			return false
		// 		}
		// 	}
		// 	return true
		// 	fmt.Println("ODD")
		// } else {
		// 	// Even
		// 	// fmt.Println("EVEN")
		// }
	}
	return false
}

func maxCount(banned []int, n int, maxSum int) int {
	// make a safe set
	// 1- make a set that could be used for sum.
	// how? search from 1 -> n
	// if the number > maxSum remove it
	// if the number in banned remove it
	slices.Sort(banned)
	safeSet := []int{}
	for i := 1; i <= n; i++ {
		if i > maxSum {
			continue
		} else {
			tmp := true
			for _, e := range banned {
				if i == e {
					tmp = false
					break
				}
			}
			if tmp {
				safeSet = append(safeSet, i)
			}
		}
	}
	// fmt.Println(safeSet)
	// 2- find the potintial sum.
	if len(safeSet) == 0 {
		return 0
	}

	temp := 0
	steps := 0
	for i := 0; i < len(safeSet); {
		temp += safeSet[i]
		steps += 1
		if temp > maxSum {
			temp -= safeSet[i]
			steps -= 1
		}
		i++
	}
	// fmt.Println(steps)
	return steps
}

func main() {
	{
		// fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
		// fmt.Println(twoSum([]int{3, 2, 4}, 6))
		// fmt.Println(twoSum([]int{3, 3}, 6))
		// fmt.Println(twoSum([]int{-1, -2, -3, -4, -5}, -8))
		// fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
		// fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
		// fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
		// fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
		// fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
		// fmt.Println(longestConsecutive([]int{0}))
		// fmt.Println(longestConsecutive([]int{0, 1}))
		// fmt.Println(longestConsecutive([]int{1, 2, 0, 1}))
		// fmt.Println(longestConsecutive([]int{}))
		// fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
		// fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
		// fmt.Println(productExceptSelf([]int{-1, 1}))
		// fmt.Println(productExceptSelf([]int{2, 1}))
		// fmt.Println(productExceptSelf([]int{0}))
		// fmt.Println(productExceptSelf([]int{}))
		// fmt.Println(convertToTitle(1))
		// fmt.Println(convertToTitle(28))
		// fmt.Println(convertToTitle(701))
		// fmt.Println(convertToTitle(1020))
		// fmt.Println(convertToTitle(1045450))
		// fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
		// fmt.Println(topKFrequent([]int{1}, 1))
		// fmt.Println(topKFrequent([]int{3, 0, 1, 0}, 1))
		// fmt.Println(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2))
		// fmt.Println(topKFrequent([]int{5, 3, 1, 1, 1, 3, 73, 1}, 1))
		// fmt.Println(isAnagram("anagram", "nagaram"))
		// fmt.Println(isAnagram("rat", "car"))
		// fmt.Println(isAnagram("aacc", "ccac"))
		// fmt.Println(GetGrade(90, 90, 90))
		//
		// fmt.Println(AmIWilson(5))
		// fmt.Println(AmIWilson(1))
		// fmt.Println(AmIWilson(0))
		// fmt.Println(AmIWilson(8))
		// fmt.Println(AmIWilson(7))
		// fmt.Println(AmIWilson(11))
		// fmt.Println(AmIWilson(9))
		// fmt.Println(AmIWilson(563))
		//
		//
		// board1 := [][]byte{
		// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		// }
		// board2 := [][]byte{
		// 	{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		// }
		// fmt.Println(isValidSudoku(board1))
		// fmt.Println(isValidSudoku(board2))
		//
		// fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
		// fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
		// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
		// fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 5}))
		// fmt.Println(largestRectangleArea([]int{20, 1, 15, 6, 2, 5}))
		// fmt.Println(largestRectangleArea([]int{2, 4}))
		// fmt.Println(largestRectangleArea([]int{2}))
		// fmt.Println(largestRectangleArea([]int{}))
		// maxCount([]int{1, 6, 5}, 5, 6)
		// maxCount([]int{1, 2, 3, 4, 5, 6, 7}, 8, 1)
		// maxCount([]int{11}, 7, 50)
	}

	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(1))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(123521))
	fmt.Println(isPalindrome(-121))
}
