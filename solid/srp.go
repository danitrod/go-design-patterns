package solid

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Single Responsibility Principle

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Saving to a file should not be a method of Journal (it's not part of its purpose)
type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.LineSeparator)), 0644)
}

func (p *Persistence) LoadFromFile(filename string) (*Journal, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	j := Journal{}
	for _, entry := range strings.Split(string(data), p.LineSeparator) {
		j.entries = append(j.entries, entry)
	}

	return &j, nil
}

// The single purpose of Journal is to manage the entries as a data structure. The purpose of the Persistence struct is to manage persistence.
