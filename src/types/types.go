package types

import (
	"encoding/json"
)

/*

	TYPES

 */

type Memory struct {
	Url string			`json:"url"`
	Title string			`json:"title"`
	Submission_timestamp int64	`json:"timestamp"`
}

type LoginDetails struct {
	Username string	`json:user`
	Password string	`json:pass`
}



/*

	METHODS

 */

// TODO: fix timestamp conversion
func (mem Memory) generate_html() string {
	return "<a href=\"" + mem.Url + "\">" + mem.Title + "</a> at time " + string(mem.Submission_timestamp)
}

func (mem Memory) toJSON() ([]byte, error) {
	return json.Marshal(mem)
}
