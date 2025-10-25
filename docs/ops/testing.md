# Testing Operations Runbook

## Purpose
Standardized testing procedures and quality gates for go-reloaded development.

## Prerequisites
- Go 1.21+ installed
- Project dependencies: `go mod download`
- Test files present in project structure

## Testing Procedures

### 1. Unit Testing
**Command**: `go test ./...`
**Purpose**: Verify individual component functionality

**Quality Gates**:
- All tests must pass
- No test should take longer than 1 second
- Test coverage must be ≥90% for core logic

### 2. Integration Testing
**Command**: `go test -tags=integration ./...`
**Purpose**: Verify component interactions

**Scope**:
- Full pipeline execution
- File I/O operations
- Error handling scenarios

### 3. Functional Testing
**Command**: `go test -run TestFunctionalCases`
**Purpose**: Verify all README examples work correctly

**Test Cases**:
- All transformation rules from README.md
- Complex interaction scenarios
- Edge cases and boundary conditions

### 4. Performance Testing
**Command**: `go test -bench=. -benchmem`
**Purpose**: Ensure acceptable performance characteristics

**Benchmarks**:
- Processing speed: >1MB/sec for typical text
- Memory usage: <2x input size
- No memory leaks in long-running tests

## Test Execution Checklist

- [ ] Run unit tests: `go test ./...`
- [ ] Check coverage: `go test -cover ./...`
- [ ] Run benchmarks: `go test -bench=.`
- [ ] Verify functional cases pass
- [ ] Test with sample files from README
- [ ] Validate error handling with malformed input

## Validation Criteria

### Pass Criteria
- All tests pass (exit code 0)
- Coverage ≥90% for transformation logic
- No performance regressions
- All functional examples work

### Fail Criteria
- Any test failure
- Coverage below threshold
- Performance degradation >20%
- Functional examples broken

## Troubleshooting

### Common Issues
1. **Test timeout**: Increase timeout or optimize slow tests
2. **Coverage gaps**: Add tests for uncovered code paths
3. **Flaky tests**: Identify and fix non-deterministic behavior
4. **Performance regression**: Profile and optimize bottlenecks

### Debug Commands
```bash
# Verbose test output
go test -v ./...

# Run specific test
go test -run TestSpecificFunction

# Profile memory usage
go test -memprofile=mem.prof

# Profile CPU usage
go test -cpuprofile=cpu.prof
```