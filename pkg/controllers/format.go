package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	FormatPath EndpointPath = "/format"
)

type Format struct {
	Base
}

type FormatInput struct {
	Base
}

func NewFormatController(input FormatInput) *Format {
	return &Format{
		Base: input.Base,
	}
}

func (c *Format) Register(r *gin.Engine) {
	r.POST(c.Base.Path(FormatPath), c.FormatHandler())
}

func (c *Format) FormatHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := c.MustGet("logger").(*logrus.Entry)

		// Extract the file from the request
		file, err := c.FormFile("file")
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
			return
		}

		// Open the file for reading
		fileReader, err := file.Open()
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open uploaded file"})
			return
		}

		defer fileReader.Close()

		// Read the file into a string
		content, err := ioutil.ReadAll(fileReader)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read uploaded file"})
			return
		}

		// Serve the file to the client
		c.Header("Content-Disposition", "attachment; filename="+"output-foo.csv")
		c.Header("Content-Type", "application/text/plain")
		c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
		c.Writer.Write([]byte(content))

	}
}
