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
func ComparePDNA(dat []byte,reportDNA []byte)  int {

	for i := 1; i <= 100; i++ {
		
		for j := 1; j <= 144; j++ {
			fmt.Print(dat[(i * 144)-j])

		}
		fmt.Println("")
		//ComparePDNA(refDNA,reportDNA)
	}


	return 0
}
func main() {

		MakeDummyfile("c:\\temp\\pdnadata100.bin",100)
		MakeDummyfile("c:\\temp\\reportpdnadata1000.bin",1000)
		MakeDummyfile("c:\\temp\\pdnadata10000.bin",10000)
		MakeDummyfile("c:\\temp\\pdnadata100000.bin",100000)


	Downloadfile("https://upload.wikimedia.org/wikipedia/commons/d/db/Patern_test.jpg", "c:\\temp\\filefromgoogle2.jpg")

	dat, err := ioutil.ReadFile("c:\\temp\\pdnadata100.bin")
	check(err)
	reportDNA, err := ioutil.ReadFile("c:\\temp\\reportpdnadata1000.bin")
	check(err)

	//fmt.Print(string(dat))
	//fmt.Print(len(dat))
	ComparePDNA(dat, reportDNA)
}