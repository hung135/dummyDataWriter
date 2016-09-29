package goConsul



import ("io/ioutil"
 "encoding/json"

	"fmt"
)

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
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func ReadConfig(file string) Services{
	var dbconn Services

	configdata,err := ioutil.ReadFile(file)
	check(err)
	json.Unmarshal(configdata,&dbconn)
	return dbconn

}

func GetConfile(filepath string) {
	s:=ReadConfig(filepath)
	fmt.Print(s)

}
