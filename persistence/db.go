package persistence

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"sync"
)

const (
	ruleSetPersistenceFileName   = "rule-set-persistence.json"
	cardPersistenceFileName      = "card-persistence.json"
	characterPersistenceFileName = "character-persistence.json"
	playerPersistenceFileName    = "player-persistence.json"
	skillPersistenceFileName     = "skill-persistence.json"
)

var (
	storagePath = ""
	loaded      = false
)

var (
// RuleSetPersistence   Persistence[RuleSet]
// CardPersistence      Persistence[Card]
// CharacterPersistence Persistence[CharacterInfo]
// PlayerPersistence    Persistence[PlayerInfo]
// SkillPersistence     Persistence[Skill]
)

func init() {
	if execPath, err := os.Executable(); err != nil {
		panic(err)
	} else if err = SetStoragePath(path.Join(execPath, "../data/persistence")); err != nil {
		//panic(err)
	} else {
		fmt.Println(path.Join(storagePath, ruleSetPersistenceFileName))
	}
}

func SetStoragePath(path string) error {
	if s, err := os.Stat(path); err == nil && s.IsDir() {
		storagePath = path
		return nil
	} else {
		return fmt.Errorf("path %s is not a directory or is not exist", path)
	}
}

// Load 从持久化文件读取信息，写入持久化模块
func Load() {
	if !loaded {
		//RuleSetPersistence.Load(path.Join(storagePath, ruleSetPersistenceFileName))
		//CardPersistence.Load(path.Join(storagePath, cardPersistenceFileName))
		//CharacterPersistence.Load(path.Join(storagePath, characterPersistenceFileName))
		//PlayerPersistence.Load(path.Join(storagePath, playerPersistenceFileName))
		//SkillPersistence.Load(path.Join(storagePath, skillPersistenceFileName))
		loaded = true
	}
}

// Quit 持久化模块退出前将缓存写入文件
func Quit() {
	//RuleSetPersistence.Flush(storagePath, ruleSetPersistenceFileName)
	//CardPersistence.Flush(storagePath, cardPersistenceFileName)
	//CharacterPersistence.Flush(storagePath, characterPersistenceFileName)
	//PlayerPersistence.Flush(storagePath, playerPersistenceFileName)
	//SkillPersistence.Flush(storagePath, skillPersistenceFileName)
}

// Persistence 持久化接口，抽象工厂集合的持久化封装
type Persistence[T any] interface {
	Load(filePath string)
	QueryByID(id uint) (has bool, result Persistent[T])
	QueryByUID(uid string) (has bool, result Persistent[T])
	Register(ctor func() T) (success bool)
	Flush(flushPath string, flushFile string)
}

// PerformanceMap 高性能的并发安全KV存储
type PerformanceMap[T any] struct {
	slices  map[byte]*performanceMapSlice[uint, Persistent[T]] // slices 存储Persistent的分表
	indexes map[byte]*performanceMapSlice[string, uint]        // indexes 存储索引的分表
	mutex   sync.Mutex                                         // mutex 控制idIndex的锁
	idIndex uint                                               // idIndex 当前的自动生成ID
}

// Load 将records加载到PerformanceMap中，耗时操作，写锁
func (p *PerformanceMap[T]) Load(records []PerformanceMapRecord) {
	for _, record := range records {
		if record.ID != 0 && record.UID != "" {
			// 更新索引和Persistent分表
			entity := NewPersistent[T](record.ID, record.UID)
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
func (p *PerformanceMap[T]) QueryByID(id uint) (has bool, result Persistent[T]) {
	slice := p.hashUint(id)
	if ok, entity := p.slices[slice].get(id); ok {
		if entity.Enable() {
			return true, entity
		} else {
			return false, NewPersistent[T](0, "")
		}
	} else {
		return false, NewPersistent[T](0, "")
	}
}

// QueryByUID 根据UID获取工厂，只返回可用结果
func (p *PerformanceMap[T]) QueryByUID(uid string) (has bool, result Persistent[T]) {
	index := p.hashString(uid)
	if ok, id := p.indexes[index].get(uid); ok {
		return p.QueryByID(id)
	} else {
		return false, NewPersistent[T](0, "")
	}
}

// Register 将一个工厂注册到PerformanceMap中
func (p *PerformanceMap[T]) Register(ctor func() T) (success bool) {
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
			record = NewPersistent[T](id, uid)
			record.set(id, uid, ctor)
			record.enable()
			return p.slices[slice].add(id, record)
		}
	} else {
		// 如果没有找到UID记录，说明是新记录，直接写入
		id := p.generateID()
		slice := p.hashUint(id)
		p.indexes[index].add(uid, id)
		record := NewPersistent[T](id, uid)
		record.set(id, uid, ctor)
		record.enable()
		return p.slices[slice].add(id, record)
	}
}

// Flush 将PerformanceMap中的数据全部取出，耗时操作，读锁
func (p *PerformanceMap[T]) Flush() (records []PerformanceMapRecord) {
	records = []PerformanceMapRecord{}
	for i := 0; i < 256; i++ {
		cache := p.indexes[byte(i)].toMap()
		for uid, id := range cache {
			records = append(records, PerformanceMapRecord{ID: id, UID: uid})
		}
	}

	return records
}

// hashUint 将一个uint类型的key哈希为byte
func (p *PerformanceMap[T]) hashUint(id uint) byte {
	return byte(id % 256)
}

// hashString 将一个string类型的key哈希为byte
func (p *PerformanceMap[T]) hashString(s string) byte {
	sum := uint(0)
	for _, b := range []byte(s) {
		sum += uint(b)
	}
	return p.hashUint(sum)
}

// generateUID 根据entity的包和类型生成其UID
func (p *PerformanceMap[T]) generateUID(entity T) (uid string) {
	entityType := reflect.TypeOf(entity)
	packagePath, entityName := entityType.PkgPath(), entityType.Name()
	uid = fmt.Sprintf("%s@%s", packagePath, entityName)
	return
}

// generateID 生成一个自增的ID
func (p *PerformanceMap[T]) generateID() (id uint) {
	p.mutex.Lock()
	p.idIndex += 1
	id = p.idIndex
	p.mutex.Unlock()
	return id
}

func NewPerformanceMap[T any]() *PerformanceMap[T] {
	entity := &PerformanceMap[T]{
		slices:  map[byte]*performanceMapSlice[uint, Persistent[T]]{},
		indexes: map[byte]*performanceMapSlice[string, uint]{},
		mutex:   sync.Mutex{},
		idIndex: 0,
	}

	for i := 0; i <= 255; i++ {
		entity.slices[byte(i)] = newPerformanceMapSlice[uint, Persistent[T]]()
		entity.indexes[byte(i)] = newPerformanceMapSlice[string, uint]()
	}

	return entity
}

// performanceMapSlice PerformanceMap的存储切片，并发安全，底层为map
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

// PerformanceMapRecord PerformanceMap的持久化结构
type PerformanceMapRecord struct {
	ID  uint
	UID string
}
