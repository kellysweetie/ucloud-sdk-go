//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMon DescribeResourceMetric

package umon

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeResourceMetricRequest is request schema for DescribeResourceMetric action
type DescribeResourceMetricRequest struct {
	request.CommonBase

	// 资源类型
	ResourceType *string `required:"true"`
}

// DescribeResourceMetricResponse is response schema for DescribeResourceMetric action
type DescribeResourceMetricResponse struct {
	response.CommonBase

	// 指标信息集合
	DataSet []MetricInfo
}

// NewDescribeResourceMetricRequest will create request of DescribeResourceMetric action.
func (c *UMonClient) NewDescribeResourceMetricRequest() *DescribeResourceMetricRequest {
	req := &DescribeResourceMetricRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeResourceMetric - 获取资源支持监控指标信息
func (c *UMonClient) DescribeResourceMetric(req *DescribeResourceMetricRequest) (*DescribeResourceMetricResponse, error) {
	var err error
	var res DescribeResourceMetricResponse

	err = c.client.InvokeAction("DescribeResourceMetric", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}