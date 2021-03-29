// Package resolver impls all query/mutation resolver
package resolver

import (
	"github.com/Akshit8/reddit-clone-api/pkg/post"
	"github.com/Akshit8/reddit-clone-api/pkg/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver contains entities service objects
type Resolver struct{
	PostService post.Service
	UserService user.Service
}
