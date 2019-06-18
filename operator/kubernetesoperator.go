package operator


type UserSubscribeService struct{
	Username     string
	Servicename  string
	Serviceurl   string
	Servicepath  string
}

func K8sCreateService(username string,servicename string,servicepath string)(UserSubscribeService,error){
		U := UserSubscribeService{username,servicename,"http://www.baidu.com",""}
		return U,nil
}

func K8sDeleteService()bool{
	return true
}