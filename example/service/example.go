package service

import (
	"errors"
	"fmt"
	"github.com/linxlib/logs"
	"github.com/robfig/cron/v3"
	"sync"
)

func Init()  {
	logs.Error(fmt.Errorf("Test Error: %+v",errors.New("error example")))
}

type WeatherJob struct {
	mtx     sync.Mutex
	running bool
}

func (w *WeatherJob) Run() {
	if !w.SetRun() {
		logs.Warn("天气服务正在运行")
		return
	}
	logs.Error(fmt.Errorf("Test Error: %+v",errors.New("error example")))

	w.doneRun()
}

func (w *WeatherJob) SetRun() bool {
	w.mtx.Lock()
	defer w.mtx.Unlock()
	if w.running {
		return false
	}
	w.running = true
	return true
}

func (w *WeatherJob) doneRun() {
	w.mtx.Lock()
	w.running = false
	defer w.mtx.Unlock()
}
type SingleJob interface {
	cron.Job

	SetRun() bool
	doneRun()
}
