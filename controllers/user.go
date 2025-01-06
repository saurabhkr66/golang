package controllers

import(
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"

)
type UserController struct {
	session *mgo.Session
}
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id:= p.ByName("id")


if !bson.IsObjectIdHex(id){
	w.WriteHeader(http.StatusNotFound)
   
}
oid:=bson.ObjectIdHex((id))
u:=models.User

if err := uc.Session.DB("mongo-golang").C("users").Findid(oid){
w.WriteHeader(404)
return
}

uj, err:=json.Marshal(u)
if err!=nil{
	fmt.println(err)
}
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.statusOk)
fmt.Fprint(w,"%s\n",uj)



}

