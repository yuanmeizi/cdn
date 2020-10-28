package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func main() {

	log.Println("运行开始")

	rs, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println("Connect to redis error", err)
		return
	}
	log.Println("连接数据库成功")

	defer rs.Close()
	// 密码
	_, err = rs.Do("AUTH", "yuan")
	if err != nil {
		log.Println("AUTH redis error", err)
		return
	}
	// 选择第几个数据库
	_, err = rs.Do("SELECT", 0)
	if err != nil {
		log.Println("select redis error", err)
		return
	}
	log.Println("选择数据库成功")

	// 存数据
	// _, err = rs.Do("SET", "name", "aaa")
	// if err != nil {
	// 	log.Println("redis set failed:", err)
	// }
	// log.Println(fmt.Sprintf("操作成功%v", 1))

	// 取数据
	// value, err := redis.Bytes(rs.Do("GET", "name"))
	// if err != nil {
	// 	log.Println("get redis error", err)
	// 	return
	// }
	// log.Println(fmt.Sprintf("value为%v", value))

	// 存列表
	// rs.Do("lpush", "mylist", "zjd")
	// rs.Do("lpush", "mylist", "jzd")
	// rs.Do("lpush", "mylist", "yuan")
	// log.Println(fmt.Sprintf("操作成功%v", 1))

	// 取列表
	// values, _ := redis.Values(rs.Do("lrange", "mylist", "0", "10"))

	// for i, v := range values {
	// 	log.Println(fmt.Sprintf("第%v项为%v", i, string(v.([]byte))))
	// }

	// 存json
	// imap := map[string]string{"username": "666", "phonenumber": "888"}
	// value, _ := json.Marshal(imap)

	// n, err := rs.Do("SETNX", "jsonstr", value)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if n == int64(1) {
	// 	fmt.Println("success")
	// }

	// 取json
	// var imapGet map[string]string

	// valueGet, err := redis.Bytes(rs.Do("GET", "jsonstr"))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = json.Unmarshal(valueGet, &imapGet)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(imapGet)

}
