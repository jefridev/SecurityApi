package main

import "github.com/gorilla/mux"

import "github.com/jefridev/securityapi/users"
import "github.com/jefridev/securityapi/roles"

import "net/http"

import "github.com/jefridev/securityapi/common"

import "log"

import "github.com/jinzhu/gorm"

const defaultPort string = ":3002"

// Migrate enables creation of initial tables.
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.UserModel{})
}

func main() {
	db := common.Init()
	Migrate(db)

	r := mux.NewRouter()
	r.HandleFunc("/users/create", users.CreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/roles/create", roles.CreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/roles/setUser", roles.SetRoleToUserHandler).Methods(http.MethodPost)

	err := http.ListenAndServe(defaultPort, r)
	log.Fatal(err)
}
