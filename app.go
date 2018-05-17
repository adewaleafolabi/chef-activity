package main

import (
	"fmt"
	"github.com/adewaleafolabi/chef-hero-activity/proto"
	"github.com/adewaleafolabi/chef-hero-activity/repository"
	"github.com/bitly/go-nsq"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net"
	"github.com/tkanos/gonfig"
	appConfig "github.com/adewaleafolabi/chef-hero-activity/config"
)

func main() {

	configuration := appConfig.Configuration{}
	//TODO: extract config file to environment variable
	err := gonfig.GetConf("config/config.json", &configuration)
	if err != nil {
		log.Fatal(err)
	}

	db, err := appConfig.CreateDBConnection(configuration)
	db.SingularTable(true)
	db.LogMode(true)
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	activityRepo := &repository.ActivityRepository{DB: db}

	lis, err := net.Listen("tcp", configuration.GRPCPort)
	if err != nil {
		log.Fatalf("failed to created server on port : %v %v",  configuration.GRPCPort, err)
	}

	config := nsq.NewConfig()
	w, _ := nsq.NewProducer(configuration.NSQAddress, config)

	s := grpc.NewServer()
	chefhero.RegisterActivityServiceServer(s, &service{repo: activityRepo, validator: validator.New(), producer: w})
	fmt.Printf("svc listening on %s\n",  configuration.GRPCPort)

	s.Serve(lis)

}
