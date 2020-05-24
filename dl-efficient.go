package main

// Derived from the algorithm from Wikipedia

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
)

// Perform n^2 + 1 mod p
func algo(x *big.Int, n *big.Int) *big.Int {
	one := big.NewInt(1)
	two := big.NewInt(2)
	x2 := new(big.Int).Exp(x, two, n) // x^2 mod n
	x2.Add(x2, one)                   // (x^2 mod n) + 1
	x2.Mod(x2, n)                     // (x^2 + 1) mod n
	return x2
}

func main() {
	fname := os.Args[1]
	dat, readerr := ioutil.ReadFile(fname)
	if readerr != nil {
		os.Stderr.WriteString("Read Error")
		os.Exit(1)
	}
	nums := bytes.Split(dat, []byte(","))
	p, _ := new(big.Int).SetString(string(bytes.Split(nums[0], []byte("("))[1]), 10)
	// Pollard-Rho
	one := big.NewInt(1)
	x, y, d := big.NewInt(2), big.NewInt(2), big.NewInt(1)

	for d.Cmp(one) == 0 {
		x = algo(x, p)
		y = algo(algo(y, p), p)
		temp := new(big.Int).Sub(x, y)
		d.GCD(nil, nil, temp.Abs(temp), p) // gcd(|x - y|, p)
	}

	if d.Cmp(p) == 0 {
		fmt.Print("algorithm failed with default parameters (x = y = 2)")
	}
	fmt.Println(d.String())
}
