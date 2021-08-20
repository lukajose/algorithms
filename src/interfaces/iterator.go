package iterator

type Collection interface {
	iterator() Iterator
}

type Iterator interface {
	hasNext() bool
	next() interface{}
}
