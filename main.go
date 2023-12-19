package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	DEFAULT_INTERVAL = 1
	INTERVAL_FLAG    = "n"
	DEFAULT_ARGSTART = 1
)

func main() {
	var nFlag = flag.Int(INTERVAL_FLAG, DEFAULT_INTERVAL, "interval (integer)")
	flag.Parse()

	if len(os.Args[1:]) == 0 {
		log.Println("Please provide one or more arguments")
		os.Exit(1)
	}

	argStart := parseArgs(os.Args)

	for {
		cmd := exec.Command(os.Args[argStart], os.Args[argStart+1:]...)
		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		time.Sleep(time.Duration(*nFlag) * time.Second)
	}
}

func parseArgs(args []string) (argStart int) {
	for i := 0; i < len(args); i++ {
		if args[i] == fmt.Sprintf("-%s", INTERVAL_FLAG) {
			_, err := strconv.Atoi(args[i+1])
			if err != nil {
				// invalid flag
				return DEFAULT_INTERVAL
			}

			return i + 2
		}
	}
	return DEFAULT_ARGSTART
}
