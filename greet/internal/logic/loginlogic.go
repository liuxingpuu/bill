package logic

import (
	"context"
	"errors"
	//"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go-zero-demo/greet/internal/logic/model"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	/*1.先检验合法性
	2.查询数据库是否存在
	还需要判断传进的密码跟号码是否跟数据库里面的值一致  登录成功(暂时不需要)
	*/
	if len(req.Mobile) != 11 || req.Password != "123456" {
		return nil, errors.New("参数错误")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	switch err {
	case model.ErrNotFound:
		return nil, errors.New("用户未注册，请先注册后在登录")
	}
	now := time.Now().Unix()
	//userInfo去拿值是直接去看usermodel_gen.go 里面User struct的
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)

	if err != nil {
		return nil, err
	}
	//TODO  待完善 验证密码
	return &types.LoginReply{
		200,
		"登录成功",
		jwtToken}, nil
}