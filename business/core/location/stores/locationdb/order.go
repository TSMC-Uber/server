package locationdb

import (
	"fmt"

	"github.com/TSMC-Uber/server/business/core/location"
	"github.com/TSMC-Uber/server/business/data/order"
)

var orderByFields = map[string]string{
	location.OrderByName: "name",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return fmt.Sprintf("%s %s", by, orderBy.Direction), nil
}
