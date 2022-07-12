package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	var target string
	{
		if len(os.Args) < 2 {
			fmt.Fprintln(os.Stderr, "ERROR: bad request: missing target")
			os.Exit(1)
			return
		}

		target = os.Args[1]
	}

	var kind string
	{
		const colon string = ":"

		index := strings.Index(target, colon)

		if index < 0 {
			fmt.Fprintf(os.Stderr, "ERROR: bad request: bad target (%q)\n", target)
			os.Exit(1)
			return
		}

		kind = target[:index]

		if " " == kind {
			fmt.Fprintf(os.Stderr, "ERROR: bad request: bad target (%q)\n", target)
			os.Exit(1)
			return
		}
	}

	var resp *http.Response
	{
		var err error

		resp, err = http.Get(target)
		if nil != err {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
			return
		}
		if nil == resp {
			fmt.Fprintln(os.Stderr, "ERROR: internal error")
			os.Exit(1)
			return
		}
	}

	{
		defer resp.Body.Close()

		_, err := io.Copy(os.Stdout, resp.Body)
		if nil != err {
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			os.Exit(1)
			return
		}
	}
}
