package kong_sdk

import (
	"gorequest"
	"fmt"
	"encoding/json"
)

type RouteRequest struct{
	Name       string `json:"name"`
	Hosts      []string `json:"hosts"`
	//service    string `json:"service"`
}

type Route struct{
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	CreatedAt     int       `json:"created_at"`
	UpdatedAt     int       `json:"updated_at"`
	Protocols     []string  `json:"protocols"`
	Methods       []string  `json:"methods"`
	Hosts         []string  `json:"hosts"`
	Paths         []string  `json:"paths"`
	RegexPriority int       `json:"regex_priority"`
	StripPath     bool      `json:"strip_path"`
	PreserveHost  bool      `json:"preserve_host"`
	Snis          []string  `json:"snis"`
	Sources       []IpPort  `json:"sources"`
	Destinations  []IpPort  `json:"destinations"`
	Service       Id        `json:"service"`
}

type IpPort struct{
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type Routes struct{
	Data []*Route  `json:"data"`
	Total int      `json:"total"`
	Next  string   `json:"next"`
}

const RoutesPath  = "/routes/"

func  CreateRoute(routerequest RouteRequest,adminhost string,serviceid string)(Route,error){
	R := Route{}
	r,body,err := gorequest.New().Post(adminhost +ServicePath+serviceid+RoutesPath).Send(routerequest).End()
	if err != nil{
		return R,fmt.Errorf("could not register the route, error")
	}
	if r.StatusCode == 401 || r.StatusCode == 403 {
		return R,fmt.Errorf("not authorised")
	}
	createroute := Route{}
	errs := json.Unmarshal([]byte(body),&createroute)

	if errs != nil{
		return R, fmt.Errorf("could not parse route get response : %v",errs)
	}

	if createroute.Id == ""{
		return R,fmt.Errorf("could not register the route")
	}

	return createroute,nil
}

func GetRouteByIdorName(endpoint string,idorname string)(Route,error){
	R := Route{}
	r,body,errs := gorequest.New().Get(endpoint + RoutesPath + idorname).End()
	if errs != nil{
		return R,fmt.Errorf("could not get the route, error : %v",errs)
	}
	if r.StatusCode == 401 || r.StatusCode == 403{
		return R, fmt.Errorf("not authorised,message from kong : %s",body)
	}
	route := Route{}
	err := json.Unmarshal([]byte(body),&route)
	 if err != nil{
	 	return R, fmt.Errorf("could not parse route get reponse,error : %v",err)
	 }

	if route.Id == ""{
		return R,nil
	}
	return route,nil
}

func GetRoutesFromServiceNameorId(endpoint string, idorname string)([]*Route,error){
	routes := []*Route{}
	data := &Routes{}
	for {
		r, body, errs := gorequest.New().Get(endpoint + ServicePath + idorname + RoutesPath).End()
		if errs != nil {
			return nil, fmt.Errorf("could not get the route, error : %v", errs)
		}

		if r.StatusCode == 401 || r.StatusCode == 403 {
			return nil, fmt.Errorf("not authorised, message from kong: %s", body)
		}
		err := json.Unmarshal([]byte(body), data)
		if err != nil {
			return nil, fmt.Errorf("could not parse route get response, error : %", err)
		}
		routes = append(routes,data.Data...)
		if data.Next == ""{
			break
		}
	}
	return routes,nil
}

func DeleteRoute(endpoint string, idorname string) error{
	r,body,errs := gorequest.New().Delete(endpoint + RoutesPath + idorname).End()

	if errs != nil{
		return fmt.Errorf("could not delete the route, result: %v error: %v",r,errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return fmt.Errorf("not authorised, message from kong : %s",body)
	}
	return nil
}

func RouteList(endpoint string)([]*Route,error){
	routes := []*Route{}
	data := &Routes{}

	for{
		r,body,errs := gorequest.New().Get(endpoint+RoutesPath).End()
		if errs != nil{
			return nil,fmt.Errorf("could not get the route,error: %v",errs)
		}

		if r.StatusCode == 401 || r.StatusCode == 403{
			return nil,fmt.Errorf("not authorised,message from kong: %s",body)
		}

		err := json.Unmarshal([]byte(body),data)
		if err != nil{
			return nil,fmt.Errorf("could not parse route get response,error: %v",err)
		}

		routes = append(routes,data.Data...)
		if data.Next == ""{
			break
		}
	}
	return routes,nil
}

