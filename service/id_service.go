package service

import "github.com/rs/xid"

type IDGenerator interface {
	GenerateID() string
}

type IDGeneratorService struct {
}

func (ids *IDGeneratorService) GenerateID() string {
	return xid.New().String()
}
