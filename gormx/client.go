package gormx

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"

	"gorm.io/gorm"
	gormLog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

type Client struct {
	Conf   Conf
	GormDb *gorm.DB
}

func MustNewClient(conf Conf) *Client {
	client, _ := NewClient(conf)
	return client
}

func NewClient(conf Conf) (*Client, error) {
	gormDb, err := GetSqlDriver(conf)
	if err != nil {
		return nil, err
	}
	return &Client{
		Conf:   conf,
		GormDb: gormDb,
	}, nil
}

func GetSqlDriver(conf Conf) (*gorm.DB, error) {
	var dbDialector gorm.Dialector
	if val, err := getDbDialector("w", conf); err == nil {
		dbDialector = val
	}

	gormDb, err := gorm.Open(dbDialector, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 gormLog.Default.LogMode(gormLog.Info),
		//Logger: redefineLog(), //拦截、接管 gorm v2 自带日志
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Write.Prefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	if conf.IsOpenReadDb {
		if val, err := getDbDialector("r", conf); err == nil {
			dbDialector = val
		}
		resolverConf := dbresolver.Config{
			Replicas: []gorm.Dialector{dbDialector}, //  读 操作库，查询类
			Policy:   dbresolver.RandomPolicy{},     // sources/replicas 负载均衡策略适用于
		}
		err = gormDb.Use(
			dbresolver.Register(resolverConf).
				SetConnMaxLifetime(time.Duration(conf.Read.SetConnMaxLifetime) * time.Second).
				SetMaxIdleConns(conf.Read.SetMaxIdleConn).
				SetMaxOpenConns(conf.Read.SetMaxIdleConn),
		)
		if err != nil {
			//gorm 数据库驱动初始化失败
			return nil, err
		}
	}

	// 查询没有数据，屏蔽 gorm v2 包中会爆出的错误
	// https://github.com/go-gorm/gorm/issues/3789  此 issue 所反映的问题就是我们本次解决掉的
	_ = gormDb.Callback().Query().Before("gorm:query").Register("disable_raise_record_not_found", MaskNotDataError)
	// https://github.com/go-gorm/gorm/issues/4838
	_ = gormDb.Callback().Create().Before("gorm:before_create").Register("CreateBeforeHook", CreateBeforeHook)
	// 为了完美支持gorm的一系列回调函数
	_ = gormDb.Callback().Update().Before("gorm:before_update").Register("UpdateBeforeHook", UpdateBeforeHook)

	// 为主连接设置连接池(43行返回的数据库驱动指针)
	if rawDb, err := gormDb.DB(); err != nil {
		return nil, err
	} else {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetConnMaxLifetime(time.Duration(conf.Read.SetConnMaxLifetime) * time.Second)
		rawDb.SetMaxIdleConns(conf.Write.SetMaxIdleConn)
		rawDb.SetMaxOpenConns(conf.Write.SetMaxIdleConn)
		return gormDb, nil
	}

}

// 获取一个数据库方言(Dialector),通俗的说就是根据不同的连接参数，获取具体的一类数据库的连接指针
func getDbDialector(rw string, conf Conf) (gorm.Dialector, error) {
	var dbDialector gorm.Dialector
	dsn := getDsn(rw, conf)
	switch conf.SqlType {
	case MysqlType:
		dbDialector = mysql.Open(dsn)
	case PostgresqlType:
		dbDialector = postgres.Open(dsn)
	case SqlServerType:
		dbDialector = sqlserver.Open(dsn)
	default:
		return nil, errors.New("ErrorsDbDriverNotExists")
	}
	return dbDialector, nil
}

func getDsn(rw string, conf Conf) string {
	Host, DataBase, User, Pass, Charset, Port := "192.168.3.88", "db", "root", "123123", "utf8", 3306

	if rw == "r" {
		Host = conf.Read.Host
		DataBase = conf.Read.DataBase
		Port = conf.Read.Port
		User = conf.Read.User
		Pass = conf.Read.Pass
		Charset = conf.Read.Charset
	} else {
		Host = conf.Write.Host
		DataBase = conf.Write.DataBase
		Port = conf.Write.Port
		User = conf.Write.User
		Pass = conf.Write.Pass
		Charset = conf.Write.Charset
	}
	switch conf.SqlType {
	case MysqlType:
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=false&loc=Local", User, Pass, Host, Port, DataBase, Charset)
	case PostgresqlType:
		return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", Host, Port, DataBase, User, Pass)
	case SqlServerType:
		return fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=disable", Host, Port, DataBase, User, Pass)
	}
	return ""
}

// 创建自定义日志模块，对 gorm 日志进行拦截、
func redefineLog() gormLog.Interface {
	return createCustomGormLog(
		SetInfoStrFormat("[info] %s\n"), SetWarnStrFormat("[warn] %s\n"), SetErrStrFormat("[error] %s\n"),
		SetTraceStrFormat("[traceStr] %s [%.3fms] [rows:%v] %s\n"), SetTracWarnStrFormat("[traceWarn] %s %s [%.3fms] [rows:%v] %s\n"),
		SetTracErrStrFormat("[traceErr] %s %s [%.3fms] [rows:%v] %s\n"))
}
