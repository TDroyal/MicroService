package logic

import (
	"captcha/dao"
	"context"
	"time"
)

var ctx = context.Background()

// 操作redis
func Rset(key, value string, expiration time.Duration) error {
	return dao.RDB.Set(ctx, key, value, expiration).Err() // 0表示没有过期时间
}

func RGet(key string) (string, error) {
	v, err := dao.RDB.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

func RDel(key string) error {
	return dao.RDB.Del(ctx, key).Err()
}
