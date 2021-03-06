// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package colony

import (
	"sync"
	"unsafe"
)

// BoolColony represents a colony of Bools.
type BoolColony struct {
	entry *colonyGroupbool
}

// NewBoolColony returns a new colony of Bools.
func NewBoolColony(size uint) *BoolColony {
	return &BoolColony{
		entry: newboolGroup(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *BoolColony) Insert(t *bool) (tp *bool) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *BoolColony) Delete(tp *bool) {
	c.entry.Delete(tp)
}

func newboolGroup(previous *colonyGroupbool, size uint) *colonyGroupbool {
	var g colonyGroupbool
	if size == 0 {
		size = 8
	}
	g.data = make([]bool, size)
	g.free = make(chan *bool, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupbool struct {
	data     []bool
	free     chan *bool
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupbool
	previous *colonyGroupbool

	l *sync.RWMutex
}

func (g *colonyGroupbool) Insert(t *bool) (tp *bool) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newboolGroup(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupbool) Delete(tp *bool) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// ByteColony represents a colony of Bytes.
type ByteColony struct {
	entry *colonyGroupbyte
}

// NewByteColony returns a new colony of Bytes.
func NewByteColony(size uint) *ByteColony {
	return &ByteColony{
		entry: newbyteGroup(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *ByteColony) Insert(t *byte) (tp *byte) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *ByteColony) Delete(tp *byte) {
	c.entry.Delete(tp)
}

func newbyteGroup(previous *colonyGroupbyte, size uint) *colonyGroupbyte {
	var g colonyGroupbyte
	if size == 0 {
		size = 8
	}
	g.data = make([]byte, size)
	g.free = make(chan *byte, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupbyte struct {
	data     []byte
	free     chan *byte
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupbyte
	previous *colonyGroupbyte

	l *sync.RWMutex
}

func (g *colonyGroupbyte) Insert(t *byte) (tp *byte) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newbyteGroup(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupbyte) Delete(tp *byte) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Complex128Colony represents a colony of Complex128s.
type Complex128Colony struct {
	entry *colonyGroupcomplex128
}

// NewComplex128Colony returns a new colony of Complex128s.
func NewComplex128Colony(size uint) *Complex128Colony {
	return &Complex128Colony{
		entry: newcomplex128Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Complex128Colony) Insert(t *complex128) (tp *complex128) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Complex128Colony) Delete(tp *complex128) {
	c.entry.Delete(tp)
}

func newcomplex128Group(previous *colonyGroupcomplex128, size uint) *colonyGroupcomplex128 {
	var g colonyGroupcomplex128
	if size == 0 {
		size = 8
	}
	g.data = make([]complex128, size)
	g.free = make(chan *complex128, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupcomplex128 struct {
	data     []complex128
	free     chan *complex128
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupcomplex128
	previous *colonyGroupcomplex128

	l *sync.RWMutex
}

func (g *colonyGroupcomplex128) Insert(t *complex128) (tp *complex128) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newcomplex128Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupcomplex128) Delete(tp *complex128) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Complex64Colony represents a colony of Complex64s.
type Complex64Colony struct {
	entry *colonyGroupcomplex64
}

// NewComplex64Colony returns a new colony of Complex64s.
func NewComplex64Colony(size uint) *Complex64Colony {
	return &Complex64Colony{
		entry: newcomplex64Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Complex64Colony) Insert(t *complex64) (tp *complex64) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Complex64Colony) Delete(tp *complex64) {
	c.entry.Delete(tp)
}

func newcomplex64Group(previous *colonyGroupcomplex64, size uint) *colonyGroupcomplex64 {
	var g colonyGroupcomplex64
	if size == 0 {
		size = 8
	}
	g.data = make([]complex64, size)
	g.free = make(chan *complex64, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupcomplex64 struct {
	data     []complex64
	free     chan *complex64
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupcomplex64
	previous *colonyGroupcomplex64

	l *sync.RWMutex
}

func (g *colonyGroupcomplex64) Insert(t *complex64) (tp *complex64) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newcomplex64Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupcomplex64) Delete(tp *complex64) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// ErrorColony represents a colony of Errors.
type ErrorColony struct {
	entry *colonyGrouperror
}

// NewErrorColony returns a new colony of Errors.
func NewErrorColony(size uint) *ErrorColony {
	return &ErrorColony{
		entry: newerrorGroup(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *ErrorColony) Insert(t *error) (tp *error) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *ErrorColony) Delete(tp *error) {
	c.entry.Delete(tp)
}

func newerrorGroup(previous *colonyGrouperror, size uint) *colonyGrouperror {
	var g colonyGrouperror
	if size == 0 {
		size = 8
	}
	g.data = make([]error, size)
	g.free = make(chan *error, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGrouperror struct {
	data     []error
	free     chan *error
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGrouperror
	previous *colonyGrouperror

	l *sync.RWMutex
}

func (g *colonyGrouperror) Insert(t *error) (tp *error) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newerrorGroup(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGrouperror) Delete(tp *error) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Float32Colony represents a colony of Float32s.
type Float32Colony struct {
	entry *colonyGroupfloat32
}

// NewFloat32Colony returns a new colony of Float32s.
func NewFloat32Colony(size uint) *Float32Colony {
	return &Float32Colony{
		entry: newfloat32Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Float32Colony) Insert(t *float32) (tp *float32) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Float32Colony) Delete(tp *float32) {
	c.entry.Delete(tp)
}

func newfloat32Group(previous *colonyGroupfloat32, size uint) *colonyGroupfloat32 {
	var g colonyGroupfloat32
	if size == 0 {
		size = 8
	}
	g.data = make([]float32, size)
	g.free = make(chan *float32, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupfloat32 struct {
	data     []float32
	free     chan *float32
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupfloat32
	previous *colonyGroupfloat32

	l *sync.RWMutex
}

func (g *colonyGroupfloat32) Insert(t *float32) (tp *float32) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newfloat32Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupfloat32) Delete(tp *float32) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Float64Colony represents a colony of Float64s.
type Float64Colony struct {
	entry *colonyGroupfloat64
}

// NewFloat64Colony returns a new colony of Float64s.
func NewFloat64Colony(size uint) *Float64Colony {
	return &Float64Colony{
		entry: newfloat64Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Float64Colony) Insert(t *float64) (tp *float64) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Float64Colony) Delete(tp *float64) {
	c.entry.Delete(tp)
}

func newfloat64Group(previous *colonyGroupfloat64, size uint) *colonyGroupfloat64 {
	var g colonyGroupfloat64
	if size == 0 {
		size = 8
	}
	g.data = make([]float64, size)
	g.free = make(chan *float64, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupfloat64 struct {
	data     []float64
	free     chan *float64
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupfloat64
	previous *colonyGroupfloat64

	l *sync.RWMutex
}

func (g *colonyGroupfloat64) Insert(t *float64) (tp *float64) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newfloat64Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupfloat64) Delete(tp *float64) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// IntColony represents a colony of Ints.
type IntColony struct {
	entry *colonyGroupint
}

// NewIntColony returns a new colony of Ints.
func NewIntColony(size uint) *IntColony {
	return &IntColony{
		entry: newintGroup(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *IntColony) Insert(t *int) (tp *int) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *IntColony) Delete(tp *int) {
	c.entry.Delete(tp)
}

func newintGroup(previous *colonyGroupint, size uint) *colonyGroupint {
	var g colonyGroupint
	if size == 0 {
		size = 8
	}
	g.data = make([]int, size)
	g.free = make(chan *int, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupint struct {
	data     []int
	free     chan *int
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupint
	previous *colonyGroupint

	l *sync.RWMutex
}

func (g *colonyGroupint) Insert(t *int) (tp *int) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newintGroup(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupint) Delete(tp *int) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Int16Colony represents a colony of Int16s.
type Int16Colony struct {
	entry *colonyGroupint16
}

// NewInt16Colony returns a new colony of Int16s.
func NewInt16Colony(size uint) *Int16Colony {
	return &Int16Colony{
		entry: newint16Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Int16Colony) Insert(t *int16) (tp *int16) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Int16Colony) Delete(tp *int16) {
	c.entry.Delete(tp)
}

func newint16Group(previous *colonyGroupint16, size uint) *colonyGroupint16 {
	var g colonyGroupint16
	if size == 0 {
		size = 8
	}
	g.data = make([]int16, size)
	g.free = make(chan *int16, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupint16 struct {
	data     []int16
	free     chan *int16
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupint16
	previous *colonyGroupint16

	l *sync.RWMutex
}

func (g *colonyGroupint16) Insert(t *int16) (tp *int16) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newint16Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupint16) Delete(tp *int16) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Int32Colony represents a colony of Int32s.
type Int32Colony struct {
	entry *colonyGroupint32
}

// NewInt32Colony returns a new colony of Int32s.
func NewInt32Colony(size uint) *Int32Colony {
	return &Int32Colony{
		entry: newint32Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Int32Colony) Insert(t *int32) (tp *int32) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Int32Colony) Delete(tp *int32) {
	c.entry.Delete(tp)
}

func newint32Group(previous *colonyGroupint32, size uint) *colonyGroupint32 {
	var g colonyGroupint32
	if size == 0 {
		size = 8
	}
	g.data = make([]int32, size)
	g.free = make(chan *int32, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupint32 struct {
	data     []int32
	free     chan *int32
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupint32
	previous *colonyGroupint32

	l *sync.RWMutex
}

func (g *colonyGroupint32) Insert(t *int32) (tp *int32) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newint32Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupint32) Delete(tp *int32) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Int64Colony represents a colony of Int64s.
type Int64Colony struct {
	entry *colonyGroupint64
}

// NewInt64Colony returns a new colony of Int64s.
func NewInt64Colony(size uint) *Int64Colony {
	return &Int64Colony{
		entry: newint64Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Int64Colony) Insert(t *int64) (tp *int64) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Int64Colony) Delete(tp *int64) {
	c.entry.Delete(tp)
}

func newint64Group(previous *colonyGroupint64, size uint) *colonyGroupint64 {
	var g colonyGroupint64
	if size == 0 {
		size = 8
	}
	g.data = make([]int64, size)
	g.free = make(chan *int64, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupint64 struct {
	data     []int64
	free     chan *int64
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupint64
	previous *colonyGroupint64

	l *sync.RWMutex
}

func (g *colonyGroupint64) Insert(t *int64) (tp *int64) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newint64Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupint64) Delete(tp *int64) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Int8Colony represents a colony of Int8s.
type Int8Colony struct {
	entry *colonyGroupint8
}

// NewInt8Colony returns a new colony of Int8s.
func NewInt8Colony(size uint) *Int8Colony {
	return &Int8Colony{
		entry: newint8Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Int8Colony) Insert(t *int8) (tp *int8) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Int8Colony) Delete(tp *int8) {
	c.entry.Delete(tp)
}

func newint8Group(previous *colonyGroupint8, size uint) *colonyGroupint8 {
	var g colonyGroupint8
	if size == 0 {
		size = 8
	}
	g.data = make([]int8, size)
	g.free = make(chan *int8, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupint8 struct {
	data     []int8
	free     chan *int8
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupint8
	previous *colonyGroupint8

	l *sync.RWMutex
}

func (g *colonyGroupint8) Insert(t *int8) (tp *int8) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newint8Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupint8) Delete(tp *int8) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// RuneColony represents a colony of Runes.
type RuneColony struct {
	entry *colonyGrouprune
}

// NewRuneColony returns a new colony of Runes.
func NewRuneColony(size uint) *RuneColony {
	return &RuneColony{
		entry: newruneGroup(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *RuneColony) Insert(t *rune) (tp *rune) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *RuneColony) Delete(tp *rune) {
	c.entry.Delete(tp)
}

func newruneGroup(previous *colonyGrouprune, size uint) *colonyGrouprune {
	var g colonyGrouprune
	if size == 0 {
		size = 8
	}
	g.data = make([]rune, size)
	g.free = make(chan *rune, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGrouprune struct {
	data     []rune
	free     chan *rune
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGrouprune
	previous *colonyGrouprune

	l *sync.RWMutex
}

func (g *colonyGrouprune) Insert(t *rune) (tp *rune) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newruneGroup(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGrouprune) Delete(tp *rune) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// StringColony represents a colony of Strings.
type StringColony struct {
	entry *colonyGroupstring
}

// NewStringColony returns a new colony of Strings.
func NewStringColony(size uint) *StringColony {
	return &StringColony{
		entry: newstringGroup(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *StringColony) Insert(t *string) (tp *string) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *StringColony) Delete(tp *string) {
	c.entry.Delete(tp)
}

func newstringGroup(previous *colonyGroupstring, size uint) *colonyGroupstring {
	var g colonyGroupstring
	if size == 0 {
		size = 8
	}
	g.data = make([]string, size)
	g.free = make(chan *string, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupstring struct {
	data     []string
	free     chan *string
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupstring
	previous *colonyGroupstring

	l *sync.RWMutex
}

func (g *colonyGroupstring) Insert(t *string) (tp *string) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newstringGroup(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupstring) Delete(tp *string) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// UintColony represents a colony of Uints.
type UintColony struct {
	entry *colonyGroupuint
}

// NewUintColony returns a new colony of Uints.
func NewUintColony(size uint) *UintColony {
	return &UintColony{
		entry: newuintGroup(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *UintColony) Insert(t *uint) (tp *uint) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *UintColony) Delete(tp *uint) {
	c.entry.Delete(tp)
}

func newuintGroup(previous *colonyGroupuint, size uint) *colonyGroupuint {
	var g colonyGroupuint
	if size == 0 {
		size = 8
	}
	g.data = make([]uint, size)
	g.free = make(chan *uint, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupuint struct {
	data     []uint
	free     chan *uint
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupuint
	previous *colonyGroupuint

	l *sync.RWMutex
}

func (g *colonyGroupuint) Insert(t *uint) (tp *uint) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newuintGroup(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupuint) Delete(tp *uint) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Uint16Colony represents a colony of Uint16s.
type Uint16Colony struct {
	entry *colonyGroupuint16
}

// NewUint16Colony returns a new colony of Uint16s.
func NewUint16Colony(size uint) *Uint16Colony {
	return &Uint16Colony{
		entry: newuint16Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Uint16Colony) Insert(t *uint16) (tp *uint16) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Uint16Colony) Delete(tp *uint16) {
	c.entry.Delete(tp)
}

func newuint16Group(previous *colonyGroupuint16, size uint) *colonyGroupuint16 {
	var g colonyGroupuint16
	if size == 0 {
		size = 8
	}
	g.data = make([]uint16, size)
	g.free = make(chan *uint16, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupuint16 struct {
	data     []uint16
	free     chan *uint16
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupuint16
	previous *colonyGroupuint16

	l *sync.RWMutex
}

func (g *colonyGroupuint16) Insert(t *uint16) (tp *uint16) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newuint16Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupuint16) Delete(tp *uint16) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Uint32Colony represents a colony of Uint32s.
type Uint32Colony struct {
	entry *colonyGroupuint32
}

// NewUint32Colony returns a new colony of Uint32s.
func NewUint32Colony(size uint) *Uint32Colony {
	return &Uint32Colony{
		entry: newuint32Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Uint32Colony) Insert(t *uint32) (tp *uint32) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Uint32Colony) Delete(tp *uint32) {
	c.entry.Delete(tp)
}

func newuint32Group(previous *colonyGroupuint32, size uint) *colonyGroupuint32 {
	var g colonyGroupuint32
	if size == 0 {
		size = 8
	}
	g.data = make([]uint32, size)
	g.free = make(chan *uint32, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupuint32 struct {
	data     []uint32
	free     chan *uint32
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupuint32
	previous *colonyGroupuint32

	l *sync.RWMutex
}

func (g *colonyGroupuint32) Insert(t *uint32) (tp *uint32) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newuint32Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupuint32) Delete(tp *uint32) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Uint64Colony represents a colony of Uint64s.
type Uint64Colony struct {
	entry *colonyGroupuint64
}

// NewUint64Colony returns a new colony of Uint64s.
func NewUint64Colony(size uint) *Uint64Colony {
	return &Uint64Colony{
		entry: newuint64Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Uint64Colony) Insert(t *uint64) (tp *uint64) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Uint64Colony) Delete(tp *uint64) {
	c.entry.Delete(tp)
}

func newuint64Group(previous *colonyGroupuint64, size uint) *colonyGroupuint64 {
	var g colonyGroupuint64
	if size == 0 {
		size = 8
	}
	g.data = make([]uint64, size)
	g.free = make(chan *uint64, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupuint64 struct {
	data     []uint64
	free     chan *uint64
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupuint64
	previous *colonyGroupuint64

	l *sync.RWMutex
}

func (g *colonyGroupuint64) Insert(t *uint64) (tp *uint64) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newuint64Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupuint64) Delete(tp *uint64) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// Uint8Colony represents a colony of Uint8s.
type Uint8Colony struct {
	entry *colonyGroupuint8
}

// NewUint8Colony returns a new colony of Uint8s.
func NewUint8Colony(size uint) *Uint8Colony {
	return &Uint8Colony{
		entry: newuint8Group(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *Uint8Colony) Insert(t *uint8) (tp *uint8) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *Uint8Colony) Delete(tp *uint8) {
	c.entry.Delete(tp)
}

func newuint8Group(previous *colonyGroupuint8, size uint) *colonyGroupuint8 {
	var g colonyGroupuint8
	if size == 0 {
		size = 8
	}
	g.data = make([]uint8, size)
	g.free = make(chan *uint8, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupuint8 struct {
	data     []uint8
	free     chan *uint8
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupuint8
	previous *colonyGroupuint8

	l *sync.RWMutex
}

func (g *colonyGroupuint8) Insert(t *uint8) (tp *uint8) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newuint8Group(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupuint8) Delete(tp *uint8) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}

// UintptrColony represents a colony of Uintptrs.
type UintptrColony struct {
	entry *colonyGroupuintptr
}

// NewUintptrColony returns a new colony of Uintptrs.
func NewUintptrColony(size uint) *UintptrColony {
	return &UintptrColony{
		entry: newuintptrGroup(nil, size),
	}
}

// Insert returns a pointer from the colony and initializes it with the provided data.
func (c *UintptrColony) Insert(t *uintptr) (tp *uintptr) {
	return c.entry.Insert(t)
}

// Delete returns a pointer to the colony.
func (c *UintptrColony) Delete(tp *uintptr) {
	c.entry.Delete(tp)
}

func newuintptrGroup(previous *colonyGroupuintptr, size uint) *colonyGroupuintptr {
	var g colonyGroupuintptr
	if size == 0 {
		size = 8
	}
	g.data = make([]uintptr, size)
	g.free = make(chan *uintptr, size)
	for i := 0; i < len(g.data); i++ {
		g.free <- &g.data[i]
	}
	g.next = nil
	g.l = &sync.RWMutex{}
	g.previous = previous
	g.minPtr = uintptr(unsafe.Pointer(&g.data[0]))
	g.maxPtr = uintptr(unsafe.Pointer(&g.data[len(g.data)-1]))
	return &g
}

type colonyGroupuintptr struct {
	data     []uintptr
	free     chan *uintptr
	maxPtr   uintptr
	minPtr   uintptr
	next     *colonyGroupuintptr
	previous *colonyGroupuintptr

	l *sync.RWMutex
}

func (g *colonyGroupuintptr) Insert(t *uintptr) (tp *uintptr) {
	select {
	case tp = <-g.free:
		return
	default:
		if g.next == nil {
			g.next = newuintptrGroup(g, uint(len(g.data)*2))
		}
		return g.next.Insert(t)
	}
}

func (g *colonyGroupuintptr) Delete(tp *uintptr) {
	tpu := uintptr(unsafe.Pointer(tp))
	if tpu < g.minPtr || tpu > g.maxPtr { // hack to determine if a pointer points to this array
		g.next.Delete(tp)
	}
	g.free <- tp
	//if !g.index.Any() {
	// TODO: if a group has no more elements, then we should de-allocate it.
	//}
	return
}
