package sharding

type ShardingRouter struct {
	ShardingCount int
}

func NewShardingRouter(shardingCount int) *ShardingRouter {
	return &ShardingRouter{
		ShardingCount: shardingCount,
	}
}

func (s *ShardingRouter) GetShard(orderID int) string {
	switch orderID % s.ShardingCount {
	case 0:
		return "shard_1"
	case 1:
		return "shard_2"
	default:
		return "shard_3"
	}
}
