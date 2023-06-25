package usecase

import (
	"log"
	"sync"
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

	g := errgroup.Group{}
	
	var dates []entity.Date
	var times []entity.Time

	g.Go(func() error {
		var err error
		dates, err = u.scheduleRepo.GetDateByCounselorId(counselorId)
		if err != nil {
			return schedule.ErrInternalServerError
		}
		return nil
	})

	g.Go(func() error {
		var err error
		times, err = u.scheduleRepo.GetTimeByCounselorId(counselorId)
		if err != nil {
			return schedule.ErrInternalServerError
		}
		return nil
	})
	
	var scheduleRes schedule.GetAllResponse

	if err := g.Wait(); err != nil {
		return scheduleRes, err
	}

	var datesRes = make([]string, len(dates))
	var timesRes = make([]string, len(times))

	wg := sync.WaitGroup{}

	for i, date := range dates {
		wg.Add(1)
		go func(i int, date entity.Date) {
			defer wg.Done()
			datesRes[i] = date.Date.Format("2006-01-02")
		}(i, date)
	}

	for i, timeDb := range times {
		wg.Add(1)
		go func(i int, timeDb entity.Time) {
			defer wg.Done()
			timeParsed, _ := time.Parse("15:04:05", timeDb.Time)
			timesRes[i] = timeParsed.Format("15:04")
		}(i, timeDb)
	}
	wg.Wait()

	scheduleRes.Dates = datesRes
	scheduleRes.Times = timesRes

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

	datesSaved, err := u.scheduleRepo.GetDateByCounselorId(input.CounselorId)

	if err != nil {
		return schedule.ErrInternalServerError
	}	

	if len(datesSaved) > 0 {
		return schedule.ErrScheduleAlreadyExist
	}

	g := errgroup.Group{}

	var times = make([]entity.Time, len(input.Times))
	var dates = make([]entity.Date, len(input.Dates))

	for i, date := range input.Dates {
		date := date
		i := i
		g.Go(func() error {
			id, err := u.UuidGenerator.GenerateUUID()

			if err != nil {
				return schedule.ErrInternalServerError
			}

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
			id, err := u.UuidGenerator.GenerateUUID()

			if err != nil {
				return schedule.ErrInternalServerError
			}


			_, err = time.Parse("15:04", timedata)
			
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

	g := errgroup.Group{}

	err = u.checkScheduleExist(counselorId, g)
	if err != nil {
		return err
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

	err = u.checkScheduleExist(input.CounselorId, g)
	if err != nil {
		return err
	}
	
	var dates = make([]entity.Date, len(input.Dates))

	for i, date := range input.Dates {
		date := date
		i := i
		g.Go(func() error {
			id, err := u.UuidGenerator.GenerateUUID()

			if err != nil {
				return schedule.ErrInternalServerError
			}

			parsedDate, err := time.Parse(time.DateOnly, date)

			if err != nil {
				return schedule.ErrDateInvalid
			}

			newDate := entity.Date{
				ID: id,
				CounselorID: input.CounselorId,
				Date: parsedDate,
			}

			dates[i] = newDate

			return nil
		})	
	}

	var times = make([]entity.Time, len(input.Times))
	for i, timedata := range input.Times {
		
		timedata := timedata
		i := i
		g.Go(func() error {
			id, err := u.UuidGenerator.GenerateUUID()

			if err != nil {
				return schedule.ErrInternalServerError
			}

			_, err = time.Parse("15:04", timedata)

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
		log.Println(err.Error())
		return schedule.ErrInternalServerError
	}


	return nil
	
}

func(u *scheduleUsecase) checkScheduleExist(counselorId string, g errgroup.Group) error {

	g.Go(func() error {
		datesSaved, err := u.scheduleRepo.GetDateByCounselorId(counselorId)
		
		if err != nil {
			return schedule.ErrInternalServerError
		}
		
		if len(datesSaved) == 0 {
			return schedule.ErrScheduleNotFound
		}
		return nil
	})

	g.Go(func() error {
		timesSaved, err := u.scheduleRepo.GetTimeByCounselorId(counselorId)
		
		if err != nil {
			return schedule.ErrInternalServerError
		}
	
		if len(timesSaved) == 0 {
			return schedule.ErrScheduleNotFound
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
