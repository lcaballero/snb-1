package sql_utils

import (
	"testing"
)

func Test_Loaded_Cache_Entries(t *testing.T) {

	if CacheEntries == nil {
		t.Error("CacheEntries should be initialized, but isn't.")
	}
}
