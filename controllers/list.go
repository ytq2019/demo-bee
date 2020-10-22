/*
 * @Time : 2020/2/23
 * @Author : tianqi.yu
 * @Software : GoLand
 * @Description :
 */
package controllers

import (
	"bytes"
	"demo-bee/dto"
	"demo-bee/helper"
	"demo-bee/inc"
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"strconv"
	"time"
)

type ListController struct {
	beego.Controller
}

func (c *ListController) Get() {

	var (
		bm        = *helper.RedisAdater()
		videoList = []*dto.KankanVideo{}
		//listResp  = &dto.ListResp{}
		page  = c.Ctx.Input.Param(":page")
		vType = c.Ctx.Input.Param(":type")
	)

	//拼接缓存key name
	redisKey := fmt.Sprintf("list-%s-%s", vType, page)
	//查看是否存在缓存
	//不存在则调用函数请求该页数据并存缓存
	pageNo, _ := strconv.Atoi(page)
	if bm.IsExist(redisKey) { //存在
		videoListTmp := &[]*dto.KankanVideo{}
		redisValue := bm.Get(redisKey)
		str := redisValue.([]byte)
		decoder := gob.NewDecoder(bytes.NewReader(str))
		decoder.Decode(videoListTmp)
		videoList = *videoListTmp
		log.Println("读取了redis", redisKey)
	} else { //不存在
		videoList = inc.GetRankTopListByMobile(vType, pageNo) //[]*dto.KankanVideo
		var buffer bytes.Buffer
		encoder := gob.NewEncoder(&buffer) //创建编码器
		err := encoder.Encode(videoList)   //编码
		if err != nil {
			log.Panic(err)
		}
		bm.Put(redisKey, buffer.String(), 3600*time.Second)
		log.Println("获取了数据并写入redis", redisKey)

	}
	//listResp.Data = videoList
	c.Data["json"] = videoList
	c.ServeJSON()
}
