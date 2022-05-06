package mysqlExample

import (
	"time"
)

type BaseModel struct {
	ID        string     `gorm:"primary_key;type:varchar(36);" json:"ID"` // 数据库ID
	CreatedAt time.Time  `json:"CreatedAt"`                               // 创建时间
	UpdatedAt time.Time  `json:"UpdatedAt"`                               // 更新时间
	DeletedAt *time.Time `sql:"index" json:"DeletedAt,omitempty"`         // 删除时间
}

type MountDisk struct {
	BaseModel
	DiskID       string `gorm:"type:VARCHAR(60)" json:"disk_id"`      //源端需要迁移的磁盘ID
	DiskType     string `json:"disk_type"`                            //源端需要迁移主机的磁盘类型
	DiskDev      string `json:"disk_dev"`                             //源端磁盘的dev
	SrcIP        string `json:"src_ip"`                               //源端主机的IP地址
	SrcUUID      string `gorm:"type:VARCHAR(60)" json:"src_uuid"`     //源端主机LongId
	DestIP       string `json:"dest_ip"`                              //目标主机的IP地址
	DestUUID     string `gorm:"type:VARCHAR(60)" json:"dest_uuid"`    //目标主机的LongId
	Status       string `json:"status"`                               //需要迁移磁盘的状态
	Percent      int    `json:"percent"`                              //任务进度百分比
	OperatorName string `gorm:"type:VARCHAR(60)" json:"OperatorName"` // 操作者名称
	Logs         string `json:"Logs" gorm:"type:LONGTEXT"`            // 日志
}

//func main() {
//	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
//	dsn := "root:@tcp(127.0.0.1:3306)/unsafe?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		//Logger:logger.Default.LogMode(logger.Info),
//	})
//	if err != nil {
//		fmt.Errorf("%s", err)
//	}
//	var mountDisk MountDisk
//	selectResult := db.Where("id = ?", "03cc5d0c-4707-4d01-9971-15cd9aae4995").Find(&mountDisk)
//	fmt.Println(mountDisk)
//	fmt.Println(selectResult)
//	mountDisk.Logs = "test2"
//	//db.Debug().Select("logs").Save(&mountDisk)
//	db.Debug().Model(&mountDisk).Update("logs", "test3")
//	fmt.Println(mountDisk)
//
//	//var mountDisks []MountDisk
//	//result := db.Debug().Model(&mountDisks).Where(
//	//		"id = ?","03cc5d0c-4707-4d01-9971-15cd9aae4995").Update(
//	//			"logs", "test1").Find(&mountDisks)
//	//fmt.Println(result)
//	//fmt.Println(mountDisks)
//}
