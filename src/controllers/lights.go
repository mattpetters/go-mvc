package controllers

import (
  "net/http"
  "viewmodels"
  "text/template"
)

type lightsController struct {
  template *template.Template
}

func (this *lightsController) get(w http.ResponseWriter, req *http.Request){
    vm := viewmodels.GetLights()

    w.Header().Add("Content Type", "text/html")
    this.template.Execute(w, vm)
}
