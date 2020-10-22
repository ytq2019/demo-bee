package task

import (
	"bytes"
	"demo-bee/dto"
	"demo-bee/helper"
	"demo-bee/inc"
	"encoding/gob"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

func init() {
	//注册定时任务
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 */20 * * * *", Index)
	c.Start()
}

func Index() {
	bm := *helper.RedisAdater()
	redisKey := "indexVideoList"
	var videoList = &dto.VideoList{}
	dianying := inc.GetRankTopListByMobile("dianying", 1)
	dianying2 := inc.GetRankTopListByMobile("dianying", 2)
	for i := 0; i < len(dianying2); i++ {
		dianying = append(dianying, dianying2[i])
	}

	dianshi := inc.GetRankTopListByMobile("dianshi", 1)
	dianshi2 := inc.GetRankTopListByMobile("dianshi", 2)
	for i := 0; i < len(dianshi2); i++ {
		dianshi = append(dianshi, dianshi2[i])
	}
	zongyi := inc.GetRankTopListByMobile("zongyi", 1)
	zongyi2 := inc.GetRankTopListByMobile("zongyi", 1)
	for i := 0; i < len(zongyi2); i++ {
		zongyi = append(zongyi, zongyi2[i])
	}
	dongman := inc.GetRankTopListByMobile("dongman", 1)
	dongman2 := inc.GetRankTopListByMobile("dongman", 1)
	for i := 0; i < len(dongman2); i++ {
		dongman = append(dongman, dongman2[i])
	}

	videoList.Dianying = dianying
	videoList.Dianshi = dianshi
	videoList.Zongyi = zongyi
	videoList.Dongman = dongman
	//序列化
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer) //创建编码器
	err := encoder.Encode(videoList)   //编码
	if err != nil {
		log.Panic(err)
	}
	bm.Put(redisKey, buffer.String(), 1800*time.Second)

	log.Println("定时任务更新了一次首页数据")
}
