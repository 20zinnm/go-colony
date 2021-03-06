// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package colony

import "testing"

func TestBoolColony(t *testing.T) {
	colony := NewBoolColony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(bool)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(bool)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var BoolBenchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkBoolColony_Insert(b *testing.B) {
	for _, bm := range BoolBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewBoolColony(count / 4) // hint
				ptrs := make([]*bool, count)
				newbool := new(bool)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newbool)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkBoolSlice(b *testing.B) {
	for _, bm := range BoolBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]bool, count/5)
				ptrs := make([]*bool, count)
				newbool := new(bool)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newbool)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestByteColony(t *testing.T) {
	colony := NewByteColony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(byte)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(byte)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var ByteBenchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkByteColony_Insert(b *testing.B) {
	for _, bm := range ByteBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewByteColony(count / 4) // hint
				ptrs := make([]*byte, count)
				newbyte := new(byte)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newbyte)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkByteSlice(b *testing.B) {
	for _, bm := range ByteBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]byte, count/5)
				ptrs := make([]*byte, count)
				newbyte := new(byte)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newbyte)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestComplex128Colony(t *testing.T) {
	colony := NewComplex128Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(complex128)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(complex128)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Complex128Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkComplex128Colony_Insert(b *testing.B) {
	for _, bm := range Complex128Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewComplex128Colony(count / 4) // hint
				ptrs := make([]*complex128, count)
				newcomplex128 := new(complex128)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newcomplex128)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkComplex128Slice(b *testing.B) {
	for _, bm := range Complex128Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]complex128, count/5)
				ptrs := make([]*complex128, count)
				newcomplex128 := new(complex128)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newcomplex128)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestComplex64Colony(t *testing.T) {
	colony := NewComplex64Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(complex64)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(complex64)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Complex64Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkComplex64Colony_Insert(b *testing.B) {
	for _, bm := range Complex64Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewComplex64Colony(count / 4) // hint
				ptrs := make([]*complex64, count)
				newcomplex64 := new(complex64)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newcomplex64)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkComplex64Slice(b *testing.B) {
	for _, bm := range Complex64Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]complex64, count/5)
				ptrs := make([]*complex64, count)
				newcomplex64 := new(complex64)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newcomplex64)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestErrorColony(t *testing.T) {
	colony := NewErrorColony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(error)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(error)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var ErrorBenchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkErrorColony_Insert(b *testing.B) {
	for _, bm := range ErrorBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewErrorColony(count / 4) // hint
				ptrs := make([]*error, count)
				newerror := new(error)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newerror)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkErrorSlice(b *testing.B) {
	for _, bm := range ErrorBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]error, count/5)
				ptrs := make([]*error, count)
				newerror := new(error)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newerror)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestFloat32Colony(t *testing.T) {
	colony := NewFloat32Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(float32)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(float32)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Float32Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkFloat32Colony_Insert(b *testing.B) {
	for _, bm := range Float32Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewFloat32Colony(count / 4) // hint
				ptrs := make([]*float32, count)
				newfloat32 := new(float32)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newfloat32)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkFloat32Slice(b *testing.B) {
	for _, bm := range Float32Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]float32, count/5)
				ptrs := make([]*float32, count)
				newfloat32 := new(float32)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newfloat32)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestFloat64Colony(t *testing.T) {
	colony := NewFloat64Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(float64)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(float64)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Float64Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkFloat64Colony_Insert(b *testing.B) {
	for _, bm := range Float64Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewFloat64Colony(count / 4) // hint
				ptrs := make([]*float64, count)
				newfloat64 := new(float64)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newfloat64)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkFloat64Slice(b *testing.B) {
	for _, bm := range Float64Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]float64, count/5)
				ptrs := make([]*float64, count)
				newfloat64 := new(float64)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newfloat64)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestIntColony(t *testing.T) {
	colony := NewIntColony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(int)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(int)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var IntBenchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkIntColony_Insert(b *testing.B) {
	for _, bm := range IntBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewIntColony(count / 4) // hint
				ptrs := make([]*int, count)
				newint := new(int)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newint)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkIntSlice(b *testing.B) {
	for _, bm := range IntBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]int, count/5)
				ptrs := make([]*int, count)
				newint := new(int)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newint)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestInt16Colony(t *testing.T) {
	colony := NewInt16Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(int16)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(int16)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Int16Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkInt16Colony_Insert(b *testing.B) {
	for _, bm := range Int16Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewInt16Colony(count / 4) // hint
				ptrs := make([]*int16, count)
				newint16 := new(int16)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newint16)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkInt16Slice(b *testing.B) {
	for _, bm := range Int16Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]int16, count/5)
				ptrs := make([]*int16, count)
				newint16 := new(int16)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newint16)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestInt32Colony(t *testing.T) {
	colony := NewInt32Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(int32)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(int32)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Int32Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkInt32Colony_Insert(b *testing.B) {
	for _, bm := range Int32Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewInt32Colony(count / 4) // hint
				ptrs := make([]*int32, count)
				newint32 := new(int32)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newint32)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkInt32Slice(b *testing.B) {
	for _, bm := range Int32Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]int32, count/5)
				ptrs := make([]*int32, count)
				newint32 := new(int32)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newint32)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestInt64Colony(t *testing.T) {
	colony := NewInt64Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(int64)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(int64)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Int64Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkInt64Colony_Insert(b *testing.B) {
	for _, bm := range Int64Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewInt64Colony(count / 4) // hint
				ptrs := make([]*int64, count)
				newint64 := new(int64)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newint64)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkInt64Slice(b *testing.B) {
	for _, bm := range Int64Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]int64, count/5)
				ptrs := make([]*int64, count)
				newint64 := new(int64)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newint64)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestInt8Colony(t *testing.T) {
	colony := NewInt8Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(int8)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(int8)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Int8Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkInt8Colony_Insert(b *testing.B) {
	for _, bm := range Int8Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewInt8Colony(count / 4) // hint
				ptrs := make([]*int8, count)
				newint8 := new(int8)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newint8)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkInt8Slice(b *testing.B) {
	for _, bm := range Int8Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]int8, count/5)
				ptrs := make([]*int8, count)
				newint8 := new(int8)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newint8)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestRuneColony(t *testing.T) {
	colony := NewRuneColony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(rune)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(rune)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var RuneBenchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkRuneColony_Insert(b *testing.B) {
	for _, bm := range RuneBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewRuneColony(count / 4) // hint
				ptrs := make([]*rune, count)
				newrune := new(rune)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newrune)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkRuneSlice(b *testing.B) {
	for _, bm := range RuneBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]rune, count/5)
				ptrs := make([]*rune, count)
				newrune := new(rune)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newrune)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestStringColony(t *testing.T) {
	colony := NewStringColony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(string)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(string)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var StringBenchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkStringColony_Insert(b *testing.B) {
	for _, bm := range StringBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewStringColony(count / 4) // hint
				ptrs := make([]*string, count)
				newstring := new(string)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newstring)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkStringSlice(b *testing.B) {
	for _, bm := range StringBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]string, count/5)
				ptrs := make([]*string, count)
				newstring := new(string)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newstring)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestUintColony(t *testing.T) {
	colony := NewUintColony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(uint)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(uint)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var UintBenchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkUintColony_Insert(b *testing.B) {
	for _, bm := range UintBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewUintColony(count / 4) // hint
				ptrs := make([]*uint, count)
				newuint := new(uint)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newuint)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkUintSlice(b *testing.B) {
	for _, bm := range UintBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]uint, count/5)
				ptrs := make([]*uint, count)
				newuint := new(uint)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newuint)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestUint16Colony(t *testing.T) {
	colony := NewUint16Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(uint16)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(uint16)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Uint16Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkUint16Colony_Insert(b *testing.B) {
	for _, bm := range Uint16Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewUint16Colony(count / 4) // hint
				ptrs := make([]*uint16, count)
				newuint16 := new(uint16)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newuint16)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkUint16Slice(b *testing.B) {
	for _, bm := range Uint16Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]uint16, count/5)
				ptrs := make([]*uint16, count)
				newuint16 := new(uint16)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newuint16)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestUint32Colony(t *testing.T) {
	colony := NewUint32Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(uint32)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(uint32)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Uint32Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkUint32Colony_Insert(b *testing.B) {
	for _, bm := range Uint32Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewUint32Colony(count / 4) // hint
				ptrs := make([]*uint32, count)
				newuint32 := new(uint32)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newuint32)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkUint32Slice(b *testing.B) {
	for _, bm := range Uint32Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]uint32, count/5)
				ptrs := make([]*uint32, count)
				newuint32 := new(uint32)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newuint32)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestUint64Colony(t *testing.T) {
	colony := NewUint64Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(uint64)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(uint64)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Uint64Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkUint64Colony_Insert(b *testing.B) {
	for _, bm := range Uint64Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewUint64Colony(count / 4) // hint
				ptrs := make([]*uint64, count)
				newuint64 := new(uint64)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newuint64)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkUint64Slice(b *testing.B) {
	for _, bm := range Uint64Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]uint64, count/5)
				ptrs := make([]*uint64, count)
				newuint64 := new(uint64)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newuint64)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestUint8Colony(t *testing.T) {
	colony := NewUint8Colony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(uint8)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(uint8)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var Uint8Benchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkUint8Colony_Insert(b *testing.B) {
	for _, bm := range Uint8Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewUint8Colony(count / 4) // hint
				ptrs := make([]*uint8, count)
				newuint8 := new(uint8)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newuint8)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkUint8Slice(b *testing.B) {
	for _, bm := range Uint8Benchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]uint8, count/5)
				ptrs := make([]*uint8, count)
				newuint8 := new(uint8)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newuint8)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}

func TestUintptrColony(t *testing.T) {
	colony := NewUintptrColony(1)
	t.Run("Insert", func(t *testing.T) {
		newT := new(uintptr)
		tp := colony.Insert(newT)
		if *tp != *newT {
			t.Fatalf("value of pointer returned from Insert does not equal the inserted value: (*newT) %v != (*tp) %v", *newT, *tp)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		newT := new(uintptr)
		tp := colony.Insert(newT)
		colony.Delete(tp)
		tp2 := colony.Insert(newT)
		if tp != tp2 { // both should be allocated to the same spot
			t.Fatalf("deletion failed")
		}
	})
}

var UintptrBenchmarks = []struct {
	name  string
	count uint
}{
	{"10", 10},
	{"100", 100},
	{"1000", 1000},
	{"10000", 10000},
	{"100000", 100000},
	{"1000000", 1000000},
}

func BenchmarkUintptrColony_Insert(b *testing.B) {
	for _, bm := range UintptrBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				colony := NewUintptrColony(count / 4) // hint
				ptrs := make([]*uintptr, count)
				newuintptr := new(uintptr)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						ptrs[j] = colony.Insert(newuintptr)
					}
				}
			}
		}(bm.count))
	}
}

func BenchmarkUintptrSlice(b *testing.B) {
	for _, bm := range UintptrBenchmarks {
		b.Run(bm.name, func(count uint) func(*testing.B) {
			return func(b *testing.B) {
				// setup
				arr := make([]uintptr, count/5)
				ptrs := make([]*uintptr, count)
				newuintptr := new(uintptr)

				b.ReportAllocs()
				b.ResetTimer()

				for i := 0; i < b.N; i++ {
					for j := 0; uint(j) < count; j++ {
						arr = append(arr, *newuintptr)
						ptrs[j] = &arr[j]
					}
				}
			}
		}(bm.count))
	}
}
