package base

import (
	"github.com/ouqiang/gocron/internal/models"
	"github.com/ouqiang/gocron/internal/modules/utils"
	"gopkg.in/macaron.v1"
)

// ParsePageAndPageSize 解析查询参数中的页数和每页数量
func ParsePageAndPageSize(ctx *macaron.Context, params models.CommonMap) {
	page := ctx.QueryInt("page")
	pageSize := ctx.QueryInt("page_size")
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = models.PageSize
	}

	params["Page"] = page
	params["PageSize"] = pageSize
}

func Ping(ctx *macaron.Context) string {
	json := utils.JsonResponse{}
	return json.Success(utils.SuccessContent, map[string]interface{}{
		"data": "pong",
	})
}
