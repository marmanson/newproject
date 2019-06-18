package kong_sdk

import (
	"gorequest"
	"fmt"
	"encoding/json"
)
type ServiceRequest struct{
	Name     string `json:"name"`
	//Protocol string `json:"protocol"`
	Url      string `json:"url,omitempty"`
	Path     string `json:"path,omitempty"`
}

type Service struct {
	Id             string `json:"id"`
	CreatedAt      int    `json:"created_at"`
	UpdatedAt      int    `json:"updated_at"`
	Protocol       string `json:"protocol"`
	Host           string `json:"host"`
	Port           int    `json:"port"`
	Path           string `json:"path"`
	Name           string `json:"name"`
	Retries        int    `json:"retries"`
	ConnectTimeout int    `json:"connect_timeout"`
	WriteTimeout   int    `json:"write_timeout"`
	ReadTimeout    int    `json:"read_timeout"`
	Url            string `json:"url"`
}

type Services struct{
	Data      []*Service  `json:"data"`
	Next      *string     `json:"next"`
}

const ServicePath = "/services/"

func CreateService(servicerequest ServiceRequest,adminhost string)(Service,error){
	S := Service{}
	r ,body,err := gorequest.New().Post(adminhost+ServicePath).Send(servicerequest).End()
	if err != nil{
		return S,fmt.Errorf("could not register the service: %v",err)
	}
	if r.StatusCode == 401 || r.StatusCode == 403{
		return S, fmt.Errorf("could not register the service: %s",body)
	}
	createService := Service{}
	errs := json.Unmarshal([]byte(body),&createService)
	if errs != nil{
		return S,fmt.Errorf("could not parse service get responce: %v",errs)
	}
	if createService.Id == ""{
		return S,fmt.Errorf("could not register the service: %v",body)
	}
	return createService,nil
}

func GetServiceByNameorId(endpoint string, idorname string)(Service,error){
	S := Service{}
	r,body,err := gorequest.New().Get(endpoint + ServicePath + idorname).End()
	if err != nil{
		return S, fmt.Errorf("cloud not get the service")
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return S, fmt.Errorf("not authorised")
	}

	service := Service{}

	errs := json.Unmarshal([]byte(body),&service)
	if errs != nil{
		return S,fmt.Errorf("could not parse service get reponse")
	}

	if service.Id == ""{
		return S,nil
	}
	return service,nil
}

func DeleteService(endpoint string,idorname string)(error){
	r,body,err := gorequest.New().Delete(endpoint + ServicePath + idorname).End()
	if err != nil{
		return fmt.Errorf("could not delete this service,result : %v error : %v",r,err)
	}
	if r.StatusCode == 401 || r.StatusCode == 403{
		return fmt.Errorf("not authorised, message from kong : %s",body)
	}

	return nil
}

func  GetServices(endpoint string)([]*Service,error){
	services := []*Service{}
	data := &Services{}

	for{
		r,body,errs := gorequest.New().Get(endpoint+ServicePath).End()
		if errs != nil{
			return nil,fmt.Errorf("could not get the service,error: %v",errs)
		}

		if r.StatusCode == 401 || r.StatusCode == 403{
			return nil,fmt.Errorf("not authorised,message from kong: %s",body)
		}

		err := json.Unmarshal([]byte(body),data)
		if err != nil{
			return nil,fmt.Errorf("could not parse service get response,error: %v",err)
		}

		services = append(services,data.Data...)

		if data.Next == nil || *data.Next == ""{
			break
		}
	}
	return services,nil
}

func  UpdateServiceByIdorName(serviceRequest ServiceRequest,endpoint string,idorname string)(Service,error){
	S := Service{}
	r,body,errs := gorequest.New().Patch(endpoint+ServicePath+idorname).Send(serviceRequest).End()
	if errs != nil{
		return S,fmt.Errorf("could not update service,error: %v",errs)
	}
	if r.StatusCode == 401 || r.StatusCode == 403{
		return S,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	s := Service{}
	err := json.Unmarshal([]byte(body),&s)
	if err != nil{
		return S,fmt.Errorf("could not parse service update response,error: %v",err)
	}

	if s.Id == ""{
		return S,fmt.Errorf("could not update service,error: %v",body)
	}

	return s,nil
}


