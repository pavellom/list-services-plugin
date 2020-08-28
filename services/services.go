package services

import (
    "encoding/json"
    "fmt"
//     "../terminal"
)

type Response struct {
    Pagination Pagination `json:"pagination"`
    Resources []Resource `json:"resources"`
}

type Pagination struct {
    TotalResults int `json:"total_results"`
    TotalPages int `json:"total_pages"`
}

type Resource struct {
    Data Data `json:"data"`
    Links Links `json:"links"`
}

type Data struct {
    Name string `json:"instance_name"`
}

type Links struct {
    Instance Link `json:"service_instance"`
}

type Link struct {
    Href string `json:"href"`
}

func ParseResponse(response string) Response {
    res := &Response{}
    err := json.Unmarshal([]byte(response), res)
    if (err != nil) {
        fmt.Println(err)
    }
    return *res
}