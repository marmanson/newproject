package kong_sdk

import (
	"gorequest"
	"fmt"
	"encoding/json"
	)

type CertificateRequest struct{
	Cert     string   `json:"cert,omitempty"`
	Key      string   `json:"key,omitempty"`
}

type  Certificate  struct{
	Id        string  `json:"id,omitempty"`
	Cert      string  `json:"cert,omitempty"`
	Key       string  `json:"key,omitempty"`
}

type Certificates  struct{
	Results    []*Certificate  `json:"data,omitempty"`
	Total      int             `json:"total,omitempty"`
}

const CertificatesPath = "/certificates/"

func GetCertificateById(endpoint string,id string)(Certificate,error){
	C := Certificate{}
	r,body,errs := gorequest.New().Get(endpoint+CertificatesPath+id).End()
	if errs != nil{
		return C,fmt.Errorf("could not get certificate,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return C,fmt.Errorf("not authorised, message from kong: %s",body)
	}

	c := Certificate{}
	err := json.Unmarshal([]byte(body),&c)

	if err != nil{
		return C,fmt.Errorf("could not parse certificate get response,error: %v",err)
	}

	if c.Id == ""{
		return C,nil
	}
	return c,nil
}

func CreateCertificate(certificateRequest CertificateRequest,endpoint string)(Certificate,error){
	C := Certificate{}
	r, body, errs := gorequest.New().Post(endpoint+CertificatesPath).Send(certificateRequest).End()
	if errs != nil{
		return C,fmt.Errorf("could not create new certificate ,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return C,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	c := Certificate{}
	err := json.Unmarshal([]byte(body),&c)
	if err != nil{
		return c,fmt.Errorf("could not parse certificate creation response,error: %v",err)
	}

	if c.Id == ""{
		return C,fmt.Errorf("could note create certificate,error: %v",body)
	}
	return c,nil
}

func DeleteCertificateById(endpoint string,id string)error{
	r,body,errs := gorequest.New().Delete(endpoint+CertificatesPath+id).End()
	if errs != nil{
		return fmt.Errorf("could not delete certificate,result: %v error: %v",r,errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return fmt.Errorf("not authorised, message from kong: %s",body)
	}
	return nil
}

func CertificateList(endpoint string)(*Certificates,error){
	r,body,errs := gorequest.New().Get(endpoint+CertificatesPath).End()
	if errs != nil{
		return nil,fmt.Errorf("could not get certification,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return nil,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	certificates := &Certificates{}
	err := json.Unmarshal([]byte(body),certificates)
	if err != nil{
		return nil,fmt.Errorf("could not parse certificates list response,error: %v",err)
	}

	return certificates,nil
}

func UpdateCertificateById(certificateRequest CertificateRequest,endpoint string,id string)(Certificate,error){
	C := Certificate{}
	r,body,errs := gorequest.New().Patch(endpoint+CertificatesPath+id).Send(certificateRequest).End()
	if errs != nil{
		return C,fmt.Errorf("could not update certificate,error: %v",errs)
	}
	if r.StatusCode == 401 || r.StatusCode == 403{
		return C,fmt.Errorf("not authorised, message from kong: %s",body)
	}

	c := Certificate{}
	err := json.Unmarshal([]byte(body),&c)
	if err != nil{
		return C,fmt.Errorf("could not parse certificate update response,error: %v",err)
	}

	if c.Id == ""{
		return C,fmt.Errorf("could not update certificate,error: %v",body)
	}
	return c,nil
}
