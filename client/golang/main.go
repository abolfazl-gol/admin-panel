package main

import (
	"adminpanel/proto"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewApiClient(conn)

	ctx := context.Background()
	// question, err := client.ListQuestion(ctx, &proto.ListQuestionRequest{})
	// // question, err := client.CreateQuestion(ctx, &proto.CreateQuestionRequest{Text: "heeell", Enabled: true, TopicId: 1})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(question)
	// For update
	topic, err := client.CreateTopic(ctx, &proto.CreateTopicRequest{Name: "", Enabled: true})
	// topic, err := client.UpdateTopic(ctx, &proto.UpdateTopicRequest{Topic: &proto.Topic{Id: 6, Name: "", Enabled: false}, UpdateMask: []string{"name"}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(topic)
	// _, err = client.Login(ctx, &proto.LoginRequest{Email: "b@gmail.ir", Password: "22adsfsdf"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
