// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	accountFieldNames          = builder.RawFieldNames(&Account{})
	accountRows                = strings.Join(accountFieldNames, ",")
	accountRowsExpectAutoSet   = strings.Join(stringx.Remove(accountFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	accountRowsWithPlaceHolder = strings.Join(stringx.Remove(accountFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheAccountIdPrefix = "cache:account:id:"
)

type (
	accountModel interface {
		Insert(ctx context.Context, data *Account) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Account, error)
		Update(ctx context.Context, data *Account) error
		Delete(ctx context.Context, id int64) error
		FindALl(ctx context.Context) ([]*Account, error)
	}

	defaultAccountModel struct {
		sqlc.CachedConn
		table string
	}

	Account struct {
		Id            int64     `db:"id"`             // 主键
		UserId        int64     `db:"user_id"`        // 用户id
		Type          string    `db:"type"`           // 资产类型
		AssetsBalance int64   `db:"assets_balance"` // 资产余额
		AccountName    string    `db:"account_name"`    // 账户名
		CreateTime   time.Time `db:"create_time"`
		UpdateTime   time.Time `db:"update_time"`
	}
)

func newAccountModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAccountModel {
	return &defaultAccountModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`account`",
	}
}

func (m *defaultAccountModel) Insert(ctx context.Context, data *Account) (sql.Result, error) {
	accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, accountRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Type, data.AssetsBalance, data.AccountName)
	}, accountIdKey)
	return ret, err
}

func (m *defaultAccountModel) FindOne(ctx context.Context, id int64) (*Account, error) {
	accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, id)
	var resp Account
	err := m.QueryRowCtx(ctx, &resp, accountIdKey,
		func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
			query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", accountRows, m.table)
			return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:

		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAccountModel) Update(ctx context.Context, data *Account) error {
	accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, accountRowsWithPlaceHolder)
		logx.Info(conn.ExecCtx(ctx, query, data.UserId, data.Type, data.AssetsBalance, data.AccountName))
		return conn.ExecCtx(ctx, query, data.UserId, data.Type, data.AssetsBalance, data.AccountName,data.Id)
	}, accountIdKey)
	return err
}

func (m *defaultAccountModel) Delete(ctx context.Context, id int64) error {
	accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, accountIdKey)
	return err
}

func (m *defaultAccountModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAccountIdPrefix, primary)
}

func (m *defaultAccountModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", accountRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAccountModel) tableName() string {
	return m.table
}

func (m *defaultAccountModel) FindALl(ctx context.Context) ([]*Account, error) {
	//accountIdKey := fmt.Sprintf("%s%v", cacheAccountIdPrefix, id)
	var resp []*Account
	query := fmt.Sprintf("select * from %s", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
