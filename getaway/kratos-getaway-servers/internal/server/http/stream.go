package http

import (
	pb "github.com/KXX747/wolf/getaway/kratos-getaway-servers/api/stream_server"
	"github.com/bilibili/kratos/pkg/ecode"
	"github.com/bilibili/kratos/pkg/net/http/blademaster"
)

/**
图片视频接口
 */

const (
	//http://127.0.0.1:39000/stream.server.v1.Upload/file
	//http://127.0.0.1:39000/stream.server.v1.Upload/listfile?id_no=5f69fcc0-6cf4-46e6-a763-0bba633b816e&start_pos=0&end_pos=10
	//http://127.0.0.1:39000/stream.server.v1.Upload/fileallevalby?vid=76d19615-4ae4-420c-902f-431cd8503d48&start_pos=5&end_pos=10&neid
	//http://127.0.0.1:39000/stream.server.v1.Upload/addevaluation?id_no=5f69fcc0-6cf4-46e6-a763-0bba633b816e&vid=76d19615-4ae4-420c-902f-431cd8503d48&n_eid=98efe528-772c-4cb4-825b-c34e154b2225&e_idno=e5533f32-6250-46ee-a6cf-1d9b339284a3&e_name=俞国&e_content=知识确实能改变命运，但是真正改变命运的却又少之又少
	StreamUploadUrl="%s:%s/stream.server.v1.Upload/file"
	FindUploadUrl="%s:%s/stream.server.v1.Upload/listfile"
	NewFileAllevalbyUrl="%s:%s/stream.server.v1.Upload/addevaluation"
	FindFileAllevalbyUrl="%s:%s/stream.server.v1.Upload/fileallevalby"
	FILEUP  = "file_up"

)


//上传stream文件
func UploadFile(c *blademaster.Context)  {
	mUploadFileReq:=new(pb.UploadFileReq)
	if err := c.Bind(mUploadFileReq); err != nil {
		c.JSON(nil,  ecode.ReqParamErr)
		return
	}

	//WmText
	//获取key获取文件名称
	file, handler, err := c.Request.FormFile(FILEUP)
	if file==nil ||err!=nil {
		err = ecode.ReqParamErr
		return
	}
	defer file.Close()

	size:=handler.Size
	c.Request.ParseMultipartForm(size)

	content := make([]byte, size)
	file.Read(content)
	mUploadFileReq.WmText = string(content[:])
	if reply,code:=streamRPCClient.File(c,mUploadFileReq);code!=nil {
		c.JSON(nil, code)
		return
	}else {
		c.JSON(reply,  nil)
	}


}

//添加用户名下指定视频的评价
func NewFileAllevalbyVid(c *blademaster.Context)  {

	mEvaluationVodieReq:=new(pb.EvaluationVodieReq)
	if err := c.Bind(mEvaluationVodieReq); err != nil {
		c.JSON(nil,  ecode.ReqParamErr)
		return
	}

	if reply,code:=streamRPCClient.Addevaluation(c,mEvaluationVodieReq);code!=nil {
		c.JSON(nil, code)
		return
	}else {
		c.JSON(reply,  nil)
	}


}


//查询用户名下所有视频
func FindListFileByIdNo(c *blademaster.Context)  {

	mFileListReq:=new(pb.FileListReq)
	if err := c.Bind(mFileListReq); err != nil {
		c.JSON(nil,  ecode.ReqParamErr)
		return
	}

	if reply,code:=streamRPCClient.Listfile(c,mFileListReq);code!=nil {
		c.JSON(nil, code)
		return
	}else {
		c.JSON(reply,  nil)
	}

}

//查询用户名下指定视频的所有评价
func FindFileAllevalbyVid(c *blademaster.Context)  {

	mEvaluationGetReq:=new(pb.EvaluationGetReq)
	if err := c.Bind(mEvaluationGetReq); err != nil {
		c.JSON(nil,  ecode.ReqParamErr)
		return
	}

	if reply,code:=streamRPCClient.Fileallevalby(c,mEvaluationGetReq);code!=nil {
		c.JSON(nil, code)
		return
	}else {
		c.JSON(reply,  nil)
	}

}


