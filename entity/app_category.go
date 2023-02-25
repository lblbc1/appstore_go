package entity

/*
厦门大学计算机专业 | 前华为工程师
专注《零基础学编程系列》  http://lblbc.cn/blog
包含：Java | 安卓 | 前端 | Flutter | iOS | 小程序 | 鸿蒙
公众号：蓝不蓝编程
*/
type AppCategory struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	App_id      string    `gorm:"type:varchar(50)" json:"app_id"`
	Categroy_id string    `gorm:"type:varchar(50)" json:"categroy_id"`
	Apps        []AppInfo `gorm:"references:App_id"`
}

func (AppCategory) TableName() string {
	return "appstore_app_category"
}
