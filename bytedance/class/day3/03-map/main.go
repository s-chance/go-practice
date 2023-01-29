package _3_map

func NoPreAlloc(size int) {
	data := make(map[int]int)
	for i := 0; i < size; i++ {
		data[i] = i
	}
}

func PreAlloc(size int) {
	data := make(map[int]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
}
