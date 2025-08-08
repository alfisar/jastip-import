package client

import (
	"context"

	pb "github.com/alfisar/jastip-import/proto/auth"
)

type authClient struct {
	profileClient pb.ProfileClient
}

func NewAuthClient(profileClient pb.ProfileClient) *authClient {
	return &authClient{
		profileClient: profileClient,
	}
}

func (s *authClient) GetAddrByID(ctx context.Context, userID, addressID int32) (*pb.ResponseAddressByID, error) {
	req := &pb.RequestAddressByID{
		UserID:   userID,
		AdressID: addressID,
	}
	return s.profileClient.AddressByID(ctx, req)
}
