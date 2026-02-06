package funcusage

import (
	"fmt"
	"go/token"
	"strconv"
	"strings"
)

const (
	_LabelName     = "Name"
	_LabelMethodOf = "Method of"
	_LabelTotal    = "Total"
)

// FunctionAnalysis describes how a single function or method is used across the module.
type FunctionAnalysis struct {
	// Key is the canonical identity of the function or method.
	// Example (function): "github.com/me/project/pkg.DoThing"
	// Example (method):   "github.com/me/project/pkg.(*User).Save"
	Key string

	// Name is the short name of the function or method (without package or receiver).
	Name string

	// MethodOf highlights the object name for which the method belongs to.
	// Alias for Object, but only populated for methods.
	MethodOf string

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

func (fa *FunctionAnalysis) updateOccurences(callerPkg, calledPkg string, callerIsTest bool) {
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

func (fa *FunctionAnalysis) getPackage() (string, error) {
	dotIndex := strings.Index(fa.Key, ".")
	if dotIndex == -1 {
		return "",
			fmt.Errorf(
				"invalid function key: no package found in %q",
				fa.Key,
			)
	}

	return fa.Key[:dotIndex], nil
}

type Analysis []FunctionAnalysis

func (a Analysis) PrintWith(printer Printer) {
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
				row = append(row, fa.MethodOf)

			case _LabelTotal:
				row = append(
					row,
					strconv.Itoa(fa.InternalCount+fa.InternalTestsCount+fa.ExternalCount+fa.ExternalTestsCount),
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
		"Name,Key,Location,Internal,InternalTests,External,ExternalTests,Total",
	)

	for _, fu := range a {
		total := fu.InternalCount +
			fu.InternalTestsCount +
			fu.ExternalCount +
			fu.ExternalTestsCount

		lines = append(
			lines,
			fmt.Sprintf(
				"%s,%s,%s,%d,%d,%d,%d,%d",

				fu.Name,
				fu.Key,
				fu.Position,
				fu.InternalCount,
				fu.InternalTestsCount,
				fu.ExternalCount,
				fu.ExternalTestsCount,
				total,
			),
		)
	}

	return strings.Join(lines, "\n")
}
