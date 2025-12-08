func nextGreaterElement(nums1 []int, arr []int) []int {
    stack := []int{}
    mp := make(map[int]int, len(arr))

    mp[arr[len(arr)-1]] = -1
    stack = append(stack, arr[len(arr)-1])

    for i := len(arr) - 2; i >= 0; i-- {
        for len(stack) > 0 && arr[i] >= stack[len(stack)-1] {
            stack = stack[:len(stack)-1]
        }

        if len(stack) == 0 {
            mp[arr[i]] = -1
        } else {
            mp[arr[i]] = stack[len(stack)-1]   
        }

        stack = append(stack, arr[i])
    }

    ans := []int{}
    for _, v := range nums1 {                  
        ans = append(ans, mp[v])
    }

    return ans
}
