package algorithm

import (
	"fmt"
)

func PairSockCount(n int, ar []int) int {
	sockCount := make(map[int]int)
	pairCount := 0

	for _, color := range ar {
		sockCount[color]++
	}

	for _, count := range sockCount {
		pairCount += count / 2
	}

	return pairCount
}

func SockMerchant() {
	var n int
	fmt.Print("Enter the number of socks: ")
	fmt.Scan(&n)

	ar := make([]int, n)
	fmt.Print("Enter the sock colors separated by space: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&ar[i])
	}

	result := PairSockCount(n, ar)
	fmt.Println("Total number of matching pairs:", result)
}
