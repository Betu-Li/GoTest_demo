package main

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

// TestShouldUpdateStats 测试 sql 执行成功的案例
func TestShouldUpdateStats(t *testing.T) {
	// mock 一个*sql.DB链接对象，不需要连接真实的数据库
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mock执行指定SQL语句时的返回结果
	mock.ExpectBegin() // 开始事务
	// 执行UPDATE语句, 返回1行受影响的记录
	mock.ExpectExec("UPDATE products").
		WillReturnResult(sqlmock.NewResult(1, 1))
	// 执行INSERT语句, 返回1行受影响的记录
	mock.ExpectExec("INSERT INTO product_viewers").
		WithArgs(2, 3).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit() // 提交事务

	// 将mock的DB对象传入我们的函数中
	if err = recordStats(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// 确保期望的结果都满足
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// TestShouldRollbackStatUpdatesOnFailure sql执行失败回滚的测试用例
func TestShouldRollbackStatUpdatesOnFailure(t *testing.T) {
	// mock 一个*sql.DB链接对象，不需要连接真实的数据库
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// mock执行指定SQL语句时的返回结果
	mock.ExpectBegin() // 开始事务
	// 执行UPDATE语句, 返回1行受影响的记录
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	// 执行INSERT语句, 返回错误
	mock.ExpectExec("INSERT INTO product_viewers").
		WithArgs(2, 3).
		WillReturnError(fmt.Errorf("some error"))
	mock.ExpectRollback() // 回滚事务

	// 将mock的DB对象传入我们的函数中
	if err = recordStats(db, 2, 3); err == nil {
		t.Errorf("was expecting an error, but there was none")
	}

	// 确保期望的结果都满足
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
