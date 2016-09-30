package makedummyfile

import (
	"os"
	"bufio"
	"encoding/binary"
	"net/http"
	"io"
	"math/rand"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func MakeDummyfile(filepath string, totalSize int) {

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	//const totalCount  = 1000000
	p := [][]byte{}
	//p :=[100]RefSet{}

	ff, err2 := os.Create(filepath)
	Check(err2)
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
	Check(err)
	defer out.Close()

	resp, err := http.Get(url)
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	Check(err)

}