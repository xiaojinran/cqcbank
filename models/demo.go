package models

type Demo struct {
	Name string `json:"name"`
	Value int `json:"value"`
}

func NewDemo(name string,value int) *Demo{
   return &Demo{
   	Name:name,
   	Value:value,
   }
}

