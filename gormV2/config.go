package gormV2

type SqlTypeEnum int

const (
	MysqlType SqlTypeEnum = iota
	PostgresqlType
	SqlServerType
)

type Conf struct {
	SqlType       SqlTypeEnum
	SlowThreshold int
	IsOpenReadDb  bool
	Write         ConfigParamsDetail
	Read          ConfigParamsDetail
}

type ConfigParamsDetail struct {
	Host               string
	DataBase           string
	Port               int
	Prefix             string
	User               string
	Pass               string
	Charset            string
	SetMaxIdleConn     int
	SetMaxOpenConn     int
	SetConnMaxLifetime int
}
