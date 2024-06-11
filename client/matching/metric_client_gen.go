// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by cmd/tools/rpcwrappers. DO NOT EDIT.

package matching

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"google.golang.org/grpc"
)

func (c *metricClient) ApplyTaskQueueUserDataReplicationEvent(
	ctx context.Context,
	request *matchingservice.ApplyTaskQueueUserDataReplicationEventRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.ApplyTaskQueueUserDataReplicationEventResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientApplyTaskQueueUserDataReplicationEvent")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.ApplyTaskQueueUserDataReplicationEvent(ctx, request, opts...)
}

func (c *metricClient) CancelOutstandingPoll(
	ctx context.Context,
	request *matchingservice.CancelOutstandingPollRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.CancelOutstandingPollResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientCancelOutstandingPoll")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.CancelOutstandingPoll(ctx, request, opts...)
}

func (c *metricClient) CreateNexusEndpoint(
	ctx context.Context,
	request *matchingservice.CreateNexusEndpointRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.CreateNexusEndpointResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientCreateNexusEndpoint")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.CreateNexusEndpoint(ctx, request, opts...)
}

func (c *metricClient) DeleteNexusEndpoint(
	ctx context.Context,
	request *matchingservice.DeleteNexusEndpointRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.DeleteNexusEndpointResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientDeleteNexusEndpoint")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.DeleteNexusEndpoint(ctx, request, opts...)
}

func (c *metricClient) DescribeTaskQueue(
	ctx context.Context,
	request *matchingservice.DescribeTaskQueueRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.DescribeTaskQueueResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientDescribeTaskQueue")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.DescribeTaskQueue(ctx, request, opts...)
}

func (c *metricClient) DescribeTaskQueuePartition(
	ctx context.Context,
	request *matchingservice.DescribeTaskQueuePartitionRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.DescribeTaskQueuePartitionResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientDescribeTaskQueuePartition")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.DescribeTaskQueuePartition(ctx, request, opts...)
}

func (c *metricClient) DispatchNexusTask(
	ctx context.Context,
	request *matchingservice.DispatchNexusTaskRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.DispatchNexusTaskResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientDispatchNexusTask")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.DispatchNexusTask(ctx, request, opts...)
}

func (c *metricClient) ForceUnloadTaskQueue(
	ctx context.Context,
	request *matchingservice.ForceUnloadTaskQueueRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.ForceUnloadTaskQueueResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientForceUnloadTaskQueue")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.ForceUnloadTaskQueue(ctx, request, opts...)
}

func (c *metricClient) GetBuildIdTaskQueueMapping(
	ctx context.Context,
	request *matchingservice.GetBuildIdTaskQueueMappingRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.GetBuildIdTaskQueueMappingResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientGetBuildIdTaskQueueMapping")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.GetBuildIdTaskQueueMapping(ctx, request, opts...)
}

func (c *metricClient) GetTaskQueueUserData(
	ctx context.Context,
	request *matchingservice.GetTaskQueueUserDataRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.GetTaskQueueUserDataResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientGetTaskQueueUserData")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.GetTaskQueueUserData(ctx, request, opts...)
}

func (c *metricClient) GetWorkerBuildIdCompatibility(
	ctx context.Context,
	request *matchingservice.GetWorkerBuildIdCompatibilityRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.GetWorkerBuildIdCompatibilityResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientGetWorkerBuildIdCompatibility")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.GetWorkerBuildIdCompatibility(ctx, request, opts...)
}

func (c *metricClient) GetWorkerVersioningRules(
	ctx context.Context,
	request *matchingservice.GetWorkerVersioningRulesRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.GetWorkerVersioningRulesResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientGetWorkerVersioningRules")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.GetWorkerVersioningRules(ctx, request, opts...)
}

func (c *metricClient) ListNexusEndpoints(
	ctx context.Context,
	request *matchingservice.ListNexusEndpointsRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.ListNexusEndpointsResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientListNexusEndpoints")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.ListNexusEndpoints(ctx, request, opts...)
}

func (c *metricClient) ListTaskQueuePartitions(
	ctx context.Context,
	request *matchingservice.ListTaskQueuePartitionsRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.ListTaskQueuePartitionsResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientListTaskQueuePartitions")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.ListTaskQueuePartitions(ctx, request, opts...)
}

func (c *metricClient) PollNexusTaskQueue(
	ctx context.Context,
	request *matchingservice.PollNexusTaskQueueRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.PollNexusTaskQueueResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientPollNexusTaskQueue")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.PollNexusTaskQueue(ctx, request, opts...)
}

func (c *metricClient) ReplicateTaskQueueUserData(
	ctx context.Context,
	request *matchingservice.ReplicateTaskQueueUserDataRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.ReplicateTaskQueueUserDataResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientReplicateTaskQueueUserData")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.ReplicateTaskQueueUserData(ctx, request, opts...)
}

func (c *metricClient) RespondNexusTaskCompleted(
	ctx context.Context,
	request *matchingservice.RespondNexusTaskCompletedRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.RespondNexusTaskCompletedResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientRespondNexusTaskCompleted")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.RespondNexusTaskCompleted(ctx, request, opts...)
}

func (c *metricClient) RespondNexusTaskFailed(
	ctx context.Context,
	request *matchingservice.RespondNexusTaskFailedRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.RespondNexusTaskFailedResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientRespondNexusTaskFailed")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.RespondNexusTaskFailed(ctx, request, opts...)
}

func (c *metricClient) RespondQueryTaskCompleted(
	ctx context.Context,
	request *matchingservice.RespondQueryTaskCompletedRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.RespondQueryTaskCompletedResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientRespondQueryTaskCompleted")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.RespondQueryTaskCompleted(ctx, request, opts...)
}

func (c *metricClient) UpdateNexusEndpoint(
	ctx context.Context,
	request *matchingservice.UpdateNexusEndpointRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.UpdateNexusEndpointResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientUpdateNexusEndpoint")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.UpdateNexusEndpoint(ctx, request, opts...)
}

func (c *metricClient) UpdateTaskQueueUserData(
	ctx context.Context,
	request *matchingservice.UpdateTaskQueueUserDataRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.UpdateTaskQueueUserDataResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientUpdateTaskQueueUserData")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.UpdateTaskQueueUserData(ctx, request, opts...)
}

func (c *metricClient) UpdateWorkerBuildIdCompatibility(
	ctx context.Context,
	request *matchingservice.UpdateWorkerBuildIdCompatibilityRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.UpdateWorkerBuildIdCompatibilityResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientUpdateWorkerBuildIdCompatibility")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.UpdateWorkerBuildIdCompatibility(ctx, request, opts...)
}

func (c *metricClient) UpdateWorkerVersioningRules(
	ctx context.Context,
	request *matchingservice.UpdateWorkerVersioningRulesRequest,
	opts ...grpc.CallOption,
) (_ *matchingservice.UpdateWorkerVersioningRulesResponse, retError error) {

	metricsHandler, startTime := c.startMetricsRecording(ctx, "MatchingClientUpdateWorkerVersioningRules")
	defer func() {
		c.finishMetricsRecording(metricsHandler, startTime, retError)
	}()

	return c.client.UpdateWorkerVersioningRules(ctx, request, opts...)
}
