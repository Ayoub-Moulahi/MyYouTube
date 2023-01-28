package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type View struct {
	Tpl    *template.Template
	Layout string
}

// NewView used to create a new View by parsing the given files alongside the "views/layout/" files
func NewView(layout string, files ...string) (*View, error) {
	layoutFile, _ := GetLayoutFiles()
	files = append(files, layoutFile...)
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}
	return &View{
		tpl,
		layout,
	}, nil

}

// RenderView used to render the view with the predefined layout and the given data
func (v *View) RenderView(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("content-type", "text/html")
	err := v.Tpl.ExecuteTemplate(w, v.Layout, data)
	if err != nil {
		panic(err)
	}

}

// ServeHTTP is assigned to the view type ,so it can be used as an http.Handler
func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.RenderView(w, r, nil)
}

// GetLayoutFiles used to return every file in the layout directory
func GetLayoutFiles() ([]string, error) {
	layoutFiles, err := filepath.Glob("views/layout/" + "*" + ".gohtml")
	if err != nil {
		return nil, err
	}
	return layoutFiles, nil
}
