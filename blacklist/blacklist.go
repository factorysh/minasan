package blacklist

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

var blacklist map[string]interface{}

func Add(mail string) {
	if blacklist == nil {
		blacklist = make(map[string]interface{})
	}
	blacklist[mail] = true
}

func IsBlackListed(mail string) bool {
	_, ok := blacklist[mail]
	if !ok {
		// Search for wildcards
		for key, value := range blacklist {
			matched, _ := filepath.Match(key, mail)
			if matched && value == true {
				ok = true
			}
		}
	}
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
