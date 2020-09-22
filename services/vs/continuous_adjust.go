package vs

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

// ContinuousAdjust invokes the vs.ContinuousAdjust API synchronously
func (client *Client) ContinuousAdjust(request *ContinuousAdjustRequest) (response *ContinuousAdjustResponse, err error) {
	response = CreateContinuousAdjustResponse()
	err = client.DoAction(request, response)
	return
}

// ContinuousAdjustWithChan invokes the vs.ContinuousAdjust API asynchronously
func (client *Client) ContinuousAdjustWithChan(request *ContinuousAdjustRequest) (<-chan *ContinuousAdjustResponse, <-chan error) {
	responseChan := make(chan *ContinuousAdjustResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ContinuousAdjust(request)
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

// ContinuousAdjustWithCallback invokes the vs.ContinuousAdjust API asynchronously
func (client *Client) ContinuousAdjustWithCallback(request *ContinuousAdjustRequest, callback func(response *ContinuousAdjustResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ContinuousAdjustResponse
		var err error
		defer close(result)
		response, err = client.ContinuousAdjust(request)
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

// ContinuousAdjustRequest is the request struct for api ContinuousAdjust
type ContinuousAdjustRequest struct {
	*requests.RpcRequest
	Focus   string           `position:"Query" name:"Focus"`
	Id      string           `position:"Query" name:"Id"`
	ShowLog string           `position:"Query" name:"ShowLog"`
	Iris    string           `position:"Query" name:"Iris"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// ContinuousAdjustResponse is the response struct for api ContinuousAdjust
type ContinuousAdjustResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Id        string `json:"Id" xml:"Id"`
}

// CreateContinuousAdjustRequest creates a request to invoke ContinuousAdjust API
func CreateContinuousAdjustRequest() (request *ContinuousAdjustRequest) {
	request = &ContinuousAdjustRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "ContinuousAdjust", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateContinuousAdjustResponse creates a response to parse from ContinuousAdjust response
func CreateContinuousAdjustResponse() (response *ContinuousAdjustResponse) {
	response = &ContinuousAdjustResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}