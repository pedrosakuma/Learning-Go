package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

var client appinsights.TelemetryClient

func Telemetry() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		log.Print(latency)

		status := c.Writer.Status()

		client.TrackRequest(c.Request.Method, c.Request.URL.String(), latency, string(status))
	}
}

func initTelemetry() {
	telemetryConfig := appinsights.NewTelemetryConfiguration("13d73cf0-6e15-4e06-9d00-341a7902863e")

	// Configure how many items can be sent in one call to the data collector:
	telemetryConfig.MaxBatchSize = 8192

	// Configure the maximum delay before sending queued telemetry:
	telemetryConfig.MaxBatchInterval = 2 * time.Second

	client = appinsights.NewTelemetryClientFromConfig(telemetryConfig)
}
