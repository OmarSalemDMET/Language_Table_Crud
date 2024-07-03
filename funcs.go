package main

import (
	"LangAssist/wordArch"
	"bufio"
	"fmt"
	"os"
)

func exportToArray(m map[string]wordArch.WordArch) []wordArch.WordArch {
	var t []wordArch.WordArch
	for _, element := range m {
		t = append(t, element)
	}

	return t
}
func loadWordTable() map[string]wordArch.WordArch {
	file, err := os.Open("WordTable.txt")
	if err != nil {
		panic(1)
	}
	var myMap = make(map[string]wordArch.WordArch)
	scanner := bufio.NewScanner(file)
	var i int = 0
	for scanner.Scan() {
		var index string
		var (
			word  string
			lang  string
			categ string
		)
		for i < len(scanner.Text()) {
			if scanner.Text()[i] == '=' {
				break
			}

			index += string(scanner.Text()[i])
			i++

		}
		i++
		for i < len(scanner.Text()) {
			if scanner.Text()[i] == '=' {
				break
			} else {
				word += string(scanner.Text()[i])
				i++
			}
		}

		i++

		for i < len(scanner.Text()) {
			if scanner.Text()[i] == '=' {
				break
			} else {
				lang += string(scanner.Text()[i])
				i++
			}
		}

		i++

		for i < len(scanner.Text()) {
			if scanner.Text()[i] == '=' {
				break
			} else {
				categ += string(scanner.Text()[i])
				i++
			}
		}

		myMap[index] = wordArch.WordArch{WordEn: index, Word: word, Lang: lang, Categ: categ}
		i = 0
	}

	defer file.Close()

	return myMap
}

func saveWordTable(wordM map[string]wordArch.WordArch) {
	file, err := os.Create("WordTable.txt")
	if err != nil {
		fmt.Println("Failed to create file ", err)
		return
	}

	defer file.Close()

	for index, word := range wordM {
		tLine := fmt.Sprintf("%s=%s=%s=%s \n", index, word.Word, word.Lang, word.Categ)
		_, err := file.WriteString(tLine)
		if err != nil {
			fmt.Println("Failed to SAVE content")
			return
		}

	}

	fmt.Println("Saved to file Successfully")

}
