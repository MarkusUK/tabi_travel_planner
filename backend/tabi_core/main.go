package main

import (
	"log"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func main() {
	//Initialise Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://6fd312a5dfc3ca65043063c6fbf21870@o4508902621184000.ingest.de.sentry.io/4508902624526416",
	})
	if err != nil {
		log.Fatal("sentry.Init %s", err)
	}

	//Ensure everything is flushed
	defer sentry.Flush(2 * time.Second)

	//Initialise Gin router
	router := gin.Default()

	//Use Sentry Gin middelware to capture errors etc.
	router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// An example route that intentionally triggers an error
	router.GET("/error", func(c *gin.Context) {
		// Capture an error manually
		sentry.CaptureMessage("Something went wrong on /error endpoint!")
		c.JSON(500, gin.H{"error": "Internal Server Error"})
	})

	// Route to confirm API connectivity
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to Tabi Core API!"})
	})

	router.Run(":8080") // Start server on port 8080
}
