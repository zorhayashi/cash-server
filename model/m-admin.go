package model

import (
	"cash-server/db"
	"cash-server/pkg/util"
	"fmt"

	"github.com/jinzhu/gorm"
)

//Platform struct
type Platform struct {
	Platformid          int    `json:"platform_id"`
	Platformaccount     string `json:"platform_account"`
	Platformpassword    string `json:"platform_password"`
	Platformname        string `json:"platform_name"`
	Platformgroupid     string `json:"platform_group_id"`
	PlatformEmail       string `json:"platform_email"`
	Platformtoken       string `json:"platform_token"`
	Platformtokensecret string `json:"platform_token_secret"`
	Status              string `json:"status"`
	Createdate          string `json:"create_date"`
	Updatedate          string `json:"update_date"`
}

//--------------------------- ALL 相關-------------------------------------

func dbErrBool(dbrut *gorm.DB) bool {
	if dbrut.Error != nil {
		return false
	}
	return true
}

//---------------------------payment_platform  表單相關-------------------------------------

//PlatformAdd Platform 註冊
func PlatformAdd(p db.Platform) {
	db.SQLDBX.Create(&p)
}

//PlatformCheck Platform 查詢
func PlatformCheck(account string) bool {
	model := &db.Platform{}
	dbrtu := db.SQLDBX.Where("platform_account = ?", account).First(&model)
	return dbErrBool(dbrtu)
}

//PlatformQueryExist 查詢帳號存在
func PlatformQueryExist(data db.Platform) db.Platform {
	//model := &db.Platform{}
	var model db.Platform
	k, v := data.DBFind()
	db.SQLDBX.Where(k+"= ?", v).First(&model)
	return model
}

//PlatformStatusEnableUpdata 更新帳號 開狀態
func PlatformStatusEnableUpdata(data db.Platform) bool {
	var model db.Platform
	k, v := data.DBFind()
	db.SQLDBX.Where(k+"= ?", v).First(&model)
	model.Status = "1"
	dbrut := db.SQLDBX.Save(model)
	return dbErrBool(dbrut)
}

//PlatformStatusDisableUpdata 更新帳號 關狀態
func PlatformStatusDisableUpdata(data db.Platform) bool {
	var model db.Platform
	k, v := data.DBFind()
	db.SQLDBX.Where(k+"= ?", v).First(&model)
	model.Status = "0"
	dbrut := db.SQLDBX.Save(model)
	return dbErrBool(dbrut)
}

//PlatformQueryOne 查詢帳號存在
func PlatformQueryOne(data db.Platform) string {
	model := &db.Platform{}
	k, v := data.DBFind()
	db.SQLDBX.Where(k+"= ?", v).First(&model)
	util.Test(model.PlatformAccount)
	return model.PlatformAccount
}

//PlatformQueryStatus 查詢帳號開通狀態
func PlatformQueryStatus(account string, password string) string {
	model := &db.Platform{}
	a := db.SQLDBX.Where("platform_account = ?", account).First(&model)
	if a.Error != nil {
		return "err"
	}
	util.Test(fmt.Sprint("PlatformQueryStatus 查詢帳號開通狀態 -> 帳號：", account, "狀態：", model.Status))
	return model.Status
}

//PlatformTokenQueryStatus 查詢帳號開通狀態
func PlatformTokenQueryStatus(token string) string {
	model := &db.Platform{}
	a := db.SQLDBX.Where("platform_token = ?", token).First(&model)
	if a.Error != nil {
		return "err"
	}
	util.Test(fmt.Sprint("PlatformTokenQueryStatus 查詢帳號狀態 -> token：", token))
	return model.Status
}

//PlatformQueryStatusUseToken  查詢帳號STAUTS資料  用TOKEN
func PlatformQueryStatusUseToken(token string) string {
	model := &db.Platform{}
	a := db.SQLDBX.Where("platform_token = ?", token).First(&model)
	if a.Error != nil {
		return "err"
	}
	util.Test(fmt.Sprint("PlatformQueryStatus 查詢帳號開通狀態 -> 帳號：", token, "狀態：", model.Status))
	return model.Status
}

//PlatformQueryInfoJSON  查詢帳號詳細資料 乖 用id查
func PlatformQueryInfoJSON(taskID string) string {
	sql := "SELECT * FROM payment_platform where platform_id =?"
	rows, err := db.GetJSON(sql, taskID)
	util.Test(fmt.Sprint("PlatformQueryInfo 查詢帳號資料 ", rows))
	if err != nil {
		util.Error(err.Error())
		return "err"
	}
	return rows
}

//PlatformQueryInfoAllJSON  查詢ALL帳號資料
func PlatformQueryInfoAllJSON() []db.Platform {
	util.Test(fmt.Sprint("查詢All Platform帳號資料 "))
	var platforms []db.Platform
	db.SQLDBX.Find(&platforms)
	return platforms
}

//PlatformQueryInfodataJSON   查詢帳號資料
func PlatformQueryInfodataJSON(token string) Platform {
	util.Test(fmt.Sprint("查詢Platform帳號資料 :", token))
	var platforms Platform
	rows, err := db.SQLDB.Query("select * from payment_platform WHERE platform_token=? ", token)
	if err != nil {
		util.Error(err.Error())
	}
	for rows.Next() {
		rows.Scan(&platforms.Platformid, &platforms.Platformaccount, &platforms.Platformpassword, &platforms.Platformname, &platforms.Platformgroupid, &platforms.PlatformEmail, &platforms.Platformtoken, &platforms.Platformtokensecret, &platforms.Status, &platforms.Createdate, &platforms.Updatedate)
		//fmt.Printf("%+v", platforms)
	}
	return platforms
}
