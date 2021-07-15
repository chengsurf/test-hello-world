package main

import (
    "database/sql"
    "fmt"
    gormSql "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "strings"
)

var (
    gControllerDB *sql.DB /*控制台的db*/

    gOrmDB *gorm.DB
)

type PolicyModel struct {
    Id            int64 `gorm:"primaryKey"`
    Name          string
    Order         int

}

type teacher struct {
    Id string `gorm:"primaryKey"`
    Name string
    Student []student

}

type student struct {
    Id uint `gorm:"primaryKey"`
    Name string
    TeacherId uint
}


type RuleExprs struct {
    PreOperator  string `json:"preOperator"`
    DictId       string  `json:"dictId"`
    DictName     string `json:"dictName"`
    PostOperator string `json:"postOperator"`
    Value        string `json:"value"`
}


func InitDB() error {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", "root", "xy_123456", "127.0.0.1", 3306, "test")
    log.Print(dsn)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Print(err)
        return err
    }

    err = db.Ping()
    if err != nil {
        log.Print(err)
        return err
    }

    db.SetMaxOpenConns(128)
    db.SetMaxIdleConns(32)

    gControllerDB = db

    gOrmDB, err = gorm.Open(gormSql.New(gormSql.Config{Conn: gControllerDB}), &gorm.Config{CreateBatchSize: 100,SkipDefaultTransaction: false,})
    if err != nil {
        return err
    }

    return nil
}

type Conditions struct {
    Operator string `json:"operator"'`
    Value    string `json:"value"'`
}
func GetBetweenStr(str, start, end string) string {
    n := strings.Index(str, start)
    if n == -1 {
       return ""
    }


    str = string([]byte(str)[n + len(start):])
    m := strings.Index(str, end)
    if m == -1 {
        m = len(str)
    }
    str = string([]byte(str)[:m])
    return str
}


func main() {
    InitDB()
    user := teacher{
        Id:   "3",
        Name: "jinzhu111",
        Student: []student{
            {Id: 4,
                Name: "jinzhu@example.com"},
        },
    }
    gOrmDB.Updates(&user)

    var t teacher

    var param map[string]interface{} = make(map[string]interface{})
    param["name"] = gorm.Expr("unix_timestamp(now(6)) * ?", 100000)

    db := gOrmDB.Table("teachers")
    res := db.Where("id in (?)", []string{"2"}).Updates(param)
    if res.Error != nil {
        log.Printf("11112%v", res.Error)
    }

    gOrmDB.Table("teachers").Where("id in (?)", []string{"2"}).Find(&t)
  ///  log.Printf("11112%v", t)

    //sql := "update teachers set id = unix_timestamp(now(6)) * 100000 where id =2"
    //gControllerDB.Exec(sql)
    //gOrmDB.Table("teachers").Where("id in (?)", []string{"2"}).Find(&t)
    //log.Printf("11133%v",t)

    //policy := PolicyModel{2, "114", 2}
    //var policyList []PolicyModel
    //policyList = append(policyList, PolicyModel{1, "111", 1})
    //policyList = append(policyList, PolicyModel{2, "112", 2})
    //policyList = append(policyList, PolicyModel{3, "113", 3})
    //
    //p := updatePolicyFilter(&policy,policyList)
    //log.Printf("11111%v",policy)
    //log.Printf("11112%v",p)

   // body := "[{\"id\":1,\"fileKey\":\"210712051905zecelv\",\"fileName\":\"main.go\",\"contentType\":\"application/octet-stream\",\"size\":3637,\"fileUrl\":\"https://webmail30.189.cn/w2/mail/previewImage.do?fileKey\\u003d210712051905zecelv%2Fmain.go\",\"code\":0}]"
    privValue :=0x00000001|0x00000020
    privValueInDB:=0x00000020

    log.Printf("11112 %v",(privValue & privValueInDB))



}



