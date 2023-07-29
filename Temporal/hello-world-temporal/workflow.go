package app

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
	"go.temporal.io/server/service/history/workflow"
)

//GreetingWorkflow is a public function
func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		startClosedTimeout: time.Second * 5,
	}

	


}