package kong_sdk

import (
	"gorequest"
	"fmt"
	"encoding/json"
)

type SnisRequest struct{
	Name           string   `json:"name,omitempty"`
	CertificateId  Id       `json:"certificate,omitempty"`
}

type Sni struct{
	Name           string    `json:"name,omitempty"`
	CertificateId  Id        `json:"certificate,omitempty"`
}

type Snis struct{
	Results []*Sni   `json:"data,omitempty"`
	Total   int      `json:"total,omitempty"`
}

const SnisPath = "/snis/"

func CreateSnis(snisRequest SnisRequest,endpoint string)(Sni,error){
	S := Sni{}
	r,body,errs := gorequest.New().Post(endpoint+SnisPath).Send(snisRequest).End()
	if errs != nil{
		return S,fmt.Errorf("could not create new sni, error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return S,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	s := Sni{}
	err := json.Unmarshal([]byte(body),&s)
	if err != nil{
		return S,fmt.Errorf("could not parse sni creation response,error: %v",err)
	}

	if s.CertificateId == ""{
		return S,fmt.Errorf("could not create sni,error : %v",body)
	}
	return s,nil
}

func GetSnisByName(endpoint string,name string)(Sni,error){
	S := Sni{}
	r,body,errs := gorequest.New().Get(endpoint+SnisPath+name).End()
	if errs != nil{
		return S,fmt.Errorf("could not get sni,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return S,fmt.Errorf("not authorised,messge from kong: %s",body)
	}

	s := Sni{}
	err := json.Unmarshal([]byte(body),&s)
	if err != nil{
		return s,fmt.Errorf("could not parse sni get response,error: %v",err)
	}
	if s.Name == ""{
		return S,nil
	}
	return s,nil
}

func SniList(endpoint string)(*Snis,error){
	r,body,errs := gorequest.New().Get(endpoint+SnisPath).End()
	if errs != nil{
		return nil,fmt.Errorf("could note get snis,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return nil,fmt.Errorf("could not parse snis parse list response, error: %v",body)
	}
	s := &Snis{}
	err := json.Unmarshal([]byte(body),s)
	if err != nil{
		return nil,fmt.Errorf("could not parse snis list response,error: %v",err)
	}
	return s,nil
}

func DeleteSnisByName(endpoint string,name string)error{
	r,body,errs := gorequest.New().Delete(endpoint+SnisPath+name).End()
	if errs != nil{
		return fmt.Errorf("could note delte sni,result: %v error: %v",r,errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return fmt.Errorf("not authorised, message from kong: %s",body)
	}
	return nil
}

func UpdateSniByName(snisRequest SnisRequest,endpoint string,name string)(Sni,error){
	S := Sni{}
	r,body,errs := gorequest.New().Patch(endpoint+SnisPath+name).Send(snisRequest).End()
	if errs != nil{
		return S,fmt.Errorf("could note update sni,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return S,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	s := Sni{}
	err := json.Unmarshal([]byte(body),&s)
	if err != nil{
		return S,fmt.Errorf("could not parse sni update response,error: %s",err)
	}

	if s.CertificateId == ""{
		return S,fmt.Errorf("could not update sni,error: %v",body)
	}
	return s,nil
}