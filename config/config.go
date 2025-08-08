package config

import (
	"context"
	"fmt"
	"strings"

	"log"
	"os"
	"strconv"
	"sync"

	"github.com/alfisar/jastip-import/database"
	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"google.golang.org/grpc"

	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

var (
	_ = godotenv.Load(".env")
)

// configuration init config
func Init() {
	// Initialize connection
	dbSQL := initDB()
	redisDB := initRedis()
	minioClient, bucketName := initMinio()

	// Initialize sync.Pool
	domain.DataPool = sync.Pool{
		New: func() interface{} {
			return &domain.Config{
				DBSql:   dbSQL,
				DBRedis: redisDB,
				SMTP: domain.SMTP{
					Host:   os.Getenv("SMTP_HOST"),
					Port:   os.Getenv("SMTP_PORT"),
					User:   os.Getenv("SMTP_USER"),
					Pass:   os.Getenv("SMTP_PASS"),
					From:   os.Getenv("SMTP_FROM"),
					Mailer: gomail.NewMessage(),
				},
				Minio: domain.Minio{
					Client:     minioClient,
					BucketName: bucketName,
				},
				Hash: domain.Hash{
					Key: os.Getenv("HASH_KEY"),
				},
				GRPC: iniGRPC(),
			}
		},
	}
}

// Function to initialize DB
func initDB() *gorm.DB {
	fmt.Println("DB_USE : " + os.Getenv("DB_USE"))
	fmt.Println("DB_HOST : " + os.Getenv("DB_HOST"))
	switch os.Getenv("DB_USE") {
	case "MySQL":
		db, err := database.NewConnSQL()
		if err != nil {
			log.Fatalf("Failed to connect to MySQL: %v", err)
		}
		return db
	default:
		log.Fatalln("Invalid DB_USE specified in environment variables")
	}
	return nil
}

// Function to initialize Redis
func initRedis() map[string]*redis.Client {
	redisDB, err := database.NewDatabaseRedis()
	if err != nil {
		log.Fatalf("Failed to initialize Redis: %v", err)
	}
	return redisDB
}

// Function to initialize MinIO
func initMinio() (*minio.Client, string) {
	client, bucketName := NewConnMinio()
	return client, bucketName
}

func NewConnMinio() (*minio.Client, string) {
	// Konfigurasi koneksi ke server MinIO
	endpoint := os.Getenv("MINIO_HOST")
	minioKey := os.Getenv("MINIO_KEY")
	minioSecret := os.Getenv("MINIO_SECRET")
	useSSL, _ := strconv.ParseBool(os.Getenv("MINIO_SSL")) // Set true jika menggunakan HTTPS

	if endpoint == "" || minioKey == "" || minioSecret == "" {
		log.Fatal("Failed to connect to Minio: Endpoint, minioKey, or minioSecret is empty")

	}

	// membuat koneksi baru ke minio
	minioClient, err := minio.New(endpoint, &minio.Options{
		Secure: useSSL,
		Creds:  credentials.NewStaticV4(minioKey, minioSecret, ""),
	})

	if err != nil {
		log.Fatalf("Failed to connect to Minio: %v", err)

	}

	// membuat bucket baru dengan nama jastip jika tidak ada bucketnya
	bucketName := "jastip"
	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			fmt.Println("Bucket already exists")
		} else {
			log.Fatalf("Failed to make bucket Minio: %v", err)
		}
	} else {
		fmt.Println("Successfully created bucket")
	}

	return minioClient, bucketName
}

func iniGRPC() (result map[string]*grpc.ClientConn) {
	gRPCAddr := strings.Split(os.Getenv("GRPC_ADDR"), ",")
	for _, v := range gRPCAddr {
		conn, err := grpc.Dial(v, grpc.WithInsecure())
		if err != nil {
			log.Fatalln("Cannot connect GRPC : " + err.Error())
		}
		defer conn.Close()
		result[consts.GrpcAuth] = conn
	}
	return
}
