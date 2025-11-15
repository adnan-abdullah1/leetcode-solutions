package main
import ("fmt"
)


func main(){
	fmt.Println(IsOneEditAway("abc","abde"))
}
func IsOneEditAway(str1,str2 string) bool{
	diff := len(str1)-len(str2)
	if diff < -1 || diff > 1{
		return false
	}
	noOfEdits:=0

	i :=0
	j:=0
	for i < len(str1) && j < len(str2){
		if str1[i] != str2[j]{
			noOfEdits++
			if len(str1)==len(str2){
				i++
				j++
			} else if len(str1) < len(str2){
				j++
			} else{
				i++
			}
		} else{
			i++
			j++
		}
	}

	// leftover
	if i < len(str1) || j < len(str2){
		noOfEdits++
	}

	return noOfEdits<=1
}
