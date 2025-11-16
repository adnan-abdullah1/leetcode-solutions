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

    if negIndex == -1{
        m := len(nums)-1
        for _,v := range nums{
            ans[m]=v
            m--
        }
        return ans
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