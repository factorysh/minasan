package blacklist

import (
	"bufio"
	"io"
	"os"
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
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		line = strings.TrimSpace(line)
		if len(line) > 0 && !strings.HasPrefix(line, "#") {
			Add(line)
		}
		if err == io.EOF {
			break
		}
	}
	return nil
}

func Lenght() int {
	return len(blacklist)
}
