package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var templateFunctions = template.FuncMap {}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	_, err := RenderTemplateSet(w)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
	}

	fmt.Printf("\n")
}

func RenderTemplateSet(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		fmt.Println("page:", page)

		name := filepath.Base(page)

		fmt.Println("name:", name)

		templateSet, err := template.New(name).Funcs(templateFunctions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		fmt.Println("templateSet:", templateSet)

		layoutGlob := "./templates/*.layouts.tmpl"

		matches, err := filepath.Glob(layoutGlob)
		if err != nil {
			return myCache, err
		}

		fmt.Println("matches:", matches)

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob(layoutGlob)
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	fmt.Println("myCache:", myCache)

	// fmt.Println("pages", pages)

	return myCache, nil
}
