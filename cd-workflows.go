package workflow

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
	"go.temporal.io/sdk/activity"
)

// Configuration Structure
type AppConfig struct {
	WorkerName     string `yaml:"worker_name"`
	ProcessingType string `yaml:"processing_type"` // "data" or "image"
}

var Config AppConfig

// ProcessingWorkflow is the main workflow that handles different types of processing
func ProcessingWorkflow(ctx workflow.Context, input string) (string, error) {
	workflow.GetLogger(ctx).Info("ProcessingWorkflow started", "input", input)

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	switch Config.ProcessingType {
	case "data":
		err := workflow.ExecuteActivity(ctx, DataProcessingActivity, input).Get(ctx, &result)
		if err != nil {
			return "", err
		}
	case "image":
		err := workflow.ExecuteActivity(ctx, ImageProcessingActivity, input).Get(ctx, &result)
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("invalid processing type: %s", Config.ProcessingType)
	}

	return fmt.Sprintf("Workflow processed with %s: %s", Config.ProcessingType, result), nil
}

// DataProcessingActivity handles data processing
func DataProcessingActivity(ctx context.Context, input string) (string, error) {
	activity.GetLogger(ctx).Info("DataProcessingActivity started", "input", input)
	return fmt.Sprintf("Data processing: %s", input), nil
}

// ImageProcessingActivity handles image processing
func ImageProcessingActivity(ctx context.Context, input string) (string, error) {
	activity.GetLogger(ctx).Info("ImageProcessingActivity started", "input", input)
	return fmt.Sprintf("Image processing: %s", input), nil
} 