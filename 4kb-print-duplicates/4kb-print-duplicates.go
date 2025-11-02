package kbprintduplicates

import "fmt"

// question on book cracking coding interview
func printDup() {
	arr := []int{1, 2, 2, 2, 5555, 2, 3, 3, 3, 4, 5, 6, 6, 444, 5555}
	mem := make([]byte, 4000)

	for _, v := range arr {
		byteIndex := (v - 1) / 8
		bitIndex := (v - 1) % 8

		if mem[byteIndex]&(1<<bitIndex) != 0 {
			fmt.Println(v)
		} else {
			mem[byteIndex] |= (1 << bitIndex)
		}
	}
}
