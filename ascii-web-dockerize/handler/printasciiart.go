package handler

import (
	"fmt"
	"html/template"
	"path/filepath"

	// "log"
	"net/http"
	"os"
	"strings"
)

type pageResult struct {
	Result string
}

func PrintAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	text := r.FormValue("text")
	banner := r.FormValue("banners")

	text = strings.ReplaceAll(text, "\\n", "\n")
	text = strings.ReplaceAll(text, "\r\n", "\n")

	if text == "" || banner == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	newSplitText := strings.Split(text, "\n")
	// bannerPath := "font/" + banner
	bannerPath := filepath.Join("font", banner)

	bannersData, err := ReadFile(bannerPath)
	if err != nil {
		// log.Println("ReadFile Error :", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	result, err := GenerateAsciiArt(newSplitText, bannersData)
	if err != nil {
		http.Error(w, "Non Printable Charater", http.StatusBadRequest)
		return
	}
	
	templ, err := template.ParseFiles(
		"template/base.tmpl.html",
		"template/home.tmpl.html",
		"template/partial/result.tmpl.html",
	)
	if err != nil {
		// log.Println("Template error:", err)
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	value := pageResult{Result: result}
	err = templ.ExecuteTemplate(w, "base", value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles(
		"template/base.tmpl.html",
		"template/home.tmpl.html",
		"template/partial/result.tmpl.html",
	)
	if err != nil {
		// log.Println("Template error:", err)
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err = templ.ExecuteTemplate(w, "base", pageResult{})
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
				if start >= len(banner) {
					return "", fmt.Errorf("Banner File Corrupted")
				}
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
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	files := strings.ReplaceAll(string(data), "\r\n", "\n")
	result := strings.Split(files, "\n")
	return result, nil
}
