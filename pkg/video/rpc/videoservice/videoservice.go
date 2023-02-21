// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package videoservice

import (
	"context"
	"douyin/pkg/video/rpc/types/video"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetAllVideoByUserIdReq  = video.GetAllVideoByUserIdReq
	GetAllVideoByUserIdResp = video.GetAllVideoByUserIdResp
	GetVideoReq             = video.GetVideoReq
	GetVideoResp            = video.GetVideoResp
	PublishVideoReq         = video.PublishVideoReq
	PublishVideoResp        = video.PublishVideoResp
	User                    = video.User
	Video                   = video.Video

	VideoService interface {
		PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*PublishVideoResp, error)
		GetVideo(ctx context.Context, in *GetVideoReq, opts ...grpc.CallOption) (*GetVideoResp, error)
		GetAllVideoByUserId(ctx context.Context, in *GetAllVideoByUserIdReq, opts ...grpc.CallOption) (*GetAllVideoByUserIdResp, error)
	}

	defaultVideoService struct {
		cli zrpc.Client
	}
)

func NewVideoService(cli zrpc.Client) VideoService {
	return &defaultVideoService{
		cli: cli,
	}
}

func (m *defaultVideoService) PublishVideo(ctx context.Context, in *PublishVideoReq, opts ...grpc.CallOption) (*PublishVideoResp, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.PublishVideo(ctx, in, opts...)
}

func (m *defaultVideoService) GetVideo(ctx context.Context, in *GetVideoReq, opts ...grpc.CallOption) (*GetVideoResp, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.GetVideo(ctx, in, opts...)
}

func (m *defaultVideoService) GetAllVideoByUserId(ctx context.Context, in *GetAllVideoByUserIdReq, opts ...grpc.CallOption) (*GetAllVideoByUserIdResp, error) {
	client := video.NewVideoServiceClient(m.cli.Conn())
	return client.GetAllVideoByUserId(ctx, in, opts...)
}
