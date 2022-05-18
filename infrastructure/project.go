package infrastructure

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/yars-ai/yars/services/action"
	"gitlab.com/yars-ai/yars/services/engine"
)

type Project interface{}

func newProject(router *mux.Router, engine engine.Engine) Project {
	p := &project{
		router: router,
		engine: engine,
	}

	router.HandleFunc("/unit/{name}/upsert", p.upsertUnit)
	router.HandleFunc("/unit/{name}/create", p.createUnit)
	router.HandleFunc("/action/{name}/upsert", p.upsertAction)
	router.HandleFunc("/recommend/{name}/get-top-candidates", p.getTopCandidates)

	return p
}

type project struct {
	router *mux.Router
	engine engine.Engine
}

func (p *project) upsertUnit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	unit, err := p.engine.UnitByName(vars["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	ans, err := unit.Upsert(r.Context(), r.Body)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithReader(w, http.StatusOK, ans)
}

func (p *project) createUnit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	unit, err := p.engine.UnitByName(vars["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	ans, err := unit.Create(r.Context(), r.Body)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithReader(w, http.StatusOK, ans)
}

func (p *project) upsertAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	act, err := p.engine.ActionByName(vars["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	actions := []action.Action{}

	if err := json.NewDecoder(r.Body).Decode(&actions); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	ans, err := act.Upsert(r.Context(), actions)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithReader(w, http.StatusOK, ans)
}

type RecommendArgs struct {
	ActorID string `json:"id"`
	Count   int    `json:"count"`
}

func (p *project) getTopCandidates(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recommend, err := p.engine.RecommendByName(vars["name"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	rargs := []RecommendArgs{}

	if err := json.NewDecoder(r.Body).Decode(&rargs); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	ans := make([][]string, len(rargs))

	for i := range rargs {
		ans[i], err = recommend.GetTopRelevant(r.Context(), rargs[i].ActorID, rargs[i].Count)
	}

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ans)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithReader(w http.ResponseWriter, code int, payload io.Reader) {
	response, _ := ioutil.ReadAll(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
