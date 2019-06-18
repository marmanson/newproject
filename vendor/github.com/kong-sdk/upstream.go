package kong_sdk

import (
	"gorequest"
	"fmt"
	"encoding/json"
)

type UpstreamRequest struct{
	Name                string   			 `json:"name"`
	Slots               int      			 `json:"slots,omitempty"`
	HashOn              string   			 `json:",omitempty"`
	HashFallback        string  		 	 `json:"hash_fallback,omitempty""`
	HashOnHeader        string   			 `json:"hash_on_header,omitempty"`
	HashFallbackHeader  string   			 `json:"hash_fallback_header,omitempty"`
	HashOnCookie        string    			 `json:"hash_on_cookie,omitempty"`
	HashOnCookiePath    string               `json:"hash_on_cookie_path,omitempty"`
	HealthChecks        *UpstreamHealthCheck `json:"health_checks,omitempty"`
}

type UpstreamHealthCheck struct{
	Active  *UpstreamHealthCheckActive  `json:"active,omitempty"`
	Passive *UpstreamHealthCheckPassive `json:"passive,omitempty"`
}

type UpstreamHealthCheckActive struct{
	Type                     string            `json:"type,omitempty"`
	Concurrency              int               `json:"concurrency,omitempty"`
	Healthy                  *ActiveHealthy    `json:"healthy,omitempty"`
	HttpPath                 string            `json:"http_path,omitempty"`
	HttpsVerifyCertificate   bool              `json:"https_verify_certificate"`
	HttpsSni                 *string           `json:"https_sni,omitempty"`
	Timeout                  int               `json:"timeout,omitempty"`
	Unhealthy                *ActiveUnhealthy  `json:"unhealthy,omitempty"`
}

type  ActiveHealthy  struct{
	HttpStatuses      []int  `json:"http_statuses,omitempty"`
	Interval          int    `json:"interval,omitempty"`
	Successes         int    `json:"successes,omitempty"`
}

type ActiveUnhealthy struct {
	HttpFailures int   `json:"http_failures,omitempty"`
	HttpStatuses []int `json:"http_statuses,omitempty"`
	Interval     int   `json:"interval,omitempty"`
	TcpFailures  int   `json:"tcp_failures,omitempty"`
	Timeouts     int   `json:"timeouts,omitempty"`
}

type UpstreamHealthCheckPassive struct {
	Type      string            `json:"type,omitempty"`
	Healthy   *PassiveHealthy   `json:"healthy,omitempty"`
	Unhealthy *PassiveUnhealthy `json:"unhealthy,omitempty"`
}

type PassiveHealthy struct {
	HttpStatuses []int `json:"http_statuses,omitempty"`
	Successes    int   `json:"successes,omitempty"`
}

type PassiveUnhealthy struct {
	HttpFailures int   `json:"http_failures,omitempty"`
	HttpStatuses []int `json:"http_statuses,omitempty"`
	TcpFailures  int   `json:"tcp_failures,omitempty"`
	Timeouts     int   `json:"timeouts,omitempty"`
}

type Upstream struct {
	Id string `json:"id,omitempty"`
	UpstreamRequest
}

type Upstreams struct {
	Results []*Upstream `json:"data,omitempty"`
	Next    string      `json:"next,omitempty"`
}

const UpstreamsPath = "/upstreams/"

func GetUpstreamByNameorId(endpoint string, idorname  string)(Upstream,error){
	U := Upstream{}
	r,body,errs := gorequest.New().Get(endpoint+UpstreamsPath+idorname).End()
	if errs != nil{
		return U,fmt.Errorf("could not get upstream,error:%v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return U,fmt.Errorf("not authorised,message from kong: %s",body)
	}
	u := Upstream{}
	err := json.Unmarshal([]byte(body),&u)
	if err != nil{
		return U,fmt.Errorf("could not parse upstream get response,error: %v",err)
	}

	if u.Id == ""{
		return U,nil
	}
	return u,nil
}

func CreateUpstream(upstreamRequest UpstreamRequest,endpoint string)(Upstream,error){
	U := Upstream{}
	r,body, errs := gorequest.New().Post(endpoint+UpstreamsPath).Send(upstreamRequest).End()
	if errs != nil{
		return U,fmt.Errorf("could not create new upstream,error : %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return U,fmt.Errorf("not authorised, message from kong:%s",body)
	}

	u := Upstream{}
	err := json.Unmarshal([]byte(body),&u)
	if err != nil{
		return U,fmt.Errorf("could not parse upstream creation response,error: %v",err)
	}

	if u.Id == ""{
		return U,fmt.Errorf("could not create update,error: %v",body)
	}
	return u,nil
}

func DeleteUpstreamByIdorName(endpoint string,idorname string) error{
	r,body,errs := gorequest.New().Delete(endpoint+UpstreamsPath+idorname).End()
	if errs != nil{
		return fmt.Errorf("could not delete upstream, result: %v error: %v",r,errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return fmt.Errorf("not authorised,message from kong: %s",body)
	}

	return nil
}

func UpstreamList(endpoint string)(*Upstreams,error){
	r,body,errs := gorequest.New().Get(endpoint+UpstreamsPath).End()
	if errs != nil{
		return nil,fmt.Errorf("could not get upstream,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return nil,fmt.Errorf("not authorised,messgae fron kong: %s",body)
	}
	upstreams := &Upstreams{}
	err := json.Unmarshal([]byte(body),upstreams)
	if err != nil{
		return nil,fmt.Errorf("could not parse upstreams list response,error: %v",err)
	}
	return upstreams,nil
}

func UpdateUpstreamByIdorName(uspstreamRequest UpstreamRequest,endpoint string,idorname string)(Upstream,error){
	U :=Upstream{}
	r,body,errs := gorequest.New().Patch(endpoint+UpstreamsPath+idorname).Send(uspstreamRequest).End()

	if errs != nil{
		return U,fmt.Errorf("could not update upstream,error: %v",errs)
	}

	if r.StatusCode == 401 || r.StatusCode == 403{
		return U,fmt.Errorf("not authorised,message from kong: %s",body)
	}

	u := Upstream{}
	err := json.Unmarshal([]byte(body),&u)
	if err != nil{
		return U,fmt.Errorf("could not parse upstream update reponse,error: %v",err)
	}
	if u.Id == ""{
		return U,fmt.Errorf("could not update upstream,error: %v",body)
	}
	return u,nil
}