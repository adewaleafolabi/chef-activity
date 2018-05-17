package repository

import (
	"fmt"
	"github.com/adewaleafolabi/chef-hero-activity/model"
	service2 "github.com/adewaleafolabi/chef-hero-activity/service"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type Repository interface {
	FindAllActivity() ([]*model.Activity, error)
	FindActivity(a *model.Activity) ([]*model.Activity, error)

	SaveActivity(a *model.Activity) (string, error)
}

type ActivityRepository struct {
	DB *gorm.DB
}

func (repo *ActivityRepository) FindAllActivity() ([]*model.Activity, error) {
	var activities []*model.Activity
	if err := repo.DB.Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

func (repo *ActivityRepository) FindActivity(activity *model.Activity) ([]*model.Activity, error) {
	activity.Data = ""
	var activities []*model.Activity
	var query = repo.DB

	date := strings.Split(activity.CreatedAt.Format(time.RFC3339), "T")
	if date[0] != "0001-01-01" {
		query = query.Where("created_at >= ?", date[0])
		activity.CreatedAt = time.Time{}
	}

	if activity.Message != "" {
		query = query.Where("message LIKE ?", fmt.Sprintf("%s%s%s", "%", activity.Message, "%"))
		activity.Message = ""
	}
	query = query.Where(&activity)

	if err := query.Find(&activities).Error; err != nil {
		logrus.Error(err)
		return nil, err
	}
	return activities, nil
}

func (repo *ActivityRepository) SaveActivity(activity *model.Activity) (string, error) {
	time := time.Now()
	idSvc := service2.IDGeneratorService{}
	activity.Id = idSvc.GenerateID()
	activity.CreatedAt = time
	err := repo.DB.Create(activity).Error
	if err != nil {
		return "", err
	}
	return activity.Id, err
}
