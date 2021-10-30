package routers

import (
	"math/rand"
	"red_envelop_server/sql"
	"strconv"

	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
)

func LoadSnatch(e *gin.Engine) {
	e.POST("/snatch", SnatchHandler)
}

func SnatchHandler(c *gin.Context) {

	//每个人能抢的最大红包数应该从配置文件读进来
	max_count := 5
	//总红包数
	total_envelope_num := 1000
	cur_envelope_num := 0

	uid, _ := c.GetPostForm("uid")
	logs.Printf("%s is snatching envelope", uid)

	int_uid, _ := strconv.ParseInt(uid, 10, 64)

	//根据uid查询用户，没有的话就创建用户
	user := sql.GetUser(int_uid)

	flag := true

	//判断用户的count是否大于个人最多抢红包数
	if user.Count >= int64(max_count) {
		flag = false
		c.JSON(200, gin.H{
			"code": -1,
			"msg":  "fail! snatch too much",
			"data": gin.H{
				"max_count": max_count,
				"cur_count": user.Count,
			},
		})
		return
	}

	//判断当前的发放的总红包数是否超过了最大的红包数
	if cur_envelope_num >= total_envelope_num {
		flag = false
		c.JSON(200, gin.H{
			"code": -2,
			"msg":  "fail! More than total number of red envelope",
			"data": gin.H{
				"total_envelope_num": total_envelope_num,
				"cur_envelope_num":   cur_envelope_num,
			},
		})
		return
	}

	//根据概率计算用户这次应不应该拿到红包，这里我想的是对所有的请求做统一的处理，直接放弃一部分请求不处理，
	//这样既满足了概率也减轻了后端的压力,只处理十分之一的请求
	rand_num := rand.Intn(10)
	if rand_num != 0 {
		flag = false
		c.JSON(200, gin.H{
			"code": -3,
			"msg":  "According to the probability, the red envelope can not be snatched this time",
		})
		return
	}

	if flag {
		//如果上述的条件都满足了，并且概率正好也轮到了，为当前用户生成红包
		envelope := sql.CreateEnvelope(user)

		//每发出一个红包
		cur_envelope_num++

		//修改当前用户的count
		sql.UpdateCount(&user)

		c.JSON(200, gin.H{
			"code": 0,
			"msg":  "success",
			"data": gin.H{
				"envelope_id": envelope.ID,
				"max_count":   max_count,
				"cur_count":   user.Count,
			},
		})
	}
}
