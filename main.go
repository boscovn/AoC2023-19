package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	key         byte
	operator    byte
	value       int
	destination string
}
type workflow struct {
	rules       []rule
	destination string
}

func checkParts(workflows map[string]workflow, keyvals map[byte]int, key string) bool {
	if key == "A" {
		return true
	}
	if key == "R" {
		return false
	}
	wf := workflows[key]
	for _, r := range wf.rules {
		checksOut := false
		switch r.operator {
		case '<':
			checksOut = keyvals[r.key] < r.value
		case '>':
			checksOut = keyvals[r.key] > r.value
		case '=':
			checksOut = keyvals[r.key] == r.value
		}
		if checksOut {
			return checkParts(workflows, keyvals, r.destination)
		}
	}
	return checkParts(workflows, keyvals, wf.destination)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	workflows := make(map[string]workflow)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		begin := strings.Index(text, "{")
		parts := strings.Split(text[begin+1:len(text)-1], ",")
		var rules []rule
		for _, v := range parts[:len(parts)-1] {
			partsAB := strings.Split(v, ":")
			value, err := strconv.Atoi(partsAB[0][2:])
			if err != nil {
				return
			}
			rules = append(rules, rule{
				key:         partsAB[0][0],
				operator:    partsAB[0][1],
				value:       value,
				destination: partsAB[1],
			})

		}
		workflows[text[:begin]] = workflow{rules: rules, destination: parts[len(parts)-1]}
	}
	sum := 0
	for scanner.Scan() {
		keyvals := make(map[byte]int)
		text := scanner.Text()
		s := strings.Split(text[1:len(text)-1], ",")
		for _, v := range s {
			value, err := strconv.Atoi(v[2:])
			if err != nil {
				return
			}
			keyvals[v[0]] = value
		}
		if checkParts(workflows, keyvals, "in") {
			for _, v := range keyvals {
				sum += v
			}
		}

	}

	fmt.Println(sum)
}
