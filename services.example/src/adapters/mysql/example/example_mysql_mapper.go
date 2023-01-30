package mysql_example

import (
	"github.com/mixarchitecture/arch/shared/i18n"

	"github.com/mixarchitecture/arch/example/src/adapters/mysql/example/entity"
	"github.com/mixarchitecture/arch/example/src/domain/example"
)

type MySqlExampleMapper interface {
	ToExample(entity *entity.MySQLExample) (*example.Example, *i18n.I18nError)
	ToExamples(entities []*entity.MySQLExample) ([]*example.Example, *i18n.I18nError)
}

type mySqlExampleMapper struct {
	exampleFactory example.Factory
}

func NewMySqlExampleMapper(factory example.Factory) MySqlExampleMapper {
	return &mySqlExampleMapper{
		exampleFactory: factory,
	}
}

func (m *mySqlExampleMapper) ToExample(entity *entity.MySQLExample) (*example.Example, *i18n.I18nError) {
	return m.exampleFactory.Unmarshal(entity.UUID, entity.Field, entity.Content)
}

func (m *mySqlExampleMapper) ToExamples(entities []*entity.MySQLExample) ([]*example.Example, *i18n.I18nError) {
	examples := make([]*example.Example, len(entities))
	for i, entity := range entities {
		ex, err := m.exampleFactory.Unmarshal(entity.UUID, entity.Field, entity.Content)
		if err != nil {
			return nil, err
		}
		examples[i] = ex
	}
	return examples, nil
}
