package persistence

import (
	"time"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	sqlite3DB *xorm.Engine = nil
)

func Sqlite3DBDelete[T any](entity T) (success bool) {
	if sqlite3DB != nil {
		_, err := sqlite3DB.Delete(entity)
		return err == nil
	} else {
		return false
	}
}

func Sqlite3DBUpdate[T any](entity T, id uint) (success bool) {
	if sqlite3DB != nil {
		_, err := sqlite3DB.Table(entity).ID(id).Update(entity)
		return err == nil
	} else {
		return false
	}
}

type Sqlite3Table[T any] struct {
	session *xorm.Session
	errChan chan error
}

func (s *Sqlite3Table[T]) Serve(flushFrequency time.Duration, flushPath, flushFile string, errChan chan error) {
}

func (s *Sqlite3Table[T]) Exit() {}

func (s *Sqlite3Table[T]) Load(filePath string) (err error) { return nil }

func (s *Sqlite3Table[T]) QueryByID(id uint) (has bool, result Persistent[T]) {
	var entity T
	if ok, err := s.session.ID(id).Get(&entity); ok && err == nil {
		return true, &persistent[T]{
			ctor:   func() T { return entity },
			status: true,
			id:     id,
			uid:    "",
		}
	} else {
		if err != nil {
			s.errChan <- err
		}
		return false, result
	}
}

func (s *Sqlite3Table[T]) QueryByUID(uid string) (has bool, result Persistent[T]) {
	return false, result
}

func (s *Sqlite3Table[T]) Register(ctor func() T) (success bool) {
	if _, err := s.session.InsertOne(ctor()); err == nil {
		return true
	} else {
		s.errChan <- err
		return false
	}
}

func (s *Sqlite3Table[T]) Flush(flushPath string, flushFile string) (err error) { return nil }

func NewSqlite3Table[T any]() Persistence[T] {
	var entity T
	return &Sqlite3Table[T]{
		session: sqlite3DB.Table(entity),
	}
}
