package pkgslist

import (
	"go/build"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

func TestFindAll(t *testing.T) {
	tests := map[string][]*build.Package{
		"": {
			{ImportPath: "a"},
			{ImportPath: "b"},
			{ImportPath: "b/c"},
		},
		"a": {
			{ImportPath: "a"},
		},
		"b": {
			{ImportPath: "b"},
			{ImportPath: "b/c"},
		},
		"x": {},
	}

	buildContext := build.Default
	buildContext.GOPATH, _ = filepath.Abs("testdata/")

	for prefix, expected := range tests {
		actual, err := FindAll(prefix, buildContext, 0)
		if err != nil {
			t.Fatalf("FindAll with prefix %#v failed: %v", err)
		}
		packageListsImportPathsAreEqual(t, actual, expected)
	}
}

func packageListsImportPathsAreEqual(t *testing.T, actual []*build.Package, expected []*build.Package) {
	apaths := packageListImportPaths(actual)
	epaths := packageListImportPaths(expected)
	sort.Strings(apaths)
	sort.Strings(epaths)
	if !reflect.DeepEqual(apaths, epaths) {
		t.Fatalf("package lists differ in import paths\nexpected: %s\ngot:      %s", epaths, apaths)
	}
}

func packageListImportPaths(ps []*build.Package) []string {
	ips := make([]string, len(ps))
	for i, p := range ps {
		ips[i] = p.ImportPath
	}
	return ips
}
