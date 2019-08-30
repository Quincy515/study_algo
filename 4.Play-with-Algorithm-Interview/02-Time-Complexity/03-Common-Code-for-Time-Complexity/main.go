package main

import "fmt"

// O(1)
func swapTwoInts(a, b int) {
	tmp := a
	a = b
	b = tmp
}

// O(n)
func sum(n int) int {
	assert(n >= 0)
	ret := 0
	for i := 0; i <= n; i++ {
		ret += i
	}
	return ret
}

func reverse(s string) string {
	n := len(s)
	ret := ""
	for i := 0; i < n/2; i++ {
		ret = swapString(s, i, n-1-i)
	}
	return ret
}

// O(n^2) Time Complexity
func selectionSort(arr []int, n int) {
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swapInt(arr, i, minIndex)
	}
}

// O(n) Time Complexity
func printInformation(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= 30; j++ {
			fmt.Println("Class ", i, " - ", "No. ", j)
		}
	}
}

// O(logn) Time Complexity
func binarySearch(arr []int, n, target int) int {
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func intToString(num int) string {
	s := ""
	sign := "+"
	if num < 0 {
		num = -num
		sign = "-"
	}

	for num != 0 {
		s += string('0' + (num % 10))
		num /= 10
	}
	if s == "" {
		s = "0"
	}
	s = reverse(s)
	if sign == "-" {
		return sign + s
	} else {
		return s
	}
}

// O(nlogn)
func hello(n int) {
	for sz := 1; sz < n; sz += sz {
		for i := 1; i < n; i++ {
			fmt.Println("Hello, Algorithm!")
		}
	}
}

// O(sqrt(n)) Time Complexity
func isPrime(num int) bool {
	for x := 2; x*x <= num; x++ {
		if num%x == 0 {
			return false
		}
	}
	return true
}

func isPrime2(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for x := 3; x*x <= num; x += 2 {
		if num%x == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("intToString(123) ", intToString(123))
	fmt.Println("intToString(0) ", intToString(0))
	fmt.Println("intToString(-123) ", intToString(-123))

	if isPrime2(137) {
		fmt.Println("137 is a prime.")
	} else {
		fmt.Println("137 is not a prime.")
	}

	if isPrime2(121) {
		fmt.Println("121 is a prime.")
	} else {
		fmt.Println("121 is not a prime.")
	}

	arr := []int{6, 9, 3, 2, 1, 8}
	selectionSort(arr, len(arr))
	fmt.Println(arr)

	fmt.Println(binarySearch(arr, len(arr), 0)) // 没找到返回-1
	fmt.Println(binarySearch(arr, len(arr), 6)) // 返回数组下标3
}

func swapString(s string, a, b int) string {
	arr := []rune(s)
	arr[a], arr[b] = arr[b], arr[a]
	return string(arr)
}

func swapInt(arr []int, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

func assert(exp bool) {
	if !exp {
		panic("断言失败!")
	}
}
