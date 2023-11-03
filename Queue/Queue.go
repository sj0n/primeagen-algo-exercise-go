package queue

type Queue struct {
	head *Queue
	tail *Queue
	value any
	length int
}
