package app

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"fmt"
	"os"
)

func Run() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	//---

	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Controller("/", new(RootController))

	app.Run(iris.Addr(":" + port), iris.WithoutVersionChecker, iris.WithConfiguration(iris.Configuration{ // default configuration:
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

	return 200
}
