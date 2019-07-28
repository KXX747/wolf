package sftp


//图片
func IsImagesType()(string){

	return REMOTE_SFTP_IMAGE
}

//视频
func IsVodieType()(string){

	return REMOTE_SFTP_VODIE
}

//其他类型文件
func IsOtherFileType()(string){

	return REMOTE_SFTP_PDF
}