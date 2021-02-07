package job

import (
	"context"
	"fmt"
	"job_server/api/user_server"
	"job_server/api/wechat_server"
	"job_server/store"
	"time"
)

var log = new(store.Log)

func BillNotice() {
	wechatServer := wechat_server.GetClient()
	userServer := user_server.GetClient()
	users := []int{1, 4, 5}
	for _, uid := range users {
		uinfo, err := userServer.ClientGetUserByUid(context.Background(), &user_server.ClientGetUserByUidReq{
			Uid: int64(uid),
		})
		if err != nil {
			log.Get().Error("bill notice err:%v", err)
			continue
		}
		data := make(map[string]*wechat_server.TemplateItem)
		data["first"] = &wechat_server.TemplateItem{Value: "您好，今天是您预约的记账提醒日，请记得记账！"}
		data["keyword1"] = &wechat_server.TemplateItem{Value: time.Now().Format("2006-01-02 15:04:05")}
		data["keyword2"] = &wechat_server.TemplateItem{Value: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")}
		data["keyword3"] = &wechat_server.TemplateItem{Value: "每天"}
		data["remark"] = &wechat_server.TemplateItem{Value: fmt.Sprintf("%s，今天是您预约的记账提醒日，请点击记账！", uinfo.Nickname)}
		wechatServer.SendTemplate(context.Background(), &wechat_server.SendTemplateReq{
			TemplateDefine: wechat_server.TemplateDefine_bill_notice,
			Account:        wechat_server.Account_momo_za_huo_pu,
			Template: &wechat_server.Template{
				Touser: uinfo.Openid,
				Url:    "http://admin.momoman.cn/accountmanage/create",
				Data:   data,
			},
		})
	}
}
