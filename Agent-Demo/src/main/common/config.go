package common

type config struct {
	RECEIVE_CMD_TOPIC string
}

var Config = &config{
	RECEIVE_CMD_TOPIC: "receive_cmd_topic",
}
