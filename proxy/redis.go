package proxy

import (
	"github.com/garyburd/redigo/redis"
	"net/http"
	"time"
	"github.com/johnlion/johnlion/config"
	"fmt"
	"io"
)

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年5月4日 下午4:36
 * Func: newPool
 * Desc: 长连接池,overwrite redis pool
 *********************************************/
func (p *Proxy) NewPool() *redis.Pool{
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.REDIS_ADDR + ":" + config.REDIS_PORT  )
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年5月4日 下午5:17
 * Func: redisServer
 * Desc: redis数据库池服务器
 *********************************************/
func ( p *Proxy ) RedisServer( res http.ResponseWriter, req *http.Request ){
	startTime := time.Now()

	// 从连接池里面获得一个连接
	c := p.RedisPool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	dbkey := "netgame:info"
	if ok, err := redis.Bool( c.Do("LPUSH", dbkey, "yangzetao") ); ok{
	}else{
		panic( err )
	}


	msg := fmt.Sprintf("用时：%s", time.Now().Sub(startTime));
	io.WriteString(res, msg+"\n\n");

}