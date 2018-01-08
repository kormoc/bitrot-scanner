package pathing

import (
	"testing"
)

var testPaths = []struct {
	path        string
	expectedErr error
}{
	{"/foo/bar/baz", nil},
	{"/foo/bar/baz", nil},
	{"/foo/bar", nil},
	{"/foo", nil},
	{"/bar/baz", nil},
	{"/baz", nil},
	{"Invalid", ErrInvalidPath},
	{"/foo/bar/baz.txt", nil},
	{"/foo/bar/baz.readme", nil},
}

func TestPath(t *testing.T) {

	for _, testPath := range testPaths {
		testPath := testPath
		t.Run(testPath.path, func(t *testing.T) {
			t.Parallel()
			// fmt.Printf("Testing %v\n", testPath.path)
			res, err := New(testPath.path)
			if err != testPath.expectedErr {
				t.Fatal(err)
			}
			if testPath.expectedErr == nil && res == nil {
				t.Fatal("Unexpected nil result")
			}
			if testPath.expectedErr == nil && res.String() != testPath.path {
				t.Fatalf("Invalid reversal: Expected '%v' received '%v'\n", testPath.path, res.String())
			}
		})
	}
}

func BenchmarkPath(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		New(testPaths[0].path)
	}
}
