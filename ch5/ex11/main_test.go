package main

import (
	"fmt"
	"testing"
)

func TestToposort(t *testing.T) {
	if _, ok := toposort(prereqs); ok {
		fmt.Errorf("unexpected result")
	}
}

func topoCheck(req map[string]map[string]bool, order []string) (string, bool) {
	seen := make(map[string]bool)
	for _, o := range order {
		if req[o] == nil {
			return "", true
		}
		for item := range req[o] {
			if !seen[item] {
				return fmt.Sprintf("unexpected order. %v, %v", o, item), false
			}
			seen[item] = true
		}
	}
	return "", true
}
