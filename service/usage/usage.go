package usage

import (
	"flag"
	"fmt"
)

var (
	nFlag   int
	bFlag   bool
	strFlag string
)

func Usage() {

	// flag.IntVar(&nFlag, "n", 0, "int flag value")
	flag.BoolVar(&bFlag, "b", false, "bool flag value")
	// flag.StringVar(&strFlag, "s", "default", "string flag value")
	flag.Parse()
	fmt.Println(nFlag, bFlag, strFlag)
}
