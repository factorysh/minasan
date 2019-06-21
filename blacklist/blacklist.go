package blacklist

import (
	"bytes"
	"io"
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

func AddIfValid(line string) {
	if len(line) > 0 && !strings.HasPrefix(line, "#") {
		Add(line)
	}
}

func ReadBlacklistFile(path string) error {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	buffer := bytes.NewBuffer(raw)
	for {
		line, err := buffer.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			AddIfValid(line)
			break
		}
		if err != nil {
			return err
		}
		AddIfValid(line)
	}
	return nil
}

func Lenght() int {
	return len(blacklist)
}
