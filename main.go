package main

import (
    "database/sql"
    "fmt"
    "github.com/bitly/go-simplejson"
    "log"

    gormSql "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var (
    gControllerDB1 *sql.DB /*控制台的db*/

    gOrmDB1 *gorm.DB
)

type PolicyModel1 struct {
    Id            int64 `gorm:"primaryKey"`
    Name          string
    Order         int

}

type teacher1 struct {
    Id string `gorm:"primaryKey"`
    Name string
    Student []student

}

type student1 struct {
    Id uint `gorm:"primaryKey"`
    Name string
    TeacherId uint
}


type RuleExprs1 struct {
    PreOperator  string `json:"preOperator"`
    DictId       string  `json:"dictId"`
    DictName     string `json:"dictName"`
    PostOperator string `json:"postOperator"`
    Value        string `json:"value"`
}


func InitDB1() error {
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

type Conditions1 struct {
    Operator string `json:"operator"'`
    Value    string `json:"value"'`
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
    log.Printf("11112%v", t)

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
    body1 := []byte{}

    log.Printf("11112%v", string(body1))

    body := `{"1":{"3":2},"2":{"1":[{"1":"22","2":{"1":"thread-a:r-6958452352180844951","2":{"14":{"1":{"1":"msg-a:r-7038016324606246701","2":{"1":1,"2":"xuyong.manyou@gmail.com","3":"yong xu","10":"xuyong.manyou@gmail.com"},"3":[{"1":1,"2":"xuyong_manyou@163.com"},{"1":1,"2":"xuyong.manyou@gmail.com","3":"yong xu"}],"7":"1624254974168","8":"ccc","9":{"2":[{"1":0,"2":"<div dir=\"ltr\">测试1111</div>"}],"7":1},"11":["^all","^pfg","^f_bt","^f_btns","^f_cl","^i","^u","^io_im","^io_imc3"],"12":[{"1":"application/octet-stream","2":"policy_order的副本.go","3":"12","5":"f_kq67gmom0","6":"https://mail.google.com/mail/?ui=2&ik=4e6b9f3294&attid=0.1&permmsgid=msg-a:r-7038016324606246701&view=att&realattid=f_kq67gmom0&zw","7":0,"8":{"1":"/gmail/att/660272317269/AAXFQFOiJrrp7XUk2bxR4g.4/796","2":"12","3":"04eb6f1b_daf51ae4_8ed5cf01_53fad859_3467b778","4":2133364169,"5":"16","7":"b64magic:NK,f,76","9":"ANGjdJ-8bgZbLbsTqwEaEZxmqJNWS1w6SjjwPxXx9qg5G0DnLjX6EO8aeC65a1cMYF2PL0QXAX0I5liTgQgWXayQtFOyYc71Z3p7Uk4frXBZTRt3wvNtRnY_R1hoOKM"},"10":"GkEvYmxvYnN0b3JlL3Byb2QvZ21haWwtdXBsb2FkLzIxZmNlNmIxLTU0ODYtNDA4Yi05NzdhLTBjNjBiNjE0MGNhMigMMiwwNGViNmYxYl9kYWY1MWFlNF84ZWQ1Y2YwMV81M2ZhZDg1OV8zNDY3Yjc3ODjJm6L5B0gQWANiEGI2NG1hZ2ljOk5LLGYsNzZoAJoBQDUxMjg0Mzg1NWZjYzkyYTUxYzgxMGIxYjU4ZTA3MzFjMDFlYWM5YTZhMjNjMTU3YmZhMDJhYWQ3MWVkZmZiZTeqAQIIAQ==","11":1}],"18":"1624254974168","36":{"6":0},"37":{"4":0},"42":0,"43":{"1":0,"2":0,"3":10,"4":0},"52":"s:24e655a6266ddca2|#msg-a:r-7038016324606246701|0"},"3":1}}}}]},"4":{"1":"1624254956210","2":1,"3":"1624254974179","4":1,"5":25},"5":2}`
    json, _ := simplejson.NewJson([]byte(body))

    attachmentsMap, _ := json.Get("2").Get("1").GetIndex(0).Get("2").Get("2").Get("14").Get("1").Get("12").Array()

    //attachmentSize:= len(attachmentsMap)

    // fmt.Println("body is %v", json.Get("2").Get("1"))
    fmt.Println("body is %v", attachmentsMap)

    tmp := make(map[string]string)
    tmp["1"] ="1"
    tmp["2"] ="2"
    fmt.Println("body is %v", tmp["3"])
}



