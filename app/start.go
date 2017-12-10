package app

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"fmt"
	"os"
	"encoding/xml"
)

func Run() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8181"
	}

	//---

	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Controller("/", new(RootController))

	app.Run(iris.Addr(":" + port), iris.WithoutVersionChecker, iris.WithConfiguration(iris.Configuration{ // default configuration:
		DisableInterruptHandler: true,
	}))
}


type RootController struct {
	mvc.C
}

type Response struct {
	XMLName 	xml.Name 	`xml:"PublishResponse"`
	Xmlns		string 		`xml:"xmlns,attr"`
	MessageId	string		`xml:"PublishResult>MessageId"`
	RequestId	string		`xml:"ResponseMetadata>RequestId"`
}

func (c *RootController) Get() string {
	return "This is my default action..."
}

func (c *RootController) Post() (string, string) {

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

	// Return the expected XML(!) response.

	resp := &Response{
		Xmlns: "http://sns.amazonaws.com/doc/2010-03-31/",
		MessageId: NewUUID(),
		RequestId: NewUUID(),
	}

	output, err := xml.MarshalIndent(resp, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return string(output), "application/xml"
}
