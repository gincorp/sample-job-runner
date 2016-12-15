package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/gincorp/gin/node"
	"github.com/gincorp/gin/taskmanager"
)

var (
	amqpURI *string
)

func init() {
	amqpURI = flag.String("amqp", "amqp://guest:guest@localhost:5671/", "URI to pass messages via")

	flag.Parse()

}

func main() {
	log.Printf("Using %q", *amqpURI)

	n := node.NewNode(*amqpURI, "", "job")

	jobManager := taskmanager.NewJobManager()
	jobManager.AddJob("write-to-file", dumpToFile)

	n.TaskManager = jobManager

	log.Fatal(n.ConsumerLoop())
}

func dumpToFile(jn taskmanager.JobNotification) (output map[string]interface{}, err error) {
	body := []byte(jn.Context["body"])
	path := jn.Context["path"]

	err = ioutil.WriteFile(path, body, 0644)
	return
}
