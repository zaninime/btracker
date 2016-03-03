package db

import (
	"math"
	"time"
)

var (
	pingChannels struct {
		time chan bool
		exit chan bool
	}
	pingBackoff = backoffManager{
		10 * time.Second,
		120 * time.Second,
		1.3591409142295225,
		0,
	}
)

type backoffManager struct {
	Min    time.Duration
	Max    time.Duration
	Factor float64
	count  int
}

func (bm *backoffManager) Next() (time.Duration, bool) {
	v := time.Duration(math.Pow(bm.Factor, float64(bm.count))) * bm.Min
	if v > bm.Max {
		return 0, true
	}
	bm.count++
	return v, false
}

func (bm *backoffManager) Reset() {
	bm.count = 0
}

func startPing() {
	logger.Info("scheduling backoff ping to postgres")
	pingChannels.time = make(chan bool)
	pingChannels.exit = make(chan bool)
	go pinger()
	go func() {
		v, _ := pingBackoff.Next()
		<-time.After(v)
		pingChannels.time <- true
	}()
}

func stopPing() {
	pingChannels.exit <- true
	pingBackoff.Reset()
}

func pinger() {
	for {
		select {
		case <-pingChannels.time:
			logger.Debug("pinging database")
			if err := DB.Ping(); err != nil {
				logger.Warn("ping failed", "err", err)
			} else {
				logger.Debug("ping successful")
				pingBackoff.Reset()
			}
			go func() {
				v, expired := pingBackoff.Next()
				if expired {
					logger.Fatal("database is not available, backoff expired")
				}
				logger.Debug("rescheduling ping", "delay", v.String())
				<-time.After(v)
				pingChannels.time <- true
			}()
		case <-pingChannels.exit:
			logger.Debug("ping stopped")
			return
		}
	}
}
