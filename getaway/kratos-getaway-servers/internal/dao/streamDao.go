package dao

import (
	"context"
	pb "github.com/KXX747/wolf/getaway/kratos-getaway-servers/api/stream_server"
)

//上传文件
func(d *dao)File(ctx context.Context, req *pb.UploadFileReq) (mUploadFileResp *pb.UploadFileResp, err error){

	return
}

//创建文件的token
func(d *dao)New(ctx context.Context, in *pb.NewTokenReq) (mNewTokenResp *pb.NewTokenResp, err error) {

	return
}

//添加评价
func(d *dao)Addevaluation(ctx context.Context, req *pb.EvaluationVodieReq)(mVodieResp *pb.EvaluationVodieResp, err error)   {

	return
}

//查询视频的多有评价
func(d *dao)Fileallevalby(ctx context.Context, req *pb.EvaluationGetReq) (mVodieResp *pb.EvaluationListByVodieResp, err error)  {

	return
}




//查询用户的所有视频列表
func(d *dao)Listfile(ctx context.Context, req *pb.FileListReq) (mFileListResp *pb.FileListResp, err error)  {

	return
}




