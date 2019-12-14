// Code generated by go run rng_gen.go; DO NOT EDIT.
package rng

type Int struct {
	start, end     int
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewInt(start, end int, startex, endex bool) *Int {
	return &Int{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Int) Start() int    { return r.start }
func (r *Int) End() int      { return r.end }
func (r *Int) Startex() bool { return r.startex }
func (r *Int) Endex() bool   { return r.endex }

func (r *Int) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Int) In(x int) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func IntInter(x, y *Int) *Int {
	r := &Int{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func IntUnion(x, y *Int) *Int {
	r := &Int{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}

type Int8 struct {
	start, end     int8
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewInt8(start, end int8, startex, endex bool) *Int8 {
	return &Int8{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Int8) Start() int8   { return r.start }
func (r *Int8) End() int8     { return r.end }
func (r *Int8) Startex() bool { return r.startex }
func (r *Int8) Endex() bool   { return r.endex }

func (r *Int8) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Int8) In(x int8) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func Int8Inter(x, y *Int8) *Int8 {
	r := &Int8{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func Int8Union(x, y *Int8) *Int8 {
	r := &Int8{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}

type Int16 struct {
	start, end     int16
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewInt16(start, end int16, startex, endex bool) *Int16 {
	return &Int16{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Int16) Start() int16  { return r.start }
func (r *Int16) End() int16    { return r.end }
func (r *Int16) Startex() bool { return r.startex }
func (r *Int16) Endex() bool   { return r.endex }

func (r *Int16) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Int16) In(x int16) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func Int16Inter(x, y *Int16) *Int16 {
	r := &Int16{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func Int16Union(x, y *Int16) *Int16 {
	r := &Int16{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}

type Int32 struct {
	start, end     int32
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewInt32(start, end int32, startex, endex bool) *Int32 {
	return &Int32{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Int32) Start() int32  { return r.start }
func (r *Int32) End() int32    { return r.end }
func (r *Int32) Startex() bool { return r.startex }
func (r *Int32) Endex() bool   { return r.endex }

func (r *Int32) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Int32) In(x int32) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func Int32Inter(x, y *Int32) *Int32 {
	r := &Int32{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func Int32Union(x, y *Int32) *Int32 {
	r := &Int32{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}

type Int64 struct {
	start, end     int64
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewInt64(start, end int64, startex, endex bool) *Int64 {
	return &Int64{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Int64) Start() int64  { return r.start }
func (r *Int64) End() int64    { return r.end }
func (r *Int64) Startex() bool { return r.startex }
func (r *Int64) Endex() bool   { return r.endex }

func (r *Int64) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Int64) In(x int64) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func Int64Inter(x, y *Int64) *Int64 {
	r := &Int64{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func Int64Union(x, y *Int64) *Int64 {
	r := &Int64{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}

type Uint struct {
	start, end     uint
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewUint(start, end uint, startex, endex bool) *Uint {
	return &Uint{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Uint) Start() uint   { return r.start }
func (r *Uint) End() uint     { return r.end }
func (r *Uint) Startex() bool { return r.startex }
func (r *Uint) Endex() bool   { return r.endex }

func (r *Uint) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Uint) In(x uint) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func UintInter(x, y *Uint) *Uint {
	r := &Uint{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func UintUnion(x, y *Uint) *Uint {
	r := &Uint{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}

type Uint8 struct {
	start, end     uint8
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewUint8(start, end uint8, startex, endex bool) *Uint8 {
	return &Uint8{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Uint8) Start() uint8  { return r.start }
func (r *Uint8) End() uint8    { return r.end }
func (r *Uint8) Startex() bool { return r.startex }
func (r *Uint8) Endex() bool   { return r.endex }

func (r *Uint8) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Uint8) In(x uint8) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func Uint8Inter(x, y *Uint8) *Uint8 {
	r := &Uint8{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func Uint8Union(x, y *Uint8) *Uint8 {
	r := &Uint8{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}

type Uint16 struct {
	start, end     uint16
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewUint16(start, end uint16, startex, endex bool) *Uint16 {
	return &Uint16{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Uint16) Start() uint16 { return r.start }
func (r *Uint16) End() uint16   { return r.end }
func (r *Uint16) Startex() bool { return r.startex }
func (r *Uint16) Endex() bool   { return r.endex }

func (r *Uint16) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Uint16) In(x uint16) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func Uint16Inter(x, y *Uint16) *Uint16 {
	r := &Uint16{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func Uint16Union(x, y *Uint16) *Uint16 {
	r := &Uint16{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}

type Uint32 struct {
	start, end     uint32
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewUint32(start, end uint32, startex, endex bool) *Uint32 {
	return &Uint32{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Uint32) Start() uint32 { return r.start }
func (r *Uint32) End() uint32   { return r.end }
func (r *Uint32) Startex() bool { return r.startex }
func (r *Uint32) Endex() bool   { return r.endex }

func (r *Uint32) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Uint32) In(x uint32) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func Uint32Inter(x, y *Uint32) *Uint32 {
	r := &Uint32{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func Uint32Union(x, y *Uint32) *Uint32 {
	r := &Uint32{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}

type Uint64 struct {
	start, end     uint64
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewUint64(start, end uint64, startex, endex bool) *Uint64 {
	return &Uint64{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Uint64) Start() uint64 { return r.start }
func (r *Uint64) End() uint64   { return r.end }
func (r *Uint64) Startex() bool { return r.startex }
func (r *Uint64) Endex() bool   { return r.endex }

func (r *Uint64) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Uint64) In(x uint64) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func Uint64Inter(x, y *Uint64) *Uint64 {
	r := &Uint64{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func Uint64Union(x, y *Uint64) *Uint64 {
	r := &Uint64{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}

type Float32 struct {
	start, end     float32
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewFloat32(start, end float32, startex, endex bool) *Float32 {
	return &Float32{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Float32) Start() float32 { return r.start }
func (r *Float32) End() float32   { return r.end }
func (r *Float32) Startex() bool  { return r.startex }
func (r *Float32) Endex() bool    { return r.endex }

func (r *Float32) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Float32) In(x float32) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func Float32Inter(x, y *Float32) *Float32 {
	r := &Float32{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func Float32Union(x, y *Float32) *Float32 {
	r := &Float32{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}

type Float64 struct {
	start, end     float64
	startex, endex bool // exclude start or end

	// If empty is true, the range must be empty. If not, we need to check it.
	empty bool
}

func NewFloat64(start, end float64, startex, endex bool) *Float64 {
	return &Float64{
		start:   start,
		end:     end,
		startex: startex,
		endex:   endex,
	}
}

func (r *Float64) Start() float64 { return r.start }
func (r *Float64) End() float64   { return r.end }
func (r *Float64) Startex() bool  { return r.startex }
func (r *Float64) Endex() bool    { return r.endex }

func (r *Float64) IsEmpty() bool {
	if r.empty {
		return true
	}
	if r.start > r.end || r.start == r.end && (r.startex || r.endex) {
		r.empty = true
		return true
	}
	return false
}

func (r *Float64) In(x float64) bool {
	if r.IsEmpty() {
		return false
	}
	if r.start > x || r.startex && r.start == x {
		return false
	}
	if r.end < x || r.endex && r.end == x {
		return false
	}
	return true
}

func Float64Inter(x, y *Float64) *Float64 {
	r := &Float64{}

	if x.IsEmpty() || y.IsEmpty() || x.end < y.start || x.start > y.end {
		r.empty = true
		return r
	}

	if x.start < y.start {
		r.start = y.start
		r.startex = y.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.start = x.start
		r.startex = x.startex
	}

	if x.end < y.end {
		r.end = x.end
		r.endex = x.endex
	} else if x.end == y.end {
		r.start = x.start
		r.startex = x.startex || y.startex
	} else {
		r.end = y.end
		r.endex = y.endex
	}

	return r
}

func Float64Union(x, y *Float64) *Float64 {
	r := &Float64{}

	if x.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = x.start, x.startex, x.end, x.endex, x.empty
		return r
	}

	if y.IsEmpty() {
		r.start, r.startex, r.end, r.endex, r.empty = y.start, y.startex, y.end, y.endex, y.empty
		return r
	}

	if x.start < y.start {
		r.start = x.start
		r.startex = x.startex
	} else if x.start == y.start {
		r.start = x.start
		r.startex = x.startex && y.startex
	} else {
		r.start = y.start
		r.startex = y.startex
	}

	if x.end < y.end {
		r.end = y.end
		r.endex = y.endex
	} else if x.end == y.end {
		r.end = x.end
		r.endex = x.endex && y.endex
	} else {
		r.end = x.end
		r.endex = x.endex
	}

	return r
}
