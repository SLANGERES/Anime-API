package utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//this function help to parse the data from the http body into it
func ParseBody(r*http.Request,x interface{}){

	//reding the http requ data and in store in byte format and check if there is no error err==nil if no error then proceed
	if body, err:=ioutil.ReadAll(r.Body); err==nil{

		//now convert the data bytes and store in x name pointed interface 
		if err:=json.Unmarshal([]byte(body), x); err!=nil{
			return
		}
	}
}