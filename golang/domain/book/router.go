package book

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type router struct {
	service Service
}

func (r *router) createBook(res http.ResponseWriter, req *http.Request) {
	var b Book
	json.NewDecoder(req.Body).Decode(&b)

	bookID := r.service.CreateBook(&b)

	json.NewEncoder(res).Encode(struct {
		ID int64 `json:"id_book"`
	}{ID: bookID})
}

func (r *router) getBooks(res http.ResponseWriter, req *http.Request) {
	books := r.service.GetBooks()

	json.NewEncoder(res).Encode(books)
}

func (r *router) getBookByID(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	b := r.service.GetBookByID(id)

	json.NewEncoder(res).Encode(b)
}

func (r *router) updateBook(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	b := &Book{ID: id}

	json.NewDecoder(req.Body).Decode(b)

	r.service.UpdateBook(b)

	res.WriteHeader(204)
	res.Write([]byte{})
}

func (r *router) deleteBook(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	r.service.DeleteBook(id)

	res.WriteHeader(204)
	res.Write([]byte{})
}

// SetupRouter will setub the SubRouter given as argumemnt
func SetupRouter(s *mux.Router) {
	r := &router{}

	s.HandleFunc("", r.createBook).Methods("POST")
	s.HandleFunc("", r.getBooks).Methods("GET")
	s.HandleFunc("/{id}", r.getBookByID).Methods("GET")
	s.HandleFunc("/{id}", r.updateBook).Methods("PUT")
	s.HandleFunc("/{id}", r.deleteBook).Methods("DELETE")
}
