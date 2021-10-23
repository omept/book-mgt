package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ong-gtp/book-mgt/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {

	sb := router.PathPrefix("/api/").Subrouter()
	sb.Use(headerMiddleware)

	sb.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	sb.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	sb.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	sb.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	sb.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	sb.HandleFunc("/", controllers.NotFound).Methods("DELETE", "PUT", "GET", "PATCH", "POST")

}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
