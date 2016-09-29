package goConsul

import (
	"io/ioutil"
	"encoding/json"

)

/*
  {
    "Node": "HungSYX",
    "Address": "172.20.2.171",
    "ServiceID": "HungSYX:filerepo0:8080",
    "ServiceName": "filerepo",
    "ServiceTags": [
      "MCD_REPO"
    ],
    "ServiceAddress": "",
    "ServicePort": 32768,
    "ServiceEnableTagOverride": false,
    "CreateIndex": 49181,
    "ModifyIndex": 53802
  }
*/

type Services struct {
	Node      string `json:"Node"`
	ServiceID string `json:"ServiceID"`
	ServiceTags    [] string `json:"Tags"`
	Address string `json:"Address"`
	ServicePort    int `json:"ServicePort"`
}
type ConsulService struct {
	Node      string
	ServiceID string
	ServiceTags    [] string
	Address string
	ServicePort    int
}/*
type ConsulService struct {
	Node      string `json:"Node"`
	ServiceID string `json:"ServiceID"`
	ServiceTags    [] string `json:"Tags"`
	Address string `json:"Address"`
	ServicePort    int `json:"ServicePort"`
}*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func ReadConfig(file string) []Services {
	var jsonData []Services

	configdata, err := ioutil.ReadFile(file)
	check(err)
	//fmt.Print(configdata)
	json.Unmarshal(configdata, &jsonData)
	//fmt.Print(jsonData)
	return jsonData

}

func GetConfile(filepath string) []Services{
	s := ReadConfig(filepath)
	return s

}
