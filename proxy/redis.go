package proxy

import (
	"github.com/garyburd/redigo/redis"
	"net/http"
	"time"
	"github.com/johnlion/sites/config"
	"fmt"

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
		MaxIdle: config.REDIS_MAXIDLE,
		MaxActive: config.REDIS_MAXACTIVE,
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
func ( p *Proxy ) RedisServer( res http.ResponseWriter, req *http.Request ,body string ){
	startTime := time.Now()

	// 从连接池里面获得一个连接
	c := p.RedisPool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	dbkey := req.Host + req.URL.RequestURI()


	exists, err := redis.Bool(c.Do("EXISTS", dbkey ))
	if err != nil {
		// handle error return from c.Do or type conversion error.
		panic( err )
	}


	msg := fmt.Sprintf("UsedTime：%s", time.Now().Sub(startTime));
	if exists {
		if strs, err := redis.String( c.Do( "GET" , dbkey ) ); err != nil{
			panic( err )
		}else {
			fmt.Fprint( res, strs )

		}
	}



	p.Debug( msg )
	p.LogText( "[redis]" + "[" + dbkey + "]" + msg  )



}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年5月5日 上午11:45
 * Func: RedisServerNoResponse
 * Desc: redis数据库池服务器,参数不带
 *********************************************/
func ( p *Proxy ) RedisServerNoResponse(  req *http.Request ,body string ){
	startTime := time.Now()

	// 从连接池里面获得一个连接
	c := p.RedisPool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	dbkey := req.Host + req.URL.RequestURI()


	exists, err := redis.Bool(c.Do("EXISTS", dbkey ))
	if err != nil {
		// handle error return from c.Do or type conversion error.
		panic( err )
	}
	if exists {
		p.Debug( "key is existed!"  )
	}else{
		redis.Values(c.Do("SET", dbkey, body ))

	}

	msg := fmt.Sprintf("用时：%s", time.Now().Sub(startTime));
	p.Debug( msg )
	//io.WriteString(res, msg+"\n\n");

}

/*********************************************
 * Author: chandlerxue
 * Email: xps_8@hotmail.com
 * Date: 2016年5月5日 下午2:27
 * Func: RedisKeyExisted
 * Desc: 验证redis 数据库中key是否存在;存在返回true,否则反回false
 *********************************************/
func ( p *Proxy ) RedisKeyExisted(  req *http.Request  ) bool{
	// 从连接池里面获得一个连接
	c := p.RedisPool.Get()
	// 连接完关闭，其实没有关闭，是放回池里，也就是队列里面，等待下一个重用
	defer c.Close()
	dbkey := req.Host + req.URL.RequestURI()

	exists, err := redis.Bool(c.Do("EXISTS", dbkey ))
	if err != nil {
		// handle error return from c.Do or type conversion error.
		panic( err )
	}
	if exists {
		return true
	}else{
		return false
	}

}