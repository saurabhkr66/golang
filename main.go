package main

import(
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"net/http"


)

func main(){

	r:= httprouter.New()
	uc:=controllers.NewController(getSession())

	r.GET("/user/:id",uc.GetUser)
	r.POST("/user",uc.CreateUser)
    r.DELETE("user/:id",uc.DeleteUser)

	http.ListenAndServe("http://localhost:9000",r);
}
func getSession() *mgo.Session{
	s,err:=mgo.Dial("mongodb+srv://sauravkum420:tLmcT4yumtsf941v@cluster0.jxadh.mongodb.net/")
    if err!=nil{
        panic(err)
    }
    return s
}