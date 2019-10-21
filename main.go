package main

import (
	"bytes"
	"encoding/json"
	. "fofo/common"
	"fofo/entity"
	"fofo/handlers"
	"github.com/go-redis/redis"
	"time"
)

var IsTerminated = false

// response 发送返回结果
func response(client *redis.Client, responseChannel chan *entity.Response) {
	for {
		resp := <-responseChannel
		if resp == nil {
			Logger.Info("response terminated")
			return
		}
		content, err := json.Marshal(resp.Content)
		if err != nil {
			Logger.Error(err)
			continue
		}
		r := client.LPush(resp.RequestID, content)
		err = r.Err()
		if err != nil {
			Logger.Error(err)
			continue
		}
	}
}

// subRedis 从 redis 队列中监听消息
func subRedis(client *redis.Client, signalChannel chan bool, responseChannel chan *entity.Response) {
	for {
		msg := client.BRPop(0, RequestQueue)
		if msg.Err() != nil {
			Logger.Error("redis channel has been closed, break.")
			signalChannel <- true
			return
		}
		res, err := msg.Result()
		if err != nil {
			Logger.Error(err)
			signalChannel <- true
			return
		}
		data := bytes.NewBufferString(res[1])
		var cmd entity.Command
		err = json.Unmarshal(data.Bytes(), &cmd)
		if err != nil {
			Logger.Error(err)
			signalChannel <- true
			return
		}
		err = handlers.CommandHandler(&cmd, responseChannel)
		if err != nil {
			Logger.Error(err)
			signalChannel <- true
			return
		}
	}
}

// healthCheck 健康检查
func healthCheck(client *redis.Client) {
	duration := time.Duration(time.Second * 5)
	t := time.NewTicker(duration)
	for {
		<-t.C
		if IsTerminated {
			Logger.Info("IsTerminated, healthCheck exit.")
			return
		}
		c := client.Publish(HealthCheckChannel, "HealthCheck")
		_, err := c.Result()
		if err != nil {
			Logger.Error(err)
			return
		}
	}
}

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		Logger.Error(err)
		return
	}
	defer redisClient.Close()
	Logger.Info("Start ...")
	signalChannel := make(chan bool)
	responseChannel := make(chan *entity.Response)
	go subRedis(redisClient, signalChannel, responseChannel)
	go response(redisClient, responseChannel)
	// 每隔 5 秒广播一个健康检查的命令
	go healthCheck(redisClient)
	<-signalChannel
	responseChannel <- nil
	IsTerminated = true
}
