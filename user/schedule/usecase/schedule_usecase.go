package usecase

import (
	"sync"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/entity"
	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/schedule"
	repo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/schedule/repository"
	transactionRepo "github.com/Kelompok-4-Capstone-Alterra/go_women_center/user/transaction/repository"
	"golang.org/x/sync/errgroup"
)

type ScheduleUsecase interface {
	GetCurrSchedule(counselorId string) (schedule.GetScheduleResponse, error)
}

// TODO tambahkan repo transaction
type scheduleUsecase struct {
	scheduleRepo repo.ScheduleRepository
	transRepo  transactionRepo.MysqlTransactionRepository
}

func NewScheduleUsecase(scheduleRepo repo.ScheduleRepository, transRepo  transactionRepo.MysqlTransactionRepository) ScheduleUsecase{
	return &scheduleUsecase{scheduleRepo, transRepo}
}

func(u *scheduleUsecase) GetCurrSchedule(counselorId string) (schedule.GetScheduleResponse, error) {

	g := errgroup.Group{}
	
	var curTransactions []entity.Transaction

	g.Go(func() error {
		var err error
		curTransactions, err = u.transRepo.GetOccurTransacTodayByCounselorId(counselorId)
		return err
	})

	var timeCounselor []entity.Time

	g.Go(func() error {
		var err error
		timeCounselor, err = u.scheduleRepo.GetTimeByCounselorId(counselorId)
		return err
	})

	var currDatesCounselor entity.Date

	g.Go(func() error {
		var err error
		currDatesCounselor, err = u.scheduleRepo.GetCurDateByCounselorId(counselorId)
		return err
	})

	if err := g.Wait(); err != nil {
		if err.Error() == "record not found" {
			return schedule.GetScheduleResponse{}, nil
		}
		return schedule.GetScheduleResponse{}, schedule.ErrInternalServerError
	}

	var scheduleTimes = make([]schedule.Time, len(timeCounselor))
	
	transactionMap := make(map[string]bool)

	wg := sync.WaitGroup{}

	// get all transaction time id, and put it in map
	for _, transaction := range curTransactions {
		wg.Add(1)
		go func(transaction entity.Transaction) {
			defer wg.Done()
			transactionMap[transaction.TimeId] = true
		}(transaction)
	}
	wg.Wait()

	// check if time is available
	for i, time := range timeCounselor {
		wg.Add(1)
		go func(i int, time entity.Time) {
			defer wg.Done()
			scheduleTime := schedule.Time{
				ID:          time.ID,
				Time:        time.Time,
				IsAvailable: !transactionMap[time.ID],
			}
			scheduleTimes[i] = scheduleTime
		}(i, time)
	}
	wg.Wait()
	

	return schedule.GetScheduleResponse{
		DateId: currDatesCounselor.ID,
		Times: scheduleTimes,
	}, nil
}