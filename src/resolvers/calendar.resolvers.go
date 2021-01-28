package resolvers

import (
	"context"
	"fmt"

	"github.com/beforesecond/fiber-gql/src/graph/model"
)

func (r *mutationResolver) AllCalendars(ctx context.Context, page *int, perPage *int, sortField *string, sortOrder *string) ([]*model.Calendar, error) {
	panic(fmt.Errorf("not implemented"))
}
