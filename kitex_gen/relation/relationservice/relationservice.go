// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	relation "tiktok/kitex_gen/relation"
)

func serviceInfo() *kitex.ServiceInfo {
	return relationServiceServiceInfo
}

var relationServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RelationService"
	handlerType := (*relation.RelationService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Follow":       kitex.NewMethodInfo(followHandler, newFollowArgs, newFollowResult, false),
		"CancelFollow": kitex.NewMethodInfo(cancelFollowHandler, newCancelFollowArgs, newCancelFollowResult, false),
		"FollowList":   kitex.NewMethodInfo(followListHandler, newFollowListArgs, newFollowListResult, false),
		"FollowerList": kitex.NewMethodInfo(followerListHandler, newFollowerListArgs, newFollowerListResult, false),
		"FriendList":   kitex.NewMethodInfo(friendListHandler, newFriendListArgs, newFriendListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "relation",
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

func followHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(relation.FollowReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(relation.RelationService).Follow(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FollowArgs:
		success, err := handler.(relation.RelationService).Follow(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FollowResult)
		realResult.Success = success
	}
	return nil
}
func newFollowArgs() interface{} {
	return &FollowArgs{}
}

func newFollowResult() interface{} {
	return &FollowResult{}
}

type FollowArgs struct {
	Req *relation.FollowReq
}

func (p *FollowArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(relation.FollowReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FollowArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FollowArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FollowArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FollowArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FollowArgs) Unmarshal(in []byte) error {
	msg := new(relation.FollowReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FollowArgs_Req_DEFAULT *relation.FollowReq

func (p *FollowArgs) GetReq() *relation.FollowReq {
	if !p.IsSetReq() {
		return FollowArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FollowArgs) IsSetReq() bool {
	return p.Req != nil
}

type FollowResult struct {
	Success *relation.FollowRes
}

var FollowResult_Success_DEFAULT *relation.FollowRes

func (p *FollowResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(relation.FollowRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FollowResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FollowResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FollowResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FollowResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FollowResult) Unmarshal(in []byte) error {
	msg := new(relation.FollowRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FollowResult) GetSuccess() *relation.FollowRes {
	if !p.IsSetSuccess() {
		return FollowResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FollowResult) SetSuccess(x interface{}) {
	p.Success = x.(*relation.FollowRes)
}

func (p *FollowResult) IsSetSuccess() bool {
	return p.Success != nil
}

func cancelFollowHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(relation.CancelFollowReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(relation.RelationService).CancelFollow(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CancelFollowArgs:
		success, err := handler.(relation.RelationService).CancelFollow(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CancelFollowResult)
		realResult.Success = success
	}
	return nil
}
func newCancelFollowArgs() interface{} {
	return &CancelFollowArgs{}
}

func newCancelFollowResult() interface{} {
	return &CancelFollowResult{}
}

type CancelFollowArgs struct {
	Req *relation.CancelFollowReq
}

func (p *CancelFollowArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(relation.CancelFollowReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CancelFollowArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CancelFollowArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CancelFollowArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CancelFollowArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CancelFollowArgs) Unmarshal(in []byte) error {
	msg := new(relation.CancelFollowReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CancelFollowArgs_Req_DEFAULT *relation.CancelFollowReq

func (p *CancelFollowArgs) GetReq() *relation.CancelFollowReq {
	if !p.IsSetReq() {
		return CancelFollowArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CancelFollowArgs) IsSetReq() bool {
	return p.Req != nil
}

type CancelFollowResult struct {
	Success *relation.CancelFollowRes
}

var CancelFollowResult_Success_DEFAULT *relation.CancelFollowRes

func (p *CancelFollowResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(relation.CancelFollowRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CancelFollowResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CancelFollowResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CancelFollowResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CancelFollowResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CancelFollowResult) Unmarshal(in []byte) error {
	msg := new(relation.CancelFollowRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CancelFollowResult) GetSuccess() *relation.CancelFollowRes {
	if !p.IsSetSuccess() {
		return CancelFollowResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CancelFollowResult) SetSuccess(x interface{}) {
	p.Success = x.(*relation.CancelFollowRes)
}

func (p *CancelFollowResult) IsSetSuccess() bool {
	return p.Success != nil
}

func followListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(relation.FollowListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(relation.RelationService).FollowList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FollowListArgs:
		success, err := handler.(relation.RelationService).FollowList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FollowListResult)
		realResult.Success = success
	}
	return nil
}
func newFollowListArgs() interface{} {
	return &FollowListArgs{}
}

func newFollowListResult() interface{} {
	return &FollowListResult{}
}

type FollowListArgs struct {
	Req *relation.FollowListReq
}

func (p *FollowListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(relation.FollowListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FollowListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FollowListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FollowListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FollowListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FollowListArgs) Unmarshal(in []byte) error {
	msg := new(relation.FollowListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FollowListArgs_Req_DEFAULT *relation.FollowListReq

func (p *FollowListArgs) GetReq() *relation.FollowListReq {
	if !p.IsSetReq() {
		return FollowListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FollowListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FollowListResult struct {
	Success *relation.FollowListRes
}

var FollowListResult_Success_DEFAULT *relation.FollowListRes

func (p *FollowListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(relation.FollowListRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FollowListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FollowListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FollowListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FollowListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FollowListResult) Unmarshal(in []byte) error {
	msg := new(relation.FollowListRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FollowListResult) GetSuccess() *relation.FollowListRes {
	if !p.IsSetSuccess() {
		return FollowListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FollowListResult) SetSuccess(x interface{}) {
	p.Success = x.(*relation.FollowListRes)
}

func (p *FollowListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func followerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(relation.FollowerListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(relation.RelationService).FollowerList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FollowerListArgs:
		success, err := handler.(relation.RelationService).FollowerList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FollowerListResult)
		realResult.Success = success
	}
	return nil
}
func newFollowerListArgs() interface{} {
	return &FollowerListArgs{}
}

func newFollowerListResult() interface{} {
	return &FollowerListResult{}
}

type FollowerListArgs struct {
	Req *relation.FollowerListReq
}

func (p *FollowerListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(relation.FollowerListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FollowerListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FollowerListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FollowerListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FollowerListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FollowerListArgs) Unmarshal(in []byte) error {
	msg := new(relation.FollowerListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FollowerListArgs_Req_DEFAULT *relation.FollowerListReq

func (p *FollowerListArgs) GetReq() *relation.FollowerListReq {
	if !p.IsSetReq() {
		return FollowerListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FollowerListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FollowerListResult struct {
	Success *relation.FollowerListRes
}

var FollowerListResult_Success_DEFAULT *relation.FollowerListRes

func (p *FollowerListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(relation.FollowerListRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FollowerListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FollowerListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FollowerListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FollowerListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FollowerListResult) Unmarshal(in []byte) error {
	msg := new(relation.FollowerListRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FollowerListResult) GetSuccess() *relation.FollowerListRes {
	if !p.IsSetSuccess() {
		return FollowerListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FollowerListResult) SetSuccess(x interface{}) {
	p.Success = x.(*relation.FollowerListRes)
}

func (p *FollowerListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func friendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(relation.FriendListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(relation.RelationService).FriendList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FriendListArgs:
		success, err := handler.(relation.RelationService).FriendList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FriendListResult)
		realResult.Success = success
	}
	return nil
}
func newFriendListArgs() interface{} {
	return &FriendListArgs{}
}

func newFriendListResult() interface{} {
	return &FriendListResult{}
}

type FriendListArgs struct {
	Req *relation.FriendListReq
}

func (p *FriendListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(relation.FriendListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FriendListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FriendListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FriendListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FriendListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FriendListArgs) Unmarshal(in []byte) error {
	msg := new(relation.FriendListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FriendListArgs_Req_DEFAULT *relation.FriendListReq

func (p *FriendListArgs) GetReq() *relation.FriendListReq {
	if !p.IsSetReq() {
		return FriendListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FriendListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FriendListResult struct {
	Success *relation.FriendListRes
}

var FriendListResult_Success_DEFAULT *relation.FriendListRes

func (p *FriendListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(relation.FriendListRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FriendListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FriendListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FriendListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FriendListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FriendListResult) Unmarshal(in []byte) error {
	msg := new(relation.FriendListRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FriendListResult) GetSuccess() *relation.FriendListRes {
	if !p.IsSetSuccess() {
		return FriendListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FriendListResult) SetSuccess(x interface{}) {
	p.Success = x.(*relation.FriendListRes)
}

func (p *FriendListResult) IsSetSuccess() bool {
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

func (p *kClient) Follow(ctx context.Context, Req *relation.FollowReq) (r *relation.FollowRes, err error) {
	var _args FollowArgs
	_args.Req = Req
	var _result FollowResult
	if err = p.c.Call(ctx, "Follow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CancelFollow(ctx context.Context, Req *relation.CancelFollowReq) (r *relation.CancelFollowRes, err error) {
	var _args CancelFollowArgs
	_args.Req = Req
	var _result CancelFollowResult
	if err = p.c.Call(ctx, "CancelFollow", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowList(ctx context.Context, Req *relation.FollowListReq) (r *relation.FollowListRes, err error) {
	var _args FollowListArgs
	_args.Req = Req
	var _result FollowListResult
	if err = p.c.Call(ctx, "FollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowerList(ctx context.Context, Req *relation.FollowerListReq) (r *relation.FollowerListRes, err error) {
	var _args FollowerListArgs
	_args.Req = Req
	var _result FollowerListResult
	if err = p.c.Call(ctx, "FollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FriendList(ctx context.Context, Req *relation.FriendListReq) (r *relation.FriendListRes, err error) {
	var _args FriendListArgs
	_args.Req = Req
	var _result FriendListResult
	if err = p.c.Call(ctx, "FriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}