package main

import (
	"testing"
)

func TestValidRepoPath(t *testing.T) {
	var tests = []struct {
		repo  string
		valid bool
	}{
		{"", false},
		{"ntns", false},
		{"ntns/gh-mirror", true},
		{"NTNS/GH-mirror", true},
		{"ntns/gh-mirror-0", true},
		{"ntns/gh!mirror", false},
		{"ntns!/gh-mirror", false},
		{"ntns/gh_mirror", true},
		{"ntns/gh.mirror", true},
	}

	for _, test := range tests {
		got := isValidRepoPath(test.repo)
		want := test.valid
		if got != want {
			t.Errorf("repo %v, got %v, wanted %v", test.repo, got, want)
		}
	}
}
