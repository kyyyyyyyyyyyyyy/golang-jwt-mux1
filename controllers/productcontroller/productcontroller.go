package productcontroller

import (
	"net/http"

	"github.com/kyyyyyyyyyyyyyy/golang-jwt-mux1/helpers"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id":    1,
			"name":  "jersey",
			"stock": 120,
		},
		{
			"id":    2,
			"name":  "polo shirt",
			"stock": 340,
		},
		{
			"id":    3,
			"name":  "T shirt",
			"stock": 299,
		},
	}

	helpers.ResponseJSON(w, http.StatusOK, data)

}
