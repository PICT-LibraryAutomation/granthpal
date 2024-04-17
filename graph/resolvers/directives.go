package resolvers

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/PICT-LibraryAutomation/granthpal/graph"
)

func IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	// state := ctx.Value(utils.AuthKey).(auth.AuthState)
	// if !state.IsAuth {
	// 	return nil, fmt.Errorf("Unauthorized")
	// }

	return next(ctx)
}

func IsKind(ctx context.Context, obj interface{}, next graphql.Resolver, kind graph.UserKind) (res interface{}, err error) {
	// db := ctx.Value(utils.RemoteKey).(*gorm.DB)
	// state := ctx.Value(utils.AuthKey).(auth.AuthState)

	// if !state.IsAuth {
	// 	return nil, fmt.Errorf("Unauthorized")
	// }

	// var user models.User
	// if err := db.First(&user, "prn = ?", state.Session.UserID).Error; err != nil {
	// 	return nil, fmt.Errorf("user not found")
	// }

	// if user.Kind != kind {
	// 	return nil, fmt.Errorf("Unauthorized")
	// }

	return next(ctx)
}
