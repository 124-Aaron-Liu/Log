package main

import (
	"os"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")
    //{"animal":"walrus","level":"info","msg":"A group of walrus emerges from the ocean","size":10,"time":"2020-09-22T17:21:21+08:00"}
//{"level":"warning","msg":"The group's number increased tremendously!","number":122,"omg":true,"time":"2020-09-22T17:21:21+08:00"}
	
}