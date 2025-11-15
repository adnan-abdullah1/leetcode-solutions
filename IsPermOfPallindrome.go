package main
 // other solution using map
 // this is without map
 // book:60
func IsPermOfPallindrome(str string) bool{
	table := [26]int{}
	oddCount:=0
	for _,v:=range str {

		if v >= 'A' && v<='Z'{
			v = v + ('a'-'A')
		}

		if v < 'a' || v > 'z'{
			continue
		}

		table[v-'a']++
		if table[v-'a'] % 2 != 0 {
			oddCount++
		} else{
			oddCount--
		}

	}

	return oddCount <= 1
}
