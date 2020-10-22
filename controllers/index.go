package controllers

import (
	"bytes"
	"demo-bee/dto"
	"demo-bee/helper"
	"demo-bee/task"
	"encoding/gob"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"log"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	bm := *helper.RedisAdater()
	var (
		redisKey  = "indexVideoList"
		videoList = &dto.VideoList{}
	)

breakHere:
	if !bm.IsExist(redisKey) {
		task.Index()
		goto breakHere
	}
	redisValue := bm.Get(redisKey)

	str := redisValue.([]byte)
	decoder := gob.NewDecoder(bytes.NewReader(str))
	err2 := decoder.Decode(videoList)
	if err2 != nil {
		log.Panic(err2)
	}

	c.Data["json"] = videoList
	c.ServeJSON()
}
