package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/YashwantSingh7062/service2/lib/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	c := controller.Controller{}
	e.GET("/", c.GetOrder)

	go func() {
		e.Logger.Fatal(e.Start(":5000"))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		fmt.Println("Forcefully shutting Down")
	}

	fmt.Println("ShutDown Successful")
}
