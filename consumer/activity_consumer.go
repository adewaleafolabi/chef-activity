package main

import (
	"github.com/bitly/go-nsq"
	"github.com/vmihailenco/msgpack"
	"log"
	"sync"
	appConfig "github.com/adewaleafolabi/chef-hero-activity/config"
)
import (
	"github.com/adewaleafolabi/chef-hero-activity/model"
	"github.com/tkanos/gonfig"
	"github.com/adewaleafolabi/chef-hero-activity/repository"
)

const topic = "activity"

func main() {
	configuration := appConfig.Configuration{}
	//TODO: extract config file to environment variable
	err := gonfig.GetConf("../config/config.json", &configuration)
	if err != nil {
		log.Fatal(err)
	}

	db,err:= appConfig.CreateDBConnection(configuration)
	db.SingularTable(true)
	db.LogMode(true)

	if err != nil {
		log.Fatal(err)
	}

	activityRepo:=repository.ActivityRepository{DB:db}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	config := nsq.NewConfig()
	q, _ := nsq.NewConsumer(topic, "ch", config)
	q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var activity model.Activity
		msgpack.Unmarshal(message.Body, &activity)
		log.Printf("Got a message: %v", activity)
		id,err:=activityRepo.SaveActivity(&activity)
		if err!=nil{
			return err
		}
		log.Printf("activity processed successfully. id:%s",id)
		return nil
	}))
	err = q.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()

}
