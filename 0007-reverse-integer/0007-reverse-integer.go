func reverse(x int) int {
    ans := 0
    sig := 1
    if x < 0{
        sig=-1
        x=-x
    }
    for x>0{
        rem := x%10
        x = x/10
        ans =( ans * 10) + rem
    }
    if ans < math.MinInt32 || ans > math.MaxInt32{
        return 0
    }
    return sig*ans
}