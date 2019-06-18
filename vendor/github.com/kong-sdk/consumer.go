package kong_sdk

import (
	"gorequest"
	"fmt"
	"encoding/json"
)

type ConsumerRequest struct{
	Username string `json:"username,omitempty"`
	CustomId string `json:"custom_id,omitempty"`
}

type Consumer struct{
	Id        string `json:"id,omitempty"`
	CustomId  string `json:"custom_id,omitempty"`
	Username  string `json:"username,omitempty"`
}

type Consumers struct{
	Results []*Consumer `json:"data,omitempty"`
	Next    string      `json:"next,omitempty"`
}

type ConsumerPluginConfig struct{
	Id      string `json:"id,omitempty"`
	Body    string
}

const ConsumerPath = "/consumers/"

func GetConsumerByIdorName(endpoint string,idorname string)(Consumer,error){
	c := Consumer{}
	r,body,errs := gorequest.New().Get(endpoint+ConsumerPath+idorname).End()
	if errs != nil{
		return c,fmt.Errorf("could not get consumer,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return c,fmt.Errorf("not authorised, message from kong: %s",body)
	}
	consumer := Consumer{}

	err := json.Unmarshal([]byte(body),&consumer)

	if err != nil{
		return c, fmt.Errorf("could not parse consumer get response, error: %v",err)
	}

	if consumer.Id == ""{
		return c,nil
	}

	return consumer,nil
}

func CreateConsumer(consumerequest ConsumerRequest,endpoint string)(Consumer,error){
	c := Consumer{}
	r,body,errs := gorequest.New().Post(endpoint+ConsumerPath).Send(consumerequest).End()
	if errs != nil{
		return c,fmt.Errorf("could not create new consumer,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return c, fmt.Errorf("not authorised, message from kong : %s",body)
	}

	createConsumer := Consumer{}
	err := json.Unmarshal([]byte(body),&createConsumer)
	if err != nil{
		return c,fmt.Errorf("could not parse consumer creation response, error : %v",err)
	}

	if createConsumer.Id == ""{
		return c,fmt.Errorf("could not create consumer,error: %v",body)
	}
	return createConsumer,nil
}

func DeleteConsumerByIdorName(endpoint string,idorname string)error{
	r,body,errs := gorequest.New().Delete(endpoint+ConsumerPath+idorname).End()
	if errs != nil{
		return fmt.Errorf("could not delete consumer,result: %v error : %v",r,errs)
	}
	if r.StatusCode == 401 || r.StatusCode == 403{
		return fmt.Errorf("not authorised, message from kong: %s",body)
	}
	return nil
}

func ConsumerList(endpoint string)(*Consumers,error){
	r,body,errs := gorequest.New().Get(endpoint+CertificatesPath).End()
	if errs != nil{
		return nil,fmt.Errorf("could not get consumers,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return nil,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	c := &Consumers{}
	err := json.Unmarshal([]byte(body),c)
	if err != nil{
		return nil,fmt.Errorf("could not parse consumers list response,error: %v",err)
	}
	return c,nil
}

func UpdateConsumerByIdorName(consumerRequest ConsumerRequest,endpoint string,idorname string)(Consumer,error){
	C := Consumer{}
	r,body,errs := gorequest.New().Patch(endpoint+CertificatesPath+idorname).Send(consumerRequest).End()
	if errs != nil{
		return C,fmt.Errorf("could not update consumer,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return C,fmt.Errorf("not authorised,message from kong: %s",body)
	}
	c := Consumer{}
	err := json.Unmarshal([]byte(body),&c)
	if err != nil{
		return C,fmt.Errorf("could note parse consumer update response,error: %v",err)
	}

	if c.Id == ""{
		return C,fmt.Errorf("could note update consumer,error: %v",body)
	}
	return c,nil
}
