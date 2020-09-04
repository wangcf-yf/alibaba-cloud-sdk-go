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

// DescribeScdnDomainQpsData invokes the scdn.DescribeScdnDomainQpsData API synchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainqpsdata.html
func (client *Client) DescribeScdnDomainQpsData(request *DescribeScdnDomainQpsDataRequest) (response *DescribeScdnDomainQpsDataResponse, err error) {
	response = CreateDescribeScdnDomainQpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainQpsDataWithChan invokes the scdn.DescribeScdnDomainQpsData API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainqpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainQpsDataWithChan(request *DescribeScdnDomainQpsDataRequest) (<-chan *DescribeScdnDomainQpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainQpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainQpsData(request)
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

// DescribeScdnDomainQpsDataWithCallback invokes the scdn.DescribeScdnDomainQpsData API asynchronously
// api document: https://help.aliyun.com/api/scdn/describescdndomainqpsdata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScdnDomainQpsDataWithCallback(request *DescribeScdnDomainQpsDataRequest, callback func(response *DescribeScdnDomainQpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainQpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainQpsData(request)
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

// DescribeScdnDomainQpsDataRequest is the request struct for api DescribeScdnDomainQpsData
type DescribeScdnDomainQpsDataRequest struct {
	*requests.RpcRequest
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	StartTime      string           `position:"Query" name:"StartTime"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	DomainName     string           `position:"Query" name:"DomainName"`
	EndTime        string           `position:"Query" name:"EndTime"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	Interval       string           `position:"Query" name:"Interval"`
}

// DescribeScdnDomainQpsDataResponse is the response struct for api DescribeScdnDomainQpsData
type DescribeScdnDomainQpsDataResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	DomainName         string             `json:"DomainName" xml:"DomainName"`
	StartTime          string             `json:"StartTime" xml:"StartTime"`
	EndTime            string             `json:"EndTime" xml:"EndTime"`
	DataInterval       string             `json:"DataInterval" xml:"DataInterval"`
	QpsDataPerInterval QpsDataPerInterval `json:"QpsDataPerInterval" xml:"QpsDataPerInterval"`
}

// CreateDescribeScdnDomainQpsDataRequest creates a request to invoke DescribeScdnDomainQpsData API
func CreateDescribeScdnDomainQpsDataRequest() (request *DescribeScdnDomainQpsDataRequest) {
	request = &DescribeScdnDomainQpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainQpsData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnDomainQpsDataResponse creates a response to parse from DescribeScdnDomainQpsData response
func CreateDescribeScdnDomainQpsDataResponse() (response *DescribeScdnDomainQpsDataResponse) {
	response = &DescribeScdnDomainQpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
