package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Type int

const (
	unknown Type = iota
	h1
	h2
	h3
	li
	p
	br
)

type Token struct {
	ty  Type
	val string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func preprocess(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func tokenize(s string) Token {
	switch {
	case strings.HasPrefix(s, "#"):
		n := strings.Count(s, "#")
		if n == 1 {
			return Token{h1, s[1:]}
		}
		if n == 2 {
			return Token{h2, s[2:]}
		}
		if n >= 3 {
			return Token{h3, s[3:]}
		}
	case strings.HasPrefix(s, "-") || strings.HasPrefix(s, "*"):
		return Token{li, s[1:]}
	case len(s) == 0:
		return Token{br, s}
	default:
		return Token{p, s}
	}
	return Token{unknown, s}
}

func generate(t Token) string {
	switch t.ty {
	case h1:
		return "<h1>" + t.val + "</h1>"
	case h2:
		return "<h2>" + t.val + "</h2>"
	case h3:
		return "<h3>" + t.val + "</h3>"
	case li:
		return "<li>" + t.val + "</li>"
	case p:
		return "<p>" + t.val + "</p>"
	case br:
		return "<br/>"
	}
	return ""
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
	for scanner.Scan() {
		text := preprocess(scanner.Text())
		token := tokenize(text)
		html := generate(token)
		fmt.Fprintln(writer, html)
	}
	err = scanner.Err()
	check(err)
	writer.Flush()
}
