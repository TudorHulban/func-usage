# Understanding the Output

The analyzer produces a CSV-style table where each row represents a function or method and each column describes how it is used across the repository. This data provides immediate, actionable insight into the structure, behavior, and quality of the codebase.

## Hotspots

Functions with high **Total** usage are called frequently and form the core execution paths of the project. These are candidates for:

- performance review  
- caching  
- careful refactoring  
- stability guarantees  

Hotspots highlight where changes have the largest impact.

## Public Surface

Functions with high **External** usage:

- are used by other packages  
- must remain stable  
- should be documented  
- should have strong test coverage  

This helps maintain a clean and intentional public surface.

## Internal Architecture

Functions with high **Internal** usage are central to the repository's internal design. They define the building blocks of the implementation and should be treated as core components.

## Test-Driven Usage

High **InternalTests** or **ExternalTests** counts indicate functions that are exercised heavily in tests. This can reveal:

- helpers used primarily in testing  
- functions with strong test coverage  
- potential over-testing of low-value helpers  

## Untested Functions

Functions with zero test counts may require:

- additional tests  
- review for necessity  
- evaluation for removal  

This helps maintain code quality and reduce risk.

## Dead or Low-Usage Code

Functions with very low **Total** usage may be unused or rarely used. These are candidates for:

- removal  
- consolidation  
- reducing accidental public exposure  

This supports ongoing cleanup and simplification.

## Standard Library Usage Patterns

High usage of standard library functions (e.g., I/O, formatting, string building) reveals patterns in how the module constructs output or processes data. This can guide optimization decisions such as:

- reducing allocations  
- pooling builders  
- consolidating write paths  

## Practical Workflows

### Refactoring Planning

Priority 1 (High Risk):

    - High Total + Low Tests → Add tests before refactoring
    - High External + Fragile → Create stable interfaces first

Priority 2 (High Impact):

    - High Internal + Complex → Simplify architecture
    - High Total + Performance issues → Optimize

Priority 3 (Cleanup):

    - Low Total + No Tests → Consider removal
    - Low External + Public → Consider making internal

### Team Onboarding

Core Functions: New engineers should understand high-Total functions first.  
Public API: Start with high-External functions for integration work.  
Test Helpers: High-Test functions show testing patterns.

### Trend Analysis

Track over time:

- Growth of public surface area
- Increasing complexity in hotspots
- Test coverage drift
- Dead code accumulation rate

## Summary

The output provides a clear, data-driven view of:

- performance hotspots  
- public surface stability  
- internal architecture  
- test coverage gaps  
- dead code  
- refactoring priorities  

It is static, deterministic, and suitable for CI integration, making it a reliable tool for ongoing maintenance and architectural clarity.
