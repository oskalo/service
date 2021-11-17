package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/oskalo/service/internal/handlers/comments"
	"github.com/oskalo/service/internal/handlers/product"
	"github.com/oskalo/service/internal/repositories"
	"github.com/oskalo/service/pkg/middlewares"
	"github.com/urfave/negroni"
	"log"
	"net/http"
)

func main() {
	//worst practice ever
	connStr := "sslmode=verify-full port=5432 user=user dbname=example password=test"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	repoComment := repositories.NewCommentRepository(db)
	repoProduct := repositories.NewProductRepository(db)
	hComment := comments.NewComment(repoComment, repoProduct)
	hProduct := product.NewProduct(repoProduct)

	//middlewares
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	rLog := negroni.NewLogger()
	rLog.SetFormat("[{{.Status}} {{.Duration}} {{.Method}}  {{.Path}}] - {{.Request.UserAgent}}")
	n.Use(rLog)

	r := mux.NewRouter()
	r.Use(middlewares.LogEndPointCalls)
	r.HandleFunc("products/{productID}/comments", hComment.AddComment).Methods(http.MethodPost, http.MethodGet)
	r.HandleFunc("products/{productID}", hProduct.GetProduct).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
