package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	r "reflect"
	"strconv"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

type Usuario struct{
	ID int `json:"id"`
	Nome string `json:"nome"`
}
func UsuarioHandler(w http.ResponseWriter, r *http.Request){
	sid := strings.TrinPrefix(r.URL.Path, '/usuarios/')
	id, _ := strconv.Atoi(sid)

	switch{
		case r.Method = "GET" && id > 0:
			usuarioPorID(w,r, id)
		case r.Method == "GET":
			usuarioTodos(w, r)
	  default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Desculpa :(")
	}
}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int){
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	var u Usuario

	db.QueryRow("select id, nome from usuarios where id =?", id).Scan(&u.ID, &u.Nome)
	json, _ := json.Marshal(u)
	w.Header().Set("Contet-Type", "application/json")
	fmt.Fprint(w, string(json))

}

func usuarioTodos(w http.ResponseWriter, r *http.Request){
	db
}
