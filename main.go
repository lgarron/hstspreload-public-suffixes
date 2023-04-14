package main

import (
	"fmt"
	"os"

	"github.com/chromium/hstspreload/chromium/preloadlist"
	"golang.org/x/net/publicsuffix"
)

func main() {
	l, err := preloadlist.NewFromLatest()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	for _, e := range l.Entries {
		publicSuffix, icann := publicsuffix.PublicSuffix(e.Name)
		if e.Name == publicSuffix {
			fmt.Printf("%s: %v\n", publicSuffix, icann)
		}
	}
}
