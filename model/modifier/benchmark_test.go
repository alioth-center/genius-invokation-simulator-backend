package modifier

import "testing"

// BenchmarkTestContextRemove
func BenchmarkTestContextRemove(b *testing.B) {
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
		handlers.Remove(128)
	}
}

// BenchmarkTestContextAppend
func BenchmarkTestContextAppend(b *testing.B) {
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
		handlers.Append(newModifier(129, add))
	}
}

// BenchmarkTestContextExecute
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
	for j := uint(0); j < 1; j++ {
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
		handlers.Execute(&data)
	}
}
