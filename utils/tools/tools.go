package tools

import (
	"github.com/bwmarrin/snowflake"
)

var IDGenerator *snowflake.Node

func InitGenerator() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	IDGenerator = node
	return
}

func GenerateIDInt64() int64 {
	return IDGenerator.Generate().Int64()
}
