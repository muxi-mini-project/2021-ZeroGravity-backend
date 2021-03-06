package idea

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
// @Summary 获取想法  
// @Description 根据用户ID获取这个ID下发布的想法
// @Tags idea
// @Accept  json
// @Produce  json
// @Param id path int true "id--根据这个id判断发布的想法是否为私密，id为1是显示想法为私密发布，你无权查看"
// @Param token header string true  "uid--用户ID"
// @Success 200 {object}  model.IdeaListItem "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Failure 400 {object} errno.Errno
// @Failure 404 {object} errno.Errno
// @Failure 500 {object} errno.Errno
// @Router /api/v1/idea/detail/:id [get]
// GetIdea ... 获取单个想法
func GetIdea(c *gin.Context) {
	log.Info("Idea getIdeaList function called.",
		zap.String("X-Request-Id", util.GetReqID(c)))

	// 获取 id 后面需要根据这个 id 判断是否为私密能获取
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		SendBadRequest(c, errno.ErrPathParam, nil, err.Error(), GetLine())
		return
	}

	uid := c.MustGet("userID").(int)

	// 调用服务
	item, err := idea.GetIdea(id, uid)
	if err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error(), GetLine())
		return
	}

	SendResponse(c, errno.OK, item)
}
