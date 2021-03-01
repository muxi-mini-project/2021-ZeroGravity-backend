package collection

import (
	"strconv"

	. "github.com/2021-ZeroGravity-backend/handler"
	"github.com/2021-ZeroGravity-backend/log"
	"github.com/2021-ZeroGravity-backend/pkg/errno"
	"github.com/2021-ZeroGravity-backend/service/idea"
	"github.com/2021-ZeroGravity-backend/util"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
// @Summary Get a user's favorite record table  
// @Description Get a user's favorite record table and return it to the front end 
// @Tags collection
// @Accept  json
// @Produce  json
// @Param userID head string true  "Get a user's favorite record table"
// @Success 200 {object} collection.IdeaResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /api/v1/collection [get]

// GetCollection ... 获取某个用户的收藏列表
func GetCollection(c *gin.Context) {
	log.Info("Idea getCollection function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	var id, limit, page int
	var err error

	// 从 token 获取 userid
	id = c.MustGet("userID").(int)

	limit, err = strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	page, err = strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, err.Error(), GetLine())
		return
	}

	// 调用服务
	item, err := idea.GetCollection(id, page*limit, limit)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	resp := &IdeaResponse{}
	resp.Count = len(item)
	resp.List = item

	SendResponse(c, errno.OK, resp)
}
