package kong_sdk

import "gorequest"
import (
	"fmt"
	"encoding/json"
)
type TargetReuqest struct{
	Target      string  `json:"target"`
	Weight      int     `json:"weight"`
}

type Target struct{
	Id         string   `json:"id,omitempty"`
	CreateAt   float32  `json:"create_at"`
	Target     string   `json:"target"`
	Weight     int      `json:"weight"`
	Upstream   Id       `json:"upstream"`
	Health     string   `json:"health"`
}

type Targets struct{
	Data     []*Target  `json:"data"`
	Total 	 int        `json:"total,omitempty"`
	Next     string     `json:"next,omitempty"`
	NodeId   string     `json:"node_id,omitempty"`
}

const  TargetsPath = "/upstreams/%s/targets"

func  CreateTargetFromUpstreamNameorId(targetRequest TargetReuqest,endpoint string,nameorid string)(Target,error){
	T := Target{}
	r,body,errs := gorequest.New().Post(endpoint+fmt.Sprintf(TargetsPath,nameorid)).Send(targetRequest).End()
	if errs != nil{
		return T,fmt.Errorf("could not register the target,error: %v",errs)
	}
	if r.StatusCode == 401 || r.StatusCode == 403{
		return T,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	t := Target{}
	err := json.Unmarshal([]byte(body),&t)
	if err != nil{
		return T,fmt.Errorf("could not parse target get response,error: %v",err)
	}

	if t.Id == ""{
		return T,fmt.Errorf("could not register the target, error: %v",body)
	}
	return t,nil
}
func GetTargetFromUpstreamNameorId(endpoint string,idorname string)([]*Target,error){
	targets := []*Target{}
	data := &Targets{}

	for {
		r,body,errs := gorequest.New().Get(endpoint+fmt.Sprintf(TargetsPath,idorname)).End()
		if errs != nil{
			return nil,fmt.Errorf("could not get targets,error: %v",errs)
		}
		if r.StatusCode == 401 || r.StatusCode == 403{
			return nil,fmt.Errorf("not authorised,message from kong: %s",body)
		}
		if r.StatusCode == 404{
			return nil,fmt.Errorf("not existent upstream: %s",idorname)
		}

		err := json.Unmarshal([]byte(body),data)
		if err != nil{
			return nil,fmt.Errorf("could not parse target get response,error: %v",err)
		}
		targets = append(targets,data.Data...)
		if data.Next == ""{
			break
		}
	}
	return targets,nil
}

func DeleteTargetFromUpstreamByNameorHostPort(endpoint string,upstreamNameorId string,idorHostPort string)error{
	r,body,errs := gorequest.New().Delete(endpoint+fmt.Sprintf(TargetsPath,upstreamNameorId)+fmt.Sprintf("%s",idorHostPort)).End()
	if errs != nil{
		return fmt.Errorf("could not delete the target,result: %v error: %v",r,errs)
	}
	if r.StatusCode == 401 || r.StatusCode == 403{
		return fmt.Errorf("not authorised,message from kong: %s",body)
	}
	if r.StatusCode != 204{
		return fmt.Errorf("Received unexpected response status code: %d,Body: %s",r.StatusCode,body)
	}
	return nil
}

func SetTargetFromUpstreamByHostPortorIdAsHealthy(endpoint string,upstreamNameorId string,idorHostPort string)error{
	r,body,errs := gorequest.New().Post(endpoint+fmt.Sprintf(TargetsPath,upstreamNameorId)+fmt.Sprintf("/%s/healthy",idorHostPort)).Send("").End()
	if errs != nil{
		return fmt.Errorf("could not set the target as healthy,result: %v error: %v",r,errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return fmt.Errorf("not authorised, message from kong: %s",body)
	}

	if r.StatusCode != 204{
		return fmt.Errorf("Received unexpected response status code: %d, Body: %s",r.StatusCode,body)
	}
	return nil
}

func GetTargetsWithHealthFromUpstreamNameorId(endpoint string, idorname string)([]*Target,error){
	targets := []*Target{}
	data := &Targets{}

	for{
		r,body,errs := gorequest.New().Get(endpoint+fmt.Sprintf("/upstreams/%s/health",idorname)).End()
		if errs != nil{
			return nil,fmt.Errorf("could not get targets,error: %v",errs)
		}

		if r.StatusCode == 401 || r.StatusCode == 403{
			return nil,fmt.Errorf("not authorised,message from kong: %s",body)
		}

		if r.StatusCode == 404{
			return nil,fmt.Errorf("not existent upstream: %s",idorname)
		}

		err := json.Unmarshal([]byte(body),data)

		if err != nil{
			return nil,fmt.Errorf("could not parse target get response,error: %v",err)
		}
		targets = append(targets,data.Data...)

		if data.Next == ""{
			break
		}
	}
	return targets,nil
}