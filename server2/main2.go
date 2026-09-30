package main
import(
	"fmt"
	"net/http"
	"encoding/json"
)

type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

func user(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")

	switch r.Method{
	case http.MethodGet:
		user:=User{
			Id: 1,
			Name: "aditya",
		}

		json.NewEncoder(w).Encode(user)
	
	case http.MethodPost:
		var user User
		err:=json.NewDecoder(r.Body).Decode(&user)
		if err!=nil{
			http.Error(w,"InvalidJson",http.StatusBadRequest)
			return
		}
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(user)
	
	default:
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

func main(){

	mux:=http.NewServeMux()

	mux.HandleFunc("/user",user)

	server:=&http.Server{
		Addr: ":8000",
		Handler: mux,
	}
	server.ListenAndServe()
	fmt.Println("server is live at port 8000")
}