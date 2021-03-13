package top

import (
	"github.com/lesovsky/pgcenter/internal/postgres"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getQueryReport(t *testing.T) {
	// Connect to Postgres and get random queryid needed for testcase.
	conn, err := postgres.NewTestConnect()
	assert.NoError(t, err)
	var queryid string
	err = conn.QueryRow(
		"SELECT left(md5(userid::text || dbid::text || queryid::text), 10) FROM pg_stat_statements LIMIT 1",
	).Scan(&queryid)
	assert.NoError(t, err)

	testcases := []struct {
		answer string
		want   string
	}{
		{answer: queryid, want: ""},
		{answer: "", want: "Report: do nothing"},
		{answer: "invalid", want: "Report: no statistics for such queryid"},
	}

	for _, tc := range testcases {
		_, got := getQueryReport(tc.answer, 130000, "public", conn)
		assert.Equal(t, tc.want, got)
	}
}
