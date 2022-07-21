/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package logstats

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"vitess.io/vitess/go/hack"

	"vitess.io/vitess/go/sqltypes"
	"vitess.io/vitess/go/streamlog"
	"vitess.io/vitess/go/vt/callinfo"
	"vitess.io/vitess/go/vt/callinfo/fakecallinfo"
	querypb "vitess.io/vitess/go/vt/proto/query"
)

func TestMain(m *testing.M) {
	hack.DisableProtoBufRandomness()
	os.Exit(m.Run())
}

func testFormat(t *testing.T, stats *LogStats, params url.Values) string {
	var b bytes.Buffer
	err := stats.Logf(&b, params)
	require.NoError(t, err)
	return b.String()
}

func TestLogStatsFormat(t *testing.T) {
	defer func() {
		*streamlog.RedactDebugUIQueries = false
		*streamlog.QueryLogFormat = "text"
	}()
	logStats := NewLogStats(context.Background(), "test", "sql1", "suuid", nil)
	logStats.StartTime = time.Date(2017, time.January, 1, 1, 2, 3, 0, time.UTC)
	logStats.EndTime = time.Date(2017, time.January, 1, 1, 2, 4, 1234, time.UTC)
	logStats.Keyspace = "ks"
	logStats.Table = "table"
	logStats.TablesUsed = []string{"ks1.tbl1", "ks2.tbl2"}
	logStats.TabletType = "PRIMARY"
	params := map[string][]string{"full": {}}
	intBindVar := map[string]*querypb.BindVariable{"intVal": sqltypes.Int64BindVariable(1)}
	stringBindVar := map[string]*querypb.BindVariable{"strVal": sqltypes.StringBindVariable("abc")}

	tests := []struct {
		name     string
		redact   bool
		format   string
		expected string
		bindVars map[string]*querypb.BindVariable
	}{
		{ //0
			redact:   false,
			format:   "text",
			expected: "test\t\t\t''\t''\t2017-01-01 01:02:03.000000\t2017-01-01 01:02:04.000001\t1.000001\t0.000000\t0.000000\t0.000000\t\t\"sql1\"\tmap[intVal:type:INT64 value:\"1\"]\t0\t0\t\"\"\t\"ks\"\t\"table\"\t\"PRIMARY\"\t\"suuid\"\tfalse\t[\"ks1.tbl1\",\"ks2.tbl2\"]\n",
			bindVars: intBindVar,
		}, { //1
			redact:   true,
			format:   "text",
			expected: "test\t\t\t''\t''\t2017-01-01 01:02:03.000000\t2017-01-01 01:02:04.000001\t1.000001\t0.000000\t0.000000\t0.000000\t\t\"sql1\"\t\"[REDACTED]\"\t0\t0\t\"\"\t\"ks\"\t\"table\"\t\"PRIMARY\"\t\"suuid\"\tfalse\t[\"ks1.tbl1\",\"ks2.tbl2\"]\n",
			bindVars: intBindVar,
		}, { //2
			redact:   false,
			format:   "json",
			expected: "{\n    \"BindVars\": {\n        \"intVal\": {\n            \"type\": \"INT64\",\n            \"value\": 1\n        }\n    },\n    \"Cached Plan\": false,\n    \"CommitTime\": 0,\n    \"Effective Caller\": \"\",\n    \"End\": \"2017-01-01 01:02:04.000001\",\n    \"Error\": \"\",\n    \"ExecuteTime\": 0,\n    \"ImmediateCaller\": \"\",\n    \"Keyspace\": \"ks\",\n    \"Method\": \"test\",\n    \"PlanTime\": 0,\n    \"RemoteAddr\": \"\",\n    \"RowsAffected\": 0,\n    \"SQL\": \"sql1\",\n    \"SessionUUID\": \"suuid\",\n    \"ShardQueries\": 0,\n    \"Start\": \"2017-01-01 01:02:03.000000\",\n    \"StmtType\": \"\",\n    \"Table\": \"table\",\n    \"TablesUsed\": [\n        \"ks1.tbl1\",\n        \"ks2.tbl2\"\n    ],\n    \"TabletType\": \"PRIMARY\",\n    \"TotalTime\": 1.000001,\n    \"Username\": \"\"\n}",
			bindVars: intBindVar,
		}, { //3
			redact:   true,
			format:   "json",
			expected: "{\n    \"BindVars\": \"[REDACTED]\",\n    \"Cached Plan\": false,\n    \"CommitTime\": 0,\n    \"Effective Caller\": \"\",\n    \"End\": \"2017-01-01 01:02:04.000001\",\n    \"Error\": \"\",\n    \"ExecuteTime\": 0,\n    \"ImmediateCaller\": \"\",\n    \"Keyspace\": \"ks\",\n    \"Method\": \"test\",\n    \"PlanTime\": 0,\n    \"RemoteAddr\": \"\",\n    \"RowsAffected\": 0,\n    \"SQL\": \"sql1\",\n    \"SessionUUID\": \"suuid\",\n    \"ShardQueries\": 0,\n    \"Start\": \"2017-01-01 01:02:03.000000\",\n    \"StmtType\": \"\",\n    \"Table\": \"table\",\n    \"TablesUsed\": [\n        \"ks1.tbl1\",\n        \"ks2.tbl2\"\n    ],\n    \"TabletType\": \"PRIMARY\",\n    \"TotalTime\": 1.000001,\n    \"Username\": \"\"\n}",
			bindVars: intBindVar,
		}, { //4
			redact:   false,
			format:   "text",
			expected: "test\t\t\t''\t''\t2017-01-01 01:02:03.000000\t2017-01-01 01:02:04.000001\t1.000001\t0.000000\t0.000000\t0.000000\t\t\"sql1\"\tmap[strVal:type:VARCHAR value:\"abc\"]\t0\t0\t\"\"\t\"ks\"\t\"table\"\t\"PRIMARY\"\t\"suuid\"\tfalse\t[\"ks1.tbl1\",\"ks2.tbl2\"]\n",
			bindVars: stringBindVar,
		}, { //5
			redact:   true,
			format:   "text",
			expected: "test\t\t\t''\t''\t2017-01-01 01:02:03.000000\t2017-01-01 01:02:04.000001\t1.000001\t0.000000\t0.000000\t0.000000\t\t\"sql1\"\t\"[REDACTED]\"\t0\t0\t\"\"\t\"ks\"\t\"table\"\t\"PRIMARY\"\t\"suuid\"\tfalse\t[\"ks1.tbl1\",\"ks2.tbl2\"]\n",
			bindVars: stringBindVar,
		}, { //6
			redact:   false,
			format:   "json",
			expected: "{\n    \"BindVars\": {\n        \"strVal\": {\n            \"type\": \"VARCHAR\",\n            \"value\": \"abc\"\n        }\n    },\n    \"Cached Plan\": false,\n    \"CommitTime\": 0,\n    \"Effective Caller\": \"\",\n    \"End\": \"2017-01-01 01:02:04.000001\",\n    \"Error\": \"\",\n    \"ExecuteTime\": 0,\n    \"ImmediateCaller\": \"\",\n    \"Keyspace\": \"ks\",\n    \"Method\": \"test\",\n    \"PlanTime\": 0,\n    \"RemoteAddr\": \"\",\n    \"RowsAffected\": 0,\n    \"SQL\": \"sql1\",\n    \"SessionUUID\": \"suuid\",\n    \"ShardQueries\": 0,\n    \"Start\": \"2017-01-01 01:02:03.000000\",\n    \"StmtType\": \"\",\n    \"Table\": \"table\",\n    \"TablesUsed\": [\n        \"ks1.tbl1\",\n        \"ks2.tbl2\"\n    ],\n    \"TabletType\": \"PRIMARY\",\n    \"TotalTime\": 1.000001,\n    \"Username\": \"\"\n}",
			bindVars: stringBindVar,
		}, { //7
			redact:   true,
			format:   "json",
			expected: "{\n    \"BindVars\": \"[REDACTED]\",\n    \"Cached Plan\": false,\n    \"CommitTime\": 0,\n    \"Effective Caller\": \"\",\n    \"End\": \"2017-01-01 01:02:04.000001\",\n    \"Error\": \"\",\n    \"ExecuteTime\": 0,\n    \"ImmediateCaller\": \"\",\n    \"Keyspace\": \"ks\",\n    \"Method\": \"test\",\n    \"PlanTime\": 0,\n    \"RemoteAddr\": \"\",\n    \"RowsAffected\": 0,\n    \"SQL\": \"sql1\",\n    \"SessionUUID\": \"suuid\",\n    \"ShardQueries\": 0,\n    \"Start\": \"2017-01-01 01:02:03.000000\",\n    \"StmtType\": \"\",\n    \"Table\": \"table\",\n    \"TablesUsed\": [\n        \"ks1.tbl1\",\n        \"ks2.tbl2\"\n    ],\n    \"TabletType\": \"PRIMARY\",\n    \"TotalTime\": 1.000001,\n    \"Username\": \"\"\n}",
			bindVars: stringBindVar,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			logStats.BindVariables = test.bindVars
			for _, variable := range logStats.BindVariables {
				fmt.Println("->" + fmt.Sprintf("%v", variable))
			}
			*streamlog.RedactDebugUIQueries = test.redact
			*streamlog.QueryLogFormat = test.format
			if test.format == "text" {
				got := testFormat(t, logStats, params)
				assert.Equal(t, test.expected, got)
				for _, variable := range logStats.BindVariables {
					fmt.Println("->" + fmt.Sprintf("%v", variable))
				}
				return
			}

			got := testFormat(t, logStats, params)
			var parsed map[string]any
			err := json.Unmarshal([]byte(got), &parsed)
			assert.NoError(t, err)
			assert.NotNil(t, parsed)
			formatted, err := json.MarshalIndent(parsed, "", "    ")
			assert.NoError(t, err)
			assert.Equal(t, test.expected, string(formatted))
		})
	}
}

func TestLogStatsFilter(t *testing.T) {
	defer func() { *streamlog.QueryLogFilterTag = "" }()

	logStats := NewLogStats(context.Background(), "test", "sql1 /* LOG_THIS_QUERY */", "", map[string]*querypb.BindVariable{"intVal": sqltypes.Int64BindVariable(1)})
	logStats.StartTime = time.Date(2017, time.January, 1, 1, 2, 3, 0, time.UTC)
	logStats.EndTime = time.Date(2017, time.January, 1, 1, 2, 4, 1234, time.UTC)
	params := map[string][]string{"full": {}}

	got := testFormat(t, logStats, params)
	want := "test\t\t\t''\t''\t2017-01-01 01:02:03.000000\t2017-01-01 01:02:04.000001\t1.000001\t0.000000\t0.000000\t0.000000\t\t\"sql1 /* LOG_THIS_QUERY */\"\tmap[intVal:type:INT64 value:\"1\"]\t0\t0\t\"\"\t\"\"\t\"\"\t\"\"\t\"\"\tfalse\t[]\n"
	assert.Equal(t, want, got)

	*streamlog.QueryLogFilterTag = "LOG_THIS_QUERY"
	got = testFormat(t, logStats, params)
	want = "test\t\t\t''\t''\t2017-01-01 01:02:03.000000\t2017-01-01 01:02:04.000001\t1.000001\t0.000000\t0.000000\t0.000000\t\t\"sql1 /* LOG_THIS_QUERY */\"\tmap[intVal:type:INT64 value:\"1\"]\t0\t0\t\"\"\t\"\"\t\"\"\t\"\"\t\"\"\tfalse\t[]\n"
	assert.Equal(t, want, got)

	*streamlog.QueryLogFilterTag = "NOT_THIS_QUERY"
	got = testFormat(t, logStats, params)
	want = ""
	assert.Equal(t, want, got)
}

func TestLogStatsRowThreshold(t *testing.T) {
	defer func() { *streamlog.QueryLogRowThreshold = 0 }()

	logStats := NewLogStats(context.Background(), "test", "sql1 /* LOG_THIS_QUERY */", "", map[string]*querypb.BindVariable{"intVal": sqltypes.Int64BindVariable(1)})
	logStats.StartTime = time.Date(2017, time.January, 1, 1, 2, 3, 0, time.UTC)
	logStats.EndTime = time.Date(2017, time.January, 1, 1, 2, 4, 1234, time.UTC)
	params := map[string][]string{"full": {}}

	got := testFormat(t, logStats, params)
	want := "test\t\t\t''\t''\t2017-01-01 01:02:03.000000\t2017-01-01 01:02:04.000001\t1.000001\t0.000000\t0.000000\t0.000000\t\t\"sql1 /* LOG_THIS_QUERY */\"\tmap[intVal:type:INT64 value:\"1\"]\t0\t0\t\"\"\t\"\"\t\"\"\t\"\"\t\"\"\tfalse\t[]\n"
	assert.Equal(t, want, got)

	*streamlog.QueryLogRowThreshold = 0
	got = testFormat(t, logStats, params)
	want = "test\t\t\t''\t''\t2017-01-01 01:02:03.000000\t2017-01-01 01:02:04.000001\t1.000001\t0.000000\t0.000000\t0.000000\t\t\"sql1 /* LOG_THIS_QUERY */\"\tmap[intVal:type:INT64 value:\"1\"]\t0\t0\t\"\"\t\"\"\t\"\"\t\"\"\t\"\"\tfalse\t[]\n"
	assert.Equal(t, want, got)
	*streamlog.QueryLogRowThreshold = 1
	got = testFormat(t, logStats, params)
	assert.Empty(t, got)
}

func TestLogStatsContextHTML(t *testing.T) {
	html := "HtmlContext"
	callInfo := &fakecallinfo.FakeCallInfo{
		Html: html,
	}
	ctx := callinfo.NewContext(context.Background(), callInfo)
	logStats := NewLogStats(ctx, "test", "sql1", "", map[string]*querypb.BindVariable{})
	if string(logStats.ContextHTML()) != html {
		t.Fatalf("expect to get html: %s, but got: %s", html, string(logStats.ContextHTML()))
	}
}

func TestLogStatsErrorStr(t *testing.T) {
	logStats := NewLogStats(context.Background(), "test", "sql1", "", map[string]*querypb.BindVariable{})
	if logStats.ErrorStr() != "" {
		t.Fatalf("should not get error in stats, but got: %s", logStats.ErrorStr())
	}
	errStr := "unknown error"
	logStats.Error = errors.New(errStr)
	if !strings.Contains(logStats.ErrorStr(), errStr) {
		t.Fatalf("expect string '%s' in error message, but got: %s", errStr, logStats.ErrorStr())
	}
}

func TestLogStatsRemoteAddrUsername(t *testing.T) {
	logStats := NewLogStats(context.Background(), "test", "sql1", "", map[string]*querypb.BindVariable{})
	addr, user := logStats.RemoteAddrUsername()
	if addr != "" {
		t.Fatalf("remote addr should be empty")
	}
	if user != "" {
		t.Fatalf("username should be empty")
	}

	remoteAddr := "1.2.3.4"
	username := "vt"
	callInfo := &fakecallinfo.FakeCallInfo{
		Remote: remoteAddr,
		User:   username,
	}
	ctx := callinfo.NewContext(context.Background(), callInfo)
	logStats = NewLogStats(ctx, "test", "sql1", "", map[string]*querypb.BindVariable{})
	addr, user = logStats.RemoteAddrUsername()
	if addr != remoteAddr {
		t.Fatalf("expected to get remote addr: %s, but got: %s", remoteAddr, addr)
	}
	if user != username {
		t.Fatalf("expected to get username: %s, but got: %s", username, user)
	}
}
