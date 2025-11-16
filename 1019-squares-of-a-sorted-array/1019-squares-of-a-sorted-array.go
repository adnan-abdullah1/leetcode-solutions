func sortedSquares(nums []int) []int {
	// find first non negative
    negIndex:=-1
	for i, v := range nums {
		if v >= 0 && negIndex == -1 {
			negIndex = i
		}

		nums[i] = v * v
	}

    ans := make([]int,len(nums))

    if negIndex == -1{ // means there where only negative numser
        // 5,4,3,2,1
        i := 0
        j := len(nums)-1
        for i <= j{
            nums[i],nums[j]=nums[j],nums[i]
            i++
            j--
        }
        return nums
    }

    if negIndex==0{
        return nums
    }

	// negIndex - 1 to 0 are sorted
    // negIndex to len -1 are sorted 
    i:=0
    j:=len(nums)-1
    k := j
    for i < negIndex && j >= negIndex{
        if nums[i] > nums[j]{
            ans[k]=nums[i]
            i++
            k--
        } else{
            ans[k]=nums[j]
            j--
            k--
        }
    }

    for i < negIndex{
        ans[k]=nums[i]
        i++
        k--
    }

     for j >= negIndex{
        ans[k]=nums[j]
        j--
        k--
    }
	return ans
}