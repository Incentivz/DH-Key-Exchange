// dh-alice2 <filename of message from Bob> <filedname to read secret key>
// Read in Bob's message and Alice's stored secrte, prints the shared secret g^ab
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Not enough args:\n\n\tUsage: dh-alice2 <filename of message from Bob> <filedname to read secret key>\n")
		os.Exit(1)
	}
	f1 := os.Args[1] //Bob's message
	f2 := os.Args[2] //Alice's stored secret

	// Parse g^b from (g^b)
	f1data, rErr1 := ioutil.ReadFile(f1)
	if rErr1 != nil {
		os.Stderr.WriteString("Error reading file")
		os.Exit(1)
	}

	temp1 := bytes.Split(f1data, []byte("("))
	temp2 := bytes.Split(temp1[1], []byte(")"))
	gExpB, _ := new(big.Int).SetString(string(temp2[0]), 10)
	f2data, rErr2 := ioutil.ReadFile(f2)
	if rErr2 != nil {
		os.Stderr.WriteString("Error Reading File")
		os.Exit(2)
	}
	nums := bytes.Split(f2data, []byte(","))
	p, _ := new(big.Int).SetString(string(bytes.Split(nums[0], []byte("("))[1]), 10)
	byteA := bytes.Split(nums[2], []byte(")"))
	a, _ := new(big.Int).SetString(string(byteA[0]), 10)

	shared := new(big.Int).Exp(gExpB, a, p)
	fmt.Printf("%s", shared)

}
