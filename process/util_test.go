package process

import (
	"testing"

	"github.com/golib/assert"
)

func TestSplits(t *testing.T) {
	assertion := assert.New(t)

	consoleContent := `
    xxxx xxxx      1 test1
    xxxx xxxx      2 test2
    xxxx xxxx      3 test3`

	processName := "test1"

	foundPids := findPids(consoleContent, processName)

	assertion.Equal(1, len(foundPids))
	assertion.Equal("1", foundPids[0])
}
