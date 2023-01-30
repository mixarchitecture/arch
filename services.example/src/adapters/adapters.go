package adapters

import (
	"github.com/mixarchitecture/arch/example/src/adapters/memory"
	"github.com/mixarchitecture/arch/example/src/adapters/mysql"
)

var (
	MySQL  = mysql.New()
	Memory = memory.New()
)
