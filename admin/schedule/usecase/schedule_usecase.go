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
	dateRepo repo.DateRepository
	timeRepo repo.TimeRepository
	UuidGenerator helper.UuidGenerator
}

func NewScheduleUsecase(
		counselorRepo Counselor.CounselorRepository,
		dateRepo repo.DateRepository,
		timeRepo repo.TimeRepository,
		uidGenerator helper.UuidGenerator,
	) ScheduleUsecase {

	return &scheduleUsecase{counselorRepo, dateRepo, timeRepo, uidGenerator}
}

func(u *scheduleUsecase) GetByCounselorId(counselorId string) (schedule.GetAllResponse, error) {
	
	_, err := u.counselorRepo.GetById(counselorId)

	if ; err != nil {
		log.Println(err)
		return schedule.GetAllResponse{}, schedule.ErrCounselorNotFound
	}

	g := errgroup.Group{}

	
	
	var schedules schedule.GetAllResponse

	
	dates, err := u.dateRepo.GetByCounselorId(counselorId)

	if err != nil {
		log.Println(err)
		return schedules, schedule.ErrInternalServerError
	}

	datesRes := make([]string, len(dates))

	for i, date := range dates {
		i := i
		date := date
		g.Go(func() error {
			datesRes[i]	= date.Date.Format(time.DateOnly)
			return nil
		})
	}

	times, err := u.timeRepo.GetByCounselorId(counselorId)

	if err != nil {
		log.Println(err)
		return schedule.GetAllResponse{}, schedule.ErrInternalServerError
	}

	timeRes := make([]string, len(times))

	for i, time := range times {
		i := i
		time := time
		g.Go(func() error {
			timeRes[i] = time.Time
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Println(err.Error())
		return schedules, schedule.ErrInternalServerError
	}

	schedules.Time = timeRes
	schedules.Date = datesRes

	return schedules, nil
}

func(u *scheduleUsecase) Create(input schedule.CreateRequest) error {

	_, err := u.counselorRepo.GetById(input.CounselorId)

	if err != nil {
		log.Println(err.Error())
		return schedule.ErrCounselorNotFound
	}
	
	g := errgroup.Group{}

	for _, date := range input.Dates {
		date := date
		g.Go(func() error {
			id, _ := u.UuidGenerator.GenerateUUID()

			date, err := time.Parse(time.DateOnly, date)

			if err != nil {
				return schedule.ErrInternalServerError
			}

			newDate := entity.Date{
				ID: id,
				CounselorID: input.CounselorId,
				Date: date,
			}

			return u.dateRepo.Create(newDate)
		})	
	}

	for _, timedata := range input.Times {
		
		timedata := timedata
		g.Go(func() error {
			id, _ := u.UuidGenerator.GenerateUUID()

			newTime := entity.Time{
				ID: id,
				CounselorID: input.CounselorId,
				Time: timedata,
			}

			return u.timeRepo.Create(newTime)
		})
	}

	if err := g.Wait(); err != nil {
		log.Println(err.Error())
		return schedule.ErrInternalServerError
	}

	return nil
}

func(u *scheduleUsecase) Delete(counselorId string) error {

	_, err := u.counselorRepo.GetById(counselorId)

	if err != nil {
		log.Println(err.Error())
		return schedule.ErrCounselorNotFound
	}
	g := errgroup.Group{}

	g.Go(func() error {
		return u.dateRepo.DeleteByCounselorId(counselorId)
	})

	g.Go(func() error {
		return u.timeRepo.DeleteByCounselorId(counselorId)
	})

	if err := g.Wait(); err != nil {
		log.Println(err.Error())
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

	g.Go(func() error {
		return u.dateRepo.DeleteByCounselorId(input.CounselorId)
	})

	g.Go(func() error {
		return u.timeRepo.DeleteByCounselorId(input.CounselorId)
	})

	g.Go(func() error {
		return u.Create(schedule.CreateRequest(input))
	})

	if err := g.Wait(); err != nil {
		log.Println(err.Error())
		return schedule.ErrInternalServerError
	}

	return nil
	
}

