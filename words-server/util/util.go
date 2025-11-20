package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"math"
	"net/http"
	"os"
	"path/filepath"
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

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Info(str string) {
	slog.Info(str)
}

func InfoFormat(format string, args ...any) {
	slog.Info(fmt.Sprintf(format, args...))
}

func Savewordvoice(sourceText string) string {
	resp, e := http.Get(Config.VoiceSourceUrl + sourceText)
	if e != nil {
		Info(e.Error())
	} else {
		body, e := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode > 299 {
			InfoFormat("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
		}
		if e != nil {
			Info(e.Error())
		}
		localPath := Config.VoiceFolder + sourceText
		e = os.WriteFile(localPath, body, 0777)
		if e != nil {
			Info(e.Error())
		} else {
			return localPath
		}
	}
	return ""
}

func AllLocalVoicePath() map[string]string {
	allLocalWord := make(map[string]string)
	dirEnrtys, err := os.ReadDir(Config.VoiceFolder)
	if err == nil {
		for _, dirEntry := range dirEnrtys {
			allLocalWord[dirEntry.Name()] = Config.VoiceFolder + dirEntry.Name()
		}
	}
	return allLocalWord
}

func LocalPath(path string) string {
	return filepath.Base(Config.VoiceFolder + path)
}
