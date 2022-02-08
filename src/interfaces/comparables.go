package interfaces

// A comparable item that has to implement the compare to function
// for the linkedlist to know where to insert and make comparisons
type ComparableItem interface {
	// CompareTo should return
	// -1 if less than item
	// 1 if more than item
	// 0 if equal to item
	CompareTo(item interface{}) int
	Print() string
}
