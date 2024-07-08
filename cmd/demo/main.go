package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/charmingruby/tynamo/pkg/aws"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := aws.NewAWSConfig()
	if err != nil {
		slog.Error(fmt.Sprintf("[AWS CONFIG] %s", err.Error()))
	}

	_ = aws.NewDynamoConnection(cfg)

}
