package main

import (
	"context"
	"github.com/go-redis/redis/v8"
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
func exam(){
	ctx := context.Background()
	iter := rdb.Scan(ctx, 0, "prefix*", 0).Iterator()
	for iter.Next(ctx){
		err := rdb.Del(ctx, iter.Val()).Err()
		if err != nil{
			panic(err)
		}
	}
	if err := iter.Err(); err != nil{
		panic(err)
	}
}
