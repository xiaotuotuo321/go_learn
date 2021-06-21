package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

// redis

// 1.连接
//var rdb *redis.Client
//// 初始化连接
//func initClient() (err error){
//	rdb = redis.NewClient(&redis.Options{
//		Addr: "127.0.0.1:6379",
//		Password: "",
//		DB: 0,
//	})
//	_, err = rdb.Ping(nil).Result()
//	if err != nil{
//		return err
//	}
//	return nil
//}

// 2.v8新版本相关 最新的go-redis库的相关命令都需要传递context.Context参数
var (
	rdb *redis.Client
)

//初始化连接
func initClient() (err error){
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
		PoolSize: 100,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	return err
}

//func V8Example(){
//	ctx := context.Background()
//	if err := initClient(); err != nil{
//		return
//	}
//	// 插入数据
//	err := rdb.Set(ctx, "key", "value", 0).Err()
//	if err != nil{
//		panic(err)
//	}
//	// 读取数据
//	val, err := rdb.Get(ctx, "key").Result()
//	if err != nil{
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := rdb.Get(ctx, "key2").Result()
//	if err == redis.Nil{
//		fmt.Println("key2 does not exist")
//	} else if err != nil{
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}
//}
//
//func main() {
//	V8Example()
//}

// 3.连接redis哨兵模式
//func initClient()(err error){
//	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
//		MasterName: "master",
//		SentinelAddrs: []string{"x.x.x.x:26379", "x.x.x.x:26379", "x.x.x.x:26379"},
//	})
//	_, err = rdb.Ping(nil).Result()
//	if err != nil{
//		return err
//	}
//	return nil
//}

// 4.连接Redis集群
//func initClient() (err error) {
//	rdb := redis.NewClusterClient(&redis.ClusterOptions{
//		Addrs: []string{":7000", ":7001", ":7002", "7003", "7004", "7005"},
//	})
//
//	_, err = rdb.Ping(nil).Result()
//	if err != nil{
//		return err
//	}
//	return nil
//}

// 5.基本使用
// 5.1.set/get示例
//func redisExample() {
//	ctx := context.Background()
//
//	if err := initClient(); err != nil{
//		return
//	}
//
//	err := rdb.Set(ctx, "score", 100, 0).Err()
//	if err != nil{
//		fmt.Printf("set score failed, err:%v\n", err)
//		return
//	}
//
//	val, err := rdb.Get(ctx, "score").Result()
//	if err != nil{
//		fmt.Printf("get score failed, err:%v\n", err)
//		return
//	}
//	fmt.Println("score", val)
//
//	val2, err := rdb.Get(ctx, "name").Result()
//	if err == redis.Nil{
//		fmt.Println("name does not exist")
//	} else if err != nil{
//		fmt.Printf("get name failed, err:%v\n", err)
//		return
//	} else {
//		fmt.Println("name", val2)
//	}
//}

// 5.2.zset示例
//func redisExample2(){
//	ctx := context.Background()
//	if err := initClient(); err != nil{
//		return
//	}
//	zsetKey := "language_rank"
//	languages := []*redis.Z{
//		&redis.Z{Score: 90.0, Member: "Golang"},
//		&redis.Z{Score: 98.0, Member: "Java"},
//		&redis.Z{Score: 95.0, Member: "Python"},
//		&redis.Z{Score: 97.0, Member: "JavaScript"},
//		&redis.Z{Score: 99.0, Member: "C/C++"},
//	}
//	// ZADD
//	num, err := rdb.ZAdd(ctx, zsetKey, languages...).Result()
//	if err != nil{
//		fmt.Printf("zadd failed, err : %v\n", err)
//	}
//	fmt.Printf("zadd %d succ.\n", num)
//
//	// 把Golang的分数加10
//	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
//	if err != nil{
//		fmt.Printf("zincrby failed err:%v\n", err)
//		return
//	}
//	fmt.Printf("Golang's score is %f now.\n", newScore)
//
//	// 取分数最高的3个
//	ret, err := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Result()
//	if err != nil{
//		fmt.Printf("zrevrange failed, err:%v\n", err)
//		return
//	}
//	for _, z := range ret{
//		fmt.Println(z.Member, z.Score)
//	}
//
//	// 取95~100分的
//	op := &redis.ZRangeBy{
//		Min: "95",
//		Max: "100",
//	}
//	ret, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, op).Result()
//	if err != nil{
//		fmt.Printf("zrangebyscore failed, err: %v\n", err)
//		return
//	}
//	for _, z := range ret{
//		fmt.Println(z.Member, z.Score)
//	}
//}
//
//func main() {
//	redisExample2()
//}

// 5.3.根据前缀获取key
// vals, err := rdb.Keys(ctx, "prefix*").Result()

// 5.4.执行自定义的命令
// res, err := rdb.Do(ctx, "set", "key", "value").Result()

// 5.5.按照通配符删除key
// 当通配符匹配的key的数量不多时，可以使用keys()得到所有的key在使用Del命令删除。如果key的数量非常多的时候，我们可以搭配使用scan命令和del命令完成删除
//func exam(){
//	ctx := context.Background()
//	iter := rdb.Scan(ctx, 0, "prefix*", 0).Iterator()
//	for iter.Next(ctx){
//		err := rdb.Del(ctx, iter.Val()).Err()
//		if err != nil{
//			panic(err)
//		}
//	}
//	if err := iter.Err(); err != nil{
//		panic(err)
//	}
//}

// 5.6.pipeline的使用
// pipeline是一种网络优化，它本质上以为着客户端缓冲一堆命令并一次性将它们发送到服务器。这些命令不能保证在事务中执行，这样节省了每个命令网络往返时间
//func examPipeline(){
//	ctx := context.Background()
//	if err := initClient(); err != nil{
//		fmt.Println("初始化数据库连接失败")
//		return
//	}
//	pipe := rdb.Pipeline()
//	incr := pipe.Incr(ctx, "pipeline_counter")
//	pipe.Expire(ctx, "pipeline_counter", time.Hour)
//
//	_, err := pipe.Exec(ctx)
//	fmt.Println(incr.Val(), err)
//}
// 上面的命令相当于给Redis发送了以下命令
/*
incr pipeline_counter
expire pipeline_counter 3600
*/

// 也可以使用pipelined
//func examplePipelined(){
//	ctx := context.Background()
//	if err := initClient(); err != nil{
//		return
//	}
//	var incr *redis.IntCmd
//	_, err := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
//		incr = pipe.Incr(ctx, "pipelined_counter")
//		pipe.Expire(ctx, "pipeline_counter", time.Hour)
//		return nil
//	})
//	fmt.Println(incr.Val(), err)
//}
// 在某些场景下，当有多条命令要执行的时候，可以考虑使用pipeline来优化

// 6.事务
/*
redis是单线程的，因此单个命令始终是原子的，但是来自不停客户端的两个给定命令可以依次执行，例如他们需要交替执行，但是，multi/exec能够确保在multi\exec两个语句之间没有其他客户端正在执行命令
在这种场景下应该使用TxPipeline。和Pipeline类似，但是内部会使用multi/exec包裹排队的命令。
*/

//func exampleTxPipeline(){
//	ctx := context.Background()
//	if err := initClient(); err != nil{
//		return
//	}
//	pipe := rdb.TxPipeline()
//	incr := pipe.Incr(ctx, "tx_pipeline_counter")
//	pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
//
//	_, err := pipe.Exec(ctx)
//	fmt.Println(incr.Val(), err)
//}
/*
上面的命令相当于执行了：
multi
incr tx_pipeline_counter
expire tx_pipeline_counter 3600
exec
*/

// 6.1.2.还有一个与上文相类似的方法
//func exampleTxPipelineFunc(){
//	ctx := context.Background()
//	if err := initClient(); err != nil{
//		return
//	}
//	var incr *redis.IntCmd
//	_, err := rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
//		incr = pipe.Incr(ctx, "tx_pipelined_counter")
//		pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
//		return nil
//	})
//
//	fmt.Println(incr.Val(), err)
//}

// 7.watch 在某些场景下，除了使用multi/exec之外，还需要配合使用watch命令。在用户使用watch命令监控某个key之后，知道用户执行exec的这段时间里，
// 如果有其他用户抢先对被监控的键进行了替换、更新、删除等操作，那么当用户尝试执行exec的时候，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或
// 放弃事务。
// Watch(fn func(*Tx) error, keys ...string) error

//func exampleWatch(){
//	ctx := context.Background()
//	if err := initClient(); err != nil{
//		return
//	}
//	// 监控watch_count的值，并在值不变得前提下将其值+1
//	key := "watch_count"
//	err := rdb.Watch(ctx, func(tx *redis.Tx) error {
//		n, err := tx.Get(ctx, key).Int()
//		if err != nil && err != redis.Nil{
//			return err
//		}
//		_, err = tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
//			pipe.Set(ctx, key, n+1, 0)
//			return nil
//		})
//		return err
//	}, key)
//	if err != nil{
//		return
//	}
//}

// 8.官方版本中的GET和SET命令以事务方式递增key的值得示例，仅当key的值不发生改变的时候提交一个事务
func transactionDemo(){
	var (
		maxRetries = 1000
		routineCount = 10
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := initClient(); err != nil{
		return
	}
	defer cancel()

	// Increment 使用Get和Set命令以事务方式递增key的值
	increment := func(key string) error{
		// 事务函数
		txf := func(tx *redis.Tx) error {
			// 获得key的当前值或零值
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil{
				return err
			}
			// 实际的操作代码，乐观锁定中的本地操作
			n++
			// 操作仅在Watch的key没发生变化的情况下提交
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error{
				pipe.Set(ctx, key, n, 0)
				return nil
			})
			return err
		}

		// 最多重试 maxRetries次
		for i := 0; i < maxRetries; i++{
			err := rdb.Watch(ctx, txf, key)
			if err == nil{
				// 成功
				return nil
			}
			if err == redis.TxFailedErr{
				// 乐观锁丢失，重试
				continue
			}
			// 返回其他的错误
			return err
		}
		return errors.New("increment reached maximum number of retries")
	}

	// 模拟多个routineCount 个并发同时去修改 counter3的值
	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++{
		go func() {
			defer wg.Done()
			if err := increment("counter3"); err != nil{
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()
	n, err := rdb.Get(context.TODO(), "counter3").Int()
	fmt.Println("end with", n, err)
}

func main() {
	transactionDemo()
}
