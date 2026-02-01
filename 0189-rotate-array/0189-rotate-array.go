func rotate(arr []int, d int)  {
	d = d % len(arr)
	for _ = range d {
		tmp := arr[len(arr)-1]

		for j := len(arr) - 1; j > 0; j-- {
			arr[j] = arr[j-1]
		}
		arr[0] = tmp
	}
}