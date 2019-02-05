package kafka

import (
	"github.com/Shopify/sarama"
	"fmt"
	"logcollect/tail"
)
var (
	kfkProductClient sarama.SyncProducer
)
func InitProduct() (err error) {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll//ack 确认服务器收到消息

	//随机分区，分片存储到多台服务器
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	config.Producer.Return.Successes = true//写成功配置

	//这里使用了同步客户端
	//还可以使用异步客户端
	kfkProductClient, err = sarama.NewSyncProducer([]string{"192.168.0.103:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	go startProduct()
	return

	//对于全局唯一，不需要频繁申请的资源，可以不关闭，
	//如果需要关闭，可以通过channel来处理
	//defer client.Close()


}
func startProduct()  {
	for{
		msgstr,topic:=tail.GetMsg()

		//构造一个消息实体
		msg := &sarama.ProducerMessage{}
		msg.Topic =topic
		msg.Value = sarama.StringEncoder(msgstr)


		//如果是同步客户端，发送完成后就成功写入到kafka了
		pid, offset, err := kfkProductClient.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}

		fmt.Printf("pid:%v offset:%v\n", pid, offset)
	}
}