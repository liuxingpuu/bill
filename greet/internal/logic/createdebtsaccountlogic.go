package logic

import (
	"context"
	"encoding/json"
	"errors"
	"go-zero-demo/greet/internal/logic/model"
	"reflect"

	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"
	"go-zero-demo/greet/internal/common"
	"github.com/zeromicro/go-zero/core/logx"
)

type Create_debts_accountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreate_debts_accountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Create_debts_accountLogic {
	return &Create_debts_accountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Create_debts_accountLogic) Create_debts_account(req *types.CreateDebtReq) (resp *types.CreateDebtReply, err error) {
	// todo: add your logic here and delete this line
	if len(req.Type) ==0 || len(req.Account_name) == 0 || req.Debt_amount == 0 {
		return nil, errors.New("参数错误")
	}
	userId := l.ctx.Value("userId")
	logx.Info("hqhqqqq")
	logx.Info(reflect.TypeOf(userId))
	int64Val, err := userId.(json.Number).Int64()
	if err != nil{
		return nil, errors.New("user_id类型转换失败")
	}
	data := model.Debt{
		Type: req.Type,
		UserId: int64Val,
		DebtAmount: common.Y2F(req.Debt_amount),
		AccountName: req.Account_name,
	}
	_, err = l.svcCtx.DebtModel.Insert(l.ctx, &data)
	if err != nil{
		logx.Info("写入数据失败")
		return nil, errors.New("创建负债账户失败")
	}
	logx.Info("创建负债账户成功")
	//TODO 等list接口,要查询所有数据然后求和debe_amout,最后获取一个负债总额
	return &types.CreateDebtReply{
		Type: req.Type,
		Debt_amount: req.Debt_amount,
		Account_name: req.Account_name,
	}, nil
}
