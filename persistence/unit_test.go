package persistence

import (
	"testing"
)

type testProductInterface interface {
	data() int
}

type emptyProductInterface = any

type testIntType struct {
	u int
}

func (t testIntType) data() int { return t.u }

type testByteType struct {
	u byte
}

func (t testByteType) data() int { return int(t.u) }

type testRuneType struct {
	u rune
}

func (t testRuneType) data() int { return int(t.u) }

var (
	testIntTypeFactory  = func() testProductInterface { return testIntType{u: 114514} }
	testByteTypeFactory = func() testProductInterface { return testByteType{u: 114} }
	testRuneTypeFactory = func() testProductInterface { return testRuneType{u: 1919810} }
)

func TestPerformanceMap(t *testing.T) {
	tests := []struct {
		name                string
		factories           []func() testProductInterface
		queries             []uint
		wantQueryResult     map[uint]testProductInterface
		wantQuerySuccess    map[uint]bool
		queriesUID          []string
		wantQueryUIDResult  map[string]testProductInterface
		wantQueryUIDSuccess map[string]bool
		flushResult         map[uint]string
		load                map[uint]string
	}{
		{
			name: "TestPerformanceMap",
			factories: []func() testProductInterface{
				testIntTypeFactory,
				testByteTypeFactory,
				testRuneTypeFactory,
			},
			queries: []uint{
				3334, 3335, 3336, 4,
			},
			wantQueryResult: map[uint]testProductInterface{
				3334: testIntTypeFactory(),
				3335: testByteTypeFactory(),
				3336: testRuneTypeFactory(),
				4:    nil,
			},
			wantQuerySuccess: map[uint]bool{
				3334: true,
				3335: true,
				3336: true,
				4:    false,
			},
			queriesUID: []string{
				"github.com/sunist-c/genius-invokation-simulator-backend/persistence@testIntType",
				"github.com/sunist-c/genius-invokation-simulator-backend/persistence@testIntType",
			},
			wantQueryUIDResult: map[string]testProductInterface{
				"github.com/sunist-c/genius-invokation-simulator-backend/persistence@testIntType": testIntTypeFactory(),
			},
			wantQueryUIDSuccess: map[string]bool{
				"github.com/sunist-c/genius-invokation-simulator-backend/persistence@testIntType": true,
			},
			flushResult: map[uint]string{
				3334: "github.com/sunist-c/genius-invokation-simulator-backend/persistence@testIntType",
				3335: "github.com/sunist-c/genius-invokation-simulator-backend/persistence@testByteType",
				3336: "github.com/sunist-c/genius-invokation-simulator-backend/persistence@testRuneType",
				2333: "2333",
				3333: "3333",
			},
			load: map[uint]string{
				2333: "2333",
				3333: "3333",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			performanceMap := newPerformanceMap[testProductInterface]()

			// 测试Performance.Load
			var loadRecords []FactoryPersistenceRecord
			for id, uid := range tt.load {
				loadRecords = append(loadRecords, FactoryPersistenceRecord{ID: id, UID: uid})
			}
			performanceMap.Load(loadRecords)

			// 测试Performance.Register
			for _, factory := range tt.factories {
				if !performanceMap.Register(factory) {
					t.Errorf("register failed")
				}
			}

			// 测试Performance.Flush
			for _, record := range performanceMap.Flush() {
				if record.UID != tt.flushResult[record.ID] {
					t.Errorf("flush failed, want uid %v, but got uid %v", tt.flushResult[record.ID], record.UID)
				}
			}

			// 测试Performance.QueryByID
			for _, query := range tt.queries {
				if success, entity := performanceMap.QueryByID(query); success != tt.wantQuerySuccess[query] {
					t.Errorf("failed to query by id, query %d want success: %v, but got %v", query, tt.wantQuerySuccess[query], success)
				} else if tt.wantQuerySuccess[query] {
					if entity.Ctor() == nil || entity.Ctor()() != tt.wantQueryResult[query] {
						t.Errorf("failed to construct entity, want %v, but got %v", tt.wantQueryResult[query], entity.Ctor()())
					}
				}
			}

			// 测试Performance.QueryByUID
			for _, query := range tt.queriesUID {
				if success, entity := performanceMap.QueryByUID(query); success != tt.wantQueryUIDSuccess[query] {
					t.Errorf("failed to query by id, query %s want success: %v, but got %v", query, tt.wantQueryUIDSuccess[query], success)
				} else if tt.wantQueryUIDSuccess[query] {
					if entity.Ctor() == nil || entity.Ctor()() != tt.wantQueryUIDResult[query] {
						t.Errorf("failed to construct entity, want %v, but got %v", tt.wantQueryUIDResult[query], entity.Ctor()())
					}
				}
			}
		})
	}
}
