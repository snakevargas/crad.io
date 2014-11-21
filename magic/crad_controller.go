package crad

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CradController struct {
	Crads map[string]*Crad
	Cmcs  map[float64][]*Crad
}

func (cc *CradController) Cmc(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	cmcString := ps.ByName("cmc")

	cmc, err := strconv.ParseFloat(cmcString, 64)


	crads, ok := cc.Cmcs[cmc]

	if !ok {
		http.Error(w, "Does not exist", http.StatusNotFound)
	}

	js, err := json.Marshal(crads)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
