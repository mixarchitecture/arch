package user

import (
	"github.com/mixarchitecture/arch/auth/src/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	userFactory user.Factory
	collection  *mongo.Collection
}

func New(userFactory user.Factory, collection *mongo.Collection) user.Repository {
	validate(userFactory, collection)
	return &repo{
		userFactory: userFactory,
	}
}

func validate(userFactory user.Factory, collection *mongo.Collection) {
	if userFactory.IsZero() {
		panic("exampleFactory is zero")
	}
	if collection == nil {
		panic("collection is nil")
	}
}
