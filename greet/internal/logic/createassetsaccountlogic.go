package logic

import (
	"context"
	"encoding/json"
	"errors"
	"go-zero-demo/greet/internal/common"
	"go-zero-demo/greet/internal/logic/model"
	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"reflect"

	"github.com/zeromicro/go-zero/core/logx"
)

type Create_assets_accountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreate_assets_accountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Create_assets_accountLogic {
	return &Create_assets_accountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Create_assets_accountLogic) Create_assets_account(req *types.CreateAssetsReq) (resp *types.CreateAssetsReply, err error) {
	// todo: add your logic here and delete this line
	/*1.检查是否缺少必传参数
	2.存入数据库
	3.返回*/
	if len(req.Type) == 0 || len(req.Account_name) == 0 {
		return nil, errors.New("参数错误")
	}

	//currentTime := time.Now()

	userId := l.ctx.Value("userId")
	logx.Info(reflect.TypeOf(userId))
	int64Val, err := userId.(json.Number).Int64()
	if err != nil{
		return nil , errors.New("user_id类型转换失败")
	}

	data := model.Account{
		Type: req.Type,
		UserId: int64Val,
		AssetsBalance:common.Y2F(req.Assets_balance),
		AccountName: req.Account_name,
	}

	_, err = l.svcCtx.AccountModel.Insert(l.ctx, &data)
	logx.Info(err)
	if err != nil {
		logx.Info("写入数据失败")
		return nil , errors.New("创建资产账户失败")
	}

	logx.Info("创建资产账户成功")
	return &types.CreateAssetsReply{
		req.Type,
		req.Assets_balance,
		req.Account_name,
	}, nil
}


type Update_assets_accountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}


func NewUpdate_assets_accountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Update_assets_accountLogic {
	return &Update_assets_accountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *Update_assets_accountLogic) Update_assets_account(req *types.UpdateAssetsReq) (resp *types.UpdateAssetsReply, err error) {
	// todo: add your logic here and delete this line

	userId := l.ctx.Value("userId")
	int64Val, err := userId.(json.Number).Int64()
	data := model.Account{
		Id: req.Id,
		UserId: int64Val,
		Type: req.Type,
		AssetsBalance: common.Y2F(req.Assets_balance),
		AccountName: req.Account_name,
	}
	_ = l.svcCtx.AccountModel.Update(l.ctx, &data)
	logx.Info(err)

	if err != nil {
		logx.Info("更新数据失败")
		return nil, errors.New("编辑资产账户信息失败")
	}

	logx.Info("编辑资产账户信息成功")
	return &types.UpdateAssetsReply{
		req.Type,
		req.Assets_balance,
		req.Account_name,
	}, nil
}

type Delete_assets_accountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelete_assets_accountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Delete_assets_accountLogic {
	return &Delete_assets_accountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Delete_assets_accountLogic) Delete_assets_account(req *types.DeleteAssetsReq) (resp *types.DeleteAssetsReply, err error) {
	// todo: add your logic here and delete this line
	_ = l.svcCtx.AccountModel.Delete(l.ctx, req.Id)
	if err != nil {
		logx.Info("删除失败", err)
		return nil, errors.New("删除失败")
	}
	logx.Info("删除成功")
	return nil, nil
}

type Get_assets_accountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGet_assets_accountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Get_assets_accountLogic {
	return &Get_assets_accountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Get_assets_accountLogic) Get_assets_account(req *types.GetAssetsReq) (resp *types.GetAssetsReqReply, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.AccountModel.FindOne(l.ctx, req.Id)

	if err != nil {
		logx.Info("查询失败", err)
		return nil, errors.New("查询失败")
	}
	logx.Info("查询成功")
	return &types.GetAssetsReqReply{
		res.Type,
		common.F2Y(res.AssetsBalance),
		res.AccountName,
	},nil
}

type Assets_accountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssets_accountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Assets_accountLogic {
	return &Assets_accountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Assets_accountLogic) Assets_account() (resp *types.ListAssetsReply, err error) {
	// todo: add your logic here and delete this line
	var as []*types.Assets
	accounts, err := l.svcCtx.AccountModel.FindALl(l.ctx)
	if err != nil{
		logx.Info("数据有误", err)
		return nil, errors.New("查询失败")
	}

	count := int32(len(accounts))
	for _, account := range accounts{
		as = append(as, &types.Assets{
			account.Type,
			common.F2Y(account.AssetsBalance),
			account.AccountName,
		})
	}
	return &types.ListAssetsReply{
		count,
	     as,
	}, nil
}

