package manager

import (
	"backend/internal/models"
	"backend/internal/repositories"
	"time"
)

type AdapterLogger struct {
	logRepo *repositories.AdapterLogRepository
}

func NewAdapterLogger(logRepo *repositories.AdapterLogRepository) *AdapterLogger {
	return &AdapterLogger{logRepo: logRepo}
}

func (l *AdapterLogger) LogRun(name string) {
	log := &models.AdapterLog{
		AdapterName: name,
		RunAt:       time.Now(),
	}
	l.logRepo.Create(log)
}
