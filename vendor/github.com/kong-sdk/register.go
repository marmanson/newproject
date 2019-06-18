package kong_sdk

//import (
//	"math/rand"
//	"time"
//)
//
//const endpoint = "http://localhost:8001"
//const repoint = "http://localhost:8000/"
//func GetRandomString(l int) string {
//	str := "0123456789abcdefghijklmnopqrstuvwxyz"
//	bytes := []byte(str)
//	result := []byte{}
//	r := rand.New(rand.NewSource(time.Now().UnixNano()))
//	for i := 0; i < l; i++ {
//		result = append(result, bytes[r.Intn(len(bytes))])
//	}
//	return string(result)
//}

//func Register(name string, url string, path string)error{
//	service := ServiceRequest{name,url,path}
//	s,err :=service.CreateService(endpoint)
//	if err != nil{
//		log.Fatal(err)
//	}
//	//fmt.Println("123123123")
//	random := GetRandomString(20)
//	h := []string{random+s.Name}
//	route := RouteRequest{random+s.Name,h}
//	r ,err1 := route.CreateRoute(endpoint,s.Name)
//	if err1 != nil{
//		log.Fatal(err1)
//		return nil
//	}
//	//fmt.Println("123123")
//	subscirbe := SubscribeTable{"wangbo",s.Name,r.Name,random}
//	subscirbe.DbInsert()
//	return nil
//}
//
//func DeleteServiceandRoute(username string,servicename string,serviceroute string){
//	subscribe := SubscribeTable{username,"",serviceroute,""}
//	err := DeleteRoute(endpoint,serviceroute)
//	if err != nil{
//		fmt.Println(err)
//	}
//	err1 := DeleteService(endpoint,servicename)
//	if err != nil{
//		fmt.Println(err1)
//	}
//	subscribe.DbDelete()
//}
//
//func RequestService(username string, serviceroute string,token string)(string,error){
//	//subscribe := SubscribeTable{"","",serviceroute,""}
//	//s := subscribe.DbQuery()
//	//if(s.Username != username){
//	//	return nil,fmt.Errorf("you can not request this service")
//	//}
//	r,body,err := gorequest.New().Get(repoint).AppendHeader("Host",serviceroute).End()
//	if err != nil{
//		return "",fmt.Errorf("cloud not get the service")
//	}
//
//	if r.StatusCode == 401 || r.StatusCode == 403{
//		return "",fmt.Errorf("not authorised")
//	}
//
//	//var service map[string]interface{}
//	//errs := json.Unmarshal([]byte(body),&service)
//	//if errs != nil{
//	//	return "",fmt.Errorf("could not parse service get reponse")
//	//}
//
//	return body,nil
//}

