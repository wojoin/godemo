package main

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"user-auth-ng-bak/dao/mysql"
	"user-auth-ng/dao"
)

// Transaction begin
type Transaction struct {
	once     sync.Once
	rollback bool
	tx       *gorm.DB
}

func NewTransaction(db *gorm.DB) (*gorm.DB, *Transaction) {
	return db, &Transaction{tx: db}
}

func (t *Transaction) Close() {
	t.once.Do(func() {
		if t.rollback {
			t.tx.Rollback()
		} else {
			t.tx.Commit()
		}
	})
}

func (t *Transaction) Failed() {
	t.rollback = true
}

// Transaction end

var DuplicatedErr = errors.New("Duplicate")

func BatchInsert(txndb *gorm.DB, data interface{}, batchSize int) (int64, error) {

	dataSlice, ok := convSliceData(data)
	if !ok {
		return 0, nil
	}

	length := len(dataSlice)

	var total int64 = 0

	if length == 0 {
		return 0, nil
	}

	tailSize := length % batchSize
	tailIndex := length / batchSize

	for index := 0; index < tailIndex+1; index++ {
		offset := index * batchSize
		limit := (index + 1) * batchSize
		if index == tailIndex {
			limit = offset + tailSize
		}

		affected, err := batchInsertAux(txndb, dataSlice[offset:limit])
		if err != nil && !strings.Contains(err.Error(), DuplicatedErr.Error()) {
			return 0, err
		}

		total = total + affected

		//fmt.Println("rows affected: ", affected)
	}

	return total, nil
}

func batchInsertAux(db *gorm.DB, data interface{}) (int64, error) {

	relations := map[string]bool{
		"belongs_to":   true,
		"has_one":      true,
		"has_many":     true,
		"many_to_many": true,
	}

	//errDatas := make(map[int]string,0)

	dataSlice, ok := convSliceData(data)
	if !ok {
		return 0, nil
	}

	length := len(dataSlice)

	if length == 0 {
		return 0, nil
	}

	obj := dataSlice[0]
	objScope := db.NewScope(obj)
	fields := objScope.Fields()
	quoted := make([]string, 0, len(fields))

	for i := range fields {
		if fields[i].IsIgnored || (fields[i].IsPrimaryKey && fields[i].IsBlank) ||
			(fields[i].Relationship != nil && relations[fields[i].Relationship.Kind]) {
			continue
		}

		quoted = append(quoted, fields[i].DBName)
	}

	placeholdersArr := make([]string, 0, length)

	for _, o := range dataSlice {
		s := db.NewScope(o)
		fs := s.Fields()
		placeholders := make([]string, 0, len(fs))

		for i := range fs {
			if fs[i].IsIgnored || (fs[i].IsPrimaryKey && fs[i].IsBlank) ||
				(fields[i].Relationship != nil && relations[fields[i].Relationship.Kind]) {
				continue
			}

			var vars interface{}
			if (fs[i].Name == "CreatedAt" || fs[i].Name == "UpdatedAt") && fs[i].IsBlank {
				vars = gorm.NowFunc()
			} else {
				vars = fs[i].Field.Interface()
			}
			placeholders = append(placeholders, s.AddToVars(vars))
		}

		placeholdersStr := "(" + strings.Join(placeholders, ",") + ")"
		placeholdersArr = append(placeholdersArr, placeholdersStr)

		objScope.SQLVars = append(objScope.SQLVars, s.SQLVars...)
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;",
		objScope.QuotedTableName(),
		strings.Join(quoted, ", "),
		strings.Join(placeholdersArr, ", "))

	objScope.Raw(sql)

	if err := objScope.Exec().DB().Error; err != nil {
		if strings.Contains(err.Error(), DuplicatedErr.Error()) {
			fmt.Println("Duplicated, sql: ", sql)
			return 0, DuplicatedErr
		}
	}

	fmt.Println("table ", objScope.QuotedTableName())

	return objScope.DB().RowsAffected, nil
}

func convSliceData(arg interface{}) (out []interface{}, ok bool) {
	s, ok := convData(arg, reflect.Slice)
	if !ok {
		ok = false
		return
	}

	count := s.Len()
	out = make([]interface{}, count)
	for i := 0; i < count; i++ {
		out[i] = s.Index(i).Interface()
	}
	return out, true
}

func convData(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

type User struct {
	ID   int    `json:"id" gorm:"primary_key" bson:"_id"`
	UID  int  `gorm:"uid"`
	Name string `gorm:"name;varchar(50);not null" bson:"name"`
}

func main() {

	users := []*User{
		&User{ID: 1, UID: 12345, Name: "name1"},
		&User{ID: 2, UID: 23456, Name: "name2"},
		&User{ID: 3, UID: 23457, Name: "name3"},
	}

	//for i, _ := range users {
	//	uid, _ := strconv.ParseInt(users[i].Name, 10, 64)
	//	u := dao.User{
	//		UID:  uid,
	//		Name: users[i].UID,
	//	}
	//}
	db := gorm.DB{}
	txndb, tx := NewTransaction(db.Begin())
	defer tx.Close()

	affected, err := dao.BatchInsert(txndb, users, 50)
	if err != nil && !strings.Contains(err.Error(), dao.DuplicatedErr.Error()) {
		tx.Failed()
		return
	}

	fmt.Println("rows affected: ", affected)

}
