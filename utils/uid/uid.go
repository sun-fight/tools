package uid

import (
	"github.com/bwmarrin/snowflake"
)

type SnowflakeNode int32

const (
	NodeDef SnowflakeNode = iota
)

var nodes map[SnowflakeNode]*snowflake.Node

func Init() {
	snowflake.NodeBits = 0  //节点数量
	snowflake.StepBits = 22 //每毫秒可产多少
	if snowflake.NodeBits+snowflake.StepBits > 22 {
		panic("snowflake NodeBits add StepBits must less than 22")
	}
	nodes = make(map[SnowflakeNode]*snowflake.Node)
	AddNode(NodeDef)
}

func SetNodeBits(bits uint8) {
	if bits > 22 {
		panic("snowflake NodeBits add StepBits must less than 22")
	}
	snowflake.NodeBits = bits
	snowflake.StepBits = 22 - bits
}
func SetStepBits(bits uint8) {
	if bits > 22 {
		panic("snowflake NodeBits add StepBits must less than 22")
	}
	snowflake.StepBits = bits
	snowflake.NodeBits = 22 - bits
}

func AddNode(nodeNum SnowflakeNode) {
	var err error
	newNode, err := snowflake.NewNode(int64(nodeNum))
	if err != nil {
		panic("add snowflake node " + err.Error())
	}
	_, ok := nodes[nodeNum]
	if ok {
		panic("repeat add snowflake node " + err.Error())
	}
	nodes[nodeNum] = newNode
}

func Gen64(nodeNum SnowflakeNode) int64 {
	return nodes[nodeNum].Generate().Int64()
}

func Gen(nodeNum SnowflakeNode) snowflake.ID {
	return nodes[nodeNum].Generate()
}

func Gen64Def() int64 {
	return nodes[NodeDef].Generate().Int64()
}
