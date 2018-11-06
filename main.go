package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/OGNeutron/golang-blockchain/blockchain/cli"
)



func main() {
	defer os.Exit(0)
	cli := CommandLine{}
	cli.run()
}
