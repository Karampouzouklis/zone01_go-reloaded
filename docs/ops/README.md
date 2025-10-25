# Operations Documentation

## Overview

This directory contains operational runbooks and procedures for the go-reloaded project.

## Documentation Structure

### Core Operations
- `testing.md` - Testing procedures and quality gates
- `release.md` - Release process and deployment checklist
- `development.md` - Development environment setup

### Templates
- `templates/` - Standardized templates for common procedures
  - `test-report.md` - Template for test execution reports
  - `release-checklist.md` - Pre-release validation checklist
  - `bug-report.md` - Standardized bug reporting format

## Quality Gates

### Pre-commit Checks
- All tests pass: `go test ./...`
- Code formatting: `gofmt -d .`
- Linting: `golint ./...`
- Test coverage: `go test -cover ./...`

### Release Criteria
- 100% test coverage for core transformation logic
- All functional test cases pass
- Performance benchmarks within acceptable limits
- Documentation updated

## Evidence Storage

Store operational evidence in:
- `docs/ops/audits/` - Test reports and quality audits
- `releases/` - Release artifacts and changelogs
- `tasks/` - Task completion evidence

## Runbook Usage

Each runbook follows the format:
1. **Purpose** - What this procedure accomplishes
2. **Prerequisites** - Required setup or conditions
3. **Steps** - Detailed execution steps
4. **Validation** - How to verify success
5. **Troubleshooting** - Common issues and solutions