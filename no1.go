func seqSearch(T tabInt, N int, X int) bool {
	var ketemu bool
	var k int

	ketemu = false
	k = 0

	for !ketemu && k < N {
		ketemu = T[k] == X
		k++
	}
	return ketemu
}