package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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