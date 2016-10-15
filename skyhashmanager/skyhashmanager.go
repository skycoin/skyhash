package skyhashmanager

import ()

type SkyhashManagerConfig struct {
}

type SkyhashManager struct {
}

func NewSkyhashManager(SkyhashManagerConfig *SkyhashManagerConfig) *SkyhashManager {
	shm := SkyhashManager{}
	return &shm
}

func (self *SkyhashManager) Start() {

}

func (self *SkyhashManager) Shutdown() {

}
