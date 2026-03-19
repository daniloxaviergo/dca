package assets

import (
	"fmt"
	"testing"
)

func TestDebugLoadMoreError(t *testing.T) {
	m := NewAssetHistoryModal()
	m.AssetTicker = "BTC"
	m.Loaded = true
	m.Visible = true

	err := m.LoadMore("non_existent_file.json")
	fmt.Printf("Error from LoadMore: %v\n", err)
	fmt.Printf("AllLoaded: %v\n", m.AllLoaded)
	fmt.Printf("Loading: %v\n", m.Loading)
	fmt.Printf("Error field: %v\n", m.Error)
}
