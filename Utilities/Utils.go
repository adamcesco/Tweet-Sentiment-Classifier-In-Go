package Utilities

import (
	"encoding/csv"
	"github.com/reiver/go-porterstemmer"
	"os"
	"strings"
)

func ReadCSVFile(filePath string) ([]string, error) {
	fileContent, err := os.Open(filePath)
	if err != nil {
		return []string{}, err
	}
	defer fileContent.Close()

	lines, err := csv.NewReader(fileContent).Read()
	if err != nil {
		return []string{}, err
	}
	return lines, nil

}

func IsUrl(text string) bool {
	return strings.Contains(text, "www.") || strings.Contains(text, "http")
}

func IsUsername(text string) bool {
	return text[0] == '@'
}

func IsHashtag(text string) bool {
	return text[0] == '#'
}

func IsStopWord(text string) bool {
	_, conatins := StopWords()[text]
	return conatins
}

func Stem(text string) string {
	return porterstemmer.StemString(text)
}
