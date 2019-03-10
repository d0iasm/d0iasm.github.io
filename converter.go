package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fname := os.Args[1]
	name := strings.Split(fname, ".")

	rfile, err := os.Open(fname)
	check(err)
	defer rfile.Close()

	wfile, err := os.Create(name[0] + ".html")
	check(err)
	defer wfile.Close()

        scanner := bufio.NewScanner(rfile)
        writer := bufio.NewWriter(wfile)
        text := ""
        for scanner.Scan() {
		text = scanner.Text()
		fmt.Fprintln(writer, text)
	}
	err = scanner.Err()
	check(err)
	writer.Flush()
}
