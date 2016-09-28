// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"os"
	// "math/rand"

	// "io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
type RefSet struct {
	pdna [] byte
}


func main() {

	// To start, here's how to dump a string (or just
	// bytes) into a file.

	p:= [][]int8{}
	//p :=[100]RefSet{}
	ff, err2 := os.Create("c:\\temp\\pdnadat2")
	check(err2)
	x := bufio.NewWriter(ff)
	for jj := 0; jj < 1000000; jj++ {
		var ranp []int8
		for ii := 0; ii < 144; ii++ {
			ranp= append(ranp,int8(ii))
			//fmt.Printf("aa")

		}
		//fmt.Printf("\n")
		p=append(p,ranp)

		fmt.Fprint(x, ranp,"\n")
	}

	x.Flush() // Don't forget to flush!

	/*

*/
	// For more granular writes, open a file for writing.
	f, err := os.Create("c:\\temp\\dat2")
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	defer f.Close()

	// You can `Write` byte slices as you'd expect.
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// A `WriteString` is also available.
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a `Sync` to flush writes to stable storage.
	f.Sync()

	// `bufio` provides buffered writers in addition
	// to the buffered readers we saw earlier.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// Use `Flush` to ensure all buffered operations have
	// been applied to the underlying writer.
	w.Flush()

}