
// O(N^2)
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums)-1; i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{-1, -1}
// }
// second solution will be sort,
// then do target - arr[i]=key
// then find that key with binary search
// sorting: nlog(n)+ loop and bs (o(log(n)))

// third solution
// insert elements to map
// then look if target-arr[i] is in map
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		k := target - nums[i]

        v,ok:=mp[k]
        if !ok {
            continue
        }

        if v != i{
            return []int{i,v}
        }
        
	}

	return []int{}
}