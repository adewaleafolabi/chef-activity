package main

import (
	"fmt"
	"github.com/adewaleafolabi/chef-hero-activity/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
)

const (
	address = ":5058"
)

func createActivity(client chefhero.ActivityServiceClient, request *chefhero.ActivityRequest) {
	resp, err := client.CreateActivity(context.Background(), request)

	if err != nil {
		log.Fatal(fmt.Sprintf("failed to create activity %v", err))
	}

	if resp.Success {
		log.Info(fmt.Sprintf("customer created with id %s", resp.Id))
	}

}

func getActivity(client chefhero.ActivityServiceClient, filter *chefhero.ActivityRequest) {
	stream, err := client.FindActivity(context.Background(), filter)
	if err != nil {
		log.Fatalf("failed to get activities %v", err)
	}
	for {
		activity, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Infof("%v.FindActivity(_) = _, %v", client, err)
		}
		log.Infof("Activity: %v", activity)
	}
}
func main() {
	fmt.Println("Init Client")
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed %v", err)
	}
	defer conn.Close()

	client := chefhero.NewActivityServiceClient(conn)

	activity := &chefhero.ActivityRequest{
		Component:   "ES-CS",
		Email:       "chef@chefhero.com",
		Message:     "ES init complete",
		Data:        `{"message":"ok"}`,
		Environment: "staging",
	}

	filter := &chefhero.ActivityRequest{
		Component:   "ES-CS",
		Message:     "complete",
		Environment: "staging",
		CreatedAt:   "2018-05-15",
		Id:          "",
	}

	createActivity(client, activity)
	getActivity(client, filter)
}
