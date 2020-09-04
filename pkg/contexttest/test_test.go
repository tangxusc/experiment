package contexttest

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := context.TODO()
	f(ctx)
	fmt.Println(ctx.Value("foo"))
}

func f(ctx context.Context) {
	context.WithValue(ctx, "foo", -6)
}
