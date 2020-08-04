package opensearch

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

// ResultItemInListDeployedAlgorithmModels is a nested struct in opensearch response
type ResultItemInListDeployedAlgorithmModels struct {
	Id           string       `json:"id" xml:"id"`
	AppGroupName string       `json:"appGroupName" xml:"appGroupName"`
	Scene        string       `json:"scene" xml:"scene"`
	Desc         string       `json:"desc" xml:"desc"`
	GmtCreate    string       `json:"gmtCreate" xml:"gmtCreate"`
	GmtModified  string       `json:"gmtModified" xml:"gmtModified"`
	Status       string       `json:"status" xml:"status"`
	Apps         []string     `json:"apps" xml:"apps"`
	Models       []ModelsItem `json:"models" xml:"models"`
}