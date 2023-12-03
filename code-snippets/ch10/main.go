package main

// 泛型
type Slice[T int | string] []T

type Map[K int | string, V float32 | float64] map[K]V

type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

// 类型约束字面量，通常可省略外层interface{}
// func min[T interface{ int | float64 }](a, b T) T {
// 	if a <= b {
// 		return a
// 	}
// 	return b
// }

// 事先定义的类型约束类型
type Value interface {
	int | float64
}

func min[T Value](a, b T) T {
	if a <= b {
		return a
	}
	return b
}
