package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func asdot2plain(asdot string) int {
	// This returns an ASPLAIN formatted ASN given an ASDOT+ format
	parts := strings.Split(asdot, ".")
	if len(parts) != 2 {
		fmt.Println("Invalid ASDOT+ format")
		os.Exit(1)
	}
	left, err1 := strconv.Atoi(parts[0])
	right, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid ASDOT+ format")
		os.Exit(1)
	}
	return left*65536 + right
}

func asplain2asdot(asplain int) string {
	// This returns an ASDOT+ formatted ASN given an ASPLAIN format,
	// unless given a 16-bit ASN
	if asplain <= 65535 {
		return strconv.Itoa(asplain)
	}
	counter := asplain / 65536
	remainder := asplain % 65536
	return fmt.Sprintf("%d.%d", counter, remainder)
}

func main() {
	// Runs as a utility to convert between formats automatically
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Printf("  %s <asn>\n\n", os.Args[0])
		fmt.Println("    <asn> - ASN to convert in ASPLAIN or ASDOT+ format")
		fmt.Println("")
		fmt.Println("Outputs ASPLAIN if given ASDOT+, and ASDOT+ if given ASPLAIN,")
		fmt.Println("unless as 16-bit ASN, then no change")
		fmt.Println("")
		os.Exit(1)
	}

	start := os.Args[1]
	if strings.Contains(start, ".") {
		// ASDOT+
		fmt.Println(asdot2plain(start))
	} else {
		// ASPLAIN
		asplain, err := strconv.Atoi(start)
		if err != nil {
			fmt.Println("Invalid ASPLAIN format")
			os.Exit(1)
		}
		fmt.Println(asplain2asdot(asplain))
	}
}

