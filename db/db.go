package db

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/iancoleman/strcase"
)

type Client struct {
	db *sql.DB
}

func New(dbuser, dbpass, dbhost, dbname, socket string, port int) (*Client, error) {
	userpass := fmt.Sprintf("%s:%s", dbuser, dbpass)
	var conn string
	if socket != "" {
		conn = fmt.Sprintf("unix(%s)", socket)
	} else {
		conn = fmt.Sprintf("tcp(%s:%d)", dbhost, port)
	}

	if dbname == "" {
		dbname = "mysql"
	}

	_db, err := sql.Open("mysql", fmt.Sprintf("%s@%s/%s", userpass, conn, dbname))

	if err != nil {
		return nil, err
	}

	return &Client{
		db: _db,
	}, nil
}

type Columns map[string]string

func NewColumns(col string) Columns {
	c := Columns{}
	c["Field"] = col

	return c
}

func (c Columns) Column() string {
	return strcase.ToCamel(c["Field"])
}

func (c *Client) execute(_sql string, args ...interface{}) ([]Columns, error) {
	var rows *sql.Rows
	var err error
	if len(args) == 0 {
		rows, err = c.db.Query(_sql)
	} else {
		rows, err = c.db.Query(_sql, args...)
	}

	if err != nil {
		return []Columns{}, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return []Columns{}, err
	}

	values := make([][]byte, len(columns))
	row := make([]interface{}, len(columns))
	for i, _ := range values {
		row[i] = &values[i]
	}

	data := make([]Columns, 0)

	for rows.Next() {
		if err := rows.Scan(row...); err != nil {
			return []Columns{}, err
		}

		r := make(Columns)
		for i, val := range values {
			v := string(val)
			r[columns[i]] = v
		}

		data = append(data, r)
	}

	return data, nil
}

func (c *Client) TableRows(dbname, sort string, reverse bool) ([][]string, error) {
	baseSql := "SELECT table_name, table_rows FROM `information_schema`.`TABLES` WHERE table_schema = ? ORDER BY %s %s"
	orderBy := "ASC"
	if reverse {
		orderBy = "DESC"
	}
	_sql := fmt.Sprintf(baseSql, sort, orderBy)

	cols, err := c.execute(_sql, dbname)
	if err != nil {
		return [][]string{}, err
	}

	data := make([][]string, 0)
	for _, col := range cols {
		data = append(data, []string{col["TABLE_NAME"], col["TABLE_ROWS"]})
	}

	return data, nil
}

func (c *Client) GetTables(dbname string, reverse bool) ([]string, error) {
	baseSql := "SELECT table_name FROM `information_schema`.`TABLES` WHERE table_schema = ? ORDER BY table_name %s"
	orderBy := "ASC"
	if reverse {
		orderBy = "DESC"
	}
	_sql := fmt.Sprintf(baseSql, orderBy)

	cols, err := c.execute(_sql, dbname)
	if err != nil {
		return []string{}, err
	}

	data := make([]string, 0)
	for _, col := range cols {
		data = append(data, col["TABLE_NAME"])
	}

	return data, nil
}

func (c *Client) GetGlobalVariable(name string) (string, error) {
	_sql := fmt.Sprintf("SHOW GLOBAL VARIABLES LIKE '%s'", name)

	cols, err := c.execute(_sql)
	if err != nil {
		return "", err
	}

	return cols[0]["Value"], nil
}

func (c *Client) SlowlogOn(datadir string, longQueryTime float64, persist bool) error {
	set := "SET GLOBAL"
	if persist {
		set = "SET PERSIST"
	}

	_, err := c.execute(set + " slow_query_log=1")
	if err != nil {
		return err
	}

	slowlogPath := filepath.Join(strings.TrimRight(datadir, "/"), "slow.log")
	_, err = c.execute(fmt.Sprintf(set+` slow_query_log_file="%s"`, slowlogPath))
	if err != nil {
		return err
	}

	_, err = c.execute(fmt.Sprintf(set+` long_query_time=%s`, fmt.Sprint(longQueryTime)))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SlowlogOff(persist bool) error {
	set := "SET GLOBAL"
	if persist {
		set = "SET PERSIST"
	}

	_, err := c.execute(set + " slow_query_log=0")
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Close() {
	c.db.Close()
}
