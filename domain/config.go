package domain

import (
	"sync"

	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
)

type Config struct {
	DBSql   *gorm.DB
	DBRedis map[string]*redis.Client
	SMTP    SMTP
	Minio   Minio
	Hash    Hash
}

type Minio struct {
	Client     *minio.Client
	BucketName string
}

type Hash struct {
	Key string
}

var (
	DataPool sync.Pool
)
