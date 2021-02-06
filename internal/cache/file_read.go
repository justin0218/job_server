package cache

import (
	"github.com/parnurzeal/gorequest"
	"job_server/api"
)

type FileRead struct {
}

func (s *FileRead) GetFile(key string) (content string) {
	client := api.Rds.Get()
	res, err := client.Get(key).Result()
	if err != nil {
		return
	}
	content = res
	if content == "" {
		gorequest.New().Get(key)
	}
	//	key := this.GetString("key")
	//	res,err := initialize.RedisClient.Get(key).Result()
	//	if err == nil{
	//		errors.JsonOK(this.Ctx,res)
	//		return
	//	}
	//	response,err := http.Get(key)
	//	if err != nil{
	//		errors.JsonError(this.Ctx,errors.ErrBase.WithMessage(err.Error()))
	//		return
	//	}
	//	defer response.Body.Close()
	//	fdata,err := ioutil.ReadAll(response.Body)
	//	if err != nil{
	//		errors.JsonError(this.Ctx,errors.ErrBase.WithMessage(err.Error()))
	//		return
	//	}
	//	initialize.RedisClient.Set(key,fdata,-1)
	//	errors.JsonOK(this.Ctx,string(fdata))
	//	return
}
