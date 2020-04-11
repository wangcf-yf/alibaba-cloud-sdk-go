package arms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryDataset invokes the arms.QueryDataset API synchronously
// api document: https://help.aliyun.com/api/arms/querydataset.html
func (client *Client) QueryDataset(request *QueryDatasetRequest) (response *QueryDatasetResponse, err error) {
	response = CreateQueryDatasetResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDatasetWithChan invokes the arms.QueryDataset API asynchronously
// api document: https://help.aliyun.com/api/arms/querydataset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDatasetWithChan(request *QueryDatasetRequest) (<-chan *QueryDatasetResponse, <-chan error) {
	responseChan := make(chan *QueryDatasetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDataset(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QueryDatasetWithCallback invokes the arms.QueryDataset API asynchronously
// api document: https://help.aliyun.com/api/arms/querydataset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDatasetWithCallback(request *QueryDatasetRequest, callback func(response *QueryDatasetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDatasetResponse
		var err error
		defer close(result)
		response, err = client.QueryDataset(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// QueryDatasetRequest is the request struct for api QueryDataset
type QueryDatasetRequest struct {
	*requests.RpcRequest
	DateStr       string                      `position:"Query" name:"DateStr"`
	MinTime       requests.Integer            `position:"Query" name:"MinTime"`
	ProxyUserId   string                      `position:"Query" name:"ProxyUserId"`
	ReduceTail    requests.Boolean            `position:"Query" name:"ReduceTail"`
	MaxTime       requests.Integer            `position:"Query" name:"MaxTime"`
	OptionalDims  *[]QueryDatasetOptionalDims `position:"Query" name:"OptionalDims"  type:"Repeated"`
	Measures      *[]string                   `position:"Query" name:"Measures"  type:"Repeated"`
	IntervalInSec requests.Integer            `position:"Query" name:"IntervalInSec"`
	IsDrillDown   requests.Boolean            `position:"Query" name:"IsDrillDown"`
	HungryMode    requests.Boolean            `position:"Query" name:"HungryMode"`
	OrderByKey    string                      `position:"Query" name:"OrderByKey"`
	Limit         requests.Integer            `position:"Query" name:"Limit"`
	DatasetId     requests.Integer            `position:"Query" name:"DatasetId"`
	RequiredDims  *[]QueryDatasetRequiredDims `position:"Query" name:"RequiredDims"  type:"Repeated"`
	Dimensions    *[]QueryDatasetDimensions   `position:"Query" name:"Dimensions"  type:"Repeated"`
}

// QueryDatasetOptionalDims is a repeated param struct in QueryDatasetRequest
type QueryDatasetOptionalDims struct {
	Type  string `name:"Type"`
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// QueryDatasetRequiredDims is a repeated param struct in QueryDatasetRequest
type QueryDatasetRequiredDims struct {
	Type  string `name:"Type"`
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// QueryDatasetDimensions is a repeated param struct in QueryDatasetRequest
type QueryDatasetDimensions struct {
	Type  string `name:"Type"`
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// QueryDatasetResponse is the response struct for api QueryDataset
type QueryDatasetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateQueryDatasetRequest creates a request to invoke QueryDataset API
func CreateQueryDatasetRequest() (request *QueryDatasetRequest) {
	request = &QueryDatasetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "QueryDataset", "", "")
	return
}

// CreateQueryDatasetResponse creates a response to parse from QueryDataset response
func CreateQueryDatasetResponse() (response *QueryDatasetResponse) {
	response = &QueryDatasetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
