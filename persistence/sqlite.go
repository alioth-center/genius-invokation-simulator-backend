package persistence

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	sqlite3DB *xorm.Engine = nil
)

type sqliteTable[PK Increasable, T any] struct {
	session *xorm.Session
	errChan chan error
}

func (s *sqliteTable[PK, T]) QueryByID(id PK) (has bool, result T) {
	if _, err := s.session.ID(id).Get(&result); err != nil {
		s.errChan <- err
		return false, result
	} else {
		return true, result
	}
}

func (s *sqliteTable[PK, T]) UpdateByID(id PK, newEntity T) (success bool) {
	var condition T
	if _, err := s.session.ID(id).Get(&condition); err != nil {
		s.errChan <- err
		return false
	} else if _, err = s.session.ID(id).Update(newEntity); err != nil {
		s.errChan <- err
		return false
	} else {
		return true
	}
}

func (s *sqliteTable[PK, T]) InsertOne(entity *T) (success bool, result *T) {
	if affected, err := s.session.InsertOne(entity); err != nil || affected != 1 {
		s.errChan <- err
		return false, nil
	} else {
		return true, entity
	}
}

func (s *sqliteTable[PK, T]) DeleteOne(id PK) (success bool) {
	var condition T
	if _, err := s.session.ID(id).Get(&condition); err != nil {
		s.errChan <- err
		return false
	} else if _, err = s.session.ID(id).Delete(nil); err != nil {
		s.errChan <- err
		return false
	} else {
		return true
	}
}

func (s *sqliteTable[PK, T]) FindOne(condition T) (has bool, result T) {
	var results []T
	if err := s.session.Find(&results, condition); err != nil {
		s.errChan <- err
		return false, result
	} else if len(results) != 1 {
		return false, result
	} else {
		return true, results[0]
	}
}

func (s *sqliteTable[PK, T]) Find(condition T) (results []T) {
	if err := s.session.Find(&results, condition); err != nil {
		s.errChan <- err
		return []T{}
	} else {
		return results
	}
}

func newDatabasePersistence[PK Increasable, T any](errChan chan error) (success bool, persistence DatabasePersistence[PK, T]) {
	var entity T
	if err := sqlite3DB.Sync2(entity); err != nil {
		errChan <- err
		return false, persistence
	} else {
		table := &sqliteTable[PK, T]{
			session: sqlite3DB.Table(entity),
			errChan: errChan,
		}
		return true, table
	}
}
