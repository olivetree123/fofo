package main

import (
	"bytes"
	"encoding/json"
	. "fofo/common"
	"github.com/go-redis/redis"
	"github.com/mitchellh/mapstructure"
)

const CommandHeartbeat int = 0
const CommandRegisterService int = 100

// 命令
type Command struct {
	Code int
	Args map[string]interface{}
}

type RegisterServiceParam struct {
	Name    string
	Group   string
	Port    int
	Address string
	Extra   interface{}
}

func registerService(cmd Command) {
	var param RegisterServiceParam
	err := mapstructure.Decode(cmd.Args, &param)
	if err != nil {
		Logger.Error(err)
		return
	}

}

// 从 redis 队列中监听消息
func subRedis(redisSubChannel <-chan *redis.Message, cmdChannel chan Command, signalChannel chan bool) {
	for {
		msg, ok := <-redisSubChannel
		if !ok {
			Logger.Error("redis channel has been closed, break.")
			break
		}
		Logger.Info(msg.Channel, msg.Payload)
		buf := bytes.NewBufferString(msg.Payload)
		var cmd Command
		err := json.Unmarshal(buf.Bytes(), &cmd)
		if err != nil {
			Logger.Error(err)
			signalChannel <- true
			return
		}
		cmdChannel <- cmd
		Logger.Info(cmd)
	}
	Logger.Info("Finish to subscribe redis.")
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
	pubSub := redisClient.Subscribe("FOFO:REQUEST")
	defer pubSub.Close()
	signalChannel := make(chan bool)
	cmdChannel := make(chan Command)
	redisSubChannel := pubSub.Channel()
	go subRedis(redisSubChannel, cmdChannel, signalChannel)
	<-signalChannel
}
