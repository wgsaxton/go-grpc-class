package main

import (
	"context"
	"log"

	"github.com/wgsaxton/go-grpc-class/module2-exercise/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := proto.NewTodoServiceClient(conn)

	task1, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "wake up"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Task 1 created: %s", task1.GetId())

	task2, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "walk the dog"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Task 2 created: %s", task2.GetId())

	tasks, err := client.ListTasks(ctx, &proto.ListTasksRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("existing tasks: %s", tasks)

	_, err = client.CompleteTask(ctx, &proto.CompleteTaskRequest{Id: task1.GetId()})
	if err != nil {
		log.Fatal(err)
	}

	task3, err := client.AddTask(ctx, &proto.AddTaskRequest{Task: "eat breakfast"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Task 3 created: %s", task3.GetId())

	tasks, err = client.ListTasks(ctx, &proto.ListTasksRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("existing tasks: %s", tasks.GetTasks())


}
