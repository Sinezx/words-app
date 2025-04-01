package util

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
)

// 30 days
var AlarmHours = 30 * 24.0

var sessionMap map[string]context.CancelFunc

func initSessionMap() {
	if sessionMap == nil {
		sessionMap = make(map[string]context.CancelFunc)
	}
}

func SetSessionFunc(id string, cf context.CancelFunc) {
	initSessionMap()
	sessionMap[id] = cf
}

func GetSessionFunc(id string) context.CancelFunc {
	initSessionMap()
	return sessionMap[id]
}

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

func Error(str string) {
	slog.Error(str)
}
