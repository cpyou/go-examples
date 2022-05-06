package mysqlExample

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGorm(t *testing.T) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/unsafe?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Errorf("%s", err)
	}
	var mountDisk MountDisk
	selectResult := db.Where("id = ?", "03cc5d0c-4707-4d01-9971-15cd9aae4995").Find(&mountDisk)
	fmt.Println(mountDisk)
	fmt.Println(selectResult)
	mountDisk.Logs = "test2"
	//db.Debug().Select("logs").Save(&mountDisk)
	//db.Debug().Model(&mountDisk).Update("logs", "test3")
	//fmt.Println(mountDisk)

	//var mountDisks []MountDisk
	//result := db.Debug().Model(&mountDisks).Where(
	//		"id = ?","03cc5d0c-4707-4d01-9971-15cd9aae4995").Update(
	//			"logs", "test1").Find(&mountDisks)
	//fmt.Println(result)
	//fmt.Println(mountDisks)

}
