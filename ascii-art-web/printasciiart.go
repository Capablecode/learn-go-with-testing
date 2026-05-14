package asciiartweb

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func PrintAscii(w http.ResponseWriter, r *http.Request) {
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Bad Request", http.StatusBadRequest)
	// }
	text := r.FormValue("text")
	banner := r.FormValue("banners")
	log.Println(text, banner)
	// text := "hello"
	// banner := "standard.txt"

	if text == "" || banner == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	newText := strings.Split(text, "\n")
	banners, err := ReadFile(banner)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	result, err := GenerateAsciiArt(newText, banners)
	if err != nil {
		http.Error(w, "Non Printable Charater", http.StatusBadRequest)
	}
	w.Write([]byte(result))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("ascii-art-web/template/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GenerateAsciiArt(text, banner []string) (string, error) {
	var result strings.Builder

	for k := 0; k < len(text); k++ {
		if text[k] == "" {
			result.WriteString("\n")
			continue
		}

		//the i handles the vertical rows of the font
		for i := 1; i <= 8; i++ {
			//the j handles the horizontal movement of characters in the word
			for j := 0; j < len(text[k]); j++ {
				if text[k][j] < 32 || text[k][j] > 126 {
					return "", fmt.Errorf("found unprintable character")
				}
				character_index := int(text[k][j] - 32)
				start := (character_index * 9) + i
				result.WriteString((banner[start]))
			}
			//Move to a new line to continue printing the block of character line by line
			// fmt.Println()
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}

func ReadFile(path string) ([]string, error) {
	data, err := os.ReadFile("ascii-art-web/standard.txt")
	if err != nil {
		return nil, err
	}
	files := string(data)
	result := strings.Split(files, "\n")
	return result, nil
}
