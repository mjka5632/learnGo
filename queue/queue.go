package queue

//{}表示任何类型
type Queue [] interface{}

//使用别名
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	//转换成int
	return head.(int)
}

func (q *Queue) IsEmpty() bool {

	return len(*q) == 0

}
