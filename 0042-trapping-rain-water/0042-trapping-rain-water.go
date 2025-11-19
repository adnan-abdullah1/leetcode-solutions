// n square solution
func trap(arr []int) int {
    trapped := 0
	for i := range arr {
		rMax := getRightMax(arr, i)
		lMax := getLeftMax(arr, i)
		trapped += min(rMax, lMax) - arr[i]
	}
	return trapped
}

func getLeftMax(arr []int, i int) int {
	maxNo := 0
	for i := i; i >= 0; i-- {
		if arr[i] > maxNo {
			maxNo = arr[i]
		}
	}
	return maxNo
}

func getRightMax(arr []int, i int) int {
	maxNo := 0
	for i := i; i < len(arr); i++ {
		if arr[i] > maxNo {
			maxNo = arr[i]
		}
	}
	return maxNo
}

// better solution can be by computing
// premax sum, and postmax sum for every element
// tc: o(n), sc: o(n)

// func trap(arr []int) int {
// 	maxLeft := calcPreMax(arr)
// 	maxRight := calcPostMax(arr)
// 	fmt.Println(maxLeft)
// 	fmt.Println(maxRight)

// 	trapped := 0
// 	for i := range arr {
// 		trapped += min(maxLeft[i], maxRight[i]) - arr[i]
// 	}
// 	return trapped
// }

func calcPreMax(arr []int) []int {
	maxNo := 0
	res := make([]int, len(arr))
	for i, v := range arr {
		if v > maxNo {
			maxNo = v
		}
		res[i] = maxNo
	}
	return res
}
func calcPostMax(arr []int) []int {
	maxNo := 0
	res := make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > maxNo {
			maxNo = arr[i]
		}
		res[i] = maxNo
	}
	return res
}

// most optimal soution is two pointer