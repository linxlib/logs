package main

import (
	"errors"
	"example/service"
	"fmt"
	"github.com/linxlib/logs"
	"github.com/robfig/cron/v3"
)

func main() {
	logs.AddFileHook("example")
	logs.Error(fmt.Errorf("Test Error: %+v", errors.New("error example")))
	service.Init()
	logs.Infoln("sdasd")
	logs.Traceln("sjahdasd")
	logs.Debugln("djahsda")
	logs.Warn("dsada")
	logs.Warn("sadhasd")

	c := cron.New(cron.WithSeconds())
	c.AddJob("0 0 6/2 * * ? ", new(service.WeatherJob))

	c.Start()
	select {}
}
