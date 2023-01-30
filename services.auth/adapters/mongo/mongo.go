package mongo

import (
	mongo_user "github.com/mixarchitecture/arch/auth/adapters/mongo/user"
	"github.com/mixarchitecture/arch/auth/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewUser(userFactory user.Factory, collection *mongo.Collection) user.Repository
}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}

func (m *mongodb) NewUser(userFactory user.Factory, collection *mongo.Collection) user.Repository {
	return mongo_user.New(userFactory, collection)
}
