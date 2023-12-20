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
type valRange struct {
	minimum int
	maximum int
}

func copyMap(origin map[byte]valRange) map[byte]valRange {
	dest := make(map[byte]valRange)
	for k, v := range origin {
		dest[k] = v
	}
	return dest
}

func checkParts(workflows map[string]workflow, keyvals map[byte]valRange, key string) int {
	if key == "A" {
		sum := 1
		for _, v := range keyvals {
			sum *= v.maximum - v.minimum + 1
		}
		return sum
	}
	if key == "R" {
		return 0
	}
	wf := workflows[key]
	sum := 0
	for _, r := range wf.rules {
		switch r.operator {
		case '<':
			if keyvals[r.key].minimum < r.value {
				newKeyvals := copyMap(keyvals)
				newKeyvals[r.key] = valRange{minimum: keyvals[r.key].minimum, maximum: r.value - 1}
				sum += checkParts(workflows, newKeyvals, r.destination)
				keyvals[r.key] = valRange{minimum: r.value, maximum: keyvals[r.key].maximum}

			}
		case '>':
			if keyvals[r.key].maximum > r.value {
				newKeyvals := copyMap(keyvals)
				newKeyvals[r.key] = valRange{minimum: r.value + 1, maximum: keyvals[r.key].maximum}
				sum += checkParts(workflows, newKeyvals, r.destination)
				keyvals[r.key] = valRange{minimum: keyvals[r.key].minimum, maximum: r.value}

			}
		}

	}
	return sum + checkParts(workflows, keyvals, wf.destination)
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
	keyvals := make(map[byte]valRange)
	keyvals['x'] = valRange{minimum: 1, maximum: 4000}
	keyvals['m'] = valRange{minimum: 1, maximum: 4000}
	keyvals['a'] = valRange{minimum: 1, maximum: 4000}
	keyvals['s'] = valRange{minimum: 1, maximum: 4000}
	fmt.Println(checkParts(workflows, keyvals, "in"))
}
