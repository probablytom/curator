package types

import "time"

type Memory struct {
	Url, Title string
	Submission_timestamp int64
}

func (mem Memory) generate_html() string {
	return "<a href=\"" + mem.Url + "\">" + mem.Title + "</a> at time " + time.Time{mem.Submission_timestamp}
}

