package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

var argDebug = flag.Bool("debug", false, "run in debug mode")
var argLogJSON = flag.Bool("logjson", false, "output log in JSON format")

// Info is the info logger
var Info *log.Logger

// Error is the error logger
var Error *log.Logger

const logo = `

___.   .__  __         .__                       __                          .__                         
\_ |__ |__|/  |_  ____ |  |   ____  ______ _____/  |_          _____  ______ |__|           ____   ____  
 | __ \|  \   __\/ ___\|  |  /  _ \/  ___// __ \   __\  ______ \__  \ \____ \|  |  ______  / ___\ /  _ \ 
 | \_\ \  ||  | \  \___|  |_(  <_> )___ \\  ___/|  |   /_____/  / __ \|  |_> >  | /_____/ / /_/  >  <_> )
 |___  /__||__|  \___  >____/\____/____  >\___  >__|           (____  /   __/|__|         \___  / \____/ 
     \/              \/                \/     \/                    \/|__|               /_____/         


`

// BitClosetIndex returns the index for the bitcloset resource
func BitClosetIndex(c *gin.Context) {
	c.String(200, "hello")

}

func main() {
	flag.Parse()

	// Print logo.
	fmt.Print(logo)

	// Initialize the Loggers.  First Info Logger
	Info = log.New()
	Info.Out = os.Stdout

	// Initialize Error Logger
	Error = log.New()
	Error.Out = os.Stderr

	if *argDebug {
		log.SetLevel(log.DebugLevel)
	}

	if *argLogJSON {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	}

	router := gin.Default()
	v1 := router.Group("/api/v1/bitcloset")
	{
		v1.GET("/", BitClosetIndex)
	}

	router.Run()

}
