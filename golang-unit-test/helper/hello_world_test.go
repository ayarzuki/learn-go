package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{name: "Eko", request: "Eko", expected: "Hello Eko"},
		{name: "Kurniawan", request: "Kurniawan", expected: "Hello Kurniawan"},
		{name: "Khannedy", request: "Khannedy", expected: "Hello Khannedy"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	})
	t.Run("Kurniawan", func(t *testing.T) {
		result := HelloWorld("Kurniawan")
		require.Equal(t, "Hello Kurniawan", result, "Result must be 'Hello Kuniawan'")

	})
}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")
	m.Run()
	fmt.Println("Setelah Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skip test on windows")
	}

	result := HelloWorld("Eko")

	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")

	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("Test Hello World with Require is done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")

	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("Test Hello World with Assert is done")
}


func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		//error
		// panic("Result is not 'Hello Eko'")
		// t.Fail()
		t.Error("Result must be 'Hello Eko'")
	}
	fmt.Println("Test Hello World is done")
}

func TestHelloWorldKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy" {
		//error
		// panic("Result is not 'Hello Khannedy'")
		// t.FailNow()
		t.Fatal("Test Hello World Khannedy is failed")
	}
	fmt.Println("Test Hello World Khannedy is done")
}