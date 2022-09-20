package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/zorlovskiy/phonebook/internal/models"
	"github.com/zorlovskiy/phonebook/internal/usecase"
)

type contactHandler struct {
	uc usecase.Contact
}

/*func NewcontactHandler(uc usecase.Contact) *contactHandler {
	return &contactHandler{
		uc: uc,
	}
}*/

func Router(uc usecase.Contact) *mux.Router {

	c := &contactHandler{uc}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/contact", c.CreateHandler).Methods("POST")
	router.HandleFunc("/contact/update/{id}", c.UpdateHandler).Methods("PUT")
	router.HandleFunc("/contact/{fname}", c.GetByFNameHandler).Methods("GET")
	router.HandleFunc("/contact/bynumber/{number}", c.GetByNumberHandler).Methods("GET")
	router.HandleFunc("/contact/{id}", c.DeleteByIDHandler).Methods("DELETE")

	return router
}

func (h *contactHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.uc.Create(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *contactHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var updateContact models.Contact

	err := json.NewDecoder(r.Body).Decode(&updateContact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err = h.uc.Update(&updateContact); err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
}

func (h *contactHandler) GetByFNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nameParams := vars["fname"]

	//type GetByFNameResponse struct {
	//	Contacts []models.Contact `json:"contacts"`
	//}

	//var response GetByFNameResponse

	response, err := h.uc.GetByFName(nameParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *contactHandler) GetByNumberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numbParams := vars["number"]

	response, err := h.uc.GetByNumber(numbParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *contactHandler) DeleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idDelete := vars["id"]

	err := h.uc.Delete(idDelete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

}
