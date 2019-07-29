package client
/**
视频和评价服务
 */
import (
	pb "github.com/KXX747/wolf/getaway/kratos-getaway-servers/api/stream_server"
	"github.com/bilibili/kratos/pkg/net/rpc/warden"
	"github.com/bilibili/kratos/pkg/log"
	"google.golang.org/grpc"
	"context"
)

type StreamServer struct {
	cfg *warden.ClientConfig
	uploadClient pb.UploadClient
}

// NewClient new member grpc client
func newStreamClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (pb.UploadClient, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), "127.0.0.1:39000")
	if err != nil {
		return nil, err
	}
	// 注意替换这里：
	// NewDemoClient方法是在"api"目录下代码生成的
	// 对应proto文件内自定义的service名字，请使用正确方法名替换
	return pb.NewUploadClient(conn), nil
}

//
func NewStreamServer(conf *warden.ClientConfig) (mStreamServer *StreamServer){

	streamRpcClient,err:=newStreamClient(conf)
	if err!=nil {
		log.Error("userRPCClient warden.ClientConfig err=%s",err)

	}

	mStreamServer =&StreamServer{
		cfg:conf,
		uploadClient:streamRpcClient,
	}
	return
}

//上传文件
func (mStreamServer *StreamServer) File(ctx context.Context, req *pb.UploadFileReq) (mUploadFileResp *pb.UploadFileResp, err error){

	if mUploadFileResp,err=mStreamServer.uploadClient.File(ctx,req);err!=nil {
		log.Warn("StreamServer File err=%s UploadFileReq=%s",err,req)
		return nil,err
	}
	return
}


//添加视频的评价
func (mStreamServer *StreamServer) Addevaluation(ctx context.Context, req *pb.EvaluationVodieReq)(mVodieResp *pb.EvaluationVodieResp, err error){

	if mVodieResp,err=mStreamServer.uploadClient.Addevaluation(ctx,req);err!=nil {
		log.Warn("StreamServer Addevaluation err=%s EvaluationVodieReq=%s",err,req)
		return nil,err
	}

	return
}

//获取视频的评价
func (mStreamServer *StreamServer) Fileallevalby(ctx context.Context, req *pb.EvaluationGetReq) (mVodieResp *pb.EvaluationListByVodieResp, err error){

	if mVodieResp,err=mStreamServer.uploadClient.Fileallevalby(ctx,req);err!=nil {
		log.Warn("StreamServer Fileallevalby err=%s EvaluationGetReq=%s",err,req)
		return nil,err
	}

	return
}

//湖区用户的视频和评价
func (mStreamServer *StreamServer)Listfile(ctx context.Context, req *pb.FileListReq) (mFileListResp *pb.FileListResp, err error){

	if mFileListResp,err=mStreamServer.uploadClient.Listfile(ctx,req);err!=nil {
		log.Warn("StreamServer Listfile err=%s FileListReq=%s",err,req)
		return nil,err
	}

	return
}







