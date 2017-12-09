package app

import (
	"encoding/xml"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"fmt"
)

func Run() {

	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Controller("/", new(RootController))

	app.Run(iris.Addr(":8080"), iris.WithoutVersionChecker, iris.WithConfiguration(iris.Configuration{ // default configuration:
		DisableInterruptHandler:           true,
	}))
}


type RootController struct {
	mvc.C
}

func (c *RootController) Get() string {
	return "This is my default action..."
}

func (c *RootController) Post() int {

	message := "SNS message received"

	for key, values := range c.Ctx.FormValues() {

		message += fmt.Sprintf("\n-------------------------\nKey: %s", key)

		for _, value := range values {
			message += fmt.Sprintf("\nValue: %s", value)
		}
	}

	message += "\n--------------------------"

	c.Ctx.Application().Logger().Info(message)

	//---

	type ExampleXML struct {
		XMLName xml.Name `xml:"example"`
		One     string   `xml:"one,attr"`
		Two     string   `xml:"two,attr"`
	}

	return 200
}
