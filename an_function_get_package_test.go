package funcusage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPackage(t *testing.T) {
	t.Parallel()

	tests := []struct {
		description string
		key         string
		wantPkg     string
		wantErr     bool
	}{
		{
			description: "01 error missing slash and missing dot",
			key:         "invalid",
			wantErr:     true,
		},
		{
			description: "02 error slash but missing dot after slash",
			key:         "github.com/user/project/pkgfile",
			wantErr:     true,
		},
		{
			description: "03 dot based strconv.ParseInt",
			key:         "strconv.ParseInt",
			wantPkg:     "strconv",
		},
		{
			description: "04 dot based fmt.Printf",
			key:         "fmt.Printf",
			wantPkg:     "fmt",
		},
		{
			description: "05 dot based bytes.*bytes.Buffer.String",
			key:         "bytes.*bytes.Buffer.String",
			wantPkg:     "bytes",
		},
		{
			description: "06 slash based github.com/TudorHulban/func-usage.something",
			key:         "github.com/TudorHulban/func-usage.something",
			wantPkg:     "funcusage",
		},
		{
			description: "07 slash based internal/pkg/foo.go",
			key:         "internal/pkg/foo.go",
			wantPkg:     "foo",
		},
	}

	for _, tc := range tests {
		t.Run(
			tc.description,
			func(t *testing.T) {
				t.Parallel()

				fa := &AnalysisFunction{
					Key: tc.key,
				}

				gotPkg, errGetPackage := fa.getPackage()
				if tc.wantErr {
					require.Error(t,
						errGetPackage,
						tc.description,
					)

					return
				}

				require.NoError(t, errGetPackage, tc.description)
				require.Equal(t, tc.wantPkg, string(gotPkg), tc.description)
			},
		)
	}
}
