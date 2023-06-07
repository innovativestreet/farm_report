package main

import (
	"flag"
	"fmt"
	"os"

	"absolutetech/farm_report/global-lib/amqputils"
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
	fmt.Println("Farm Report Generation")
	flg := getFlags()

	amqpConnection, connectionErr := amqputils.NewAMQPConnection(flg.environment)
	if connectionErr != nil {
		panic(connectionErr)
	}
	defer amqpConnection.Close() //nolint: errcheck

	channel, channelErr := amqputils.OpenAMQPChannel(amqpConnection)
	if channelErr != nil {
		panic(channelErr)
	}
	defer channel.Close() //nolint: errcheck

	amqpQueue, queueErr := amqputils.DeclareAMQPQueue(channel, amqputils.QueueFarmReport)
	if queueErr != nil {
		panic(queueErr)
	}

	data := &amqputils.AMQPMsg{
		FarmID: 100,
	}

	messageSendingErr := amqputils.SendMessage(channel, amqpQueue, data)
	if messageSendingErr != nil {
		panic(queueErr)
	}

	// dbClient, closeDBClient, err := mysqldb.NewIgrowDBClient(flg.environment, "farm_reprt")
	// if err != nil {
	// 	print(closeDBClient)
	// }

	// var value int
	// err = dbClient.QueryRow("select 1 from DUAL").Scan(&value)
	// if err != nil {
	// 	log.Println("query failed:", err)
	// 	return
	// }

	// log.Println("value:", value)
}
