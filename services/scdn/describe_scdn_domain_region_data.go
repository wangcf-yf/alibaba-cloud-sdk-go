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

// DescribeScdnDomainRegionData invokes the scdn.DescribeScdnDomainRegionData API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainregiondata.html
func (client *Client) DescribeScdnDomainRegionData(request *DescribeScdnDomainRegionDataRequest) (response *DescribeScdnDomainRegionDataResponse, err error) {
	response = CreateDescribeScdnDomainRegionDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainRegionDataWithChan invokes the scdn.DescribeScdnDomainRegionData API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainregiondata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainRegionDataWithChan(request *DescribeScdnDomainRegionDataRequest) (<-chan *DescribeScdnDomainRegionDataResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainRegionDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainRegionData(request)
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

// DescribeScdnDomainRegionDataWithCallback invokes the scdn.DescribeScdnDomainRegionData API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainregiondata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainRegionDataWithCallback(request *DescribeScdnDomainRegionDataRequest, callback func(response *DescribeScdnDomainRegionDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainRegionDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainRegionData(request)
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

// DescribeScdnDomainRegionDataRequest is the request struct for api DescribeScdnDomainRegionData
type DescribeScdnDomainRegionDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnDomainRegionDataResponse is the response struct for api DescribeScdnDomainRegionData
type DescribeScdnDomainRegionDataResponse struct {
	*responses.BaseResponse
	RequestId    string                              `json:"RequestId" xml:"RequestId"`
	DomainName   string                              `json:"DomainName" xml:"DomainName"`
	DataInterval string                              `json:"DataInterval" xml:"DataInterval"`
	StartTime    string                              `json:"StartTime" xml:"StartTime"`
	EndTime      string                              `json:"EndTime" xml:"EndTime"`
	Value        ValueInDescribeScdnDomainRegionData `json:"Value" xml:"Value"`
}

// CreateDescribeScdnDomainRegionDataRequest creates a request to invoke DescribeScdnDomainRegionData API
func CreateDescribeScdnDomainRegionDataRequest() (request *DescribeScdnDomainRegionDataRequest) {
	request = &DescribeScdnDomainRegionDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainRegionData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnDomainRegionDataResponse creates a response to parse from DescribeScdnDomainRegionData response
func CreateDescribeScdnDomainRegionDataResponse() (response *DescribeScdnDomainRegionDataResponse) {
	response = &DescribeScdnDomainRegionDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
