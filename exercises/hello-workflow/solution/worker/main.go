package main

import (
	"log"
	hello "temporal101/exercises/hello-workflow/solution"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	serverAddress := "localhost:7233"

	c, err := client.Dial(client.Options{
		HostPort: serverAddress,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "greeting-tasks", worker.Options{})

	w.RegisterWorkflow(hello.GreetSomeone)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
