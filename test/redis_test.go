package test

import (
	"encoding/json"
	"fmt"
	"github.com/hwUltra/fb-tools/redisx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"testing"
)

type UserInfo struct {
	Uid      int64  `json:"uid"`
	UserName string `json:"username"`
}

func TestRedisBuilder(t *testing.T) {

	redisClient, err := redis.NewRedis(redis.RedisConf{Host: "192.168.3.88:6379", Type: "node"})

	//userInfo := map[string]any{
	//	"name": "kyle",
	//	"age":  10,
	//}
	userInfo := UserInfo{1, "kyle"}
	bytes, _ := json.Marshal(userInfo)
	err = redisClient.Setex("user_1", string(bytes), 1200)
	if err != nil {
		return
	}

	exists, err := redisClient.Exists("user_1")
	if err != nil {
		fmt.Println("exists err:", err)
	}
	if exists {
		infos, err := redisClient.Get("user_1")
		if err == nil {
			//var result map[string]any
			//if err = json.Unmarshal([]byte(infos), &result); err != nil {
			//	fmt.Println("exists err:", err)
			//}
			//fmt.Println("result:", result, result["age"], result["name"])
			var result UserInfo
			if err = json.Unmarshal([]byte(infos), &result); err != nil {
				fmt.Println("exists err:", err)
			}
			fmt.Println("result:", result, result.Uid, result.UserName)

		}

	}

}

func TestRedisDelBuilder(t *testing.T) {

	redisClient, err := redis.NewRedis(redis.RedisConf{Host: "127.0.0.1:6379", Type: "node"})
	if err != nil {
		fmt.Println("err:", err)
	}
	del, err := redisClient.Del("user_1")
	if err != nil {
		return
	}
	fmt.Println(del)
}

func TestRedisSetBuilder(t *testing.T) {
	redisClient, err := redis.NewRedis(redis.RedisConf{Host: "127.0.0.1:6379", Type: "node"})
	if err != nil {
		fmt.Println("err:", err)
	}
	onLine, err := redisClient.Sadd("onLine", "1", "2", "3")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("onLine:", onLine)

	sismember, err := redisClient.Sismember("onLine", "1")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("onLine:", sismember)

	//redisClient.Srem("onLine", "1")

}

func TestRedisZSetBuilder(t *testing.T) {
	redisClient, err := redis.NewRedis(redis.RedisConf{Host: "192.168.3.88:116379", Type: "node"})
	if err != nil {
		fmt.Println("err:", err)
	}

	key := "string:zset"
	pairs := []redis.Pair{
		{Score: 80, Key: "Java"},
		{Score: 90, Key: "Python"},
		{Score: 95, Key: "Golang"},
		{Score: 98, Key: "PHP"},
	}
	redisClient.Zadds(key, pairs...)
	if err != nil {
		fmt.Println(err)
	}

	incr, err := redisClient.Zincrby(key, 29, "Java")
	if err != nil {
		return
	}
	fmt.Println("incr", incr)

	strings, err := redisClient.ZrevrangebyscoreWithScores(key, 0, 1000)
	if err != nil {
		return
	}
	fmt.Println("strings", strings)

}

func TestGeoAdd(t *testing.T) {
	redisClient, err := redis.NewRedis(redis.RedisConf{Host: "192.168.3.88:116379", Type: "node"})
	if err != nil {
		fmt.Println("err:", err)
	}
	key := "address"
	gens := []*redis.GeoLocation{
		{Name: "shanghai", Latitude: 38.115557, Longitude: 120.361389},
		{Name: "beijing", Latitude: 37.502668, Longitude: 133.613893},
	}
	geoAdd, err := redisClient.GeoAdd(key, gens...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(geoAdd)
}

func TestGeoFind(t *testing.T) {
	redisClient, err := redis.NewRedis(redis.RedisConf{Host: "192.168.3.88:116379", Type: "node"})
	if err != nil {
		fmt.Println("err:", err)
	}
	key := "address"

	geoPos, err := redisClient.GeoPos(key, "shanghai")
	if err != nil {
		return
	}
	fmt.Println("geoPos", geoPos[0].Latitude)

}

func TestGeoDist(t *testing.T) {
	redisClient, err := redis.NewRedis(redis.RedisConf{Host: "192.168.3.88:116379", Type: "node"})
	if err != nil {
		fmt.Println("err:", err)
	}
	key := "address"

	dist, err := redisClient.GeoDist(key, "shanghai", "beijing", "km")
	if err != nil {
		return
	}
	fmt.Println("dist", dist)
	//GEORADIUS address 120 38 199 km
}

func TestGeoRadius(t *testing.T) {
	redisClient, err := redis.NewRedis(redis.RedisConf{Host: "192.168.3.88:116379", Type: "node"})
	if err != nil {
		fmt.Println("err:", err)
	}
	key := "address"
	radius, err := redisClient.GeoRadius(key, 120.313, 38.115557, &redis.GeoRadiusQuery{
		Radius:   2000,
		WithDist: true,
		Unit:     "km",
		Sort:     "ASC",
	})
	if err != nil {
		return
	}
	fmt.Println("radius", radius)

	//Radius float64
	//// Can be m, km, ft, or mi. Default is km.
	//Unit        string
	//WithCoord   bool
	//WithDist    bool
	//WithGeoHash bool
	//Count       int
	//// Can be ASC or DESC. Default is no sort order.
	//Sort      string
	//Store     string
	//StoreDist string

}

func TestScanDel(t *testing.T) {
	redisClient, err := redis.NewRedis(redis.RedisConf{Host: "127.0.0.1:6379", Type: "node"})
	if err != nil {
		fmt.Println("err:", err)
	}

	//redisClient.Set("x:1", "x1")
	//redisClient.Set("x:2", "x2")
	//redisClient.Set("x:3", "x3")
	//redisClient.Set("x:4", "x4")

	redisx.DelPrefix(redisClient, "x:*")

}
