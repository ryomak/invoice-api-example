package helper

import (
	"fmt"
	"github.com/ryomak/invoice-api-example/infrastructure/client/db"
	"github.com/tanimutomo/sqlfile"
	"path"
	"runtime"
)

const (
	mockDataFilePathFormat = "%s/mock/mock_data.sql"
)

func TestDBSetup(conn *db.Conn) error {

	//schema
	_, pwd, _, _ := runtime.Caller(0)

	mockDataFilePath := fmt.Sprintf(mockDataFilePathFormat, path.Dir(pwd))
	if err := executeFromSqlFile(conn, mockDataFilePath); err != nil {
		return err
	}
	return nil
}

func executeFromSqlFile(conn *db.Conn, filepath string) error {
	s := sqlfile.New()
	if err := s.File(filepath); err != nil {
		return err
	}
	_, err := s.Exec(conn.DB)
	return err
}
