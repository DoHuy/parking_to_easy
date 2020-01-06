package redis

import(
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/gomodule/redigo/redis"
	"time"
	"github.com/DoHuy/parking_to_easy/config"
)
// init singleton instance redis
type Redis struct {
	pool	*redis.Pool
}

func NewRedis() *Redis{
	var pool *redis.Pool
	pool = initPoolConnectionRedis()
	return &Redis{pool: pool}
}

func initPoolConnectionRedis() *redis.Pool {
	var pool *redis.Pool
	var redisConfig model.Redis
	redisConfig = config.GetConfigRedis()
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(redisConfig.Networks, redisConfig.Address)
		},
	}
	return pool
}


// thêm 1 token vào trong redis hoặc ghi đè token trong reddis
func (this *Redis)SetJWTTokenToRedis(key, jwt string) error{
	conn := this.pool.Get()
	var redisConfig model.Redis
	redisConfig = config.GetConfigRedis()
	_, err := conn.Do("HMSET", redisConfig.Topic[0], key, jwt)
	if err != nil {
		return fmt.Errorf("Lỗi thêm dữ liêu vào Redis", err)
	}
	defer conn.Close()
	return nil
}

func (this *Redis)DelJWTToken(key string) error{
	conn := this.pool.Get()
	var redisConfig model.Redis
	redisConfig = config.GetConfigRedis()
	_, err := conn.Do("HDEL", redisConfig.Topic[0], key)
	if err != nil {
		return fmt.Errorf("Lỗi xóa dữ liệu trong redis", err)
	}
	fmt.Println("Co chay ao HDEL")
	defer conn.Close()
	return nil
}

func (this *Redis)GetJWTTokenFromRedis(key string) (model.JWTOfUser, error){
	conn := this.pool.Get()
	defer conn.Close()
	var redisConfig model.Redis
	redisConfig = config.GetConfigRedis()
	jwt, err := redis.String(conn.Do("HGET", redisConfig.Topic[0], key))
	if err != nil {
		return model.JWTOfUser{}, fmt.Errorf("Lỗi lấy dữ liệu từ Redis %s", err.Error())
	}
	return model.JWTOfUser{Key: key, Jwt: jwt}, nil
}

func (this *Redis)SetDevicesTokenToRedis(key string, tokenDevice string) error {
	conn := this.pool.Get()
	var redisConfig model.Redis
	redisConfig = config.GetConfigRedis()
	_, err := conn.Do("HMSET", redisConfig.Topic[1], key, tokenDevice)
	if err != nil {
		return fmt.Errorf("Lỗi thêm dữ liêu vào Redis", err)
	}
	defer conn.Close()
	return nil
}

func (this *Redis)GetAllTokenDevicesFromRedis(){

}



