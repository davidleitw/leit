package internal

import (
	"fmt"
	"strings"
	"time"
)

var YEAR int = time.Now().Year()

// like "12-21"
func getDateStrByMonth_and_Day(d string) string {
	md := strings.Split(d, "-")

	if len(md) != 2 {
		return ""
	}

	date := fmt.Sprintf("%d-%s-%s", YEAR, md[0], md[1])
	if !VaildDateStr(date) {
		return ""
	}

	return date
}

func VaildDateStr(d string) bool {
	_, err := time.Parse("2006-01-02", d)
	return err == nil
}
