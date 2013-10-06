package sql_utils

import (
	"sql_utils/caching"
	"testing"
)

func Test_Loaded_Cache_Entries(t *testing.T) {

	if caching.Cache() == nil {
		t.Error("CacheEntries should be initialized, but isn't.")
	}
}
