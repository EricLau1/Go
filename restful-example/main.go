package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/emicklei/go-restful"
)

type User struct {
	Id   uint64
	Name string
}

var users = []User{
	User{Id: 1, Name: "John Doe"},
}

func NewService() *restful.WebService {
	service := new(restful.WebService)
	service.
		Path("/users").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	service.Route(service.GET("").To(GetAll))
	service.Route(service.GET("/{id}").To(Find))
	service.Route(service.POST("").To(Create))
	service.Route(service.PUT("/{id}").To(Update))
	service.Route(service.DELETE("/{id}").To(Delete))

	return service
}

func GetAll(req *restful.Request, res *restful.Response) {
	res.WriteEntity(users)
}

func Find(req *restful.Request, res *restful.Response) {
	id, err := strconv.ParseUint(req.PathParameter("id"), 10, 64)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}
	for _, u := range users {
		if u.Id == id {
			res.WriteEntity(u)
			return
		}
	}

	res.WriteError(http.StatusBadRequest, errors.New("User not found"))
}

func Create(req *restful.Request, res *restful.Response) {
	u := User{}
	err := req.ReadEntity(&u)
	if err != nil {
		res.WriteError(http.StatusUnprocessableEntity, err)
		return
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	u.Id = r.Uint64()
	users = append(users, u)
	res.Header().Set("Location", fmt.Sprintf("%s%s/%d", req.Request.Host, req.Request.RequestURI, u.Id))
	res.WriteHeaderAndEntity(http.StatusCreated, u)
}

func Update(req *restful.Request, res *restful.Response) {
	u := User{}
	err := req.ReadEntity(&u)
	if err != nil {
		res.WriteError(http.StatusUnprocessableEntity, err)
		return
	}
	id, err := strconv.ParseUint(req.PathParameter("id"), 10, 64)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}
	for i, _ := range users {
		if users[i].Id == id {
			users[i].Name = u.Name
			res.WriteEntity(users[i])
			return
		}
	}
	res.WriteError(http.StatusBadRequest, errors.New("User not found"))
}

func Delete(req *restful.Request, res *restful.Response) {
	id, err := strconv.ParseUint(req.PathParameter("id"), 10, 64)
	if err != nil {
		res.WriteError(http.StatusBadRequest, err)
		return
	}
	for i, _ := range users {
		if users[i].Id == id {
			users = append(users[:i], users[i+1:]...)
			res.Header().Set("Entity", fmt.Sprintf("%v", id))
			res.WriteHeaderAndEntity(http.StatusNoContent, "")
			return
		}
	}
	res.WriteError(http.StatusBadRequest, errors.New("User not found"))
}

func main() {
	port := ":8080"
	fmt.Println("Server on", port)
	restful.Add(NewService())
	log.Fatal(http.ListenAndServe(port, nil))
}
