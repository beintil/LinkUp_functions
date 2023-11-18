package gohelp

// Ptr returns a pointer to the value
func Ptr[T any](v T) *T { return &v }

// RemoveElement removes the element from the slice for according to the received index
func RemoveElement[T any](slice []T, index int) []T {
	if index < 0 || index > len(slice) {
		return slice
	}
	if index == len(slice) {
		slice = nil
		return slice
	}
	if index+1 >= len(slice) {
		return append(slice[:index])
	}
	return append(slice[:index], slice[index+1:]...)
}
