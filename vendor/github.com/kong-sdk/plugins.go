package kong_sdk

import (
	"gorequest"
	"fmt"
	"encoding/json"
)

type PluginRequest struct{
	Name        string   				`json:"name"`
	ConsumerId  *Id      				`json:"consumer,omitempty"`
	ServiceId   *Id      				`json:"service,omitempty"`
	RouteId     *Id      				`json:"route,omitempty"`
	RunOn       string   				`json:"run_on,omitempty"`
	Config      map[string]interface{}	`json:"config,omitempty"`
}

type Plugin struct{
	Id           string					`json:"id"`
	Name         string					`json:"name"`
	ConsumerId   *Id					`json:"consumer,omitempty"`
	ServiceId    *Id					`json:"service,omitempty"`
	RouteId      *Id					`json:"route,omitempty"`
	RunOn        string					`json:"run_on,omitempty"`
	Config       map[string]interface{}	`json:"config,omitempty"`
	Enabled      bool					`json:"enabled,omitempty"`
}

type  Plugins struct{
	Results []*Plugin	`json:"data,omitempty"`
	Next    string		`json:"next,omitempty"`
}

const PluginsPath = "/plugins/"

func GetPluginById(endpoint string,id string)(Plugin,error){
	P := Plugin{}
	r,body,errs := gorequest.New().Get(endpoint+PluginsPath+id).End()
	if errs != nil{
		return P,fmt.Errorf("could not get plugin, error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return P,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	p := Plugin{}
	err := json.Unmarshal([]byte(body),&p)

	if err != nil{
		return P,fmt.Errorf("could not parse plugin response,error: %v",err)
	}
	if p.Id == ""{
		return P,nil
	}

	return p,nil
}

func PluginList(endpoint string)(*Plugins,error){
	r,body,errs := gorequest.New().Get(endpoint+PluginsPath).End()
	if errs != nil{
		return nil,fmt.Errorf("could not get plugins,error: %v",errs)
	}
	if r.StatusCode == 401 || r.StatusCode == 403{
		return nil,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	plugins := &Plugins{}
	err := json.Unmarshal([]byte(body),plugins)
	if err != nil{
		return nil,fmt.Errorf("could not parse plugins list response,error: %v",err)
	}
	return plugins,nil
}

func CreatePlugin(pluginRequest PluginRequest,endpoint string)(Plugin,error){
	P := Plugin{}
	r,body,errs := gorequest.New().Post(endpoint+PluginsPath).Send(pluginRequest).End()
	if errs != nil{
		return P,fmt.Errorf("could not create new plugin,error : %v",errs)
	}
	if r.StatusCode == 401 || r.StatusCode == 403{
		return P,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	p := Plugin{}
	err := json.Unmarshal([]byte(body),&p)

	if err != nil{
		return P,fmt.Errorf("could note parse plugin creation response,error: %v kong response: %s",err,body)
	}
	if p.Id == ""{
		return P,fmt.Errorf("could not create plugin,error: %v",body)
	}

	return p,nil
}

func UpdatePluginById(pluginRequest PluginRequest,endpoint string,id string)(Plugin,error){
	P := Plugin{}
	r,body,errs := gorequest.New().Patch(endpoint+PluginsPath+id).Send(pluginRequest).End()
	if errs != nil{
		return P,fmt.Errorf("coudl note update plugin,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return P,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	p := Plugin{}
	err := json.Unmarshal([]byte(body),p)
	if err != nil{
		return P,fmt.Errorf("could not parse plugin update response,error: %v kong response: %s",err,body)
	}

	if p.Id == ""{
		return P,fmt.Errorf("could not update plugin,error: %v",body)
	}
	return p,nil
}

func DeletePluginById(endpoint string,id string)error{
	r,body,errs := gorequest.New().Delete(endpoint+PluginsPath+id).End()
	if errs != nil{
		return fmt.Errorf("could not delete plugin,result: %v error: %v",r,errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return fmt.Errorf("not authorised,message from kong: %s",body)
	}
	return nil
}
