package slicelist

import (
	//"errors"
	"fmt"
)

type SliceList[T comparable] struct {
	slice []T
}

func NewSliceList[T comparable]() *SliceList[T] {

	return &SliceList[T]{}
}

func (sl *SliceList[T]) Append(value T) {

	sl.slice = append(sl.slice, value)
}

func (sl *SliceList[T]) Prepend(value T) {

	sl.slice = append([]T{value}, sl.slice...)
}

func (sl *SliceList[T]) InsertAt(value T, position int) {

	sl.slice = append(sl.slice[:position], append([]T{value}, sl.slice[position:]...)...)
}

func (sl *SliceList[T]) Remove(value T) bool {

	for i, v := range sl.slice {
		
		if v == value{

			sl.slice = append(sl.slice[:i], sl.slice[i+1:]...)
			return true
		}
	}

	return false
}

func (sl *SliceList[T]) Search (value T) int {

	for i, v := range sl.slice {
		
		if value == v {
			
			return i
		}
	}

	return -1
}

func (sl *SliceList[T]) Get(position int) T{

	return sl.slice[position]
}

func (sl *SliceList[T]) Size() int {

	return len(sl.slice)
}

func (sl *SliceList[T]) String() string {

	if len(sl.slice) == 0 {
		
		return "[]"
	}

	result := "["

	for _, v := range sl.slice {
		
		result += fmt.Sprintf("%v", v)
	}

	result += "]"

	return result
}
