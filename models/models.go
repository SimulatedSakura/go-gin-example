package models

import (
	"fmt"
	"log"
	"time"

	"github.com/SimulatedSakura/go-gin-example/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// 创建时的回调函数
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	//检查是否有错误
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		//获取所有的字段检查是否存在所需要的字段
		if createTimeFiled, ok := scope.FieldByName("CreatedOn"); ok {
			//IsBlank 可以用来判断所选的区块是不是空值
			if createTimeFiled.IsBlank {
				//给该字段设置值
				createTimeFiled.Set(nowTime)
			}
		}
		//同时第一次修改时间应当与第一次创建的时间相同
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 修改调用的函数
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	/*
		scope.Get(...)会根据参数获取设置了字面值的参数,比方说在本文中
		"gorm:update_column" 没有对应的参数,所以OK是false
		scope.SetColumn(...)假设没有指定update_column的字段,我们默认在更新回调设置ModifiedOn的值
	*/
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// 删除时使用的回调函数
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {

		//可能会有需要执行的更长的sql语句
		var extraOption string

		//检查是否手动配置了 delete_operation
		if str, ok := scope.Get("gorm:delete_operation"); ok {
			extraOption = fmt.Sprint(str)
		}

		//获取我们约定的删除字段,若存在UPDATE软删除,若不存在DELETE硬删除
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		if !scope.Search.Unscoped && hasDeletedOnField {
			//此处实在以fmt格式化后的SQL语句执行
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),                            //获取当前的表的名字
				scope.Quote(deletedOnField.DBName),                 //获取要删除的标记列在数据表中的名字
				scope.AddToVars(time.Now().Unix()),                 //获取系统当前的时间
				addExtraSpaceIfExist(scope.CombinedConditionSql()), //返回组合好的sql语句
				addExtraSpaceIfExist(extraOption),                  //如果你想做其他sql附属条件
			)).Exec()
			fmt.Println("11111")
		} else {
			fmt.Println("22222")
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	//创建和修改的回调
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)

	//删除的回调
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
