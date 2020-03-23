package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rakhmadbudiono/compare-api-framework/golang/models"
	"github.com/rakhmadbudiono/compare-api-framework/golang/services"
)

type API struct {
	Service *services.BookService
	Router  *mux.Router
}

func NewAPI() *API {
	api := API{
		Service: services.NewBookService(),
		Router:  Router(),
	}

	api.Router.HandleFunc("", api.createBook).Methods("POST")
	api.Router.HandleFunc("", api.getBooks).Methods("GET")
	api.Router.HandleFunc("/{id}", api.getBookByID).Methods("GET")
	api.Router.HandleFunc("/{id}", api.updateBook).Methods("PUT")
	api.Router.HandleFunc("/{id}", api.deleteBook).Methods("DELETE")

	return &api
}

func Router() *mux.Router {
	r := mux.NewRouter()

	return r
}

// func (a *API) createBook(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)

// 	message := params["message"]

// 	err := a.Service.InsertMessage(message)
// 	if err != nil {
// 		handleError(w, NewErrorNoMessage(400))
// 		return
// 	}

// 	var data struct {
// 		Data struct {
// 			Message string `json:"message"`
// 			Status  string `json:"status"`
// 		} `json:"data"`
// 	}

// 	data.Data.Message = message
// 	data.Data.Status = "Success"

// 	handleJSONResponse(w, data)
// }

// func (a *API) getAllMessages(w http.ResponseWriter, r *http.Request) {
// 	messages := a.Service.GetAllMessages()

// 	var data struct {
// 		Data struct {
// 			Length   int `json:"length"`
// 			Messages []struct {
// 				MessageID string `json:"id_message"`
// 				Body      string `json:"body"`
// 			} `json:"messages"`
// 		} `json:"data"`
// 	}

// 	for _, v := range messages {
// 		var message struct {
// 			MessageID string `json:"id_message"`
// 			Body      string `json:"body"`
// 		}

// 		message.MessageID = v.GetID().String()
// 		message.Body = v.GetBody()

// 		data.Data.Messages = append(data.Data.Messages, message)
// 	}

// 	data.Data.Length = len(messages)

// 	handleJSONResponse(w, data)
// }


// wingman

func (a *API) createBook(res http.ResponseWriter, req *http.Request) {
	var b models.Book
	json.NewDecoder(req.Body).Decode(&b)

	bookID := a.Service.CreateBook(&b)

	json.NewEncoder(res).Encode(struct {
		ID int64 `json:"id"`
	}{ID: bookID})
}

func (a *API) getBooks(res http.ResponseWriter, req *http.Request) {
	books := a.Service.GetBooks()

	json.NewEncoder(res).Encode(books)
}

func (a *API) getBookByID(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	b := a.Service.GetBookByID(id)

	json.NewEncoder(res).Encode(b)
}

func (a *API) updateBook(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	b := &models.Book{ID: id}

	json.NewDecoder(req.Body).Decode(b)

	a.Service.UpdateBook(b)

	res.WriteHeader(204)
	res.Write([]byte{})
}

func (a *API) deleteBook(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	a.Service.DeleteBook(id)

	res.WriteHeader(204)
	res.Write([]byte{})
}