package entity

/*
厦门大学计算机专业 | 前华为工程师
专注《零基础学编程系列》  http://lblbc.cn/blog
包含：Java | 安卓 | 前端 | Flutter | iOS | 小程序 | 鸿蒙
公众号：蓝不蓝编程
*/
type AppInfo struct {
	ID              uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name            string `gorm:"type:varchar(200)" json:"name"`
	Logo_url        string `gorm:"type:varchar(200)" json:"logoUrl"`
	Screenshot_urls string `gorm:"type:varchar(1000)" json:"screenShotUrls"`
	Description     string `gorm:"type:varchar(1000)" json:"description"`
	Apk_url         string `gorm:"type:varchar(500)" json:"apkUrl"`
	File_size       string `gorm:"type:varchar(50)" json:"fileSize"`
	Download_count  string `gorm:"type:varchar(50)" json:"downloadCount"`
}

func (AppInfo) TableName() string {
	return "appstore_app"
}
