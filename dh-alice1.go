package main

// dh-alice1 <filename for message to Bob> <file name to store secret key>
// Simulate Alice's initial message to Bob
// Outputs decimal formatted (p,g,g^a) to Bob, writes (p,g,a) to a second file

import (
	"crypto/rand"
	"math/big"
	"os"
	"strings"
)

// Generate an n-bit safe prime number
func generateSafePrime(n int) (*big.Int, *big.Int) {
	for {
		// Generate some prime
		q, _ := rand.Prime(rand.Reader, n-1)
		p := new(big.Int).Set(q)
		// Check if p=2q+1 is prime
		p.Mul(p, big.NewInt(2))
		p.Add(p, big.NewInt(1))
		if p.ProbablyPrime(20) {
			return q, p
		}
	}
	return big.NewInt(0), big.NewInt(0)
}

func generateGenerator(q *big.Int, p *big.Int) *big.Int {
	//    This section gives an algorithm (derived from [FIPS-186]) for generating g.
	//
	// 1. Let j = (p - 1)/q.
	//
	// 2. Set h = any integer, where 1 < h < p - 1 and h differs from any value previously tried.
	//
	// 3. Set g = h^j mod p
	//
	// 4. If g = 1 go to step 2
	h, errH := rand.Int(rand.Reader, p)
	if errH != nil {
		os.Stderr.WriteString("Bad P Chosen.")
		os.Exit(1)
	}
	j := new(big.Int).Sub(p, big.NewInt(1))
	j.Div(j, q)
	for {
		g := new(big.Int).Exp(h, j, p)
		if g.Cmp(big.NewInt(1)) == 0 {
			// Unlucky
			h.Add(h, big.NewInt(1))
			if h.Cmp(p) != -1 {
				h = big.NewInt(2)
			}
			continue
		}
		return g
	}
	return big.NewInt(0)
}

func main() {
	bits := 1024
	if len(os.Args) < 3 {
		os.Stderr.WriteString("Not enough args:\n\n\tUsage: dh-alice1 <filename for message to Bob> <file name to store secret key>\n")
		os.Exit(1)
	}
	f1 := os.Args[1] //Bob's message
	f2 := os.Args[2] //Alice's stored secret

	//Generate params (q, p, g, b, g^a)
	q, p := generateSafePrime(bits)
	g := generateGenerator(q, p)
	// Calc g^a mod p
	a, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(1)))
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	a.Add(a, big.NewInt(1)) // Alice's secret integer [1,p)
	gExpA := new(big.Int).Exp(g, a, p)

	secret := new(strings.Builder)
	secret.WriteString("(" + p.String() + "," + g.String() + "," + a.String() + ")")
	share := new(strings.Builder)
	share.WriteString("(" + p.String() + "," + g.String() + "," + gExpA.String() + ")")

	// Write deimal-formatted (p,g,g^a) to f1 and (p,g,a) to f2
	fptr1, err1 := os.Create(f1)
	defer fptr1.Close()
	if err1 != nil {
		os.Stderr.WriteString("Error Creating File: " + f1)
		os.Exit(1)
	}
	fptr1.WriteString(share.String())

	fptr2, err2 := os.Create(f2)
	defer fptr2.Close()
	if err2 != nil {
		os.Stderr.WriteString("Error Creating File: " + f2)
		os.Exit(1)
	}
	fptr2.WriteString(secret.String())

}
