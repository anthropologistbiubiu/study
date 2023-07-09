package day10

import "time"

func (l *Lock) TryLock(ctx context.Context) error {
	success, err := l.client.SetNX(ctx, l.resource, l.randomValue, ttl).Result()
	if err != nil {
		return err
	}
	// 加锁失败
	if !success {
		return ErrLockFailed
	}
	// 加锁成功
	l.randomValue = randomValue
	return nil
}

func (l *Lock) Unlock(ctx context.Context) error {
	return l.script.Run(ctx, l.client, []string{l.resource}, l.randomValue).Err()
}

func (l *Lock) Lock(ctx context.Context) error {
	// 尝试加锁
	err := l.TryLock(ctx)
	if err == nil {
		return nil
	}
	if !errors.Is(err, ErrLockFailed) {
		return err
	}
	// 加锁失败，不断尝试
	ticker := time.NewTicker(l.tryLockInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			// 超时
			return ErrTimeout
		case <-ticker.C:
			// 重新尝试加锁
			err := l.TryLock(ctx)
			if err == nil {
				return nil
			}
			if !errors.Is(err, ErrLockFailed) {
				return err
			}
		}
	}
}
