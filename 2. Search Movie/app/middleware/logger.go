package middleware

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// LogJSON godoc.
type LogJSON struct {
	// TimeStamp shows the time after the server returns a response.
	TimeStamp string `json:"timestamp"`
	// StatusCode is HTTP response code.
	StatusCode int `json:"statuscode"`
	// Latency is how much time the server cost to process a certain request.
	Latency float64 `json:"latency"`
	// ClientIP equals Context's ClientIP method.
	ClientIP string `json:"clientip"`
	// Method is the HTTP method given to the request.
	Method string `json:"method"`
	// Path is a path the client requests.
	Path string `json:"path"`
	// BodySize is the size of the Response Body
	BodySize int      `json:"bodysize"`
	Error    []string `json:"error"`
	// Request body
	Request map[string]interface{} `json:"request"`
	// Response body
	Response interface{} `json:"response"`
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write godoc.
func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

// Logger godoc.
func Logger(ctx *gin.Context) {
	// Start Timer
	start := time.Now()
	path := ctx.Request.URL.Path
	raw := ctx.Request.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}

	var logJSON = &LogJSON{
		ClientIP: ctx.ClientIP(),
		Method:   ctx.Request.Method,
	}

	// Get request
	var body string
	if ctx.Request.Method != http.MethodGet && !strings.Contains(ctx.GetHeader("Content-Type"), "form-data") {
		reqBody, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			body = ""
		} else {
			body = string(reqBody)
			ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		}
	}

	logJSON.Request = parseBody([]byte(body))

	// Get response
	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
	ctx.Writer = w

	ctx.Next()

	bytesJSON := w.body.Bytes()
	if bytesJSON != nil {
		_ = json.Unmarshal(bytesJSON, &logJSON.Response)
	}

	logJSON.StatusCode = ctx.Writer.Status()
	logJSON.BodySize = ctx.Writer.Size()
	logJSON.Path = path

	logJSON.Error = ctx.Errors.Errors()
	logJSON.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	logJSON.Latency = GetDurationInMilliseconds(start)

	printLog(logJSON)
}

func printLog(logJSON *LogJSON) {
	logByte, _ := json.MarshalIndent(logJSON, "", " ")
	logString := string(logByte)
	log.SetFlags(0)
	log.Print(logString)
}

func parseBody(body []byte) (res map[string]interface{}) {
	_ = json.Unmarshal(body, &res)
	return
}

// GetDurationInMilliseconds takes a start time and returns a duration in milliseconds
func GetDurationInMilliseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}
