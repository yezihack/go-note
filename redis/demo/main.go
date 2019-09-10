package queue

type Queue struct {
	items  []int
	unused int
	Header int
}

func (q *Queue) Append(val int) bool {
	return true
}
