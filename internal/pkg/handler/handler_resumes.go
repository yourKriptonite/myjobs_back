package handler

import (
	. "2019_2_IBAT/internal/pkg/interfaces"
	"encoding/json"
	"log"

	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) CreateResume(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// log.Println("Create resume req:")
	// log.Println(r)
	authInfo, ok := FromContext(r.Context())

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		errJSON, _ := json.Marshal(Error{Message: UnauthorizedMsg})
		w.Write([]byte(errJSON))
		return
	}

	id, err := h.UserService.CreateResume(r.Body, authInfo)
	if err != nil {
		var code int
		switch err.Error() {
		case ForbiddenMsg:
			code = http.StatusForbidden
		case UnauthorizedMsg:
			code = http.StatusUnauthorized
		case InternalErrorMsg:
			code = http.StatusInternalServerError
		default:
			code = http.StatusBadRequest
		}
		w.WriteHeader(code)

		errJSON, _ := json.Marshal(Error{Message: err.Error()})
		w.Write([]byte(errJSON))
		return
	}

	idJSON, err := json.Marshal(Id{Id: id.String()})

	if err != nil {
		errJSON, _ := json.Marshal(Error{Message: err.Error()})
		w.Write([]byte(errJSON))
		return
	}
	log.Printf("Returning id: %s", id.String())

	w.Write([]byte(idJSON))
}

func (h *Handler) DeleteResume(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	authInfo, ok := FromContext(r.Context())

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		errJSON, _ := json.Marshal(Error{Message: UnauthorizedMsg})
		w.Write([]byte(errJSON))
		return
	}

	strId := mux.Vars(r)["id"]
	resId, err := uuid.Parse(strId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errJSON, _ := json.Marshal(Error{Message: InvalidIdMsg})
		w.Write([]byte(errJSON))
		return
	}
	err = h.UserService.DeleteResume(resId, authInfo)

	if err != nil {
		var code int
		switch err.Error() {
		case ForbiddenMsg:
			code = http.StatusForbidden
		case UnauthorizedMsg:
			code = http.StatusUnauthorized
		case InternalErrorMsg:
			code = http.StatusInternalServerError
		default:
			code = http.StatusBadRequest
		}
		w.WriteHeader(code)

		errJSON, _ := json.Marshal(Error{Message: err.Error()})
		w.Write([]byte(errJSON))
		return
	}
}

func (h *Handler) GetResume(w http.ResponseWriter, r *http.Request) { //+
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	log.Println("Handle GetResume: start")

	resId, err := uuid.Parse(mux.Vars(r)["id"])

	if err != nil {
		log.Println("Handle GetResume: invalid id")
		w.WriteHeader(http.StatusBadRequest)
		errJSON, _ := json.Marshal(Error{Message: InvalidIdMsg})
		w.Write([]byte(errJSON))
		return
	}

	resume, err := h.UserService.GetResume(resId)

	if err != nil {
		log.Println("Handle GetResume: failed to get resume")
		w.WriteHeader(http.StatusBadRequest)
		errJSON, _ := json.Marshal(Error{Message: InvalidIdMsg})
		w.Write([]byte(errJSON))
		return
	}

	resumeJSON, _ := json.Marshal(resume)

	w.Write([]byte(resumeJSON))
}

func (h *Handler) PutResume(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	defer r.Body.Close()
	authInfo, ok := FromContext(r.Context())

	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		errJSON, _ := json.Marshal(Error{Message: UnauthorizedMsg})
		w.Write([]byte(errJSON))
		return
	}

	resId, err := uuid.Parse(mux.Vars(r)["id"])

	err = h.UserService.PutResume(resId, r.Body, authInfo)

	if err != nil {
		var code int
		switch err.Error() {
		case ForbiddenMsg:
			code = http.StatusForbidden
		case UnauthorizedMsg:
			code = http.StatusUnauthorized
		case InternalErrorMsg:
			code = http.StatusInternalServerError
		default:
			code = http.StatusBadRequest
		}
		w.WriteHeader(code)

		errJSON, _ := json.Marshal(Error{Message: err.Error()})
		w.Write([]byte(errJSON))
		return
	}
}

func (h *Handler) GetResumes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	resumes, _ := h.UserService.GetResumes() //error handling

	resumesJSON, _ := json.Marshal(resumes)

	w.Write([]byte(resumesJSON))
}

func (h *Handler) ParseResumesQuery(query url.Values) map[string]interface{} {
	params := make(map[string]interface{})

	if query.Get("region") != "" {
		params["region"] = query.Get("region")
	}
	if query.Get("wage_from") != "" {
		params["wage_from"] = query.Get("wage_from")
	}
	if query.Get("wage_to") != "" {
		params["wage_to"] = query.Get("wage_to")
	}
	if query.Get("experience") != "" {
		params["experience"] = query.Get("experience")
	}
	if query.Get("type_of_employment") != "" {
		params["type_of_employment"] = query.Get("type_of_employment")
	}
	if query.Get("work_schedule") != "" {
		params["work_schedule"] = query.Get("work_schedule")
	}

	return params
}
