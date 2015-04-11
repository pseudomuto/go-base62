package base62

import (
	"fmt"
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, "0", Encode(0))
	assert.Equal(t, "1B", Encode(99))
}

func TestDecode(t *testing.T) {
	assert.Equal(t, int64(0), Decode("0"))
	assert.Equal(t, int64(99), Decode("1B"))
}

func ExampleEncode() {
	fmt.Println(Encode(99))
	// Output:
	// 1B
}

func ExampleDecode() {
	fmt.Println(Decode("1B"))
	// Output:
	// 99
}
