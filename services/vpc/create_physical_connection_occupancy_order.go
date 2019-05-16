package vpc

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

// CreatePhysicalConnectionOccupancyOrder invokes the vpc.CreatePhysicalConnectionOccupancyOrder API synchronously
// api document: https://help.aliyun.com/api/vpc/createphysicalconnectionoccupancyorder.html
func (client *Client) CreatePhysicalConnectionOccupancyOrder(request *CreatePhysicalConnectionOccupancyOrderRequest) (response *CreatePhysicalConnectionOccupancyOrderResponse, err error) {
	response = CreateCreatePhysicalConnectionOccupancyOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePhysicalConnectionOccupancyOrderWithChan invokes the vpc.CreatePhysicalConnectionOccupancyOrder API asynchronously
// api document: https://help.aliyun.com/api/vpc/createphysicalconnectionoccupancyorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePhysicalConnectionOccupancyOrderWithChan(request *CreatePhysicalConnectionOccupancyOrderRequest) (<-chan *CreatePhysicalConnectionOccupancyOrderResponse, <-chan error) {
	responseChan := make(chan *CreatePhysicalConnectionOccupancyOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePhysicalConnectionOccupancyOrder(request)
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

// CreatePhysicalConnectionOccupancyOrderWithCallback invokes the vpc.CreatePhysicalConnectionOccupancyOrder API asynchronously
// api document: https://help.aliyun.com/api/vpc/createphysicalconnectionoccupancyorder.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePhysicalConnectionOccupancyOrderWithCallback(request *CreatePhysicalConnectionOccupancyOrderRequest, callback func(response *CreatePhysicalConnectionOccupancyOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePhysicalConnectionOccupancyOrderResponse
		var err error
		defer close(result)
		response, err = client.CreatePhysicalConnectionOccupancyOrder(request)
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

// CreatePhysicalConnectionOccupancyOrderRequest is the request struct for api CreatePhysicalConnectionOccupancyOrder
type CreatePhysicalConnectionOccupancyOrderRequest struct {
	*requests.RpcRequest
	Period               requests.Integer `position:"Query" name:"Period"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	PhysicalConnectionId string           `position:"Query" name:"PhysicalConnectionId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceChargeType   string           `position:"Query" name:"InstanceChargeType"`
	PricingCycle         string           `position:"Query" name:"PricingCycle"`
}

// CreatePhysicalConnectionOccupancyOrderResponse is the response struct for api CreatePhysicalConnectionOccupancyOrder
type CreatePhysicalConnectionOccupancyOrderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreatePhysicalConnectionOccupancyOrderRequest creates a request to invoke CreatePhysicalConnectionOccupancyOrder API
func CreateCreatePhysicalConnectionOccupancyOrderRequest() (request *CreatePhysicalConnectionOccupancyOrderRequest) {
	request = &CreatePhysicalConnectionOccupancyOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreatePhysicalConnectionOccupancyOrder", "vpc", "openAPI")
	return
}

// CreateCreatePhysicalConnectionOccupancyOrderResponse creates a response to parse from CreatePhysicalConnectionOccupancyOrder response
func CreateCreatePhysicalConnectionOccupancyOrderResponse() (response *CreatePhysicalConnectionOccupancyOrderResponse) {
	response = &CreatePhysicalConnectionOccupancyOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}