package adapters

import (
	"github.com/mixarchitecture/arch/auth/adapters/memory"
	"github.com/mixarchitecture/arch/auth/adapters/mongo"
	"github.com/mixarchitecture/arch/auth/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
	Mongo  = mongo.New()
)
