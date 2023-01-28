package contextWithValue

import (
	"context"
	"log"
)

const (
	key   = "CtxValueKey"
	value = "CtxValue"
)

func ContextWithValueDemonstration() {
	// context.Background() initializes an empty context.
	ctx := context.Background()
	procedureA(ctx)
}

func procedureA(ctx context.Context) {
	// assign a value to the context and sent it to procedureB
	ctx = context.WithValue(ctx, key, value)
	procedureB(ctx)
}

func procedureB(ctx context.Context) {
	ctxValue := ctx.Value(key)
	log.Printf("Got %v from Context", ctxValue)
}
