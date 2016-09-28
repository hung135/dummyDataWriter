package goConsul

/*
HungSYX:elastic:9200: {
ID: "HungSYX:elastic:9200",
Service: "elasticsearch-9200",
Tags: null,
Address: "",
Port: 9200,
EnableTagOverride: false,
CreateIndex: 0,
ModifyIndex: 0
},
*/

type Services struct {
	services [] ConsulService

}
type ConsulService struct {
	ID string `json:"ID"`
	Service string `json:"Service"`
	Tags [] string `json:"Tags"`
	Address string `json:"Address"`
	Port int `json:"Port"`


}
func getConfile() {
}
