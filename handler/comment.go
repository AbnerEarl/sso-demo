/**
 * @author: yangchangjia
 * @email 1320259466@qq.com
 * @date: 2024/4/8 14:26
 * @desc: about the role of class.
 */

package handler

import (
	"github.com/AbnerEarl/goutils/gins"
	"sso-demo/base"
)

// DocsComment
// @Summary 字段解释
// @Description 功能描述
// @Tags API接口管理
// @Schemes http https
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200 {object} gins.Response "{"code":0,"data":{},"msg":"请求成功"}"
// @Failure 200 {object} gins.Response "{"code":1,"data":{},"msg":"内部错误"}"
// @Failure 200 {object} gins.Response "{"code":2,"data":{},"msg":"业务错误"}"
// @Router /docs [get]
func DocsComment(c *gins.Context) {
	gins.SendResponse(c, nil, base.Comment)
}
