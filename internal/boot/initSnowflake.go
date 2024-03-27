package boot

import (
	"context"
	"github.com/bwmarrin/snowflake"
)

var GID *snowflake.Node

func InitSnowflake(ctx context.Context) error {
	Node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}
	GID = Node
	return nil
}
