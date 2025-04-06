package util

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
)

// 30 days
var AlarmHours = 30 * 24.0

// param x unit: hour
func TheEbbinghausForgettingCurve(x float64) float64 {
	return 1 - 0.56*math.Pow(x, 0.06)
}

func JsonString(str any) string {
	bytes, _ := json.Marshal(&str)
	return string(bytes)
}

func Info(str string) {
	slog.Info(str)
}

func InfoFormat(format string, args ...any) {
	slog.Info(fmt.Sprintf(format, args...))
}
