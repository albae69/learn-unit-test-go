package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bae")
	}
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bae")
	require.Equal(t, "Hello Bae", result, "Result must be Hello Bae")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Bae")
	assert.Equal(t, "Hello Bae", result, "Result must be Hello Bae")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bae")

	if result != "Hello Bae" {
		t.Error("Result must be Hello Bae")
	}

	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldHaqi(t *testing.T) {
	result := HelloWorld("Haqi")

	if result != "Hello Haqi" {
		t.Fatal("Result must be Hello Haqi")
	}

	fmt.Println("TestHelloWorldHaqi Done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("this unit test is not working on mac OS")
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Bae", func(t *testing.T) {
		result := HelloWorld("Bae")
		assert.Equal(t, "Hello Bae", result, "Result must be Hello Bae")
	})

	t.Run("Haqi", func(t *testing.T) {
		result := HelloWorld("Haqi")
		assert.Equal(t, "Hello Haqi", result, "Result must be Hello Haqi")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before Unit Test")

	m.Run()

	// after
	fmt.Println("After Unit Test")

}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Bae)",
			request:  "Bae",
			expected: "Hello Bae",
		},
		{
			name:     "HelloWorld(Haqi)",
			request:  "Haqi",
			expected: "Hello Haqi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
