package task

import (
	"os"
	"os/signal"
	"time"
)

// 同步执行任务
type RunnerAsync struct {
	interrupt chan os.Signal   //操作系统的信号检测
	complte   chan error       // 记录执行完成的状态
	timeout   <-chan time.Time // 超时检测
	tasks     []func(id int)   //保存所有要执行的任务，顺序执行
}

// new一个RunnerAsync对象
func NewRunnerAsync(d time.Duration) *RunnerAsync {
	return &RunnerAsync{
		interrupt: make(chan os.Signal, 1),
		complte:   make(chan error),
		timeout:   time.After(d),
	}
}

//添加一个任务
func (this *RunnerAsync) Add(tasks ...func(id int)) {
	this.tasks = append(this.tasks, tasks...)
}

// 判断是否接收到操作系统中断信号
func (this *RunnerAsync) gotInterrupt() bool {
	select {
	case <-this.interrupt:
		// 停止接收别的信号
		signal.Stop(this.interrupt)
		return true

	// 正好执行
	default:
		return false
	}
}

//顺序执行所有任务
func (this *RunnerAsync) Run() error {
	for id, task := range this.tasks {
		if this.gotInterrupt() {
			return ErrInterrupt
		}
		// 执行任务
		task(id)
	}
	return nil
}

//启动RunnerAsync， 监听错误信息
func (this *RunnerAsync) Start() error {
	// 接收操作系统信号
	signal.Notify(this.interrupt, os.Interrupt)
	// 执行任务
	go func() {
		this.complte <- this.Run()
	}()

	select {
	// 返回执行结果
	case err := <-this.complte:
		return err
	//超时返回
	case <-this.timeout:
		return ErrTimeout

	}
}
