package http

import "github.com/bilibili/kratos/pkg/net/http/blademaster"

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


)


//上传stream文件
func UploadFile(c *blademaster.Context)  {

}


//查询用户名下所有视频
func FindListFileByIdNo(c *blademaster.Context)  {

}

//查询用户名下指定视频的所有评价
func FindFileAllevalbyVid(c *blademaster.Context)  {

}

//添加用户名下指定视频的评价
func NewFileAllevalbyVid(c *blademaster.Context)  {

}
