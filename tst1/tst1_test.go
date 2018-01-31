package tst

// To run test cases here, you can execute:
// `go test`
// `go test <name of test file>`, such as go test tst_test.go
// It looks like the name of the test file must end with _test.go when using `go test`.
// Here we are leveraging testify for assertions: https://github.com/stretchr/testify
// `go get github.com/stretchr/testify` (so remember to fetch this dependency first)
// `go get -u github.com/stretchr/testify` (updates testify)

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestXYZ (t *testing.T) {
    var a string = "Hello"
    var b string = "Hello"  
    assert.Equal(t, a, b, "The two words should be the same.")
}

