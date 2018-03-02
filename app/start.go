package app

import (
	"os"
	"github.com/gin-gonic/gin"
	"encoding/xml"
	"fmt"
)

type Response struct {
	XMLName 	xml.Name 	`xml:"PublishResponse"`
	Xmlns		string 		`xml:"xmlns,attr"`
	MessageId	string		`xml:"PublishResult>MessageId"`
	RequestId	string		`xml:"ResponseMetadata>RequestId"`
}

func Run() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8181"
	}

	//---

	router := gin.Default()

	router.POST("/", func(c *gin.Context) {

		message := "SNS message received"

		c.Request.ParseForm()

		for key, value := range c.Request.PostForm {
			message += fmt.Sprintf("\n-------------------------\nKey: %s", key)
			message += fmt.Sprintf("\nValue: %s", value)
		}

		message += "\n--------------------------"

		fmt.Println(message)

		resp := &Response{
			Xmlns: "http://sns.amazonaws.com/doc/2010-03-31/",
			MessageId: NewUUID(),
			RequestId: NewUUID(),
		}

		output, _ := xml.MarshalIndent(resp, "", "    ")

		c.Data(200, "application/xml; charset=utf-8", output)
	})

	//---

	router.Run()

}


