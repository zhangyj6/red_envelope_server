package routers

import (
	"red_envelop_server/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func LoadOpen(e *gin.Engine) {
	e.POST("/open", OpenHandler)
}

func OpenHandler(c *gin.Context) {
	uid, _ := c.GetPostForm("uid")
	envelope_id, _ := c.GetPostForm("envelope_id")

	logs.Printf("envelope %d opened by %d", envelope_id, uid)

	int_uid, _ := strconv.ParseInt(uid, 10, 64)
	int_envelope_id, _ := strconv.ParseInt(envelope_id, 10, 64)

	//开红包的业务反应到程序中是
	//1、找到envelope_id对应的红包，用uid和它检验一下（安全性之一）
	//2、是opened = true,
	//3、如果这个红包已经开过了，判断一下直接返回就行
	//因为其它的值在正常来说在创建红包时应该已经确定了，预热数据库，把除了uid之外的有关envelope的值全部确定
	envelope := sql.GetEnvelopeByEevelopeID(int_envelope_id)

	flag := true

	//需要判断一下查的envelope_id不存在,因为envelop_id是从1开始自增长的
	if envelope.ID == int64(0) {
		flag = false
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "envelope_id does not exist",
			"data": gin.H{
				"envelope_id": envelope_id,
			},
		})
		return
	}

	//用传入的uid和envelope查询到的uid值检验一下
	if envelope.UID != int_uid {
		flag = false
		c.JSON(200, gin.H{
			"code": -2,
			"msg":  "Inconsistent uid",
			"data": gin.H{
				"uid": uid,
			},
		})
		return
	}

	//如果这个红包已经开过了，判断一下直接返回就行
	if envelope.Opened == true {
		flag = false
		c.JSON(200, gin.H{
			"code": -3,
			"msg":  "The red envelope has already been opened",
			"data": gin.H{
				"opened": envelope.Opened,
			},
		})
		return
	}

	//打开一个未打开的红包，更改红包的状态为opened，返回value
	if flag {

		sql.UpdateState(int_envelope_id)
		value := envelope.Value

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"value": value,
			},
		})
	}
}
