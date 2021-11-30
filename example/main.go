package main

import (
	"errors"
	"example/service"
	"fmt"
	"github.com/linxlib/logs"
	"github.com/robfig/cron/v3"
)



func main() {
	logs.Error(fmt.Errorf("Test Error: %+v",errors.New("error example")))
	service.Init()

	c := cron.New(cron.WithSeconds())
	c.AddJob("0 0 6/2 * * ? ", new(service.WeatherJob))

	c.Start()
	select {

	}
}

