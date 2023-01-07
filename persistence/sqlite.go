package persistence

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var (
	sqlite3DB *xorm.Engine
)

type Sqlite3Table[T any] struct {
}

func (s *Sqlite3Table[T]) Serve(flushFrequency time.Duration, flushPath, flushFile string, errChan chan error) {

}

func (s *Sqlite3Table[T]) Exit() {

}

func (s *Sqlite3Table[T]) Load(filePath string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Sqlite3Table[T]) QueryByID(id uint) (has bool, result Persistent[T]) {
	//TODO implement me
	panic("implement me")
}

func (s *Sqlite3Table[T]) QueryByUID(uid string) (has bool, result Persistent[T]) {
	//TODO implement me
	panic("implement me")
}

func (s *Sqlite3Table[T]) Register(ctor func() T) (success bool) {
	//TODO implement me
	panic("implement me")
}

func (s *Sqlite3Table[T]) Flush(flushPath string, flushFile string) (err error) {
	//TODO implement me
	panic("implement me")
}

func NewSqlite3Table[T any]() Persistence[T] {
	return &Sqlite3Table[T]{}
}
