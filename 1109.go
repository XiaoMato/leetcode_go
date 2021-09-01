package main

func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n+1)
	for _, book := range bookings {
		for i := book[0]; i <= book[1]; i++ {
			res[i] += book[2]
		}
	}
	return res
}
