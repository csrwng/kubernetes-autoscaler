// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

//go:build go1.16 && integration
// +build go1.16,integration

package health_test

import (
	"context"
	"testing"
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/awserr"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/awstesting/integration"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/health"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_DescribeEntityAggregates(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-east-1")
	svc := health.New(sess)
	params := &health.DescribeEntityAggregatesInput{}
	_, err := svc.DescribeEntityAggregatesWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
