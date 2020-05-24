// dl-brute <filename for inputs>.
// On input a file containing decimal-formatted (p,g,h), prints x to standard output
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

func main() {
	fname := os.Args[1]
	dat, readerr := ioutil.ReadFile(fname)
	if readerr != nil {
		os.Stderr.WriteString("Read Error")
		os.Exit(1)
	}
	nums := bytes.Split(dat, []byte(","))
	p, _ := new(big.Int).SetString(string(bytes.Split(nums[0], []byte("("))[1]), 10)
	g, _ := new(big.Int).SetString(string(nums[1]), 10)
	h, _ := new(big.Int).SetString(string(bytes.Split(nums[2], []byte(")"))[0]), 10)
	x := big.NewInt(1)
	for {
		if p.Cmp(x) == 0 {
			os.Stderr.WriteString("Exhausted all p")
			os.Exit(-1)
		}
		gExpX := new(big.Int).Exp(g, x, p)
		if h.Cmp(gExpX) == 0 {
			break
		}
		x.Add(x, big.NewInt(1))
	}
	fmt.Printf("%s\n", x.String())
}
