package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/httplib"
)

type RequestController struct {
	beego.Controller
}

type RequestResult struct {
	Error string  `json:"err"`
	Result interface{} `json:"result"`
}

func RequestServer(name string) *RequestResult{
	iniconf, err := config.NewConfig("ini","requests/"+name+".ini")
	if err != nil{
	  return &RequestResult{
	  	Error:err.Error(),
	  	Result:"",
	  }
	}
	
    url := iniconf.String("url")
	req := httplib.Get(url)
	str, err := req.String()
	if err != nil{
		return &RequestResult{
			Error:err.Error(),
			Result:"",
		}
	}
	
	return  &RequestResult{
		Error:"no errors",
		Result: str,
	}
}


