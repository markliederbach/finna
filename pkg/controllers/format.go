package controllers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markliederbach/finna/pkg/command"
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

		// Read the file
		transactions, err := command.ReadCSVToVanguardRecords(fileReader)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse uploaded file"})
			return
		}

		// Format to Stock Events
		stockEvents := transactions.ToStockEvents()
		csvData := stockEvents.ToCSVSlice()

		// Write to in-memory buffer
		var buf bytes.Buffer
		csvWriter := csv.NewWriter(&buf)
		err = csvWriter.WriteAll(csvData)
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to format uploaded file"})
			return
		}
		csvWriter.Flush()
		if err := csvWriter.Error(); err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to write formatted file"})
			return
		}

		// Serve the file to the client
		c.Header("Content-Disposition", "attachment; filename="+"output-foo.csv")
		c.Header("Content-Type", "application/text/plain")
		c.Header("Accept-Length", fmt.Sprintf("%d", len(buf.Bytes())))
		_, err = c.Writer.Write(buf.Bytes())
		if err != nil {
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "CSV writer error"})
			return
		}

		logger.Info("file formatted successfully")

	}
}
