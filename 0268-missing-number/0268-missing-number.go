func missingNumber(nums []int) int {
    sum := 0
    for _,v := range nums{
        sum+=v
    }

    n := len(nums)
    totalSum := (n*(n+1))/2
    
    return totalSum-sum
}