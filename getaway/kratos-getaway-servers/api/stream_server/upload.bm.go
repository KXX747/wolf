// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: upload.proto

/*
Package stream_server_v1 is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

It is generated from these files:
	upload.proto
*/
package stream_server_v1

import (
	"context"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathTokenNew = "/stream.server.v1.Token/new"

var PathUploadFile = "/stream.server.v1.Upload/file"
var PathUploadAddevaluation = "/stream.server.v1.Upload/addevaluation"
var PathUploadFileallevalby = "/stream.server.v1.Upload/fileallevalby"
var PathUploadListfile = "/stream.server.v1.Upload/listfile"

// TokenBMServer is the server API for Token service.
type TokenBMServer interface {
	// Request for a token for upload.
	// `method:"POST" internal:"true"`
	New(ctx context.Context, req *NewTokenReq) (resp *NewTokenResp, err error)
}

var TokenSvc TokenBMServer

func tokenNew(c *bm.Context) {
	p := new(NewTokenReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := TokenSvc.New(c, p)
	c.JSON(resp, err)
}

// RegisterTokenBMServer Register the blademaster route
func RegisterTokenBMServer(e *bm.Engine, server TokenBMServer) {
	TokenSvc = server
	e.POST("/stream.server.v1.Token/new", tokenNew)
}

// UploadBMServer is the server API for Upload service.
type UploadBMServer interface {
	// `method:"POST" content-type:"multipart/form-data" midware:"cors,guest"`
	File(ctx context.Context, req *UploadFileReq) (resp *UploadFileResp, err error)

	// 添加视频的评价
	Addevaluation(ctx context.Context, req *EvaluationVodieReq) (resp *EvaluationVodieResp, err error)

	// 获取视频所有评价
	Fileallevalby(ctx context.Context, req *EvaluationGetReq) (resp *EvaluationListByVodieResp, err error)

	// 获取指定视频
	Listfile(ctx context.Context, req *FileListReq) (resp *FileListResp, err error)
}

var UploadSvc UploadBMServer

func uploadFile(c *bm.Context) {
	p := new(UploadFileReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := UploadSvc.File(c, p)
	c.JSON(resp, err)
}

func uploadAddevaluation(c *bm.Context) {
	p := new(EvaluationVodieReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := UploadSvc.Addevaluation(c, p)
	c.JSON(resp, err)
}

func uploadFileallevalby(c *bm.Context) {
	p := new(EvaluationGetReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := UploadSvc.Fileallevalby(c, p)
	c.JSON(resp, err)
}

func uploadListfile(c *bm.Context) {
	p := new(FileListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := UploadSvc.Listfile(c, p)
	c.JSON(resp, err)
}

// RegisterUploadBMServer Register the blademaster route
func RegisterUploadBMServer(e *bm.Engine, server UploadBMServer) {
	UploadSvc = server
	e.POST("/stream.server.v1.Upload/file", uploadFile)
	e.GET("/stream.server.v1.Upload/addevaluation", uploadAddevaluation)
	e.GET("/stream.server.v1.Upload/fileallevalby", uploadFileallevalby)
	e.GET("/stream.server.v1.Upload/listfile", uploadListfile)
}
