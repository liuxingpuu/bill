package logic

import (
	"context"
	"errors"
	"time"

	"go-zero-demo/greet/internal/logic/model"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)


type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterReply, err error) {
	// todo: add your logic here and delete this line
	/*1.先检验输入注册的参数是否合法
	2.查询数据库是否存在
	3.存入数据库,获取当前时间存入
	 */


	if len(req.Mobile) != 11 || req.VerifyCode != "111111"{
		return nil, errors.New("参数错误")
	}
	UserInfo, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	switch err {
	case model.ErrNotFound:
		logx.Info("用户手机号不存在")
	}
	if UserInfo != nil {
		return nil, errors.New("手机号已存在")
	}

	currentTime := time.Now()
	//data := model.User{Mobile: req.Mobile, CreatedTime: currentTime}
	data := model.User{
		Mobile: req.Mobile,
		Password: "123456",
		LastActiveTime: currentTime,
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &data)

	if err !=nil {
		logx.Info("存入数据失败", err)
		return nil, errors.New("注册失败")
	}
	logx.Info("注册成功")
	return nil, nil
}
