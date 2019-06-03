package blacklist

import (
	"io/ioutil"
	"strings"
)

var blacklist map[string]interface{}

func Add(mail string) {
	blacklist[mail] = true
}

func IsBlackListed(mail string) bool {
	_, ok := blacklist[mail]
	return ok
}

func ReadBlacklistFile(path string) error {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(raw), "\n") {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}
		Add(line)
	}
	return nil
}

func Lenght() int {
	return len(blacklist)
}
