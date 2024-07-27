package sequence

import (
	"database/sql"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const sqlReplaceIntoStub = `replace into sequence (stub) values ('a')`

// 建立一个mysql连接，执行replace insert
type Mysql struct {
	conn sqlx.SqlConn
}

func NewMysql(dsn string) *Mysql {
	return &Mysql{conn: sqlx.NewMysql(dsn)}
}

func (m *Mysql) Next() (seq uint64, err error) {
	var stmt sqlx.StmtSession
	stmt, err = m.conn.Prepare(sqlReplaceIntoStub)
	if err != nil {
		logx.Errorw("conn.Prepare failed: ", logx.LogField{
			Key:   "err",
			Value: err.Error(),
		})
	}
	defer stmt.Close()

	// 执行sql语句
	var rest sql.Result
	rest, err = stmt.Exec()
	if err != nil {
		logx.Errorw("stmt.Exec() failed: ", logx.LogField{
			Key:   "err",
			Value: err.Error(),
		})
		return 0, err
	}

	lid, err := rest.LastInsertId()
	if err != nil {
		logx.Errorw("rest.LastInsertId failed: ", logx.LogField{
			Key:   "err",
			Value: err.Error(),
		})
		return 0, err
	}
	return uint64(lid), nil
}
