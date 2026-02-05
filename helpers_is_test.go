package funcusage

import (
	"testing"
)

func TestIsExternalPackage(t *testing.T) {
	tests := []struct {
		desc       string
		pkgPath    string
		modulePath string
		isExternal bool
	}{
		{
			desc:       "1. stdlib fmt",
			pkgPath:    "fmt",
			modulePath: "github.com/tudor/project",
			isExternal: true,
		},
		{
			desc:       "2. stdlib strings",
			pkgPath:    "strings",
			modulePath: "github.com/tudor/project",
			isExternal: true,
		},
		{
			desc:       "3. internal package exact match",
			pkgPath:    "github.com/tudor/project",
			modulePath: "github.com/tudor/project",
			isExternal: false,
		},
		{
			desc:       "4. internal subpackage",
			pkgPath:    "github.com/tudor/project/internal/foo",
			modulePath: "github.com/tudor/project",
			isExternal: false,
		},
		{
			desc:       "5. external different module",
			pkgPath:    "github.com/other/module",
			modulePath: "github.com/tudor/project",
			isExternal: true,
		},
		{
			desc:       "6. external deeper module",
			pkgPath:    "github.com/tudor/other",
			modulePath: "github.com/tudor/project",
			isExternal: true,
		},
		{
			desc:       "7. local test module (single segment) should be internal",
			pkgPath:    "simple",
			modulePath: "simple",
			isExternal: false,
		},
		{
			desc:       "8. local test module subpackage",
			pkgPath:    "simple/foo",
			modulePath: "simple",
			isExternal: false,
		},
		{
			desc:       "9. local call in same package (empty pkgPath)",
			pkgPath:    "",
			modulePath: "simple",
			isExternal: false,
		},
	}

	for _, tt := range tests {
		got := isExternalPackage(
			tt.pkgPath,
			tt.modulePath,
		)

		if got != tt.isExternal {
			t.Errorf(
				"%s: isExternalPackage(%q, %q) = %v, want %v",

				tt.desc,
				tt.pkgPath,
				tt.modulePath,
				got,
				tt.isExternal,
			)
		}
	}
}
