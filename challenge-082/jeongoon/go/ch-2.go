/*
 * note: this is different approach when compared to my other solutions
 *
 * checking just interleaving or not (non-recursive version)
 *
 * Tested with:

  *  go run ch-2.go ABCDEFGHIJKLMNOPQRSTUVWXYZ ABCDEFGHIJKLMNOPQRSTUVWXYZ ABCABCDDEFGHIJEFKLGHIJKLMNMNOOPPQRQRSTUVSTWXYUVWXYZZ #  -> 1
  *  go run ch-2.go XY XX XYXX #  -> 1
  *  go run ch-2.go YX X XXY   #  -> 0
  */

package main

import (
	"os"
	"fmt"
)

func isInterleaving (A string, B string, C string) bool {
	Alen, Blen, Clen := len(A), len(B), len(C)
	if Alen + Blen != Clen {
		return false
	}

	Apin, Bpin := -1, -1
	checkpingPlanB := false

	for Ai, Bi, Ci := 0, 0, 0 ;; Ci = Ai + Bi {
		if checkpingPlanB {
			if Bpin > 0 {
				// note: it was A[Ai] == B[Bi] == C[Ci]
				// and tried A already.
				Bi = Bpin + 1
				Ai = Apin
				Apin, Bpin = -1, -1
				checkpingPlanB = false
			} else {
				// there is no plan B ...
				return false
			}
			continue
		} else {
			if Ci == Clen {
				return true
			}
			if Ai == Alen {
				if B[Bi:] == C[Ci:] {
					return true
				} else {
					checkpingPlanB = true
					continue
				}
			} else if Bi == Blen {
				return A[Ai:] == C[Ci:]
			}

			if A[Ai] == B[Bi] {

				if A[Ai] != C[Ci] {
					checkpingPlanB = true
				}  else {
					// remember this node
					Apin, Bpin = Ai, Bi
					// try plan A first
					Ai++
				}
			} else {

				if A[Ai] == C[Ci] {
					Ai++
				} else if B[Bi] == C[Ci] {
					Bi++
				} else {
					checkpingPlanB = true
				}
			}
		}
	}
}

func usage() {
	fmt.Println( "Usage: go run ch-2.go " +
		"<string> <string> <may be interleaved string>\n" )
}

func main() {
	if len(os.Args[1:]) != 3 {
		usage();
		os.Exit(1);
	}
	A, B, C := os.Args[1], os.Args[2], os.Args[3]

	if isInterleaving(A, B, C) {
		fmt.Println( "1" )
	} else {
		fmt.Println( "0" )
	}
}
