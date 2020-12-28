package modules
import (
"time"
"strconv"
)

func GetTodayDate() string {
	var year, month, day = time.Now().Date()

	var todaysDate = month.String() + " " +  strconv.Itoa(day) + ", " + strconv.Itoa(year)
    return todaysDate	
}
