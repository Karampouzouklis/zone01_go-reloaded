# Performance Optimization Documentation

## String Builder Optimization (TASK-013)

### Problem
Original implementation used string concatenation to reconstruct text from tokens:
```go
// BEFORE: Inefficient string concatenation
result := ""
for _, token := range tokens {
    result += token.Value  // Creates new string each time
}
```

### Solution
Optimized implementation using `strings.Builder` with pre-allocation:
```go
// AFTER: Efficient string builder with pre-allocation
var builder strings.Builder
builder.Grow(len(content)) // Pre-allocate based on input size
for _, token := range tokens {
    builder.WriteString(token.Value)
}
result := builder.String()
```

### Performance Results

#### Benchmark Comparison
```
BenchmarkStringBuilder/StringConcatenation-12    7003    183093 ns/op    601004 B/op    589 allocs/op
BenchmarkStringBuilder/StringBuilder-12        342087      3128 ns/op      2304 B/op      1 allocs/op
```

#### Performance Improvement
- **Speed**: 54x faster (183μs → 3.1μs)
- **Memory**: 260x less memory (601KB → 2.3KB)
- **Allocations**: 589x fewer allocations (589 → 1)

### Why String Builder is Better
1. **Single Allocation**: Pre-allocates buffer, grows as needed
2. **No String Copying**: Appends directly to internal buffer
3. **Memory Efficient**: Reuses same buffer instead of creating new strings
4. **Predictable Performance**: O(n) instead of O(n²) for concatenation

### Implementation Location
- **File**: `main.go` lines 35-41
- **Function**: `run()` function
- **Optimization**: Added in TASK-013 performance optimization phase

### Usage Commands
```bash
# Run string builder comparison benchmark
go test -bench=BenchmarkStringBuilder -benchmem ./transform/

# Run all performance benchmarks
go test -bench=. -benchmem ./transform/

# Generate memory profile
go test -bench=BenchmarkMemoryUsage -memprofile=mem.prof ./transform/
```

### Key Learning
String concatenation in Go creates a new string object each time, leading to O(n²) performance. String builder maintains an internal buffer and only allocates when necessary, achieving O(n) performance.