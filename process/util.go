package process

import (
	"bufio"
	"strings"
	"time"
)

func findPids(content, processName string) (pIds []string) {
	scanner := bufio.NewScanner(strings.NewReader(content))

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if len(words) > 0 && words[len(words)-1] == processName {
			pIds = append(pIds, words[len(words)-2])
		}
	}

	return
}

func parseDuration(interval int64) time.Duration {
	if interval < 5 {
		interval = 5
	}

	return time.Duration(interval) * time.Second
}

func isEmpty(pids []string) bool {
	return pids == nil || len(pids) == 0
}
