package pgext

import (
	"context"
	"fmt"

	"github.com/AirGateway/pg"
)

// DebugHook is a query hook that logs an error with a query if there are any.
// It can be installed with:
//
//   db.AddQueryHook(pgext.DebugHook{})
type DebugHook struct {
	// Verbose causes hook to print all queries (even those without an error).
	Verbose bool
}

var _ pg.QueryHook = (*DebugHook)(nil)

func (h DebugHook) BeforeQuery(ctx context.Context, evt *pg.QueryEvent) (context.Context, error) {
	q, err := evt.FormattedQuery()
	if err != nil {
		return nil, err
	}

	if evt.Err != nil {
		fmt.Printf("%s executing a query:\n%s\n", evt.Err, q)
	} else if h.Verbose {
		fmt.Println(string(q))
	}

	return ctx, nil
}

func (DebugHook) AfterQuery(context.Context, *pg.QueryEvent) error {
	return nil
}
