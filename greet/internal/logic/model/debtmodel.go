package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DebtModel = (*customDebtModel)(nil)

type (
	// DebtModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDebtModel.
	DebtModel interface {
		debtModel
	}

	customDebtModel struct {
		*defaultDebtModel
	}
)

// NewDebtModel returns a model for the database table.
func NewDebtModel(conn sqlx.SqlConn, c cache.CacheConf) DebtModel {
	return &customDebtModel{
		defaultDebtModel: newDebtModel(conn, c),
	}
}
