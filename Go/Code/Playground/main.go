package main

import (
	"context"
	"fmt"
)

// Services for Temporal test
type Services struct {
	abc0 string
	abc1 []interface{}
	abc2 []interface{}
	Abc4 string
}

// Config struct
type Config struct {
	abc4 string
}

// New Is the struct
func New(config Config) *Services {
	return &Services{
		Abc4: config.abc4,
		abc0: "init",
		abc1: make([]interface{}, 0),
		abc2: make([]interface{}, 0),
	}
}

// GetName return a Hello World
func (s *Services) GetName() string {
	return "Hello World"
}

func (s *Services) addAbc1(s1 string) (string, error) {
	return s1, nil

}

func (s *Services) addAbc2(s2 string) (string, error) {
	return s2, nil
}

func (s *Services) start(ctx context.Context, cancel context.CancelCauseFunc) error {

	fmt.Println("In the start function")
	return nil

}

func main() {

	var ctx context.Context
	var cancel context.CancelCauseFunc
	var s Services

	err := s.start(ctx, cancel)

	if err != nil {
		fmt.Println("There was an error")
	}

}
