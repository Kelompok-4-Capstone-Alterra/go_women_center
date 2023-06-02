package repository

import (
	"log"
	"sync"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
)

type Otp struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type cachedOtp struct {
	Otp
	expireAtTimestamp int64
}

type LocalCache interface {
	StopCleanup()
	Update(o Otp, expireAtTimestamp int64)
	Read(email string) (Otp, error)
	Delete(email string)
}

type localCache struct {
	stop chan struct{}

	wg    sync.WaitGroup
	mu    sync.RWMutex
	codes map[string]cachedOtp
}

// create an instance of local cache,
// that will do cleanup at set interval time
// defined by the function parameter
func NewLocalCache(cleanupInterval time.Duration) *localCache {
	lc := &localCache{
		codes: make(map[string]cachedOtp),
		stop:  make(chan struct{}),
	}

	lc.wg.Add(1)
	go func(cleanupInterval time.Duration) {
		defer lc.wg.Done()
		lc.cleanupLoop(cleanupInterval)
	}(cleanupInterval)

	return lc
}

func (lc *localCache) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for {
		select {
		case <-lc.stop:
			return
		case <-t.C:
			lc.mu.Lock()
			for otpEmail, cu := range lc.codes {
				if cu.expireAtTimestamp <= time.Now().Unix() {
					delete(lc.codes, otpEmail)
				}
			}
			log.Println("cleanup executed")
			lc.mu.Unlock()
		}
	}
}

func (lc *localCache) StopCleanup() {
	close(lc.stop)
	lc.wg.Wait()
}

func (lc *localCache) Update(o Otp, expireAtTimestamp int64) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	lc.codes[o.Email] = cachedOtp{
		Otp:               o,
		expireAtTimestamp: expireAtTimestamp,
	}
}

func (lc *localCache) Read(email string) (Otp, error) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()

	co, ok := lc.codes[email]
	if !ok {
		return Otp{}, constant.ErrInvalidOtp
	}

	if time.Now().Unix() >= co.expireAtTimestamp {
		return Otp{}, constant.ErrExpiredOtp
	}

	return co.Otp, nil
}

func (lc *localCache) Delete(email string) {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	delete(lc.codes, email)
}
