package utils

import (
	"html/template"
	"net/http"
)


func PaerseTemplate(w http.ResponseWriter,
	r *http.Request,
	files []string,
	tplName string,
	data interface{},
	){
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	err = tpl.ExecuteTemplate(w, tplName,data)
	if err != nil {
		panic(err)
	}
}
