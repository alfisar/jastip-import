package client

import (
	"context"

	pb "github.com/alfisar/jastip-import/proto/auth"
)

type AuthClientContract interface {
	GetAddrByID(ctx context.Context, userID, addressID int32) (*pb.ResponseAddressByID, error)
}
