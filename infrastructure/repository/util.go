package repository

import (
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func InnerJoin(joinTable, baseTable, joinTableColumn, baseTableColumn string) qm.QueryMod {
	return qm.InnerJoin(fmt.Sprintf("%s ON %s.%s = %s.%s",
		joinTable,
		joinTable,
		joinTableColumn,
		baseTable,
		baseTableColumn,
	))
}
