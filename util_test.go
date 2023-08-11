package sfconfig

import "testing"

var testData = map[string]string{
	"TestToSnakeCase": "TEST_TO_SNAKE_CASE",
	"testToSnakeCase": "TEST_TO_SNAKE_CASE",
	//"_TestToSnake":    "_TEST_TO_SNAKE",
	"TestToSnake_": "TEST_TO_SNAKE_",
	"TestToSnake1": "TEST_TO_SNAKE1",
}

func TestToSnakeCase(t *testing.T) {
	for key, value := range testData {
		if ToSnakeCase(key) != value {
			t.Errorf("ToSnakeCase value is wrong: %s, want: %s", ToSnakeCase(key), value)
		}
	}
}
