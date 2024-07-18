package redis

type RedisMutex interface {
	Lock() error
	UnLock() error
}
