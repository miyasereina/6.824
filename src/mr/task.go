package mr

type task interface {
	Do()   //Do 执行
	Pong() //Pong 响应
}
