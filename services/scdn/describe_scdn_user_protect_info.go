package scdn

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

// DescribeScdnUserProtectInfo invokes the scdn.DescribeScdnUserProtectInfo API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdnuserprotectinfo.html
func (client *Client) DescribeScdnUserProtectInfo(request *DescribeScdnUserProtectInfoRequest) (response *DescribeScdnUserProtectInfoResponse, err error) {
	response = CreateDescribeScdnUserProtectInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnUserProtectInfoWithChan invokes the scdn.DescribeScdnUserProtectInfo API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdnuserprotectinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnUserProtectInfoWithChan(request *DescribeScdnUserProtectInfoRequest) (<-chan *DescribeScdnUserProtectInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnUserProtectInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnUserProtectInfo(request)
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

// DescribeScdnUserProtectInfoWithCallback invokes the scdn.DescribeScdnUserProtectInfo API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdnuserprotectinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnUserProtectInfoWithCallback(request *DescribeScdnUserProtectInfoRequest, callback func(response *DescribeScdnUserProtectInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnUserProtectInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnUserProtectInfo(request)
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

// DescribeScdnUserProtectInfoRequest is the request struct for api DescribeScdnUserProtectInfo
type DescribeScdnUserProtectInfoRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnUserProtectInfoResponse is the response struct for api DescribeScdnUserProtectInfo
type DescribeScdnUserProtectInfoResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	ServiceDDoS int    `json:"ServiceDDoS" xml:"ServiceDDoS"`
}

// CreateDescribeScdnUserProtectInfoRequest creates a request to invoke DescribeScdnUserProtectInfo API
func CreateDescribeScdnUserProtectInfoRequest() (request *DescribeScdnUserProtectInfoRequest) {
	request = &DescribeScdnUserProtectInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnUserProtectInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeScdnUserProtectInfoResponse creates a response to parse from DescribeScdnUserProtectInfo response
func CreateDescribeScdnUserProtectInfoResponse() (response *DescribeScdnUserProtectInfoResponse) {
	response = &DescribeScdnUserProtectInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
