package modifier

import (
	"testing"
)

type testModifier struct {
	info    Info
	handler func(ctx *Context[int])
}

func (t testModifier) Info() Info { return t.info }

func (t testModifier) Handler() func(ctx *Context[int]) { return t.handler }

func (t testModifier) Clone() Modifier[int] { return nil }

func (t testModifier) RoundReset() {}

func (t testModifier) Effected() bool { return true }

func (t testModifier) EffectLeft() uint { return 0 }

func newModifier(id uint, handler func(ctx *Context[int])) Modifier[int] {
	return &testModifier{info: Info{ID: id}, handler: handler}
}

func TestContext(t *testing.T) {
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
			ctx := NewContext[int](&data, handlers)
			ctx.Next()

			if *ctx.data != tt.want {
				t.Errorf("incorrect data: want %v, got %v", tt.want, *ctx.data)
			}
		})
	}
}

// BenchmarkTestContextAdd
// 3handler, 110ns/op; 12handler, 440ns/op; 24handler, 1100ns/op; 128handler, 17000ns/op
func BenchmarkTestContextAdd(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		handlers := NewChain[int]()
		for j := uint(0); j <= 24; j++ {
			switch j % 3 {
			case 0:
				handlers.Append(newModifier(j, add))
			case 1:
				handlers.Append(newModifier(j, addTwice))
			case 2:
				handlers.Append(newModifier(j, addBack))
			}
		}
	}
}

// BenchmarkTestContextExecute
// 3handler, 70ns/op; 12handler, 120ns/op; 24handler, 180ns/op; 128handler, 1400ns/op
func BenchmarkTestContextExecute(b *testing.B) {
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
	handlers := NewChain[int]()
	for j := uint(0); j <= 128; j++ {
		switch j % 3 {
		case 0:
			handlers.Append(newModifier(j, add))
		case 1:
			handlers.Append(newModifier(j, addTwice))
		case 2:
			handlers.Append(newModifier(j, addBack))
		}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data := 0
		ctx := NewContext(&data, handlers)
		ctx.Next()
	}
}
