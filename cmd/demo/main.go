package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/tynamo/internal/factory"
	"github.com/charmingruby/tynamo/internal/repository/dynamo"
	"github.com/charmingruby/tynamo/pkg/aws"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := aws.NewAWSConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("[AWS CONFIG] %s", err.Error()))
		os.Exit(1)
	}

	dummyExample := factory.MakeExample("dummy title")

	db := aws.NewDynamoConnection(cfg)

	exampleRepo := dynamo.NewDynamoExampleRepository(db)

	if err := exampleRepo.Store(dummyExample); err != nil {
		slog.Error(fmt.Sprintf("[DYNAMO] %s", err.Error()))
		os.Exit(1)
	}
}
