package adapters

import (
	"github.com/mixarchitecture/arch/auth/src/adapters/memory"
	"github.com/mixarchitecture/arch/auth/src/adapters/mongo"
	"github.com/mixarchitecture/arch/auth/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)
