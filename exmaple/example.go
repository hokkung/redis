package example

import (
	"context"
	"time"

	v9 "github.com/hokkung/redis/v9"
)

func main() {
	ctx := context.Background()
	rds := v9.NewRedisConnection()
	_, _ = rds.Set(ctx, "", "", time.Second*2).Result()
}
