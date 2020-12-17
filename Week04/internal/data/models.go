package models

import (
	"Go-000/Week04/internal/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/unknwon/com"
	"log"
	_  "github.com/go-sql-driver/mysql"
	"time"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn int `json:"deleted_on"`
}

func Setup() {

	var (
		err error
	)

	databaseSetting := setting.DatabaseSetting

	address := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		databaseSetting.User,databaseSetting.Password,databaseSetting.Host,databaseSetting.Name)

	db, err = gorm.Open(databaseSetting.Type,address)
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return databaseSetting.TablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForCreateCallback)
	db.Callback().Create().Replace("gorm:update_time_stamp",updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

}

var updateTimeFieldList = []string{"CreatedOn", "ModifiedOn"}

func updateTimeStampForCreateCallback(scop *gorm.Scope) {
	if scop.HasError() {
		return
	}
	now := time.Now()
	for _, fieldName := range updateTimeFieldList {
		if timeField, ok := scop.FieldByName(fieldName); ok {
			if timeField.IsBlank {
				timeField.Set(now)
			}
		}

	}
}

// updateTimeStampForUpdateCallback will set `ModifyTime` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func deleteCallback(scope *gorm.Scope)  {
	if scope.HasError() {
		return
	}
	var extraOption string
	if str, ok := scope.Get("gorm:delete_option"); ok{
		extraOption = com.ToStr(str)
	}
	deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
	//scope.AddToVars 该方法可以添加值作为SQL的参数，也可用于防范SQL注入
	var sprintSql string
	if !scope.Search.Unscoped && hasDeletedOnField {
		sprintSql = fmt.Sprintf(
			"UPDATE %v SET %v=%v%v%v",
			scope.QuotedTableName(),
			scope.Quote(deletedOnField.DBName),
			scope.AddToVars(time.Now().Unix()),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
		)
	}else{
		sprintSql = fmt.Sprintf(
			"DELETE FROM %v%v%v",
			scope.QuotedTableName(),
			addExtraSpaceIfExist(scope.CombinedConditionSql()),
			addExtraSpaceIfExist(extraOption),
		)
	}
	scope.Raw(sprintSql).Exec()

}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

