package persistence

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
	"strings"
	"sync"
	"time"
)

// performanceMap 高性能的并发安全KV存储
type performanceMap[T any] struct {
	slices    map[byte]*performanceMapSlice[uint, Factory[T]] // slices 存储Persistent的分表
	indexes   map[byte]*performanceMapSlice[string, uint]     // indexes 存储索引的分表
	subTables uint                                            // subTables 子表数量，不得大于256
	mutex     sync.Mutex                                      // mutex 控制idIndex的锁
	idIndex   uint                                            // idIndex 当前的自动生成ID
}

// Load 将records加载到performanceMap中，耗时操作，写锁
func (p *performanceMap[T]) Load(records []FactoryPersistenceRecord) {
	for _, record := range records {
		if record.ID != 0 && record.UID != "" {
			// 更新索引和Persistent分表
			entity := NewFactory[T](record.ID, record.UID)
			slice, index := p.hashUint(record.ID), p.hashString(record.UID)
			p.slices[slice].add(record.ID, entity)
			p.indexes[index].add(record.UID, record.ID)

			// 自动更新idIndex
			p.mutex.Lock()
			if p.idIndex < record.ID {
				p.idIndex = record.ID
			}
			p.mutex.Unlock()
		}
	}
}

// QueryByID 根据ID获取工厂，只返回可用结果
func (p *performanceMap[T]) QueryByID(id uint) (has bool, result Factory[T]) {
	slice := p.hashUint(id)
	if ok, entity := p.slices[slice].get(id); ok {
		if entity.Enable() {
			return true, entity
		} else {
			return false, NewFactory[T](0, "")
		}
	} else {
		return false, NewFactory[T](0, "")
	}
}

// QueryByUID 根据UID获取工厂，只返回可用结果
func (p *performanceMap[T]) QueryByUID(uid string) (has bool, result Factory[T]) {
	index := p.hashString(uid)
	if ok, id := p.indexes[index].get(uid); ok {
		return p.QueryByID(id)
	} else {
		return false, NewFactory[T](0, "")
	}
}

// Register 将一个工厂注册到performanceMap中
func (p *performanceMap[T]) Register(ctor func() T) (success bool) {
	entity := ctor()
	uid := p.generateUID(entity)
	index := p.hashString(uid)

	if ok, id := p.indexes[index].get(uid); ok {
		slice := p.hashUint(id)
		if has, record := p.slices[slice].get(id); has {
			if !record.Enable() {
				// 如果能找到UID，并且没有enable，说明是持久化文件中读取的，需要进行注册且不改变ID
				record.set(id, uid, ctor)
				record.enable()
				return true
			} else {
				// 如果能找到UID，但是enable了，说明注册过了，拒绝此次注册
				return false
			}
		} else {
			// 如果找到UID，但是找不到对应的记录，说明记录可能丢了，重新按照ID和UID记录一次
			record = NewFactory[T](id, uid)
			record.set(id, uid, ctor)
			record.enable()
			return p.slices[slice].add(id, record)
		}
	} else {
		// 如果没有找到UID记录，说明是新记录，直接写入
		id := p.generateID()
		slice := p.hashUint(id)
		p.indexes[index].add(uid, id)
		record := NewFactory[T](id, uid)
		record.set(id, uid, ctor)
		record.enable()
		return p.slices[slice].add(id, record)
	}
}

// Flush 将performanceMap中的数据全部取出，耗时操作，读锁
func (p *performanceMap[T]) Flush() (records []FactoryPersistenceRecord) {
	records = []FactoryPersistenceRecord{}
	for i := uint(0); i < p.subTables; i++ {
		cache := p.indexes[byte(i)].toMap()
		for uid, id := range cache {
			records = append(records, FactoryPersistenceRecord{ID: id, UID: uid})
		}
	}

	return records
}

// hashUint 将一个uint类型的key哈希为byte
func (p *performanceMap[T]) hashUint(id uint) byte {
	return byte(id % p.subTables)
}

// hashString 将一个string类型的key哈希为byte
func (p *performanceMap[T]) hashString(s string) byte {
	sum := uint(0)
	for _, b := range []byte(s) {
		sum += uint(b)
	}
	return p.hashUint(sum)
}

// generateUID 根据entity的包和类型生成其UID
func (p *performanceMap[T]) generateUID(entity T) (uid string) {
	entityType := reflect.TypeOf(entity)
	packagePath, entityName := entityType.PkgPath(), entityType.Name()
	uid = fmt.Sprintf("%s@%s", packagePath, entityName)
	return
}

// generateID 生成一个自增的ID
func (p *performanceMap[T]) generateID() (id uint) {
	p.mutex.Lock()
	p.idIndex += 1
	id = p.idIndex
	p.mutex.Unlock()
	return id
}

func newPerformanceMap[T any]() *performanceMap[T] {
	entity := &performanceMap[T]{
		slices:    map[byte]*performanceMapSlice[uint, Factory[T]]{},
		indexes:   map[byte]*performanceMapSlice[string, uint]{},
		subTables: 256,
		mutex:     sync.Mutex{},
		idIndex:   0,
	}

	for i := uint(0); i < entity.subTables; i++ {
		entity.slices[byte(i)] = newPerformanceMapSlice[uint, Factory[T]]()
		entity.indexes[byte(i)] = newPerformanceMapSlice[string, uint]()
	}

	return entity
}

// newPerformanceMapWithOpts 使用指定的子表数量新建performanceMap，subTables的取值范围为[1,256]
func newPerformanceMapWithOpts[T any](subTables uint) (success bool, entity *performanceMap[T]) {
	if subTables > 256 || subTables == 0 {
		return false, nil
	}

	entity = &performanceMap[T]{
		slices:    map[byte]*performanceMapSlice[uint, Factory[T]]{},
		indexes:   map[byte]*performanceMapSlice[string, uint]{},
		subTables: subTables,
		mutex:     sync.Mutex{},
		idIndex:   0,
	}

	for i := uint(0); i < entity.subTables; i++ {
		entity.slices[byte(i)] = newPerformanceMapSlice[uint, Factory[T]]()
		entity.indexes[byte(i)] = newPerformanceMapSlice[string, uint]()
	}

	return true, entity
}

// performanceMapSlice performanceMap的存储切片，并发安全，底层为map
type performanceMapSlice[index comparable, cache any] struct {
	mutex   sync.RWMutex
	storage map[index]cache
}

func (p *performanceMapSlice[index, cache]) exist(id index) bool {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	_, has := p.storage[id]
	return has
}

func (p *performanceMapSlice[index, cache]) get(id index) (has bool, result cache) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	result, has = p.storage[id]
	return has, result
}

func (p *performanceMapSlice[index, cache]) add(id index, entity cache) (success bool) {
	if p.exist(id) {
		return false
	} else {
		p.mutex.Lock()
		defer p.mutex.Unlock()
		p.storage[id] = entity
		return true
	}
}

func (p *performanceMapSlice[index, cache]) remove(id index) (success bool) {
	if !p.exist(id) {
		return false
	} else {
		p.mutex.Lock()
		defer p.mutex.Unlock()
		delete(p.storage, id)
		return true
	}
}

func (p *performanceMapSlice[index, cache]) toMap() map[index]cache {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	result := map[index]cache{}
	for id, entity := range p.storage {
		result[id] = entity
	}
	return result
}

func (p *performanceMapSlice[index, cache]) attachToMap(original map[index]cache) map[index]cache {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	for id, entity := range p.storage {
		original[id] = entity
	}
	return original
}

func newPerformanceMapSlice[index comparable, cache any]() *performanceMapSlice[index, cache] {
	return &performanceMapSlice[index, cache]{
		mutex:   sync.RWMutex{},
		storage: map[index]cache{},
	}
}

// Factory 可持久化的实体接口，本质是一个生产entity的工厂
type Factory[entity any] interface {
	Ctor() func() entity
	Enable() bool
	ID() uint
	UID() string

	set(id uint, uid string, ctor func() entity)
	enable()
	disable()
}

func NewFactory[entity any](id uint, uid string) Factory[entity] {
	return &factory[entity]{
		id:     id,
		uid:    uid,
		status: false,
		ctor:   nil,
	}
}

// factory Factory的实现
type factory[entity any] struct {
	ctor   func() entity
	status bool
	id     uint
	uid    string
}

func (p *factory[entity]) Ctor() func() entity {
	return p.ctor
}

func (p *factory[entity]) Enable() bool {
	return p.status
}

func (p *factory[entity]) ID() uint {
	return p.id
}

func (p *factory[entity]) UID() string {
	return p.uid
}

func (p *factory[entity]) set(id uint, uid string, ctor func() entity) {
	p.id, p.uid, p.ctor = id, uid, ctor
}

func (p *factory[entity]) disable() {
	p.status = false
}

func (p *factory[entity]) enable() {
	p.status = true
}

// factoryPersistence 持久化接口的实现
type factoryPersistence[T any] struct {
	impl *performanceMap[T]
	exit chan struct{}
}

func (p *factoryPersistence[T]) Serve(flushFrequency time.Duration, flushPath, flushFile string, errChan chan error) {
	go func() {
		// 执行定时任务
		exitCh := make(chan struct{})
		go func(ch chan struct{}) {
			for {
				select {
				case <-exitCh:
					return
				default:
					// 每隔flushFrequency，将缓存写入文件
					if err := p.Flush(flushPath, flushFile); err != nil {
						errChan <- err
					}
					time.Sleep(flushFrequency)
				}
			}
		}(exitCh)

		// 监听exit信号
		<-p.exit
		exitCh <- struct{}{}
		// 收到退出信号，立即将缓存写入文件
		if err := p.Flush(flushPath, strings.Join([]string{flushFile, "quit"}, ".")); err != nil {
			errChan <- err
		}
	}()
}

func (p *factoryPersistence[T]) Exit() {
	p.exit <- struct{}{}
}

func (p *factoryPersistence[T]) Load(filePath string) (err error) {
	if file, err := os.Open(filePath); err != nil {
		return err
	} else if fileContent, err := io.ReadAll(file); err != nil {
		return err
	} else {
		var persistenceEntities []FactoryPersistenceRecord
		if err = json.Unmarshal(fileContent, &persistenceEntities); err != nil {
			return err
		} else {
			p.impl.Load(persistenceEntities)
			return nil
		}
	}
}

func (p *factoryPersistence[T]) QueryByID(id uint) (has bool, result Factory[T]) {
	return p.impl.QueryByID(id)
}

func (p *factoryPersistence[T]) QueryByUID(uid string) (has bool, result Factory[T]) {
	return p.impl.QueryByUID(uid)
}

func (p *factoryPersistence[T]) Register(ctor func() T) (success bool) {
	return p.impl.Register(ctor)
}

func (p *factoryPersistence[T]) Flush(flushPath string, flushFile string) (err error) {
	persistenceEntities := p.impl.Flush()
	if fileContent, err := json.Marshal(&persistenceEntities); err != nil {
		return err
	} else if file, err := os.OpenFile(path.Join(flushPath, flushFile), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755); err != nil {
		return err
	} else {
		_, err = file.Write(fileContent)
		return err
	}
}

func newFactoryPersistence[T any]() FactoryPersistence[T] {
	return &factoryPersistence[T]{
		impl: newPerformanceMap[T](),
		exit: make(chan struct{}, 1),
	}
}
