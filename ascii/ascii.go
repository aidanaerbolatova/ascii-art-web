package ascii

import (
	"io/ioutil"
	"strings"
)

func Ascii(text string, banner string) (string, int) {
	for _, letter := range text {
		if letter != 10 && letter != 13 && letter < 32 || letter > 126 {
			return "", 400
		}
	}
	filename := ""
	switch banner {
	case "Standard":
		filename = "ascii/source/standard.txt"
	case "Shadow":
		filename = "ascii/source/shadow.txt"
	case "Thinkertoy":
		filename = "ascii/source/thinkertoy.txt"
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", 500
	}
	k := map[string]string{
		"ascii/source/shadow.txt":     "d44671e556d138171774efbababfc135",
		"ascii/source/standard.txt":   "ac85e83127e49ec42487f272d9b9db8b",
		"ascii/source/thinkertoy.txt": "8efd138877a4b281312f6dd1cbe84add",
	}
	hash := Hash(data, k[filename])
	if !hash {
		return "", 500
	}
	ascii := string(data)
	splitText := strings.Split(text, "\r\n")
	asciiMap := MakeMap(ascii)
	result := ""
	for _, w := range splitText {
		result += GetToStr(w, asciiMap)
	}
	return result, 200
}
