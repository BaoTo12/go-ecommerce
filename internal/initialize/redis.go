package initialize

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"time"

	"github.com/BaoTo12/go-ecommerce/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background() // Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline.

// TODO: The context package defines an interface that centralizes control and metadata for goroutines and function calls. Its key traits include:

// Cancellation: Signal work to stop gracefully.
// Timeouts and Deadlines: Automatically cancel work after a set time.
// Value propagation: Pass request-scoped data like request IDs or auth tokens.

/*
	type Context interface {
	Deadline() (deadline time.Time, ok bool) --> Return the time when the context expires.
	Done() <-chan struct{} --> A channel closed when the context is canceled or timeout is reached.
	Err() error --> Error indicating why context ended (Canceled or DeadlineExceeded).
	Value(key interface{}) interface{} --> Retrieve data stored in the context.
}
*/

// TODO: Root Contexts: Background() & TODO()
// context.Background(): The base, empty context—used at the top of call trees.
// context.TODO(): Placeholder when unsure which context to use, but be sure to replace later.

// TODO:  Child Contexts
// WithCancel(parent) → returns a cancelable child and a cancel function.
// WithDeadline(parent, time) → auto-cancels at a set time.
// WithTimeout(parent, duration) → convenience wrapper for WithDeadline.
// WithValue(parent, key, val) → embeds data in the context.

func openRedisWithRetry(maxElapsed time.Duration, initialDelay time.Duration) error {
	deadline := time.Now().Add(maxElapsed)
	attempt := 0
	for {
		attempt++
		// TODO: Get redis config
		redisConfig := global.Config.REDIS
		rdb := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%v", redisConfig.Host, redisConfig.Port),
			Password: redisConfig.Password,
			DB:       redisConfig.Database,
			PoolSize: redisConfig.PoolSize,
		})
		context, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		_, err := rdb.Ping(context).Result()
		if err == nil {
			fmt.Println("Redis Initialization Successfully")
			global.Rdb = rdb
			cancel()
			return nil
		}
		if time.Now().After(deadline) {
			cancel()
			return fmt.Errorf("could not connect to REDIS after %s: last error: %w", maxElapsed, err)
		}

		capDelay := 5 * time.Second
		backoff := float64(initialDelay) * math.Pow(2, float64(attempt-1))
		if backoff > float64(capDelay) {
			backoff = float64(capDelay)
		}
		jitter := (rand.Float64()*0.6 - 0.3) * backoff
		sleep := time.Duration(backoff + jitter)

		log.Printf("redis not ready (attempt %d): %v. retrying in %v", attempt, err, sleep)
		time.Sleep(sleep)
	}
}

func InitRedis() {
	// 6379
	// redisConfig := global.Config.REDIS
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     fmt.Sprintf("%s:%v", redisConfig.Host, redisConfig.Port),
	// 	Password: redisConfig.Password,
	// 	DB:       redisConfig.Database,
	// 	PoolSize: redisConfig.PoolSize,
	// })
	// _, err := rdb.Ping(ctx).Result()
	// if err != nil {
	// 	global.Logger.Error("Redis Initialization Error", zap.Error(err))
	// }
	// fmt.Println("Redis Initialization Successfully")
	// global.Rdb = rdb

	err := openRedisWithRetry(60*time.Second, 500*time.Millisecond)
	if err != nil {
		global.Logger.Error("Redis Initialization Error", zap.Error(err))
		panic(err)
	}

	// redisExample()
}
func redisExample() {
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Error redis setting:", zap.Error(err))
		return
	}

	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Error redis setting:", zap.Error(err))
		return
	}

	global.Logger.Info("value score is::", zap.String("score", value))
}
