// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"bufio"
	//"io/ioutil"
	"os"
	// "math/rand"

	// "io/ioutil"
	"encoding/binary"
	//"fmt"
	"math/rand"
	"io/ioutil"
	"fmt"
	"net/http"
	"io"
	"goConsul"

	"runtime"

	"time"
	"log"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type RefSet struct {
	pdna [] byte
}

func MakeDummyfile(filepath string, totalSize int) {

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	//const totalCount  = 1000000
	p := [][]byte{}
	//p :=[100]RefSet{}

	ff, err2 := os.Create(filepath)
	check(err2)
	x := bufio.NewWriter(ff)
	for jj := 0; jj < totalSize; jj++ {
		var ranp []byte
		for ii := 0; ii < 144; ii++ {
			ranp = append(ranp, byte(rand.Int()))
			//fmt.Printf("aa")

		}
		//fmt.Printf("\n")
		p = append(p, ranp)

		binary.Write(x, binary.BigEndian, ranp)
		//fmt.Print( ranp)
	}
	x.Flush() // Don't forget to flush!

}

func Downloadfile(url string, filename string) {

	out, err := os.Create(filename)
	check(err)
	defer out.Close()

	resp, err := http.Get(url)
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	check(err)

}
func pdnaCompare(ref_dna []byte, reporDNA []byte) bool {
	//var maxTotal = 2000
	var runnintTotal int
	x :=  int(uint(ref_dna[1]) - uint(reporDNA[0]))
	//fmt.Println(int(x),uint(ref_dna[1]) , uint(reporDNA[0]))
	runnintTotal=(runnintTotal+x)
	if (x > 200) {
		return false
	} else
	{
		return true
	}
}
func ComparePDNA(dat []byte, reportDNA []byte) int {
	var datlen int
	var replen int
	datlen = len(dat) / 144
	replen = len(reportDNA) / 144

	fmt.Println(datlen, " Refset")
	fmt.Println(replen, " report ")
	t:=time.Now()
	fmt.Println(t)
	for i := 0; i < datlen; i++ {
		startIdx := i * 144
		ref_dna := dat[startIdx:startIdx + 144]
		for j := 0; j < replen; j++ {
			startIdx2 := j * 144
			report_dna := reportDNA[startIdx2:startIdx2 + 144]
			_ = pdnaCompare(ref_dna, report_dna)
			//fmt.Print((result))
		}
		/*
		for j := 1; j <= 144; j++ {
			fmt.Print(dat[(i * 144) - j])

		}*/
		//fmt.Println("")
		//ComparePDNA(refDNA,reportDNA)
	}
	elapsed := time.Since(t)
	log.Printf("Compare took %s", elapsed)
	return 0
}

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}
func main() {
	var path string

	if runtime.GOOS == "windows" {

		path = "c:/temp/godump/"
	} else {
		path = "/Volumes/3tb/godump/"
	}

	os.MkdirAll(path, 777)

	MakeDummyfile((path + "pdnadata100.bin"), 100)
	/*	MakeDummyfile((path + "reportpdnadata10000.bin"), 100000)
				MakeDummyfile((path+"pdnadata10000.bin"),10000)
				MakeDummyfile((path+"pdnadata100000.bin"),100000)
				MakeDummyfile((path+"pdnadata1MIL.bin"),1000000)
		*/
	Downloadfile("https://upload.wikimedia.org/wikipedia/commons/d/db/Patern_test.jpg", path + "filefromgoogle2.jpg")
		//Downloadfile("http://172.20.2.171:8500/v1/catalog/service/filerepo", path + "consulservies.json")
	dat, err := ioutil.ReadFile(path + "pdnadata1MIL.bin")
	check(err)
	reportDNA, err := ioutil.ReadFile(path + "reportpdnadata10000.bin")
	//reportDNA, err := ioutil.ReadFile(path + "pdnadata100.bin")


	check(err)

	//fmt.Print(string(dat))
	//fmt.Print(len(dat))
	/*x:=len(reportDNA)
	fmt.Print(x)
	s := make([]uint, x, x)
	fmt.Println(s[0])
	for i:=0;i<x;i++{
		//fmt.Print(i)
		s[i]=uint(reportDNA[i])
	}
	y:=len(dat)
	r := make([]uint, y, y)
	fmt.Println(y)
	for i:=0;i<y;i++{
		//fmt.Print(i)
		r[i]=uint(dat[i])
	}*/
fmt.Println(MaxParallelism(),"Number CPU")
	cpu:=MaxParallelism()-2
	for i:=0;i<cpu-1;i++{
		s:=dat[i:500000*144]
		go ComparePDNA(s, reportDNA)
		time.Sleep(1000 * time.Millisecond)
	}
	x:=dat[500000*144:]
	ComparePDNA(x, reportDNA)
	services := goConsul.GetConfile(path + "consulservies.json")
	fmt.Print(services)

}