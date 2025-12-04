
func evalRPN(tokens []string) int {
    stack := []int{}

    for _, v := range tokens {
        if v == "*" || v == "/" || v == "-" || v == "+" {

            var op1, op2 int
            stack, op1 = pop(stack)
            stack, op2 = pop(stack)

            res := 0

            if v == "+" {
                res = op2 + op1
            } else if v == "-" {
                res = op2 - op1
            } else if v == "*" {
                res = op2 * op1
            } else {
                res = op2 / op1
            }

            stack = append(stack, res)

        } else {
            // safe
            val, _ := strconv.Atoi(v)
            stack = append(stack, val)
        }
    }

    return stack[len(stack)-1]
}

func pop(stack []int) ([]int, int) {
    el := stack[len(stack)-1]
    return stack[:len(stack)-1], el
}
