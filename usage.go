package funcusage

import (
	"fmt"
	"go/token"
	"strings"
)

// FunctionUsage describes how a single function or method is used across the module.
type FunctionUsage struct {
	// Key is the canonical identity of the function or method.
	// Example (function): "github.com/me/project/pkg.DoThing"
	// Example (method):   "github.com/me/project/pkg.(*User).Save"
	Key string

	// Name is the short name of the function or method (without package or receiver).
	Name string

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

func (u *FunctionUsage) updateOccurences(callerPkg, calledPkg string, callerIsTest bool) {
	if strings.SplitN(callerPkg, " ", 2)[0] == calledPkg {
		if callerIsTest {
			u.InternalTestsCount++
		} else {
			u.InternalCount++
		}

		return
	}

	if callerIsTest {
		u.ExternalTestsCount++
	} else {
		u.ExternalCount++
	}
}

type Usage []FunctionUsage

func (u Usage) String() string {
	lines := make([]string, 0, 1+len(u))

	lines = append(
		lines,
		"Name,Key,Location,Internal,InternalTests,External,ExternalTests,Total",
	)

	for _, fu := range u {
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
