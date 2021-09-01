package corpFlightBookings

/*
lc:https://leetcode-cn.com/problems/corporate-flight-bookings/
*/

func corpFlightBookings(bookings [][]int, n int) []int {
	var counter = make([]int, n, n)
	for _, booking := range bookings {
		counter[booking[0]-1] += booking[2]
		if booking[1] < n {
			counter[booking[1]] -= booking[2]
		}
	}

	for i := range counter {
		if i == 0 {
			continue
		}
		counter[i] += counter[i-1]
	}

	return counter
}
