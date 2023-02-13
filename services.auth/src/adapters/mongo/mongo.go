package mongo

import (
	mongo_user "github.com/mixarchitecture/arch/auth/src/adapters/mongo/user"
	"github.com/mixarchitecture/arch/auth/src/domain/user"
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
