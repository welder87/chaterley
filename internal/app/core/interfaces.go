package core

import "fmt"

type Comparable[T any] interface {
	// сравнение на равенство с объектом такого же типа
	Equal(other T) bool
}

type Value[T any] interface {
	// возвращает значение типа T
	Val() T
}

type ValueObject[T any] interface {
	Value[T]
	Comparable[T]
	fmt.Stringer
}
