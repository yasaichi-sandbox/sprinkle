package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets " + otherWord,
}

func readTransformRulesFrom(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	transforms = []string{}
	s := bufio.NewScanner(file)

	for s.Scan() {
		transforms = append(transforms, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return transforms, nil
}

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	if filename != "" {
		rules, err := readTransformRulesFrom(filename)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		transforms = rules
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
