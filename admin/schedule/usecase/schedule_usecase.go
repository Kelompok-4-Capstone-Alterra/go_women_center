package usecase

import (
	"log"
	"time"

	Counselor "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/counselor/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule"
	repo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/admin/schedule/repository"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"golang.org/x/sync/errgroup"
)

type ScheduleUsecase interface {
	GetByCounselorId(counselorId string) (schedule.GetAllResponse, error)
	Create(input schedule.CreateRequest) error
	Delete(counselorId string) error
	Update(input schedule.UpdateRequest) error
}

type scheduleUsecase struct {
	counselorRepo Counselor.CounselorRepository
	scheduleRepo repo.ScheduleRepository
	UuidGenerator helper.UuidGenerator
}

func NewScheduleUsecase(
		counselorRepo Counselor.CounselorRepository,
		scheduleRepo repo.ScheduleRepository,
		uidGenerator helper.UuidGenerator,
	) ScheduleUsecase {

	return &scheduleUsecase{counselorRepo, scheduleRepo, uidGenerator}
}

func(u *scheduleUsecase) GetByCounselorId(counselorId string) (schedule.GetAllResponse, error) {
	
	_, err := u.counselorRepo.GetById(counselorId)

	if err != nil {
		log.Println(err)
		return schedule.GetAllResponse{}, schedule.ErrCounselorNotFound
	}

	scheduleRes, err := u.scheduleRepo.GetByCounselorId(counselorId)

	if err != nil {
		log.Println(err.Error())
		return scheduleRes, schedule.ErrInternalServerError
	}

	return scheduleRes, nil
}

func(u *scheduleUsecase) Create(input schedule.CreateRequest) error {

	_, err := u.counselorRepo.GetById(input.CounselorId)

	if err != nil {
		log.Println(err.Error())
		return schedule.ErrCounselorNotFound
	}
	
	g := errgroup.Group{}

	var times = make([]entity.Time, len(input.Times))
	var dates = make([]entity.Date, len(input.Dates))

	for i, date := range input.Dates {
		date := date
		i := i
		g.Go(func() error {
			id, _ := u.UuidGenerator.GenerateUUID()

			date, err := time.Parse(time.DateOnly, date)

			if err != nil {
				return schedule.ErrDateInvalid
			}
			
			newDate := entity.Date{
				ID: id,
				CounselorID: input.CounselorId,
				Date: date,
			}

			dates[i] = newDate
			return nil
		})	
	}

	for i, timedata := range input.Times {
		
		timedata := timedata
		i := i
		g.Go(func() error {
			id, _ := u.UuidGenerator.GenerateUUID()

			_, err := time.Parse(time.TimeOnly, timedata)

			if err != nil {
				return schedule.ErrTimeInvalid
			}

			newTime := entity.Time{
				ID: id,
				CounselorID: input.CounselorId,
				Time: timedata,
			}

			times[i] = newTime

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Println(err.Error())
		return err
	}

	err = u.scheduleRepo.Create(dates, times)

	if err != nil {
		return schedule.ErrCounselorNotFound
	}

	return nil
}

func(u *scheduleUsecase) Delete(counselorId string) error {

	_, err := u.counselorRepo.GetById(counselorId)

	if err != nil {
		log.Println(err.Error())
		return schedule.ErrCounselorNotFound
	}

	err = u.scheduleRepo.DeleteByCounselorId(counselorId)

	if err != nil {
		return schedule.ErrInternalServerError
	}

	return nil
	
}

func(u *scheduleUsecase) Update(input schedule.UpdateRequest) error {

	_, err := u.counselorRepo.GetById(input.CounselorId)

	if err != nil {
		log.Println(err.Error())
		return schedule.ErrCounselorNotFound
	}

	g := errgroup.Group{}

	var times = make([]entity.Time, len(input.Times))
	var dates = make([]entity.Date, len(input.Dates))

	for i, date := range input.Dates {
		date := date
		i := i
		g.Go(func() error {
			id, _ := u.UuidGenerator.GenerateUUID()

			date, err := time.Parse(time.DateOnly, date)

			if err != nil {
				return schedule.ErrDateInvalid
			}
			
			newDate := entity.Date{
				ID: id,
				CounselorID: input.CounselorId,
				Date: date,
			}

			dates[i] = newDate
			return nil
		})	
	}

	for i, timedata := range input.Times {
		
		timedata := timedata
		i := i
		g.Go(func() error {
			id, _ := u.UuidGenerator.GenerateUUID()

			_, err := time.Parse(time.TimeOnly, timedata)

			if err != nil {
				return schedule.ErrTimeInvalid
			}

			newTime := entity.Time{
				ID: id,
				CounselorID: input.CounselorId,
				Time: timedata,
			}

			times[i] = newTime

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Println(err.Error())
		return err
	}

	err = u.scheduleRepo.Update(input.CounselorId, dates, times)

	if err != nil {
		return schedule.ErrInternalServerError
	}


	return nil
	
}

