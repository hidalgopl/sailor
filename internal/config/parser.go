package config

import (
	"github.com/hidalgopl/sailor/internal/messages"
	"github.com/pkg/errors"
	"strings"
)

// Used if test config contains tests key
func Intersection(allTests, requestedTests []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range allTests {
		m[item] = true
	}

	for _, item := range requestedTests {
		key := strings.ToUpper(item)
		if _, ok := m[key]; ok {
			c = append(c, key)
		}
	}
	return
}

// Used if test config contains exclude key
func Difference(a, b []string) (c []string) {
	return
}

type TestConfParser struct {
	config   *Config
	allTests []string
}

func NewTestParser(conf *Config) *TestConfParser {
	tp := &TestConfParser{
		conf,
		messages.TestNames,
	}
	return tp
}

func (tp *TestConfParser) GetTestList() ([]string, error) {
	if len(tp.config.Tests) == 0 {
		return nil, errors.New("")
	}
	tests := Intersection(tp.allTests, tp.config.Tests)
	if len(tests) == 0 {
		return nil, errors.New("")
	}
	return tests, nil
}
