// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	//"io/ioutil"
	"os"
	"io/ioutil"
	"runtime"
	"makedummyfile"
	"msgqueue"
)






// Clears the bit at pos in n.
func clearBit(n int, pos uint) int {
	n &^= (1 << pos)
	return n
}
func main() {
	var path string

	if runtime.GOOS == "windows" {

		path = "c:/temp/godump/"
	} else {
		path = "/Volumes/3tb/godump/"
	}

	os.MkdirAll(path, 777)

	makedummyfile.MakeDummyfile((path + "pdnadata1MIL.bin"), 1000)
	makedummyfile.MakeDummyfile((path + "reportpdnadata10000.bin"), 1000)
	/*
	makedummyfile.Downloadfile("https://upload.wikimedia.org/wikipedia/commons/d/db/Patern_test.jpg", path + "filefromgoogle2.jpg")
	*/
	//Downloadfile("http://172.20.2.171:8500/v1/catalog/service/filerepo", path + "consulservies.json")
	//randomin pdna file generator
	dat, err := ioutil.ReadFile(path + "pdnadata1MIL.bin")
	makedummyfile.Check(err)
	reportDNA, err := ioutil.ReadFile(path + "reportpdnadata10000.bin")
	makedummyfile.Check(err)


	//creating slices which are pointers instead of passing in full array by value
	refSet := dat[0:]
	reportSet := reportDNA[0:]

	//pdnacompare.DoPDNAWORK(dat,reportDNA)
	//msgqueue.CreateQueue("amqp://guest:guest@172.20.2.209:5672/","PDNAQueue")
	//msgqueue.Producer("amqp://guest:guest@172.20.2.209:5672/", "PDNAQueue", "Go Do Some Work!!!")
	msgqueue.Consumer("amqp://guest:guest@172.20.2.196:5672/", "PDNAQueue", "Workerx", refSet, reportSet)

	//services := goConsul.GetConfile(path + "consulservies.json")
	//fmt.Print(services)

}