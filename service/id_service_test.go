package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	gophers = 100
	entries = 100000
)

var idSvc = IDGeneratorService{}
var ids map[string]string = make(map[string]string)
var c chan string = make(chan string)

func TestIDGeneratorService_GenerateIDUnique(t *testing.T) {
	for i := 0; i < gophers; i++ {
		go func() {
			for j := 0; j < entries; j++ {
				c <- idSvc.GenerateID()
			}
		}()
	}

	for i := 0; i < gophers; i++ {
		for j := 0; j < entries; j++ {
			x := <-c
			ids[x] = ""
		}
	}

	assert.Equal(t, gophers*entries, len(ids))
}
