// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
)

// WaitUntilAddonActive uses the Amazon EKS API operation
// DescribeAddon to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EKS) WaitUntilAddonActive(input *DescribeAddonInput) error {
	return c.WaitUntilAddonActiveWithContext(aws.BackgroundContext(), input)
}

// WaitUntilAddonActiveWithContext is an extended version of WaitUntilAddonActive.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EKS) WaitUntilAddonActiveWithContext(ctx aws.Context, input *DescribeAddonInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilAddonActive",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(10 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "addon.status",
				Expected: "CREATE_FAILED",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "addon.status",
				Expected: "DEGRADED",
			},
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "addon.status",
				Expected: "ACTIVE",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeAddonInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeAddonRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilAddonDeleted uses the Amazon EKS API operation
// DescribeAddon to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EKS) WaitUntilAddonDeleted(input *DescribeAddonInput) error {
	return c.WaitUntilAddonDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilAddonDeletedWithContext is an extended version of WaitUntilAddonDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EKS) WaitUntilAddonDeletedWithContext(ctx aws.Context, input *DescribeAddonInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilAddonDeleted",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(10 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "addon.status",
				Expected: "DELETE_FAILED",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeAddonInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeAddonRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilClusterActive uses the Amazon EKS API operation
// DescribeCluster to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EKS) WaitUntilClusterActive(input *DescribeClusterInput) error {
	return c.WaitUntilClusterActiveWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterActiveWithContext is an extended version of WaitUntilClusterActive.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EKS) WaitUntilClusterActiveWithContext(ctx aws.Context, input *DescribeClusterInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilClusterActive",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "cluster.status",
				Expected: "DELETING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "cluster.status",
				Expected: "FAILED",
			},
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "cluster.status",
				Expected: "ACTIVE",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeClusterInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeClusterRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilClusterDeleted uses the Amazon EKS API operation
// DescribeCluster to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EKS) WaitUntilClusterDeleted(input *DescribeClusterInput) error {
	return c.WaitUntilClusterDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterDeletedWithContext is an extended version of WaitUntilClusterDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EKS) WaitUntilClusterDeletedWithContext(ctx aws.Context, input *DescribeClusterInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilClusterDeleted",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "cluster.status",
				Expected: "ACTIVE",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "cluster.status",
				Expected: "CREATING",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "cluster.status",
				Expected: "PENDING",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeClusterInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeClusterRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilFargateProfileActive uses the Amazon EKS API operation
// DescribeFargateProfile to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EKS) WaitUntilFargateProfileActive(input *DescribeFargateProfileInput) error {
	return c.WaitUntilFargateProfileActiveWithContext(aws.BackgroundContext(), input)
}

// WaitUntilFargateProfileActiveWithContext is an extended version of WaitUntilFargateProfileActive.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EKS) WaitUntilFargateProfileActiveWithContext(ctx aws.Context, input *DescribeFargateProfileInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilFargateProfileActive",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(10 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "fargateProfile.status",
				Expected: "CREATE_FAILED",
			},
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "fargateProfile.status",
				Expected: "ACTIVE",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeFargateProfileInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeFargateProfileRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilFargateProfileDeleted uses the Amazon EKS API operation
// DescribeFargateProfile to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EKS) WaitUntilFargateProfileDeleted(input *DescribeFargateProfileInput) error {
	return c.WaitUntilFargateProfileDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilFargateProfileDeletedWithContext is an extended version of WaitUntilFargateProfileDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EKS) WaitUntilFargateProfileDeletedWithContext(ctx aws.Context, input *DescribeFargateProfileInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilFargateProfileDeleted",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "fargateProfile.status",
				Expected: "DELETE_FAILED",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeFargateProfileInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeFargateProfileRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilNodegroupActive uses the Amazon EKS API operation
// DescribeNodegroup to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EKS) WaitUntilNodegroupActive(input *DescribeNodegroupInput) error {
	return c.WaitUntilNodegroupActiveWithContext(aws.BackgroundContext(), input)
}

// WaitUntilNodegroupActiveWithContext is an extended version of WaitUntilNodegroupActive.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EKS) WaitUntilNodegroupActiveWithContext(ctx aws.Context, input *DescribeNodegroupInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilNodegroupActive",
		MaxAttempts: 80,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "nodegroup.status",
				Expected: "CREATE_FAILED",
			},
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "nodegroup.status",
				Expected: "ACTIVE",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeNodegroupInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeNodegroupRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilNodegroupDeleted uses the Amazon EKS API operation
// DescribeNodegroup to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *EKS) WaitUntilNodegroupDeleted(input *DescribeNodegroupInput) error {
	return c.WaitUntilNodegroupDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilNodegroupDeletedWithContext is an extended version of WaitUntilNodegroupDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EKS) WaitUntilNodegroupDeletedWithContext(ctx aws.Context, input *DescribeNodegroupInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilNodegroupDeleted",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(30 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "nodegroup.status",
				Expected: "DELETE_FAILED",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeNodegroupInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeNodegroupRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
