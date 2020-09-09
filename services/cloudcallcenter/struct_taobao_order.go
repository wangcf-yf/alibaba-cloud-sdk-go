package cloudcallcenter

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

// TaobaoOrder is a nested struct in cloudcallcenter response
type TaobaoOrder struct {
	OutcomingAccount         int     `json:"outcomingAccount" xml:"outcomingAccount"`
	FactMoney                float64 `json:"factMoney" xml:"factMoney"`
	ProdFee                  float64 `json:"prodFee" xml:"prodFee"`
	ArticleItemCode          string  `json:"articleItemCode" xml:"articleItemCode"`
	OrderId                  int64   `json:"OrderId" xml:"OrderId"`
	StartDate                int64   `json:"startDate" xml:"startDate"`
	ExpiresIn                int64   `json:"expiresIn" xml:"expiresIn"`
	Id                       int64   `json:"id" xml:"id"`
	LastCalculateTime        int64   `json:"lastCalculateTime" xml:"lastCalculateTime"`
	Status                   int     `json:"status" xml:"status"`
	ConsumedIncomingAccount  int     `json:"consumedIncomingAccount" xml:"consumedIncomingAccount"`
	OrderRecordId            int64   `json:"orderRecordId" xml:"orderRecordId"`
	ConsumedOutcomingAccount int     `json:"consumedOutcomingAccount" xml:"consumedOutcomingAccount"`
	IncomingAccount          int     `json:"incomingAccount" xml:"incomingAccount"`
	PlanId                   int64   `json:"planId" xml:"planId"`
	TaobaoNick               string  `json:"taobaoNick" xml:"taobaoNick"`
	ConfirmedAccount         int     `json:"confirmedAccount" xml:"confirmedAccount"`
	TaobaoUid                int64   `json:"TaobaoUid" xml:"TaobaoUid"`
	PayDate                  int64   `json:"payDate" xml:"payDate"`
	ArticleCode              string  `json:"articleCode" xml:"articleCode"`
	Type                     int     `json:"type" xml:"type"`
	ParentOrderId            int64   `json:"ParentOrderId" xml:"ParentOrderId"`
}