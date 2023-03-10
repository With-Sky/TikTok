// Code generated by Kitex v0.4.4. DO NOT EDIT.

package commentservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	comment "tiktok/kitex_gen/comment"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

var commentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*comment.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CommentAction": kitex.NewMethodInfo(commentActionHandler, newCommentActionArgs, newCommentActionResult, false),
		"DeleteComment": kitex.NewMethodInfo(deleteCommentHandler, newDeleteCommentArgs, newDeleteCommentResult, false),
		"CommentList":   kitex.NewMethodInfo(commentListHandler, newCommentListArgs, newCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.CommentReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).CommentAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CommentActionArgs:
		success, err := handler.(comment.CommentService).CommentAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CommentActionResult)
		realResult.Success = success
	}
	return nil
}
func newCommentActionArgs() interface{} {
	return &CommentActionArgs{}
}

func newCommentActionResult() interface{} {
	return &CommentActionResult{}
}

type CommentActionArgs struct {
	Req *comment.CommentReq
}

func (p *CommentActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.CommentReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CommentActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CommentActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CommentActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CommentActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CommentActionArgs) Unmarshal(in []byte) error {
	msg := new(comment.CommentReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CommentActionArgs_Req_DEFAULT *comment.CommentReq

func (p *CommentActionArgs) GetReq() *comment.CommentReq {
	if !p.IsSetReq() {
		return CommentActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CommentActionArgs) IsSetReq() bool {
	return p.Req != nil
}

type CommentActionResult struct {
	Success *comment.CommentRes
}

var CommentActionResult_Success_DEFAULT *comment.CommentRes

func (p *CommentActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.CommentRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CommentActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CommentActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CommentActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CommentActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CommentActionResult) Unmarshal(in []byte) error {
	msg := new(comment.CommentRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CommentActionResult) GetSuccess() *comment.CommentRes {
	if !p.IsSetSuccess() {
		return CommentActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CommentActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.CommentRes)
}

func (p *CommentActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.DeleteCommentReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).DeleteComment(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteCommentArgs:
		success, err := handler.(comment.CommentService).DeleteComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteCommentResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteCommentArgs() interface{} {
	return &DeleteCommentArgs{}
}

func newDeleteCommentResult() interface{} {
	return &DeleteCommentResult{}
}

type DeleteCommentArgs struct {
	Req *comment.DeleteCommentReq
}

func (p *DeleteCommentArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.DeleteCommentReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteCommentArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteCommentArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteCommentArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteCommentArgs) Unmarshal(in []byte) error {
	msg := new(comment.DeleteCommentReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteCommentArgs_Req_DEFAULT *comment.DeleteCommentReq

func (p *DeleteCommentArgs) GetReq() *comment.DeleteCommentReq {
	if !p.IsSetReq() {
		return DeleteCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteCommentResult struct {
	Success *comment.DeleteCommentRes
}

var DeleteCommentResult_Success_DEFAULT *comment.DeleteCommentRes

func (p *DeleteCommentResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.DeleteCommentRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteCommentResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteCommentResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteCommentResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteCommentResult) Unmarshal(in []byte) error {
	msg := new(comment.DeleteCommentRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteCommentResult) GetSuccess() *comment.DeleteCommentRes {
	if !p.IsSetSuccess() {
		return DeleteCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.DeleteCommentRes)
}

func (p *DeleteCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(comment.CommentListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(comment.CommentService).CommentList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CommentListArgs:
		success, err := handler.(comment.CommentService).CommentList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CommentListResult)
		realResult.Success = success
	}
	return nil
}
func newCommentListArgs() interface{} {
	return &CommentListArgs{}
}

func newCommentListResult() interface{} {
	return &CommentListResult{}
}

type CommentListArgs struct {
	Req *comment.CommentListReq
}

func (p *CommentListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(comment.CommentListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CommentListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CommentListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CommentListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CommentListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CommentListArgs) Unmarshal(in []byte) error {
	msg := new(comment.CommentListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CommentListArgs_Req_DEFAULT *comment.CommentListReq

func (p *CommentListArgs) GetReq() *comment.CommentListReq {
	if !p.IsSetReq() {
		return CommentListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CommentListArgs) IsSetReq() bool {
	return p.Req != nil
}

type CommentListResult struct {
	Success *comment.CommentListRes
}

var CommentListResult_Success_DEFAULT *comment.CommentListRes

func (p *CommentListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(comment.CommentListRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CommentListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CommentListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CommentListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CommentListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CommentListResult) Unmarshal(in []byte) error {
	msg := new(comment.CommentListRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CommentListResult) GetSuccess() *comment.CommentListRes {
	if !p.IsSetSuccess() {
		return CommentListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CommentListResult) SetSuccess(x interface{}) {
	p.Success = x.(*comment.CommentListRes)
}

func (p *CommentListResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CommentAction(ctx context.Context, Req *comment.CommentReq) (r *comment.CommentRes, err error) {
	var _args CommentActionArgs
	_args.Req = Req
	var _result CommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteComment(ctx context.Context, Req *comment.DeleteCommentReq) (r *comment.DeleteCommentRes, err error) {
	var _args DeleteCommentArgs
	_args.Req = Req
	var _result DeleteCommentResult
	if err = p.c.Call(ctx, "DeleteComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, Req *comment.CommentListReq) (r *comment.CommentListRes, err error) {
	var _args CommentListArgs
	_args.Req = Req
	var _result CommentListResult
	if err = p.c.Call(ctx, "CommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
