package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func parseBody(r *http.Request,x interface{}){

	if body , err :=ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal(data[] byte,x); err == nill{
			return ;
		}
	}
}