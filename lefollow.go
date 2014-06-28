package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/bsphere/le_go"
	"log"
	"os"
	"os/exec"
)

func main() {
	var tokenflag string

	flag.StringVar(&tokenflag, "token", "", "-token=<logentries_token>")

	flag.Parse()

	if tokenflag == "" || len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: lefollow -token=<logentries_token> <filename>")
		flag.PrintDefaults()
		os.Exit(2)
	}

	le, err := le_go.Connect(tokenflag)
	if err != nil {
		log.Fatal(err)
	}

	defer le.Close()

	cmd := exec.Command("/usr/bin/tail", "-F", flag.Arg(0))

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	stdoutScanner := bufio.NewScanner(stdout)

	for stdoutScanner.Scan() {
		le.Println(stdoutScanner.Text())
	}
}
