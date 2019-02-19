package controllers

import (
	"encoding/json"
	"go-contacts/models"
	u "go-contacts/utils"
	"net/http"
)

var CreateTax = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
	tax := &models.Tax{}

	err := json.NewDecoder(r.Body).Decode(tax)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	tax.UserId = user
	resp := tax.Create()
	u.Respond(w, resp)
}

var GetTaxFor = func(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("user").(uint)
	data := models.GetTaxes(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetTribeFor = func(w http.ResponseWriter, r *http.Request) {

	//d := r.Context().Value("user").(uint)
	data := models.GetTaxes(0) // Only allowed to anon get tribe stats
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
var UpdateTax = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user").(uint)
	tax := &models.Tax{}

	tax.UserId = user
	err := json.NewDecoder(r.Body).Decode(tax)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := tax.Update(tax.Name, tax.Value)
	u.Respond(w, resp)

}
