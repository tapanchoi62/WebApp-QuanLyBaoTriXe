package services

import (
	"log"

	"github.com/robfig/cron/v3"
)

// Hàm setup cronjob
func StartCronJobs() {
	c := cron.New()

	// chạy mỗi giờ
	_, err := c.AddFunc("* * 30 * *", func() {
		log.Println("Cron: dọn dẹp file rác...")
		CleanupUnusedFiles()
	})

	if err != nil {
		log.Println("Không thể thêm cronjob:", err)
	}

	c.Start()
}
