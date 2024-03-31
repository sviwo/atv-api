package statistics

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	utils "sviwo/pkg/utility/date"
	"testing"
	"time"
)

func TestGetDays(t *testing.T) {
	year := time.Now().Year()
	month := time.Now().Month()
	t.Log(utils.CalcDaysFromYearMonth(year, int(month)))
}
