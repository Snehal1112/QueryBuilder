package ddl

import "testing"

func TestNewBuilder(t *testing.T) {
	NewCreateQuery(&Builder{})
}
