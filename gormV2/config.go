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
	SetMaxIdleConns    int
	SetMaxOpenConns    int
	SetConnMaxLifetime int
}

//type Conf struct {
//	SlowThreshold int
//	IsOpenReadDb  bool
//	Write         struct {
//		Host               string
//		Port               int
//		User               string
//		Pass               string
//		Charset            string
//		DataBase           string
//		Prefix             string
//		SetMaxIdleConns    int
//		SetMaxOpenConns    int
//		SetConnMaxLifetime int
//	}
//	Read struct {
//		Host               string
//		Port               int
//		User               string
//		Pass               string
//		Charset            string
//		DataBase           string
//		Prefix             string
//		SetMaxIdleConns    int
//		SetMaxOpenConns    int
//		SetConnMaxLifetime int
//	}
//}
