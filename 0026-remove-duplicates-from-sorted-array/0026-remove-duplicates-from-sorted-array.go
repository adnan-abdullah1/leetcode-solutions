func removeDuplicates(nums []int) int {
    k :=1
    j:=0
	for i:=0;i<len(nums)-1;i++{ 

        
        for j<len(nums) && nums[i]==nums[j]{
            j++
        }

        if j>=len(nums) {
            return k
        }
        
        nums[i+1]=nums[j]
        k++
    }
    return k
}