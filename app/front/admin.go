package front

import (
	"html/template"
	"net/http"
	"store/app/api/coupons"
	"store/app/api/products"
	"store/types"
	"store/util"

	"github.com/gorilla/csrf"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(util.GetTeamplte("admin")...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := types.Payload{
		Title:     "Admin",
		Prods:     products.GetRawProds(r.URL.Path, csrf.Token(r)),
		Coupons:   coupons.GetRawCopons(),
		Path:      r.URL.Path,
		IsAdmin:   true,
		CsrfField: csrf.TemplateField(r),
		CsrfToken: csrf.Token(r),
	}
	err = tmpl.Execute(w, data)

	if err != nil {
		panic(err)
	}
}
