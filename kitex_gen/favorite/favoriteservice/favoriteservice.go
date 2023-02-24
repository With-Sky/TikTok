// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoriteservice

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	favorite "tiktok/kitex_gen/favorite"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*favorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Favorite":       kitex.NewMethodInfo(favoriteHandler, newFavoriteArgs, newFavoriteResult, false),
		"DeleteFavorite": kitex.NewMethodInfo(deleteFavoriteHandler, newDeleteFavoriteArgs, newDeleteFavoriteResult, false),
		"FavoriteList":   kitex.NewMethodInfo(favoriteListHandler, newFavoriteListArgs, newFavoriteListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favorite",
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

func favoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.FavoriteReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).Favorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteArgs:
		success, err := handler.(favorite.FavoriteService).Favorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteArgs() interface{} {
	return &FavoriteArgs{}
}

func newFavoriteResult() interface{} {
	return &FavoriteResult{}
}

type FavoriteArgs struct {
	Req *favorite.FavoriteReq
}

func (p *FavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.FavoriteReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteArgs) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteArgs_Req_DEFAULT *favorite.FavoriteReq

func (p *FavoriteArgs) GetReq() *favorite.FavoriteReq {
	if !p.IsSetReq() {
		return FavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type FavoriteResult struct {
	Success *favorite.FavoriteRes
}

var FavoriteResult_Success_DEFAULT *favorite.FavoriteRes

func (p *FavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.FavoriteRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteResult) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteResult) GetSuccess() *favorite.FavoriteRes {
	if !p.IsSetSuccess() {
		return FavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.FavoriteRes)
}

func (p *FavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deleteFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.DeleteFavoriteReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).DeleteFavorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeleteFavoriteArgs:
		success, err := handler.(favorite.FavoriteService).DeleteFavorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteFavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newDeleteFavoriteArgs() interface{} {
	return &DeleteFavoriteArgs{}
}

func newDeleteFavoriteResult() interface{} {
	return &DeleteFavoriteResult{}
}

type DeleteFavoriteArgs struct {
	Req *favorite.DeleteFavoriteReq
}

func (p *DeleteFavoriteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.DeleteFavoriteReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteFavoriteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteFavoriteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteFavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeleteFavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteFavoriteArgs) Unmarshal(in []byte) error {
	msg := new(favorite.DeleteFavoriteReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteFavoriteArgs_Req_DEFAULT *favorite.DeleteFavoriteReq

func (p *DeleteFavoriteArgs) GetReq() *favorite.DeleteFavoriteReq {
	if !p.IsSetReq() {
		return DeleteFavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteFavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeleteFavoriteResult struct {
	Success *favorite.DeleteFavoriteRes
}

var DeleteFavoriteResult_Success_DEFAULT *favorite.DeleteFavoriteRes

func (p *DeleteFavoriteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.DeleteFavoriteRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteFavoriteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteFavoriteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteFavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeleteFavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteFavoriteResult) Unmarshal(in []byte) error {
	msg := new(favorite.DeleteFavoriteRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteFavoriteResult) GetSuccess() *favorite.DeleteFavoriteRes {
	if !p.IsSetSuccess() {
		return DeleteFavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteFavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.DeleteFavoriteRes)
}

func (p *DeleteFavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.FavoriteListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteService).FavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteListArgs:
		success, err := handler.(favorite.FavoriteService).FavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteListResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteListArgs() interface{} {
	return &FavoriteListArgs{}
}

func newFavoriteListResult() interface{} {
	return &FavoriteListResult{}
}

type FavoriteListArgs struct {
	Req *favorite.FavoriteListReq
}

func (p *FavoriteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(favorite.FavoriteListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteListArgs_Req_DEFAULT *favorite.FavoriteListReq

func (p *FavoriteListArgs) GetReq() *favorite.FavoriteListReq {
	if !p.IsSetReq() {
		return FavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FavoriteListResult struct {
	Success *favorite.FavoriteListRes
}

var FavoriteListResult_Success_DEFAULT *favorite.FavoriteListRes

func (p *FavoriteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(favorite.FavoriteListRes)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteListResult) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteListRes)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteListResult) GetSuccess() *favorite.FavoriteListRes {
	if !p.IsSetSuccess() {
		return FavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.FavoriteListRes)
}

func (p *FavoriteListResult) IsSetSuccess() bool {
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

func (p *kClient) Favorite(ctx context.Context, Req *favorite.FavoriteReq) (r *favorite.FavoriteRes, err error) {
	var _args FavoriteArgs
	_args.Req = Req
	var _result FavoriteResult
	if err = p.c.Call(ctx, "Favorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteFavorite(ctx context.Context, Req *favorite.DeleteFavoriteReq) (r *favorite.DeleteFavoriteRes, err error) {
	var _args DeleteFavoriteArgs
	_args.Req = Req
	var _result DeleteFavoriteResult
	if err = p.c.Call(ctx, "DeleteFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, Req *favorite.FavoriteListReq) (r *favorite.FavoriteListRes, err error) {
	var _args FavoriteListArgs
	_args.Req = Req
	var _result FavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
