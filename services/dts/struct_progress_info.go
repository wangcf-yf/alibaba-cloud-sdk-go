package dts

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

// ProgressInfo is a nested struct in dts response
type ProgressInfo struct {
	BootTime     string   `json:"BootTime" xml:"BootTime"`
	CanSkip      bool     `json:"CanSkip" xml:"CanSkip"`
	DelaySeconds int      `json:"DelaySeconds" xml:"DelaySeconds"`
	FinishTime   string   `json:"FinishTime" xml:"FinishTime"`
	IgnoreFlag   string   `json:"IgnoreFlag" xml:"IgnoreFlag"`
	Item         string   `json:"Item" xml:"Item"`
	JobId        string   `json:"JobId" xml:"JobId"`
	Names        string   `json:"Names" xml:"Names"`
	OrderNum     int      `json:"OrderNum" xml:"OrderNum"`
	Skip         bool     `json:"Skip" xml:"Skip"`
	State        string   `json:"State" xml:"State"`
	Sub          string   `json:"Sub" xml:"Sub"`
	RepairMethod string   `json:"RepairMethod" xml:"RepairMethod"`
	TargetNames  string   `json:"TargetNames" xml:"TargetNames"`
	Total        int      `json:"Total" xml:"Total"`
	SourceSchema string   `json:"SourceSchema" xml:"SourceSchema"`
	ParentObj    string   `json:"ParentObj" xml:"ParentObj"`
	DiffRow      int64    `json:"DiffRow" xml:"DiffRow"`
	DestSchema   string   `json:"DestSchema" xml:"DestSchema"`
	ErrDetail    string   `json:"ErrDetail" xml:"ErrDetail"`
	ErrMsg       string   `json:"ErrMsg" xml:"ErrMsg"`
	DdlSql       string   `json:"DdlSql" xml:"DdlSql"`
	Logs         []JobLog `json:"Logs" xml:"Logs"`
}
