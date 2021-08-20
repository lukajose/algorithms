package queue

type QueueIterator struct {
	index int
	arr   []interface{}
}

func (qit *QueueIterator) HasNext() bool {

	return qit.index < len(qit.arr)
}

func (qit *QueueIterator) Next() interface{} {
	if qit.HasNext() {
		index := qit.index
		qit.index++
		return &(qit.arr[index])
	}
	return nil
}
