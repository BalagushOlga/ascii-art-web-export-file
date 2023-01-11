package main

import(
	"os"
	"strings"
	"net/http"
)

func asciiArt(inputText string, nameOfFile string, w http.ResponseWriter) string {
	if inputText == "" {
		w.WriteHeader(http.StatusNotFound)
		return "No input"
	}

	if inputText[0] == '"' && inputText[len(inputText)-1] == '"'{
		inputText = inputText[1:len(inputText)-1]
	}

	input := strings.Replace(string(inputText), "\r\n", "\\n", -1)
	inputStringByPart := strings.Split(input, "\\n")
	inputAscii, _ := os.ReadFile("./banners/" + nameOfFile + ".txt")
	if inputAscii == nil {
		w.WriteHeader(http.StatusNotFound)
		return "No ascii font"
	}

	ASCIIcharactersByLines := strings.Split(string(inputAscii[1:]), "\n\n")
	w.WriteHeader(http.StatusOK)
	return separateInputString(ASCIIcharactersByLines, inputStringByPart)
}

func separateInputString(ASCIIcharactersByLines []string, inputString []string) string{
	str := ""
	for i, element := range inputString {
		if inputString[0] == "" && (len(inputString) == 1 || i == len(inputString) - 1){
			break
		}else if element == "" {
			str += "\n"
		} else {
			list := createArrayToWriteALine(ASCIIcharactersByLines, element)
			for index, line := range list {
				str += line
				if !(index == len(list) - 1  && i == len(inputString) - 1){
					str += "\n"
				}
			}
		}
	}
	return str
}

func createArrayToWriteALine(s []string, inp string) [8]string {
	list := [8]string{}
	for _, val := range inp {
		for i := range s {
			if val == rune(i+32) {
				character := strings.Split(s[i], "\n")
				for k := 0; k < 8; k++ {
					list[k] += character[k]
				}
			}
		}
	}
	return list
}