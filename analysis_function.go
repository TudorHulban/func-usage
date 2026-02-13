package funcusage

import (
	"fmt"
	"go/token"
	"strconv"
	"strings"
)

type (
	namePackage string
	nameObject  string
)

// AnalysisFunction describes how a single function or method is used across the module.
type AnalysisFunction struct {
	// Key is the canonical identity of the function or method.
	// Example (function): "github.com/me/project/pkg.DoThing"
	// Example (method):   "github.com/me/project/pkg.(*User).Save"
	Key string

	// Name is the short name of the function or method (without package or receiver).
	Name string

	// MethodOf highlights the object name for which the method belongs to.
	// Alias for Object, but only populated for methods.
	MethodOf nameObject

	TypesParams  []string
	TypesResults []string

	// Position is the source position of the function declaration.
	Position token.Position

	// InternalCount is the number of calls from within the same package.
	// Does not include InternalTestsCount.
	InternalCount int

	// InternalTestsCount is the number of calls from within the same package tests.
	InternalTestsCount int

	// ExternalCount is the number of calls from other packages.
	// Does not include ExternalTestsCount.
	ExternalCount int

	// ExternalTestsCount is the number of calls from other packages tests.
	ExternalTestsCount int
}

func (fa *AnalysisFunction) updateOccurences(callerPkg, calledPkg string, callerIsTest bool) {
	if strings.SplitN(callerPkg, " ", 2)[0] == calledPkg {
		if callerIsTest {
			fa.InternalTestsCount++
		} else {
			fa.InternalCount++
		}

		return
	}

	if callerIsTest {
		fa.ExternalTestsCount++
	} else {
		fa.ExternalCount++
	}
}

type Analysis []AnalysisFunction

func (a Analysis) PrintWith(printer *Printer) {
	fmt.Println(
		strings.Join(printer.columns, ", "),
	)

	for _, fa := range a {
		var row []string

		for _, col := range printer.columns {
			switch col {
			case _LabelName:
				row = append(row, fa.Name)

			case "Key":
				row = append(row, fa.Key)

			case _LabelMethodOf:
				row = append(row, string(fa.MethodOf))

			case _LabelTotal:
				row = append(
					row,
					strconv.Itoa(fa.InternalCount+fa.InternalTestsCount+fa.ExternalCount+fa.ExternalTestsCount),
				)

			case _LabelTypesParams:
				row = append(
					row,
					"\"<"+strings.Join(fa.TypesParams, ",")+">\"",
				)

			case _LabelTypesResults:
				row = append(
					row,
					"\"<"+strings.Join(fa.TypesResults, ", ")+">\"",
				)
			}
		}

		fmt.Println(strings.Join(row, ", "))
	}
}

func (a Analysis) String() string {
	lines := make([]string, 0, 1+len(a))

	lines = append(
		lines,
		"Name,Key,Method of,Location,Internal,InternalTests,External,ExternalTests,Total,TypesParams,TypesResults",
	)

	for _, fa := range a {
		total := fa.InternalCount +
			fa.InternalTestsCount +
			fa.ExternalCount +
			fa.ExternalTestsCount

		lines = append(
			lines,
			fmt.Sprintf(
				"%s,%s,%s,%s,%d,%d,%d,%d,%d,%q,%q",

				fa.Name,
				fa.Key,
				fa.MethodOf,
				fa.Position,
				fa.InternalCount,
				fa.InternalTestsCount,
				fa.ExternalCount,
				fa.ExternalTestsCount,
				total,
				fa.TypesParams,
				fa.TypesResults,
			),
		)
	}

	return strings.Join(lines, "\n")
}
