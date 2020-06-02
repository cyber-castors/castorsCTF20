package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Yes function to respond to request Yes
func (cc Controller) Yes(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "%s", "Well that sucks, the admin hasn't given me any more tasks.")
	return
}

//No funcion to respond to request No
func (cc Controller) No(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "%s", "Well that's all folks!")
	return
}
