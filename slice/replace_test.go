package cxbytesslice

import (
        "fmt"
        "testing"
)

// TestReplaceByteSlice tests the ReplaceByteSlice function with various scenarios.
func TestReplaceByteSlice(t *testing.T) {
        tests := []struct {
                name   string
                target []byte
                source []byte
                want   []byte
        }{
                {"Target smaller than source", []byte{1, 2, 3}, []byte{4, 5, 6, 7}, []byte{4, 5, 6, 7}},
                {"Target larger than source", []byte{1, 2, 3, 4}, []byte{5, 6}, []byte{5, 6}},
                {"Target equal to source", []byte{1, 2, 3}, []byte{4, 5, 6}, []byte{4, 5, 6}},
        }

        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        targetCopy := make([]byte, len(tt.target))
                        copy(targetCopy, tt.target)
                        ReplaceByteSlice(&targetCopy, tt.source)
                        if len(targetCopy) != len(tt.want) || !equal(targetCopy, tt.want) {
                                t.Errorf("ReplaceByteSlice() = %v, want %v", targetCopy, tt.want)
                        }
                })
        }
}

// equal checks if two byte slices are equal.
func equal(a, b []byte) bool {
        if len(a) != len(b) {
                return false
        }
        for i := range a {
                if a[i] != b[i] {
                        return false
                }
        }
        return true
}

// BenchmarkReplaceByteSlice benchmarks the ReplaceByteSlice function.
func BenchmarkReplaceByteSlice(b *testing.B) {
        benchmarks := []struct {
                name   string
                target []byte
                source []byte
        }{
                {"Target smaller than source", make([]byte, 100), make([]byte, 200)},
                {"Target larger than source", make([]byte, 200), make([]byte, 100)},
                {"Target equal to source", make([]byte, 150), make([]byte, 150)},
        }

        for _, bm := range benchmarks {
                b.Run(bm.name, func(b *testing.B) {
                        targetCopy := make([]byte, len(bm.target))
                        copy(targetCopy, bm.target)
                        b.ResetTimer()
                        for i := 0; i < b.N; i++ {
                                ReplaceByteSlice(&targetCopy, bm.source)
                        }
                        b.ReportAllocs()
                })
        }
}

// ExampleReplaceByteSlice provides an example usage of ReplaceByteSlice.
func ExampleReplaceByteSlice() {
        target := []byte{1, 2, 3}
        source := []byte{4, 5, 6, 7}
        ReplaceByteSlice(&target, source)
        fmt.Println(target)
        // Output: [4 5 6 7]
}
