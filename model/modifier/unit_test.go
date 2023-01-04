package modifier

import (
	"testing"
)

type testModifier struct {
	info      uint
	innerData *int
	handler   func(ctx *Context[int])
	effective bool
}

func (t testModifier) ID() uint { return t.info }

func (t *testModifier) Handler() func(ctx *Context[int]) {
	return func(ctx *Context[int]) {
		*t.innerData = 114514
		t.handler(ctx)
	}
}

func (t testModifier) Clone() Modifier[int] {
	return &testModifier{
		info:      t.info,
		innerData: new(int),
		handler:   t.handler,
	}
}

func (t testModifier) RoundReset() {}

func (t testModifier) Effective() bool { return t.effective }

func (t testModifier) EffectLeft() uint { return 0 }

func newModifier(id uint, handler func(ctx *Context[int])) Modifier[int] {
	return &testModifier{info: id, handler: handler, innerData: new(int), effective: true}
}

func newModifierWithInnerData(id uint, handler func(ctx *Context[int]), data *int) Modifier[int] {
	return &testModifier{info: id, handler: handler, innerData: data, effective: true}
}

func newModifierWithEffective(id uint, handler func(ctx *Context[int]), effective bool) Modifier[int] {
	return &testModifier{info: id, handler: handler, innerData: new(int), effective: effective}
}

func TestContextExecute(t *testing.T) {
	add := func(ctx *Context[int]) {
		*ctx.data += 1
	}
	subAndAbort := func(ctx *Context[int]) {
		*ctx.data -= 1
		ctx.Abort()
	}
	backSub := func(ctx *Context[int]) {
		ctx.Next()
		*ctx.data -= 1
	}
	abortedJudge := func(ctx *Context[int]) {
		ctx.Next()
		if ctx.IsAborted() {
			*ctx.data = 114514
		}
	}

	testCases := []struct {
		name     string
		handlers []func(ctx *Context[int])
		want     int
	}{
		{
			name: "TestContext-Append",
			handlers: []func(ctx *Context[int]){
				add, add, add,
			},
			want: 3,
		},
		{
			name: "TestContext-Abort",
			handlers: []func(ctx *Context[int]){
				add, subAndAbort, add, add, add, add,
			},
			want: 0,
		},
		{
			name: "TestContext-BackCall",
			handlers: []func(ctx *Context[int]){
				add, backSub, subAndAbort, add,
			},
			want: -1,
		},
		{
			name: "TestContext-Aborted",
			handlers: []func(ctx *Context[int]){
				add, abortedJudge, add, subAndAbort,
			},
			want: 114514,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			handlers := NewChain[int]()
			for i, f := range tt.handlers {
				handlers.Append(newModifier(uint(i), f))
			}
			data := 0
			handlers.Execute(&data)

			if data != tt.want {
				t.Errorf("incorrect data: want %v, got %v", tt.want, data)
			}
		})
	}
}

func TestContextRemove(t *testing.T) {
	add := func(ctx *Context[int]) {
		*ctx.data += 1
	}
	addTwice := func(ctx *Context[int]) {
		*ctx.data += 1
		ctx.Next()
		*ctx.data += 1
	}
	addBack := func(ctx *Context[int]) {
		ctx.Next()
		*ctx.data += 1
	}

	testCases := []struct {
		name           string
		addHandlers    []func(ctx *Context[int])
		removeHandlers []uint
		want           int
	}{
		{
			name:           "TestContextRemove-1",
			addHandlers:    []func(ctx *Context[int]){add, addTwice, addBack},
			removeHandlers: []uint{1},
			want:           2,
		},
		{
			name:           "TestContextRemove-2",
			addHandlers:    []func(ctx *Context[int]){add, addTwice, addTwice, addBack},
			removeHandlers: []uint{1, 1, 1},
			want:           4,
		},
		{
			name:           "TestContextRemove-3",
			addHandlers:    []func(ctx *Context[int]){},
			removeHandlers: []uint{1, 1, 4},
			want:           0,
		},
		{
			name:           "TestContextRemove-4",
			addHandlers:    []func(ctx *Context[int]){add, add, add, add, add},
			removeHandlers: []uint{4, 2, 1, 3},
			want:           1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			handlers := NewChain[int]()
			for i, f := range tt.addHandlers {
				handlers.Append(newModifier(uint(i), f))
			}
			for _, id := range tt.removeHandlers {
				handlers.Remove(id)
			}
			data := 0
			handlers.Execute(&data)

			if data != tt.want {
				t.Errorf("incorrect data: want %v, got %v", tt.want, data)
			}
		})
	}
}

func TestContextClone(t *testing.T) {
	add := func(ctx *Context[int]) {
		*ctx.data += 1
	}
	addTwice := func(ctx *Context[int]) {
		*ctx.data += 1
		ctx.Next()
		*ctx.data += 1
	}

	testCases := []struct {
		name       string
		wantResult int
		wantInner  int
	}{
		{
			name:       "TestContextClone",
			wantResult: 3,
			wantInner:  0,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			handlers := NewChain[int]()
			innerData := 0
			handlers.Append(newModifierWithInnerData(0, add, &innerData))
			handlers.Append(newModifier(1, add))
			handlers.Append(newModifier(1, addTwice))

			data := 0
			handlers.Preview(&data)

			if data != tt.wantResult {
				t.Errorf("incorrect result: want %v, got %v", tt.wantResult, data)
			}

			if innerData != tt.wantInner {
				t.Errorf("incorrect data: want %v, got %v", tt.wantResult, data)
			}
		})
	}
}

func TestContextClear(t *testing.T) {
	t.Run("TestContextClear", func(t *testing.T) {
		add := func(ctx *Context[int]) {
			*ctx.data += 1
		}
		addTwice := func(ctx *Context[int]) {
			*ctx.data += 1
			ctx.Next()
			*ctx.data += 1
		}
		handlers := NewChain[int]()
		handlers.Append(newModifierWithEffective(0, add, false))
		handlers.Append(newModifierWithEffective(1, addTwice, false))
		data := 0
		handlers.Execute(&data)

		if handlers.size != 0 {
			t.Errorf("incorrect clear result")
		}
	})
}
