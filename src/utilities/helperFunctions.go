package utilities

import (
	"net/url"
	"time"
	"types"
	"encoding/json"
)

func GetPagePath(page string) string {
	return "pages/" + page
}

func ExtractMemory(memoryForm url.Values) types.Memory {

	return types.Memory {
		memoryForm.Get("memory_url"),
		memoryForm.Get("memory_title"),
		time.Now().Unix(),
	}

}

func ExtractJSONMemory(body []byte) (types.Memory, error) {
	var content types.Memory
	err := json.Unmarshal(body, content)
	return content, err
}
