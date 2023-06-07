package main

import (
	"flag"
	"fmt"
	"os"

	genratereport "absolutetech/farm_report/genrate-report"
	"absolutetech/farm_report/global-lib/envutils"
)

type flags struct {
	environment      envutils.Env
	structureLogDir  string
	structureLogFile string
	debug            string
}

func getFlags() *flags {
	flg := &flags{
		environment: envutils.Testing,
	}
	fs := flag.NewFlagSet("", flag.ExitOnError)
	envutils.SetFlag(fs, &flg.environment)
	fs.StringVar(&flg.structureLogDir, "structure-dir", flg.structureLogDir, "Structure Log Directory")
	fs.StringVar(&flg.structureLogFile, "structure-file", flg.structureLogFile, "Structure Log File")
	fs.StringVar(&flg.debug, "debug", flg.debug, "Debug address")
	_ = fs.Parse(os.Args[1:]) // Ignore error, because it exits on error
	return flg
}

func main() {
	fmt.Println("Farm Report Generation - Consumer")
	flg := getFlags()

	// amqpConnection, connectionErr := amqputils.NewAMQPConnection(flg.environment)
	// if connectionErr != nil {
	// 	panic(connectionErr)
	// }
	// defer amqpConnection.Close() //nolint: errcheck

	// ch, channelErr := amqputils.OpenAMQPChannel(amqpConnection)
	// if channelErr != nil {
	// 	panic(channelErr)
	// }
	// defer ch.Close() //nolint: errcheck

	// amqpQueue, queueErr := amqputils.DeclareAMQPQueue(ch, amqputils.QueueFarmReport)
	// if queueErr != nil {
	// 	panic(queueErr)
	// }

	// msgs, err := ch.Consume(
	// 	amqpQueue.Name,
	// 	"",
	// 	false,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	// if err != nil {
	// 	panic(channelErr)
	// }

	// // go func() {
	// for d := range msgs {
	// 	amqpMessage := &amqputils.AMQPMsg{}
	// 	err = json.Unmarshal(d.Body, amqpMessage)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("Recieved Message: %d %d %d\n", amqpMessage.FarmID, amqpMessage.TemplateID, amqpMessage.Retries)
	// 	err = d.Ack(false)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// // }()

	// fmt.Println("Successfully Connected to our RabbitMQ Instance")
	// fmt.Println(" [*] - Waiting for messages")
	// forever := make(chan bool)
	// <-forever

	// dbClient, dbClientErr := absmysql.ConnectIgrowlDB(context.Background(), flg.environment, "farm_report")
	// if dbClientErr != nil {
	// 	print(dbClientErr)
	// }

	farm_id := int64(54951)
	genratereport.GenerateReport(farm_id, flg.environment)
	// var value int
	// err := dbClient.QueryRow("select 1 from DUAL").Scan(&value)
	// if err != nil {
	// 	log.Println("query failed:", err)
	// 	return
	// }

	// log.Println("value:", value)
}
