package utilities

import (
	"net/url"
	"time"
	"types"
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
