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

// AddGrafana invokes the arms.AddGrafana API synchronously
// api document: https://help.aliyun.com/api/arms/addgrafana.html
func (client *Client) AddGrafana(request *AddGrafanaRequest) (response *AddGrafanaResponse, err error) {
	response = CreateAddGrafanaResponse()
	err = client.DoAction(request, response)
	return
}

// AddGrafanaWithChan invokes the arms.AddGrafana API asynchronously
// api document: https://help.aliyun.com/api/arms/addgrafana.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddGrafanaWithChan(request *AddGrafanaRequest) (<-chan *AddGrafanaResponse, <-chan error) {
	responseChan := make(chan *AddGrafanaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddGrafana(request)
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

// AddGrafanaWithCallback invokes the arms.AddGrafana API asynchronously
// api document: https://help.aliyun.com/api/arms/addgrafana.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddGrafanaWithCallback(request *AddGrafanaRequest, callback func(response *AddGrafanaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddGrafanaResponse
		var err error
		defer close(result)
		response, err = client.AddGrafana(request)
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

// AddGrafanaRequest is the request struct for api AddGrafana
type AddGrafanaRequest struct {
	*requests.RpcRequest
	Integration string `position:"Query" name:"Integration"`
	ClusterId   string `position:"Query" name:"ClusterId"`
}

// AddGrafanaResponse is the response struct for api AddGrafana
type AddGrafanaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateAddGrafanaRequest creates a request to invoke AddGrafana API
func CreateAddGrafanaRequest() (request *AddGrafanaRequest) {
	request = &AddGrafanaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "AddGrafana", "", "")
	return
}

// CreateAddGrafanaResponse creates a response to parse from AddGrafana response
func CreateAddGrafanaResponse() (response *AddGrafanaResponse) {
	response = &AddGrafanaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
