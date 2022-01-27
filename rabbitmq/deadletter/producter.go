package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"godemo/rabbitmq/deadletter/constant"
	"godemo/rabbitmq/util"
	"strconv"
	"time"
)

func main() {
	// # ========== 1.创建连接 ==========
	mq := util.NewRabbitMQ()
	defer mq.Close()
	mqCh := mq.Channel

	// # ========== 2.设置队列（队列、交换机、绑定） ==========
	// 声明队列
	var err error
	_, err = mqCh.QueueDeclare(constant.NormalQueue, true, false, false, false, amqp.Table{
		"x-message-ttl": 60000, // 消息过期时间（队列级别）,毫秒
		"x-dead-letter-exchange": constant.DeadExchange, // 指定死信交换机
		"x-dead-letter-routing-key": constant.DeadRoutingKey, // 指定死信routing-key
	})
	util.FailOnError(err, "创建normal队列失败")

	// 声明交换机
	err = mqCh.ExchangeDeclare(constant.NormalExchange, amqp.ExchangeDirect, true, false, false, false, nil)
	util.FailOnError(err, "创建normal交换机失败")

	// 队列绑定（将队列、routing-key、交换机三者绑定到一起）
	err = mqCh.QueueBind(constant.NormalQueue, constant.NormalRoutingKey, constant.NormalExchange, false, nil)
	util.FailOnError(err, "normal：队列、交换机、routing-key 绑定失败")

	// # ========== 3.设置死信队列（队列、交换机、绑定） ==========
	// 声明死信队列
	// args 为 nil。切记不要给死信队列设置消息过期时间,否则失效的消息进入死信队列后会再次过期。
	_, err = mqCh.QueueDeclare(constant.DeadQueue, true, false, false, false, nil)
	util.FailOnError(err, "创建dead队列失败")

	// 声明交换机
	err = mqCh.ExchangeDeclare(constant.DeadExchange, amqp.ExchangeDirect, true, false, false, false, nil)
	util.FailOnError(err, "创建dead队列失败")

	// 队列绑定（将队列、routing-key、交换机三者绑定到一起）
	err = mqCh.QueueBind(constant.DeadQueue, constant.DeadRoutingKey, constant.DeadExchange, false, nil)
	util.FailOnError(err, "dead：队列、交换机、routing-key 绑定失败")

	// # ========== 4.发布消息 ==========
	message := "msg" + strconv.Itoa(int(time.Now().Unix()))
	fmt.Println(message)

	// 发布消息
	err = mqCh.Publish(constant.NormalExchange, constant.NormalRoutingKey, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
		Expiration: "10000", // 消息过期时间,毫秒
	})
	util.FailOnError(err, "消息发布失败")
}