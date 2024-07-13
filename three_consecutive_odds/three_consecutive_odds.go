package main

func threeConsecutiveOdds(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i]%2 == 1 && i+1 < len(arr) {
			if arr[i+1]%2 == 1 && i-1 >= 0 {
				if arr[i-1]%2 == 1 {
					return true
				}
			}
		}
	}
	return false
}