package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type CodeLog struct {
	Log_id 		int64
	Log_code	string
	Type		string
	Severity	int
	Source		string
	Filepath	string
	Line_no		int
	Message		string
}

var conStr string = "ashim:ashim@tcp(127.0.0.1:3306)/ersolv"


func GetCodeLog() (CodeLog, error) {
	var db *sql.DB
	var err error
	var log CodeLog
	var query string = 
`select logs.log_id, logs.log_code, type, severity, source, filepath, line_no, message_string
from logs, log_types, code_logs, messages
where logs.log_code = log_types.log_code and
logs.log_id = code_logs.log_id and
code_logs.message_id = messages.message_id;`

	db, err = sql.Open("mysql", conStr)
	if err != nil {
		return CodeLog{}, err
	}
	defer db.Close()

	err = db.QueryRow(query).Scan(
		&log.Log_id,
		&log.Log_code,
		&log.Type,
		&log.Severity,
		&log.Source,
		&log.Filepath,
		&log.Line_no,
		&log.Message,
	)
	if err != nil {
		return CodeLog{}, err
	}
	/* Fill the message read from the database with appropriate data.
	 * TODO: Can we use localisation to format the message?
	 */
	log.Message = fmt.Sprintf(log.Message, log.Line_no, log.Filepath)

	return log, nil
}

func CreateCodeLog(log CodeLog) (int64, error) {
	var db *sql.DB
	var err error
	var tx *sql.Tx
	var query string
	var res sql.Result
	var log_id int64

	var queryInsertLogs string =
`insert into logs(log_code)
values('%s')`

	var queryInsertCodeLogs string =
`insert into code_logs(log_id, filepath, line_no, message_id)
values(%d, '%s', %d, %d)`


	db, err = sql.Open("mysql", conStr)
	if err != nil {
		return -1, err
	}
	defer db.Close()

	tx, err = db.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	query = fmt.Sprintf(queryInsertLogs, log.Log_code)
	res, err = tx.Exec(query)
	if err != nil {
		return -1, err
	}

	log_id, err = res.LastInsertId()
	if err != nil {
		return -1, err
	}

	// Code smell: magic number
	query = fmt.Sprintf(queryInsertCodeLogs,
		log_id,
		log.Filepath,
		log.Line_no,
		1)
	res, err = tx.Exec(query)
	if err != nil {
		return -1, err
	}

	// tx.Commit()
	return log_id, nil
}
