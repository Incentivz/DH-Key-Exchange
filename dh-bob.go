package main

// dh-bob <filename of message from Alice> <file name of message back to Alice>
// Reads in Alice’s message, outputs (g^b) to Alice, prints the shared secret g^ab.

import (
    "fmt"
	"math/big"
	"os"
    "crypto/rand"
    "io/ioutil"
    "bytes"
)



func main() {
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Not enough args:\n\n\tUsage: dh-bob <filename of message from Alice> <file name of message back to Alice>\n")
		os.Exit(1)
	}
	f1 := os.Args[1] //Alice's Message
	f2 := os.Args[2] //Message to Alice

    //Generate params (p, g, b, g^a, g^b)
    data, errRead := ioutil.ReadFile(f1)
    if errRead != nil {
        os.Stderr.WriteString("Error Reading from File")
        os.Exit(1)
    }
    nums := bytes.Split(data, []byte(","))
    p,_ := new(big.Int).SetString(string(bytes.Split(nums[0], []byte("("))[1]), 10)
    g,_ := new(big.Int).SetString(string(nums[1]),10)
    bytegExpA := bytes.Split(nums[2], []byte(")"))
    gExpA,_ := new(big.Int).SetString(string(bytegExpA[0]),10)

    // Calc g^a mod p
	b, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(1)))
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	b.Add(b, big.NewInt(1)) // Bobs secret integer [1,p)
	gExpB := new(big.Int).Exp(g, b, p)

    // Print g^ab
	pk := new(big.Int).Exp(gExpA, b, p)
    fmt.Printf("%s\n", pk.String())
    // Reads in Alice’s message, outputs (g^b) to Alice
    fptr, errWrite := os.Create(f2)
    if errWrite != nil {
        os.Exit(1)
    }
    fptr.WriteString("("+ gExpB.String() +")")
}
