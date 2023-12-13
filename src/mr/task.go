package mr

type task interface {
	Do()   //Do 执行
	Done() //Done 完成任务
	Pong() //Pong 响应
}

const (
	AvailableIndex = iota
	workingIndex
	DongIndex
)

const (
	MapIndex = iota
	ReduceIndex
)

type WorkerPool [2][3]map[string]*TaskWorker

type TaskWorker struct {
	State int8
}

func (p *WorkerPool) AvailableMap() map[string]*TaskWorker {
	return p[MapIndex][AvailableIndex]
}
func (p *WorkerPool) WorkingMap() map[string]*TaskWorker {
	return p[MapIndex][workingIndex]
}
func (p *WorkerPool) DoneMap() map[string]*TaskWorker {
	return p[MapIndex][DongIndex]
}

func (p *WorkerPool) AvailableReduce() map[string]*TaskWorker {
	return p[ReduceIndex][AvailableIndex]
}
func (p *WorkerPool) WorkingReduce() map[string]*TaskWorker {
	return p[ReduceIndex][workingIndex]
}
func (p *WorkerPool) DoneReduce() map[string]*TaskWorker {
	return p[ReduceIndex][DongIndex]
}
