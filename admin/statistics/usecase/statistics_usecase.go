package usecase

import (
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/statistics"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/statistics/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
)

type StatisticUsecase interface {
	GetStatistic() (statistics.GetStatisticsResponse, error) 
}

type statisticUseCase struct {
	Repo repository.StatisticRepo
}

func NewAuthUsecase(repo repository.StatisticRepo) StatisticUsecase {
	return &statisticUseCase{
		Repo: repo,
	}
}

func (au *statisticUseCase) GetStatistic() (statistics.GetStatisticsResponse, error) {
	data := statistics.GetStatisticsResponse{}
	
	userCount, err := au.Repo.GetRowCount(entity.User{})
	if err != nil {
		return statistics.GetStatisticsResponse{}, err
	}
	data.TotalUser = userCount

	counselorCount, err := au.Repo.GetRowCount(entity.Counselor{})
	if err != nil {
		return statistics.GetStatisticsResponse{}, err
	}
	data.TotalCounselor = counselorCount

	transactionCount, err := au.Repo.GetRowCount(entity.Transaction{})
	if err != nil {
		return statistics.GetStatisticsResponse{}, err
	}
	data.TotalTransaction = transactionCount

	return data, nil
}