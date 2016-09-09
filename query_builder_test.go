package sqlm

import (
	"testing"
	"fmt"
	"errors"
)

type TestFielder struct {
	Field1 string
	Field2 string
}

func (t *TestFielder) FieldForName(name string) interface{} {
	switch name {
	case "field1":
		return &t.Field1
	case "field2":
		return &t.Field2
	default:
		panic(errors.New("hh"))
	}
}

func TestQueryBuilder(t *testing.T) {
	{
		sql, arguments := Build(Exp(
			"SELECT abc, def FROM what WHERE",
			Not(
				And(
					Exp("user_id >", P(12345)),
					And(
						Exp("media_id <", 12345),
						Exp("time_uuid", "=", 12345),
					),
				),
			),
		))

		fmt.Println(sql)
		fmt.Printf("len: %d, %v\n", len(arguments), arguments)
	}

	{
		mapper := NewFieldsMapper([]string{"field1", "field2"})
		values := []DBFielder {
			&TestFielder{"1","2"},
			&TestFielder{"3","4"},
		}
		sql, arguments := Build(
			Exp(
				"INSERT INTO table2 (", mapper.ColumnString(), ") VALUES ",
				Join(",",
					Exp("(", P(mapper.Fields(values[0])...), ")"),
					Exp("(", P(mapper.Fields(values[0])...), ")"),
					Exp("(", P(mapper.Fields(values[0])...), ")"),
					Exp("(", P(mapper.Fields(values[0])...), ")"),
				),
			),
		)

		fmt.Println(sql)
		fmt.Printf("len: %d, %v\n", len(arguments), arguments)
	}

	{
		i := 30
		sql, arguments := Build(
			Exp("UPDATE table2 SET ",
			    "a =", P("300"), ",",
			    "b =", P("300"), ",",
			    "c =", P("300"), ",",
			    "d =", P(i),
			    "WHERE a = ", P(300),
			),
		)

		fmt.Println(sql)
		fmt.Printf("len: %d, %v\n", len(arguments), arguments)
	}
}