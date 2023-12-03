package main

// 2-3 常量

const (
	pi = 3.1415
	e  = 2.7182
)

const (
	n1 = 100
	n2
	n3
)

const (
	level0 = iota // 0
	level1 = iota // 1
	level2 = iota // 2
	level3 = iota // 3
)

const (
	_  = iota
	KB = 1 << (10 * iota) // 1 << (10 * 1)
	MB = 1 << (10 * iota) // 1 << (10 * 2)
	GB = 1 << (10 * iota) // 1 << (10 * 3)
	TB = 1 << (10 * iota) // 1 << (10 * 4)
	PB = 1 << (10 * iota) // 1 << (10 * 5)
)

const (
	a, b = iota + 1, iota + 2 // 1, 2
	c, d                      // 2, 3
)
