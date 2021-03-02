package job

import (
	"context"
	"job_server/api/mall_server"
)

func AutoCloseOrder() {
	mallServer := mall_server.GetClient()
	mallServer.AutoCloseOrder(context.Background(), &mall_server.AutoCloseOrderReq{})
}
