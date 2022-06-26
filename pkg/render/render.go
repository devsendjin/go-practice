package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templateFunctions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	templateCache, err := CreateTemplateCache()
	if err != nil {
		fmt.Println("Error parsing template:", err)
		log.Fatal(err)
	}

	currentTemplate, ok := templateCache[tmpl]
	fmt.Println("RenderTemplate - currentTemplate:", currentTemplate)
	if !ok {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)

	currentTemplate.Execute(buffer, nil)

	_, err = buffer.WriteTo(w)
	if err != nil {
		fmt.Println("RenderTemplate - Error writing template to browser", err)
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	// err = parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("Error parsing template:", err)
	// }

	fmt.Printf("\n\n")
	fmt.Printf("----------------------------")
	fmt.Printf("\n\n")
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		fmt.Println("CreateTemplateCache - page:", page)

		name := filepath.Base(page)

		fmt.Println("CreateTemplateCache - name:", name)

		templateSet, err := template.New(name).Funcs(templateFunctions).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		fmt.Println("CreateTemplateCache - templateSet:", templateSet)

		layoutGlob := "./templates/*.layout.tmpl"

		matches, err := filepath.Glob(layoutGlob)
		if err != nil {
			return templateCache, err
		}

		fmt.Println("CreateTemplateCache - matches:", matches)

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob(layoutGlob)
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = templateSet
	}

	fmt.Println("CreateTemplateCache - myCache:", templateCache)

	// fmt.Println("pages", pages)

	return templateCache, nil
}
