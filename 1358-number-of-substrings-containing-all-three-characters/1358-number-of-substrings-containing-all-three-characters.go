
// // suffers TLE
// func numberOfSubstrings(str string) int {
// 	var count int = 0

// 	for i := 0; i < len(str); i++ {
// 		var A, B, C = false, false, false
// 		for j := i; j < len(str); j++ {
// 			ch := str[j]

// 			if ch == 'a' {
// 				A = true
// 			} else if ch == 'b' {
// 				B = true
// 			} else if ch == 'c' {
// 				C = true
// 			}

// 			// when i found a,b,c the all longer sub strings contain it also
// 			if A && B && C {
// 				count += len(str) - j
// 				break
// 			}
// 		}
// 	return count
// }
// 	}

// sliding widnow
func numberOfSubstrings(str string) int {
	var count int = 0
	var A, B, C = -1, -1, -1
	for j := 0; j < len(str); j++ {
		ch := str[j]

		if ch == 'a' {
			A = j
		} else if ch == 'b' {
			B = j
		} else if ch == 'c' {
			C = j
		}

		if A > -1 && B > -1 && C > -1 {
			count += min(A, B, C) + 1

		}
	}

	return count
}

