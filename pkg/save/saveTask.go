package save

import (
	"context"
	"time"
)

func StartTask(saveService SaveService) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()
	doEvery(ctx, 30*time.Second, SaveData, saveService)
}

func doEvery(ctx context.Context, d time.Duration, f func(time.Time, SaveService), saveService SaveService) error {
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case x := <-ticker.C:
			go f(x, saveService)
		}
	}
}

func SaveData(t time.Time, saveService SaveService) {
	saveService.Save()
}