package main

import (
	"context"
	"fmt"
	"github.com/adewaleafolabi/chef-hero-activity/model"
	"github.com/adewaleafolabi/chef-hero-activity/proto"
	"github.com/adewaleafolabi/chef-hero-activity/repository"
	"github.com/bitly/go-nsq"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type service struct {
	repo      repository.Repository
	validator *validator.Validate
	producer  *nsq.Producer
}

func (srv *service) CreateActivity(ctx context.Context, in *chefhero.ActivityRequest) (*chefhero.ActivityResponse, error) {
	res := &chefhero.ActivityResponse{Success: false, Id: ""}
	activity := convertToActivity(in)
	err := srv.validator.Struct(activity)

	if err != nil {
		return res, err
	}

	data, _ := msgpack.Marshal(activity)
	_ = srv.producer.Publish("activity", data)

	res.Success = true
	return res, nil
}

func (srv *service) FindActivity(in *chefhero.ActivityRequest, stream chefhero.ActivityService_FindActivityServer) error {

	filter := convertToActivity(in)
	activities, err := srv.repo.FindActivity(filter)

	if err != nil {
		return err
	}
	for _, a := range activities {
		if err := stream.Send(convertToActivityRequest(a)); err != nil {
			return err
		}
	}
	return nil
}

func convertToActivityRequest(in *model.Activity) *chefhero.ActivityRequest {
	return &chefhero.ActivityRequest{
		Environment: in.Environment,
		Email:       in.Email,
		Message:     in.Message,
		Data:        in.Data,
		Component:   in.Component,
		Id:          in.Id,
		CreatedAt:   in.CreatedAt.Format(time.RFC3339),
	}

}
func convertToActivity(in *chefhero.ActivityRequest) *model.Activity {

	output := model.Activity{
		Environment: in.Environment,
		Email:       in.Email,
		Message:     in.Message,
		Data:        in.Data,
		Component:   in.Component,
		Id:          in.Id,
	}
	if len(in.CreatedAt) > 9 {
		time, err := time.Parse(time.RFC3339, fmt.Sprintf("%s%s", in.CreatedAt[0:10], "T00:00:00Z"))
		if err == nil {
			output.CreatedAt = time
		}
	}
	return &output

}
