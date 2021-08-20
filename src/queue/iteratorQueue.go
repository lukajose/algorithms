package queue

type queueIterator struct {
	index int
	arr   []interface{}
}

func (iter *queueIterator) hasNext() bool {

	return true
}

func (iter *queueIterator) next() interface{} {

	return nil
}
