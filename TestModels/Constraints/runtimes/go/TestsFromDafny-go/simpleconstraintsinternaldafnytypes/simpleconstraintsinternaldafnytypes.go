// Package simpleconstraintsinternaldafnytypes
// Dafny module simpleconstraintsinternaldafnytypes compiled into Go

package simpleconstraintsinternaldafnytypes

import (
	os "os"

	_System "github.com/dafny-lang/DafnyRuntimeGo/System_"
	_dafny "github.com/dafny-lang/DafnyRuntimeGo/dafny"
	StandardLibrary "github.com/dafny-lang/DafnyStandardLibGo/StandardLibrary"
	StandardLibraryInterop "github.com/dafny-lang/DafnyStandardLibGo/StandardLibraryInterop"
	StandardLibrary_UInt "github.com/dafny-lang/DafnyStandardLibGo/StandardLibrary_UInt"
	UTF8 "github.com/dafny-lang/DafnyStandardLibGo/UTF8"
	Wrappers "github.com/dafny-lang/DafnyStandardLibGo/Wrappers"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ _System.Dummy__
var _ Wrappers.Dummy__
var _ StandardLibrary_UInt.Dummy__
var _ StandardLibrary.Dummy__
var _ StandardLibraryInterop.Dummy__

type Dummy__ struct{}

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "simpleconstraintsinternaldafnytypes.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) IsValid__BlobLessThanOrEqualToTen(x _dafny.Sequence) bool {
	return (_dafny.IntOfUint32((x).Cardinality())).Cmp(_dafny.IntOfInt64(10)) <= 0
}
func (_static *CompanionStruct_Default___) IsValid__GreaterThanOne(x int32) bool {
	return (int32(1)) <= (x)
}
func (_static *CompanionStruct_Default___) IsValid__LessThanTen(x int32) bool {
	return (x) <= (int32(10))
}
func (_static *CompanionStruct_Default___) IsValid__ListLessThanOrEqualToTen(x _dafny.Sequence) bool {
	return (_dafny.IntOfUint32((x).Cardinality())).Cmp(_dafny.IntOfInt64(10)) <= 0
}
func (_static *CompanionStruct_Default___) IsValid__ListOfUtf8Bytes(x _dafny.Sequence) bool {
	return ((_dafny.One).Cmp(_dafny.IntOfUint32((x).Cardinality())) <= 0) && ((_dafny.IntOfUint32((x).Cardinality())).Cmp(_dafny.IntOfInt64(2)) <= 0)
}
func (_static *CompanionStruct_Default___) IsValid__MapLessThanOrEqualToTen(x _dafny.Map) bool {
	return ((x).Cardinality()).Cmp(_dafny.IntOfInt64(10)) <= 0
}
func (_static *CompanionStruct_Default___) IsValid__MyBlob(x _dafny.Sequence) bool {
	return ((_dafny.One).Cmp(_dafny.IntOfUint32((x).Cardinality())) <= 0) && ((_dafny.IntOfUint32((x).Cardinality())).Cmp(_dafny.IntOfInt64(10)) <= 0)
}
func (_static *CompanionStruct_Default___) IsValid__MyList(x _dafny.Sequence) bool {
	return ((_dafny.One).Cmp(_dafny.IntOfUint32((x).Cardinality())) <= 0) && ((_dafny.IntOfUint32((x).Cardinality())).Cmp(_dafny.IntOfInt64(10)) <= 0)
}
func (_static *CompanionStruct_Default___) IsValid__MyMap(x _dafny.Map) bool {
	return ((_dafny.One).Cmp((x).Cardinality()) <= 0) && (((x).Cardinality()).Cmp(_dafny.IntOfInt64(10)) <= 0)
}
func (_static *CompanionStruct_Default___) IsValid__MyString(x _dafny.Sequence) bool {
	return ((_dafny.One).Cmp(_dafny.IntOfUint32((x).Cardinality())) <= 0) && ((_dafny.IntOfUint32((x).Cardinality())).Cmp(_dafny.IntOfInt64(10)) <= 0)
}
func (_static *CompanionStruct_Default___) IsValid__NonEmptyBlob(x _dafny.Sequence) bool {
	return (_dafny.One).Cmp(_dafny.IntOfUint32((x).Cardinality())) <= 0
}
func (_static *CompanionStruct_Default___) IsValid__NonEmptyList(x _dafny.Sequence) bool {
	return (_dafny.One).Cmp(_dafny.IntOfUint32((x).Cardinality())) <= 0
}
func (_static *CompanionStruct_Default___) IsValid__NonEmptyMap(x _dafny.Map) bool {
	return (_dafny.One).Cmp((x).Cardinality()) <= 0
}
func (_static *CompanionStruct_Default___) IsValid__NonEmptyString(x _dafny.Sequence) bool {
	return (_dafny.One).Cmp(_dafny.IntOfUint32((x).Cardinality())) <= 0
}
func (_static *CompanionStruct_Default___) IsValid__OneToTen(x int32) bool {
	return ((int32(1)) <= (x)) && ((x) <= (int32(10)))
}
func (_static *CompanionStruct_Default___) IsValid__StringLessThanOrEqualToTen(x _dafny.Sequence) bool {
	return (_dafny.IntOfUint32((x).Cardinality())).Cmp(_dafny.IntOfInt64(10)) <= 0
}
func (_static *CompanionStruct_Default___) IsValid__TenToTen(x int64) bool {
	return ((int64(-10)) <= (x)) && ((x) <= (int64(10)))
}
func (_static *CompanionStruct_Default___) IsValid__Utf8Bytes(x _dafny.Sequence) bool {
	return ((_dafny.One).Cmp(_dafny.IntOfUint32((x).Cardinality())) <= 0) && ((_dafny.IntOfUint32((x).Cardinality())).Cmp(_dafny.IntOfInt64(10)) <= 0)
}

// End of class Default__

// Definition of datatype DafnyCallEvent
type DafnyCallEvent struct {
	Data_DafnyCallEvent_
}

func (_this DafnyCallEvent) Get_() Data_DafnyCallEvent_ {
	return _this.Data_DafnyCallEvent_
}

type Data_DafnyCallEvent_ interface {
	isDafnyCallEvent()
}

type CompanionStruct_DafnyCallEvent_ struct {
}

var Companion_DafnyCallEvent_ = CompanionStruct_DafnyCallEvent_{}

type DafnyCallEvent_DafnyCallEvent struct {
	Input  interface{}
	Output interface{}
}

func (DafnyCallEvent_DafnyCallEvent) isDafnyCallEvent() {}

func (CompanionStruct_DafnyCallEvent_) Create_DafnyCallEvent_(Input interface{}, Output interface{}) DafnyCallEvent {
	return DafnyCallEvent{DafnyCallEvent_DafnyCallEvent{Input, Output}}
}

func (_this DafnyCallEvent) Is_DafnyCallEvent() bool {
	_, ok := _this.Get_().(DafnyCallEvent_DafnyCallEvent)
	return ok
}

func (CompanionStruct_DafnyCallEvent_) Default(_default_I interface{}, _default_O interface{}) DafnyCallEvent {
	return Companion_DafnyCallEvent_.Create_DafnyCallEvent_(_default_I, _default_O)
}

func (_this DafnyCallEvent) Dtor_input() interface{} {
	return _this.Get_().(DafnyCallEvent_DafnyCallEvent).Input
}

func (_this DafnyCallEvent) Dtor_output() interface{} {
	return _this.Get_().(DafnyCallEvent_DafnyCallEvent).Output
}

func (_this DafnyCallEvent) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case DafnyCallEvent_DafnyCallEvent:
		{
			return "SimpleConstraintsTypes.DafnyCallEvent.DafnyCallEvent" + "(" + _dafny.String(data.Input) + ", " + _dafny.String(data.Output) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this DafnyCallEvent) Equals(other DafnyCallEvent) bool {
	switch data1 := _this.Get_().(type) {
	case DafnyCallEvent_DafnyCallEvent:
		{
			data2, ok := other.Get_().(DafnyCallEvent_DafnyCallEvent)
			return ok && _dafny.AreEqual(data1.Input, data2.Input) && _dafny.AreEqual(data1.Output, data2.Output)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this DafnyCallEvent) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(DafnyCallEvent)
	return ok && _this.Equals(typed)
}

func Type_DafnyCallEvent_(Type_I_ _dafny.TypeDescriptor, Type_O_ _dafny.TypeDescriptor) _dafny.TypeDescriptor {
	return type_DafnyCallEvent_{Type_I_, Type_O_}
}

type type_DafnyCallEvent_ struct {
	Type_I_ _dafny.TypeDescriptor
	Type_O_ _dafny.TypeDescriptor
}

func (_this type_DafnyCallEvent_) Default() interface{} {
	Type_I_ := _this.Type_I_
	_ = Type_I_
	Type_O_ := _this.Type_O_
	_ = Type_O_
	return Companion_DafnyCallEvent_.Default(Type_I_.Default(), Type_O_.Default())
}

func (_this type_DafnyCallEvent_) String() string {
	return "simpleconstraintsinternaldafnytypes.DafnyCallEvent"
}
func (_this DafnyCallEvent) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = DafnyCallEvent{}

// End of datatype DafnyCallEvent

// Definition of class BlobLessThanOrEqualToTen
type BlobLessThanOrEqualToTen struct {
}

func New_BlobLessThanOrEqualToTen_() *BlobLessThanOrEqualToTen {
	_this := BlobLessThanOrEqualToTen{}

	return &_this
}

type CompanionStruct_BlobLessThanOrEqualToTen_ struct {
}

var Companion_BlobLessThanOrEqualToTen_ = CompanionStruct_BlobLessThanOrEqualToTen_{}

func (*BlobLessThanOrEqualToTen) String() string {
	return "simpleconstraintsinternaldafnytypes.BlobLessThanOrEqualToTen"
}

// End of class BlobLessThanOrEqualToTen

func Type_BlobLessThanOrEqualToTen_() _dafny.TypeDescriptor {
	return type_BlobLessThanOrEqualToTen_{}
}

type type_BlobLessThanOrEqualToTen_ struct {
}

func (_this type_BlobLessThanOrEqualToTen_) Default() interface{} {
	return _dafny.EmptySeq
}

func (_this type_BlobLessThanOrEqualToTen_) String() string {
	return "simpleconstraintsinternaldafnytypes.BlobLessThanOrEqualToTen"
}
func (_this *CompanionStruct_BlobLessThanOrEqualToTen_) Is_(__source _dafny.Sequence) bool {
	var _0_x _dafny.Sequence = (__source)
	_ = _0_x
	return Companion_Default___.IsValid__BlobLessThanOrEqualToTen(_0_x)
}

// Definition of datatype GetConstraintsInput
type GetConstraintsInput struct {
	Data_GetConstraintsInput_
}

func (_this GetConstraintsInput) Get_() Data_GetConstraintsInput_ {
	return _this.Data_GetConstraintsInput_
}

type Data_GetConstraintsInput_ interface {
	isGetConstraintsInput()
}

type CompanionStruct_GetConstraintsInput_ struct {
}

var Companion_GetConstraintsInput_ = CompanionStruct_GetConstraintsInput_{}

type GetConstraintsInput_GetConstraintsInput struct {
	MyString                   Wrappers.Option
	NonEmptyString             Wrappers.Option
	StringLessThanOrEqualToTen Wrappers.Option
	MyBlob                     Wrappers.Option
	NonEmptyBlob               Wrappers.Option
	BlobLessThanOrEqualToTen   Wrappers.Option
	MyList                     Wrappers.Option
	NonEmptyList               Wrappers.Option
	ListLessThanOrEqualToTen   Wrappers.Option
	MyMap                      Wrappers.Option
	NonEmptyMap                Wrappers.Option
	MapLessThanOrEqualToTen    Wrappers.Option
	OneToTen                   Wrappers.Option
	MyTenToTen                 Wrappers.Option
	GreaterThanOne             Wrappers.Option
	LessThanTen                Wrappers.Option
	MyUtf8Bytes                Wrappers.Option
	MyListOfUtf8Bytes          Wrappers.Option
}

func (GetConstraintsInput_GetConstraintsInput) isGetConstraintsInput() {}

func (CompanionStruct_GetConstraintsInput_) Create_GetConstraintsInput_(MyString Wrappers.Option, NonEmptyString Wrappers.Option, StringLessThanOrEqualToTen Wrappers.Option, MyBlob Wrappers.Option, NonEmptyBlob Wrappers.Option, BlobLessThanOrEqualToTen Wrappers.Option, MyList Wrappers.Option, NonEmptyList Wrappers.Option, ListLessThanOrEqualToTen Wrappers.Option, MyMap Wrappers.Option, NonEmptyMap Wrappers.Option, MapLessThanOrEqualToTen Wrappers.Option, OneToTen Wrappers.Option, MyTenToTen Wrappers.Option, GreaterThanOne Wrappers.Option, LessThanTen Wrappers.Option, MyUtf8Bytes Wrappers.Option, MyListOfUtf8Bytes Wrappers.Option) GetConstraintsInput {
	return GetConstraintsInput{GetConstraintsInput_GetConstraintsInput{MyString, NonEmptyString, StringLessThanOrEqualToTen, MyBlob, NonEmptyBlob, BlobLessThanOrEqualToTen, MyList, NonEmptyList, ListLessThanOrEqualToTen, MyMap, NonEmptyMap, MapLessThanOrEqualToTen, OneToTen, MyTenToTen, GreaterThanOne, LessThanTen, MyUtf8Bytes, MyListOfUtf8Bytes}}
}

func (_this GetConstraintsInput) Is_GetConstraintsInput() bool {
	_, ok := _this.Get_().(GetConstraintsInput_GetConstraintsInput)
	return ok
}

func (CompanionStruct_GetConstraintsInput_) Default() GetConstraintsInput {
	return Companion_GetConstraintsInput_.Create_GetConstraintsInput_(Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default())
}

func (_this GetConstraintsInput) Dtor_MyString() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).MyString
}

func (_this GetConstraintsInput) Dtor_NonEmptyString() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).NonEmptyString
}

func (_this GetConstraintsInput) Dtor_StringLessThanOrEqualToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).StringLessThanOrEqualToTen
}

func (_this GetConstraintsInput) Dtor_MyBlob() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).MyBlob
}

func (_this GetConstraintsInput) Dtor_NonEmptyBlob() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).NonEmptyBlob
}

func (_this GetConstraintsInput) Dtor_BlobLessThanOrEqualToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).BlobLessThanOrEqualToTen
}

func (_this GetConstraintsInput) Dtor_MyList() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).MyList
}

func (_this GetConstraintsInput) Dtor_NonEmptyList() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).NonEmptyList
}

func (_this GetConstraintsInput) Dtor_ListLessThanOrEqualToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).ListLessThanOrEqualToTen
}

func (_this GetConstraintsInput) Dtor_MyMap() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).MyMap
}

func (_this GetConstraintsInput) Dtor_NonEmptyMap() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).NonEmptyMap
}

func (_this GetConstraintsInput) Dtor_MapLessThanOrEqualToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).MapLessThanOrEqualToTen
}

func (_this GetConstraintsInput) Dtor_OneToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).OneToTen
}

func (_this GetConstraintsInput) Dtor_myTenToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).MyTenToTen
}

func (_this GetConstraintsInput) Dtor_GreaterThanOne() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).GreaterThanOne
}

func (_this GetConstraintsInput) Dtor_LessThanTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).LessThanTen
}

func (_this GetConstraintsInput) Dtor_MyUtf8Bytes() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).MyUtf8Bytes
}

func (_this GetConstraintsInput) Dtor_MyListOfUtf8Bytes() Wrappers.Option {
	return _this.Get_().(GetConstraintsInput_GetConstraintsInput).MyListOfUtf8Bytes
}

func (_this GetConstraintsInput) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case GetConstraintsInput_GetConstraintsInput:
		{
			return "SimpleConstraintsTypes.GetConstraintsInput.GetConstraintsInput" + "(" + _dafny.String(data.MyString) + ", " + _dafny.String(data.NonEmptyString) + ", " + _dafny.String(data.StringLessThanOrEqualToTen) + ", " + _dafny.String(data.MyBlob) + ", " + _dafny.String(data.NonEmptyBlob) + ", " + _dafny.String(data.BlobLessThanOrEqualToTen) + ", " + _dafny.String(data.MyList) + ", " + _dafny.String(data.NonEmptyList) + ", " + _dafny.String(data.ListLessThanOrEqualToTen) + ", " + _dafny.String(data.MyMap) + ", " + _dafny.String(data.NonEmptyMap) + ", " + _dafny.String(data.MapLessThanOrEqualToTen) + ", " + _dafny.String(data.OneToTen) + ", " + _dafny.String(data.MyTenToTen) + ", " + _dafny.String(data.GreaterThanOne) + ", " + _dafny.String(data.LessThanTen) + ", " + _dafny.String(data.MyUtf8Bytes) + ", " + _dafny.String(data.MyListOfUtf8Bytes) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GetConstraintsInput) Equals(other GetConstraintsInput) bool {
	switch data1 := _this.Get_().(type) {
	case GetConstraintsInput_GetConstraintsInput:
		{
			data2, ok := other.Get_().(GetConstraintsInput_GetConstraintsInput)
			return ok && data1.MyString.Equals(data2.MyString) && data1.NonEmptyString.Equals(data2.NonEmptyString) && data1.StringLessThanOrEqualToTen.Equals(data2.StringLessThanOrEqualToTen) && data1.MyBlob.Equals(data2.MyBlob) && data1.NonEmptyBlob.Equals(data2.NonEmptyBlob) && data1.BlobLessThanOrEqualToTen.Equals(data2.BlobLessThanOrEqualToTen) && data1.MyList.Equals(data2.MyList) && data1.NonEmptyList.Equals(data2.NonEmptyList) && data1.ListLessThanOrEqualToTen.Equals(data2.ListLessThanOrEqualToTen) && data1.MyMap.Equals(data2.MyMap) && data1.NonEmptyMap.Equals(data2.NonEmptyMap) && data1.MapLessThanOrEqualToTen.Equals(data2.MapLessThanOrEqualToTen) && data1.OneToTen.Equals(data2.OneToTen) && data1.MyTenToTen.Equals(data2.MyTenToTen) && data1.GreaterThanOne.Equals(data2.GreaterThanOne) && data1.LessThanTen.Equals(data2.LessThanTen) && data1.MyUtf8Bytes.Equals(data2.MyUtf8Bytes) && data1.MyListOfUtf8Bytes.Equals(data2.MyListOfUtf8Bytes)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GetConstraintsInput) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GetConstraintsInput)
	return ok && _this.Equals(typed)
}

func Type_GetConstraintsInput_() _dafny.TypeDescriptor {
	return type_GetConstraintsInput_{}
}

type type_GetConstraintsInput_ struct {
}

func (_this type_GetConstraintsInput_) Default() interface{} {
	return Companion_GetConstraintsInput_.Default()
}

func (_this type_GetConstraintsInput_) String() string {
	return "simpleconstraintsinternaldafnytypes.GetConstraintsInput"
}
func (_this GetConstraintsInput) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = GetConstraintsInput{}

// End of datatype GetConstraintsInput

// Definition of datatype GetConstraintsOutput
type GetConstraintsOutput struct {
	Data_GetConstraintsOutput_
}

func (_this GetConstraintsOutput) Get_() Data_GetConstraintsOutput_ {
	return _this.Data_GetConstraintsOutput_
}

type Data_GetConstraintsOutput_ interface {
	isGetConstraintsOutput()
}

type CompanionStruct_GetConstraintsOutput_ struct {
}

var Companion_GetConstraintsOutput_ = CompanionStruct_GetConstraintsOutput_{}

type GetConstraintsOutput_GetConstraintsOutput struct {
	MyString                   Wrappers.Option
	NonEmptyString             Wrappers.Option
	StringLessThanOrEqualToTen Wrappers.Option
	MyBlob                     Wrappers.Option
	NonEmptyBlob               Wrappers.Option
	BlobLessThanOrEqualToTen   Wrappers.Option
	MyList                     Wrappers.Option
	NonEmptyList               Wrappers.Option
	ListLessThanOrEqualToTen   Wrappers.Option
	MyMap                      Wrappers.Option
	NonEmptyMap                Wrappers.Option
	MapLessThanOrEqualToTen    Wrappers.Option
	OneToTen                   Wrappers.Option
	ThatTenToTen               Wrappers.Option
	GreaterThanOne             Wrappers.Option
	LessThanTen                Wrappers.Option
	MyUtf8Bytes                Wrappers.Option
	MyListOfUtf8Bytes          Wrappers.Option
}

func (GetConstraintsOutput_GetConstraintsOutput) isGetConstraintsOutput() {}

func (CompanionStruct_GetConstraintsOutput_) Create_GetConstraintsOutput_(MyString Wrappers.Option, NonEmptyString Wrappers.Option, StringLessThanOrEqualToTen Wrappers.Option, MyBlob Wrappers.Option, NonEmptyBlob Wrappers.Option, BlobLessThanOrEqualToTen Wrappers.Option, MyList Wrappers.Option, NonEmptyList Wrappers.Option, ListLessThanOrEqualToTen Wrappers.Option, MyMap Wrappers.Option, NonEmptyMap Wrappers.Option, MapLessThanOrEqualToTen Wrappers.Option, OneToTen Wrappers.Option, ThatTenToTen Wrappers.Option, GreaterThanOne Wrappers.Option, LessThanTen Wrappers.Option, MyUtf8Bytes Wrappers.Option, MyListOfUtf8Bytes Wrappers.Option) GetConstraintsOutput {
	return GetConstraintsOutput{GetConstraintsOutput_GetConstraintsOutput{MyString, NonEmptyString, StringLessThanOrEqualToTen, MyBlob, NonEmptyBlob, BlobLessThanOrEqualToTen, MyList, NonEmptyList, ListLessThanOrEqualToTen, MyMap, NonEmptyMap, MapLessThanOrEqualToTen, OneToTen, ThatTenToTen, GreaterThanOne, LessThanTen, MyUtf8Bytes, MyListOfUtf8Bytes}}
}

func (_this GetConstraintsOutput) Is_GetConstraintsOutput() bool {
	_, ok := _this.Get_().(GetConstraintsOutput_GetConstraintsOutput)
	return ok
}

func (CompanionStruct_GetConstraintsOutput_) Default() GetConstraintsOutput {
	return Companion_GetConstraintsOutput_.Create_GetConstraintsOutput_(Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default(), Wrappers.Companion_Option_.Default())
}

func (_this GetConstraintsOutput) Dtor_MyString() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).MyString
}

func (_this GetConstraintsOutput) Dtor_NonEmptyString() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).NonEmptyString
}

func (_this GetConstraintsOutput) Dtor_StringLessThanOrEqualToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).StringLessThanOrEqualToTen
}

func (_this GetConstraintsOutput) Dtor_MyBlob() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).MyBlob
}

func (_this GetConstraintsOutput) Dtor_NonEmptyBlob() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).NonEmptyBlob
}

func (_this GetConstraintsOutput) Dtor_BlobLessThanOrEqualToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).BlobLessThanOrEqualToTen
}

func (_this GetConstraintsOutput) Dtor_MyList() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).MyList
}

func (_this GetConstraintsOutput) Dtor_NonEmptyList() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).NonEmptyList
}

func (_this GetConstraintsOutput) Dtor_ListLessThanOrEqualToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).ListLessThanOrEqualToTen
}

func (_this GetConstraintsOutput) Dtor_MyMap() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).MyMap
}

func (_this GetConstraintsOutput) Dtor_NonEmptyMap() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).NonEmptyMap
}

func (_this GetConstraintsOutput) Dtor_MapLessThanOrEqualToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).MapLessThanOrEqualToTen
}

func (_this GetConstraintsOutput) Dtor_OneToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).OneToTen
}

func (_this GetConstraintsOutput) Dtor_thatTenToTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).ThatTenToTen
}

func (_this GetConstraintsOutput) Dtor_GreaterThanOne() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).GreaterThanOne
}

func (_this GetConstraintsOutput) Dtor_LessThanTen() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).LessThanTen
}

func (_this GetConstraintsOutput) Dtor_MyUtf8Bytes() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).MyUtf8Bytes
}

func (_this GetConstraintsOutput) Dtor_MyListOfUtf8Bytes() Wrappers.Option {
	return _this.Get_().(GetConstraintsOutput_GetConstraintsOutput).MyListOfUtf8Bytes
}

func (_this GetConstraintsOutput) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case GetConstraintsOutput_GetConstraintsOutput:
		{
			return "SimpleConstraintsTypes.GetConstraintsOutput.GetConstraintsOutput" + "(" + _dafny.String(data.MyString) + ", " + _dafny.String(data.NonEmptyString) + ", " + _dafny.String(data.StringLessThanOrEqualToTen) + ", " + _dafny.String(data.MyBlob) + ", " + _dafny.String(data.NonEmptyBlob) + ", " + _dafny.String(data.BlobLessThanOrEqualToTen) + ", " + _dafny.String(data.MyList) + ", " + _dafny.String(data.NonEmptyList) + ", " + _dafny.String(data.ListLessThanOrEqualToTen) + ", " + _dafny.String(data.MyMap) + ", " + _dafny.String(data.NonEmptyMap) + ", " + _dafny.String(data.MapLessThanOrEqualToTen) + ", " + _dafny.String(data.OneToTen) + ", " + _dafny.String(data.ThatTenToTen) + ", " + _dafny.String(data.GreaterThanOne) + ", " + _dafny.String(data.LessThanTen) + ", " + _dafny.String(data.MyUtf8Bytes) + ", " + _dafny.String(data.MyListOfUtf8Bytes) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this GetConstraintsOutput) Equals(other GetConstraintsOutput) bool {
	switch data1 := _this.Get_().(type) {
	case GetConstraintsOutput_GetConstraintsOutput:
		{
			data2, ok := other.Get_().(GetConstraintsOutput_GetConstraintsOutput)
			return ok && data1.MyString.Equals(data2.MyString) && data1.NonEmptyString.Equals(data2.NonEmptyString) && data1.StringLessThanOrEqualToTen.Equals(data2.StringLessThanOrEqualToTen) && data1.MyBlob.Equals(data2.MyBlob) && data1.NonEmptyBlob.Equals(data2.NonEmptyBlob) && data1.BlobLessThanOrEqualToTen.Equals(data2.BlobLessThanOrEqualToTen) && data1.MyList.Equals(data2.MyList) && data1.NonEmptyList.Equals(data2.NonEmptyList) && data1.ListLessThanOrEqualToTen.Equals(data2.ListLessThanOrEqualToTen) && data1.MyMap.Equals(data2.MyMap) && data1.NonEmptyMap.Equals(data2.NonEmptyMap) && data1.MapLessThanOrEqualToTen.Equals(data2.MapLessThanOrEqualToTen) && data1.OneToTen.Equals(data2.OneToTen) && data1.ThatTenToTen.Equals(data2.ThatTenToTen) && data1.GreaterThanOne.Equals(data2.GreaterThanOne) && data1.LessThanTen.Equals(data2.LessThanTen) && data1.MyUtf8Bytes.Equals(data2.MyUtf8Bytes) && data1.MyListOfUtf8Bytes.Equals(data2.MyListOfUtf8Bytes)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this GetConstraintsOutput) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(GetConstraintsOutput)
	return ok && _this.Equals(typed)
}

func Type_GetConstraintsOutput_() _dafny.TypeDescriptor {
	return type_GetConstraintsOutput_{}
}

type type_GetConstraintsOutput_ struct {
}

func (_this type_GetConstraintsOutput_) Default() interface{} {
	return Companion_GetConstraintsOutput_.Default()
}

func (_this type_GetConstraintsOutput_) String() string {
	return "simpleconstraintsinternaldafnytypes.GetConstraintsOutput"
}
func (_this GetConstraintsOutput) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = GetConstraintsOutput{}

// End of datatype GetConstraintsOutput

// Definition of class GreaterThanOne
type GreaterThanOne struct {
}

func New_GreaterThanOne_() *GreaterThanOne {
	_this := GreaterThanOne{}

	return &_this
}

type CompanionStruct_GreaterThanOne_ struct {
}

var Companion_GreaterThanOne_ = CompanionStruct_GreaterThanOne_{}

func (*GreaterThanOne) String() string {
	return "simpleconstraintsinternaldafnytypes.GreaterThanOne"
}

// End of class GreaterThanOne

func Type_GreaterThanOne_() _dafny.TypeDescriptor {
	return type_GreaterThanOne_{}
}

type type_GreaterThanOne_ struct {
}

func (_this type_GreaterThanOne_) Default() interface{} {
	return int32(0)
}

func (_this type_GreaterThanOne_) String() string {
	return "simpleconstraintsinternaldafnytypes.GreaterThanOne"
}
func (_this *CompanionStruct_GreaterThanOne_) Is_(__source int32) bool {
	var _1_x int32 = (__source)
	_ = _1_x
	if true {
		return Companion_Default___.IsValid__GreaterThanOne(_1_x)
	}
	return false
}

// Definition of class LessThanTen
type LessThanTen struct {
}

func New_LessThanTen_() *LessThanTen {
	_this := LessThanTen{}

	return &_this
}

type CompanionStruct_LessThanTen_ struct {
}

var Companion_LessThanTen_ = CompanionStruct_LessThanTen_{}

func (*LessThanTen) String() string {
	return "simpleconstraintsinternaldafnytypes.LessThanTen"
}

// End of class LessThanTen

func Type_LessThanTen_() _dafny.TypeDescriptor {
	return type_LessThanTen_{}
}

type type_LessThanTen_ struct {
}

func (_this type_LessThanTen_) Default() interface{} {
	return int32(0)
}

func (_this type_LessThanTen_) String() string {
	return "simpleconstraintsinternaldafnytypes.LessThanTen"
}
func (_this *CompanionStruct_LessThanTen_) Is_(__source int32) bool {
	var _2_x int32 = (__source)
	_ = _2_x
	if true {
		return Companion_Default___.IsValid__LessThanTen(_2_x)
	}
	return false
}

// Definition of class ListLessThanOrEqualToTen
type ListLessThanOrEqualToTen struct {
}

func New_ListLessThanOrEqualToTen_() *ListLessThanOrEqualToTen {
	_this := ListLessThanOrEqualToTen{}

	return &_this
}

type CompanionStruct_ListLessThanOrEqualToTen_ struct {
}

var Companion_ListLessThanOrEqualToTen_ = CompanionStruct_ListLessThanOrEqualToTen_{}

func (*ListLessThanOrEqualToTen) String() string {
	return "simpleconstraintsinternaldafnytypes.ListLessThanOrEqualToTen"
}

// End of class ListLessThanOrEqualToTen

func Type_ListLessThanOrEqualToTen_() _dafny.TypeDescriptor {
	return type_ListLessThanOrEqualToTen_{}
}

type type_ListLessThanOrEqualToTen_ struct {
}

func (_this type_ListLessThanOrEqualToTen_) Default() interface{} {
	return _dafny.EmptySeq
}

func (_this type_ListLessThanOrEqualToTen_) String() string {
	return "simpleconstraintsinternaldafnytypes.ListLessThanOrEqualToTen"
}
func (_this *CompanionStruct_ListLessThanOrEqualToTen_) Is_(__source _dafny.Sequence) bool {
	var _3_x _dafny.Sequence = (__source)
	_ = _3_x
	return Companion_Default___.IsValid__ListLessThanOrEqualToTen(_3_x)
}

// Definition of class ListOfUtf8Bytes
type ListOfUtf8Bytes struct {
}

func New_ListOfUtf8Bytes_() *ListOfUtf8Bytes {
	_this := ListOfUtf8Bytes{}

	return &_this
}

type CompanionStruct_ListOfUtf8Bytes_ struct {
}

var Companion_ListOfUtf8Bytes_ = CompanionStruct_ListOfUtf8Bytes_{}

func (*ListOfUtf8Bytes) String() string {
	return "simpleconstraintsinternaldafnytypes.ListOfUtf8Bytes"
}

// End of class ListOfUtf8Bytes

func Type_ListOfUtf8Bytes_() _dafny.TypeDescriptor {
	return type_ListOfUtf8Bytes_{}
}

type type_ListOfUtf8Bytes_ struct {
}

func (_this type_ListOfUtf8Bytes_) Default() interface{} {
	return _dafny.EmptySeq
}

func (_this type_ListOfUtf8Bytes_) String() string {
	return "simpleconstraintsinternaldafnytypes.ListOfUtf8Bytes"
}
func (_this *CompanionStruct_ListOfUtf8Bytes_) Is_(__source _dafny.Sequence) bool {
	var _4_x _dafny.Sequence = (__source)
	_ = _4_x
	return Companion_Default___.IsValid__ListOfUtf8Bytes(_4_x)
}

// Definition of class MapLessThanOrEqualToTen
type MapLessThanOrEqualToTen struct {
}

func New_MapLessThanOrEqualToTen_() *MapLessThanOrEqualToTen {
	_this := MapLessThanOrEqualToTen{}

	return &_this
}

type CompanionStruct_MapLessThanOrEqualToTen_ struct {
}

var Companion_MapLessThanOrEqualToTen_ = CompanionStruct_MapLessThanOrEqualToTen_{}

func (*MapLessThanOrEqualToTen) String() string {
	return "simpleconstraintsinternaldafnytypes.MapLessThanOrEqualToTen"
}

// End of class MapLessThanOrEqualToTen

func Type_MapLessThanOrEqualToTen_() _dafny.TypeDescriptor {
	return type_MapLessThanOrEqualToTen_{}
}

type type_MapLessThanOrEqualToTen_ struct {
}

func (_this type_MapLessThanOrEqualToTen_) Default() interface{} {
	return _dafny.EmptyMap
}

func (_this type_MapLessThanOrEqualToTen_) String() string {
	return "simpleconstraintsinternaldafnytypes.MapLessThanOrEqualToTen"
}
func (_this *CompanionStruct_MapLessThanOrEqualToTen_) Is_(__source _dafny.Map) bool {
	var _5_x _dafny.Map = (__source)
	_ = _5_x
	return Companion_Default___.IsValid__MapLessThanOrEqualToTen(_5_x)
}

// Definition of class MyBlob
type MyBlob struct {
}

func New_MyBlob_() *MyBlob {
	_this := MyBlob{}

	return &_this
}

type CompanionStruct_MyBlob_ struct {
}

var Companion_MyBlob_ = CompanionStruct_MyBlob_{}

func (*MyBlob) String() string {
	return "simpleconstraintsinternaldafnytypes.MyBlob"
}

// End of class MyBlob

func Type_MyBlob_() _dafny.TypeDescriptor {
	return type_MyBlob_{}
}

type type_MyBlob_ struct {
}

func (_this type_MyBlob_) Default() interface{} {
	return _dafny.EmptySeq
}

func (_this type_MyBlob_) String() string {
	return "simpleconstraintsinternaldafnytypes.MyBlob"
}
func (_this *CompanionStruct_MyBlob_) Is_(__source _dafny.Sequence) bool {
	var _6_x _dafny.Sequence = (__source)
	_ = _6_x
	return Companion_Default___.IsValid__MyBlob(_6_x)
}

// Definition of class MyList
type MyList struct {
}

func New_MyList_() *MyList {
	_this := MyList{}

	return &_this
}

type CompanionStruct_MyList_ struct {
}

var Companion_MyList_ = CompanionStruct_MyList_{}

func (*MyList) String() string {
	return "simpleconstraintsinternaldafnytypes.MyList"
}

// End of class MyList

func Type_MyList_() _dafny.TypeDescriptor {
	return type_MyList_{}
}

type type_MyList_ struct {
}

func (_this type_MyList_) Default() interface{} {
	return _dafny.EmptySeq
}

func (_this type_MyList_) String() string {
	return "simpleconstraintsinternaldafnytypes.MyList"
}
func (_this *CompanionStruct_MyList_) Is_(__source _dafny.Sequence) bool {
	var _7_x _dafny.Sequence = (__source)
	_ = _7_x
	return Companion_Default___.IsValid__MyList(_7_x)
}

// Definition of class MyMap
type MyMap struct {
}

func New_MyMap_() *MyMap {
	_this := MyMap{}

	return &_this
}

type CompanionStruct_MyMap_ struct {
}

var Companion_MyMap_ = CompanionStruct_MyMap_{}

func (*MyMap) String() string {
	return "simpleconstraintsinternaldafnytypes.MyMap"
}

// End of class MyMap

func Type_MyMap_() _dafny.TypeDescriptor {
	return type_MyMap_{}
}

type type_MyMap_ struct {
}

func (_this type_MyMap_) Default() interface{} {
	return _dafny.EmptyMap
}

func (_this type_MyMap_) String() string {
	return "simpleconstraintsinternaldafnytypes.MyMap"
}
func (_this *CompanionStruct_MyMap_) Is_(__source _dafny.Map) bool {
	var _8_x _dafny.Map = (__source)
	_ = _8_x
	return Companion_Default___.IsValid__MyMap(_8_x)
}

// Definition of class MyString
type MyString struct {
}

func New_MyString_() *MyString {
	_this := MyString{}

	return &_this
}

type CompanionStruct_MyString_ struct {
}

var Companion_MyString_ = CompanionStruct_MyString_{}

func (*MyString) String() string {
	return "simpleconstraintsinternaldafnytypes.MyString"
}

// End of class MyString

func Type_MyString_() _dafny.TypeDescriptor {
	return type_MyString_{}
}

type type_MyString_ struct {
}

func (_this type_MyString_) Default() interface{} {
	return _dafny.EmptySeq.SetString()
}

func (_this type_MyString_) String() string {
	return "simpleconstraintsinternaldafnytypes.MyString"
}
func (_this *CompanionStruct_MyString_) Is_(__source _dafny.Sequence) bool {
	var _9_x _dafny.Sequence = (__source)
	_ = _9_x
	return Companion_Default___.IsValid__MyString(_9_x)
}

// Definition of class NonEmptyBlob
type NonEmptyBlob struct {
}

func New_NonEmptyBlob_() *NonEmptyBlob {
	_this := NonEmptyBlob{}

	return &_this
}

type CompanionStruct_NonEmptyBlob_ struct {
}

var Companion_NonEmptyBlob_ = CompanionStruct_NonEmptyBlob_{}

func (*NonEmptyBlob) String() string {
	return "simpleconstraintsinternaldafnytypes.NonEmptyBlob"
}

// End of class NonEmptyBlob

func Type_NonEmptyBlob_() _dafny.TypeDescriptor {
	return type_NonEmptyBlob_{}
}

type type_NonEmptyBlob_ struct {
}

func (_this type_NonEmptyBlob_) Default() interface{} {
	return _dafny.EmptySeq
}

func (_this type_NonEmptyBlob_) String() string {
	return "simpleconstraintsinternaldafnytypes.NonEmptyBlob"
}
func (_this *CompanionStruct_NonEmptyBlob_) Is_(__source _dafny.Sequence) bool {
	var _10_x _dafny.Sequence = (__source)
	_ = _10_x
	return Companion_Default___.IsValid__NonEmptyBlob(_10_x)
}

// Definition of class NonEmptyList
type NonEmptyList struct {
}

func New_NonEmptyList_() *NonEmptyList {
	_this := NonEmptyList{}

	return &_this
}

type CompanionStruct_NonEmptyList_ struct {
}

var Companion_NonEmptyList_ = CompanionStruct_NonEmptyList_{}

func (*NonEmptyList) String() string {
	return "simpleconstraintsinternaldafnytypes.NonEmptyList"
}

// End of class NonEmptyList

func Type_NonEmptyList_() _dafny.TypeDescriptor {
	return type_NonEmptyList_{}
}

type type_NonEmptyList_ struct {
}

func (_this type_NonEmptyList_) Default() interface{} {
	return _dafny.EmptySeq
}

func (_this type_NonEmptyList_) String() string {
	return "simpleconstraintsinternaldafnytypes.NonEmptyList"
}
func (_this *CompanionStruct_NonEmptyList_) Is_(__source _dafny.Sequence) bool {
	var _11_x _dafny.Sequence = (__source)
	_ = _11_x
	return Companion_Default___.IsValid__NonEmptyList(_11_x)
}

// Definition of class NonEmptyMap
type NonEmptyMap struct {
}

func New_NonEmptyMap_() *NonEmptyMap {
	_this := NonEmptyMap{}

	return &_this
}

type CompanionStruct_NonEmptyMap_ struct {
}

var Companion_NonEmptyMap_ = CompanionStruct_NonEmptyMap_{}

func (*NonEmptyMap) String() string {
	return "simpleconstraintsinternaldafnytypes.NonEmptyMap"
}

// End of class NonEmptyMap

func Type_NonEmptyMap_() _dafny.TypeDescriptor {
	return type_NonEmptyMap_{}
}

type type_NonEmptyMap_ struct {
}

func (_this type_NonEmptyMap_) Default() interface{} {
	return _dafny.EmptyMap
}

func (_this type_NonEmptyMap_) String() string {
	return "simpleconstraintsinternaldafnytypes.NonEmptyMap"
}
func (_this *CompanionStruct_NonEmptyMap_) Is_(__source _dafny.Map) bool {
	var _12_x _dafny.Map = (__source)
	_ = _12_x
	return Companion_Default___.IsValid__NonEmptyMap(_12_x)
}

// Definition of class NonEmptyString
type NonEmptyString struct {
}

func New_NonEmptyString_() *NonEmptyString {
	_this := NonEmptyString{}

	return &_this
}

type CompanionStruct_NonEmptyString_ struct {
}

var Companion_NonEmptyString_ = CompanionStruct_NonEmptyString_{}

func (*NonEmptyString) String() string {
	return "simpleconstraintsinternaldafnytypes.NonEmptyString"
}

// End of class NonEmptyString

func Type_NonEmptyString_() _dafny.TypeDescriptor {
	return type_NonEmptyString_{}
}

type type_NonEmptyString_ struct {
}

func (_this type_NonEmptyString_) Default() interface{} {
	return _dafny.EmptySeq.SetString()
}

func (_this type_NonEmptyString_) String() string {
	return "simpleconstraintsinternaldafnytypes.NonEmptyString"
}
func (_this *CompanionStruct_NonEmptyString_) Is_(__source _dafny.Sequence) bool {
	var _13_x _dafny.Sequence = (__source)
	_ = _13_x
	return Companion_Default___.IsValid__NonEmptyString(_13_x)
}

// Definition of class OneToTen
type OneToTen struct {
}

func New_OneToTen_() *OneToTen {
	_this := OneToTen{}

	return &_this
}

type CompanionStruct_OneToTen_ struct {
}

var Companion_OneToTen_ = CompanionStruct_OneToTen_{}

func (*OneToTen) String() string {
	return "simpleconstraintsinternaldafnytypes.OneToTen"
}

// End of class OneToTen

func Type_OneToTen_() _dafny.TypeDescriptor {
	return type_OneToTen_{}
}

type type_OneToTen_ struct {
}

func (_this type_OneToTen_) Default() interface{} {
	return int32(0)
}

func (_this type_OneToTen_) String() string {
	return "simpleconstraintsinternaldafnytypes.OneToTen"
}
func (_this *CompanionStruct_OneToTen_) Is_(__source int32) bool {
	var _14_x int32 = (__source)
	_ = _14_x
	if true {
		return Companion_Default___.IsValid__OneToTen(_14_x)
	}
	return false
}

// Definition of class ISimpleConstraintsClientCallHistory
type ISimpleConstraintsClientCallHistory struct {
	dummy byte
}

func New_ISimpleConstraintsClientCallHistory_() *ISimpleConstraintsClientCallHistory {
	_this := ISimpleConstraintsClientCallHistory{}

	return &_this
}

type CompanionStruct_ISimpleConstraintsClientCallHistory_ struct {
}

var Companion_ISimpleConstraintsClientCallHistory_ = CompanionStruct_ISimpleConstraintsClientCallHistory_{}

func (_this *ISimpleConstraintsClientCallHistory) Equals(other *ISimpleConstraintsClientCallHistory) bool {
	return _this == other
}

func (_this *ISimpleConstraintsClientCallHistory) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*ISimpleConstraintsClientCallHistory)
	return ok && _this.Equals(other)
}

func (*ISimpleConstraintsClientCallHistory) String() string {
	return "simpleconstraintsinternaldafnytypes.ISimpleConstraintsClientCallHistory"
}

func Type_ISimpleConstraintsClientCallHistory_() _dafny.TypeDescriptor {
	return type_ISimpleConstraintsClientCallHistory_{}
}

type type_ISimpleConstraintsClientCallHistory_ struct {
}

func (_this type_ISimpleConstraintsClientCallHistory_) Default() interface{} {
	return (*ISimpleConstraintsClientCallHistory)(nil)
}

func (_this type_ISimpleConstraintsClientCallHistory_) String() string {
	return "simpleconstraintsinternaldafnytypes.ISimpleConstraintsClientCallHistory"
}
func (_this *ISimpleConstraintsClientCallHistory) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &ISimpleConstraintsClientCallHistory{}

// End of class ISimpleConstraintsClientCallHistory

// Definition of trait ISimpleConstraintsClient
type ISimpleConstraintsClient interface {
	String() string
	GetConstraints(input GetConstraintsInput) Wrappers.Result
}
type CompanionStruct_ISimpleConstraintsClient_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_ISimpleConstraintsClient_ = CompanionStruct_ISimpleConstraintsClient_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_ISimpleConstraintsClient_) CastTo_(x interface{}) ISimpleConstraintsClient {
	var t ISimpleConstraintsClient
	t, _ = x.(ISimpleConstraintsClient)
	return t
}

// End of trait ISimpleConstraintsClient

// Definition of datatype SimpleConstraintsConfig
type SimpleConstraintsConfig struct {
	Data_SimpleConstraintsConfig_
}

func (_this SimpleConstraintsConfig) Get_() Data_SimpleConstraintsConfig_ {
	return _this.Data_SimpleConstraintsConfig_
}

type Data_SimpleConstraintsConfig_ interface {
	isSimpleConstraintsConfig()
}

type CompanionStruct_SimpleConstraintsConfig_ struct {
}

var Companion_SimpleConstraintsConfig_ = CompanionStruct_SimpleConstraintsConfig_{}

type SimpleConstraintsConfig_SimpleConstraintsConfig struct {
}

func (SimpleConstraintsConfig_SimpleConstraintsConfig) isSimpleConstraintsConfig() {}

func (CompanionStruct_SimpleConstraintsConfig_) Create_SimpleConstraintsConfig_() SimpleConstraintsConfig {
	return SimpleConstraintsConfig{SimpleConstraintsConfig_SimpleConstraintsConfig{}}
}

func (_this SimpleConstraintsConfig) Is_SimpleConstraintsConfig() bool {
	_, ok := _this.Get_().(SimpleConstraintsConfig_SimpleConstraintsConfig)
	return ok
}

func (CompanionStruct_SimpleConstraintsConfig_) Default() SimpleConstraintsConfig {
	return Companion_SimpleConstraintsConfig_.Create_SimpleConstraintsConfig_()
}

func (_ CompanionStruct_SimpleConstraintsConfig_) AllSingletonConstructors() _dafny.Iterator {
	i := -1
	return func() (interface{}, bool) {
		i++
		switch i {
		case 0:
			return Companion_SimpleConstraintsConfig_.Create_SimpleConstraintsConfig_(), true
		default:
			return SimpleConstraintsConfig{}, false
		}
	}
}

func (_this SimpleConstraintsConfig) String() string {
	switch _this.Get_().(type) {
	case nil:
		return "null"
	case SimpleConstraintsConfig_SimpleConstraintsConfig:
		{
			return "SimpleConstraintsTypes.SimpleConstraintsConfig.SimpleConstraintsConfig"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this SimpleConstraintsConfig) Equals(other SimpleConstraintsConfig) bool {
	switch _this.Get_().(type) {
	case SimpleConstraintsConfig_SimpleConstraintsConfig:
		{
			_, ok := other.Get_().(SimpleConstraintsConfig_SimpleConstraintsConfig)
			return ok
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this SimpleConstraintsConfig) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(SimpleConstraintsConfig)
	return ok && _this.Equals(typed)
}

func Type_SimpleConstraintsConfig_() _dafny.TypeDescriptor {
	return type_SimpleConstraintsConfig_{}
}

type type_SimpleConstraintsConfig_ struct {
}

func (_this type_SimpleConstraintsConfig_) Default() interface{} {
	return Companion_SimpleConstraintsConfig_.Default()
}

func (_this type_SimpleConstraintsConfig_) String() string {
	return "simpleconstraintsinternaldafnytypes.SimpleConstraintsConfig"
}
func (_this SimpleConstraintsConfig) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = SimpleConstraintsConfig{}

// End of datatype SimpleConstraintsConfig

// Definition of class StringLessThanOrEqualToTen
type StringLessThanOrEqualToTen struct {
}

func New_StringLessThanOrEqualToTen_() *StringLessThanOrEqualToTen {
	_this := StringLessThanOrEqualToTen{}

	return &_this
}

type CompanionStruct_StringLessThanOrEqualToTen_ struct {
}

var Companion_StringLessThanOrEqualToTen_ = CompanionStruct_StringLessThanOrEqualToTen_{}

func (*StringLessThanOrEqualToTen) String() string {
	return "simpleconstraintsinternaldafnytypes.StringLessThanOrEqualToTen"
}

// End of class StringLessThanOrEqualToTen

func Type_StringLessThanOrEqualToTen_() _dafny.TypeDescriptor {
	return type_StringLessThanOrEqualToTen_{}
}

type type_StringLessThanOrEqualToTen_ struct {
}

func (_this type_StringLessThanOrEqualToTen_) Default() interface{} {
	return _dafny.EmptySeq.SetString()
}

func (_this type_StringLessThanOrEqualToTen_) String() string {
	return "simpleconstraintsinternaldafnytypes.StringLessThanOrEqualToTen"
}
func (_this *CompanionStruct_StringLessThanOrEqualToTen_) Is_(__source _dafny.Sequence) bool {
	var _15_x _dafny.Sequence = (__source)
	_ = _15_x
	return Companion_Default___.IsValid__StringLessThanOrEqualToTen(_15_x)
}

// Definition of class TenToTen
type TenToTen struct {
}

func New_TenToTen_() *TenToTen {
	_this := TenToTen{}

	return &_this
}

type CompanionStruct_TenToTen_ struct {
}

var Companion_TenToTen_ = CompanionStruct_TenToTen_{}

func (*TenToTen) String() string {
	return "simpleconstraintsinternaldafnytypes.TenToTen"
}

// End of class TenToTen

func Type_TenToTen_() _dafny.TypeDescriptor {
	return type_TenToTen_{}
}

type type_TenToTen_ struct {
}

func (_this type_TenToTen_) Default() interface{} {
	return int64(0)
}

func (_this type_TenToTen_) String() string {
	return "simpleconstraintsinternaldafnytypes.TenToTen"
}
func (_this *CompanionStruct_TenToTen_) Is_(__source int64) bool {
	var _16_x int64 = (__source)
	_ = _16_x
	if true {
		return Companion_Default___.IsValid__TenToTen(_16_x)
	}
	return false
}

// Definition of class Utf8Bytes
type Utf8Bytes struct {
}

func New_Utf8Bytes_() *Utf8Bytes {
	_this := Utf8Bytes{}

	return &_this
}

type CompanionStruct_Utf8Bytes_ struct {
}

var Companion_Utf8Bytes_ = CompanionStruct_Utf8Bytes_{}

func (*Utf8Bytes) String() string {
	return "simpleconstraintsinternaldafnytypes.Utf8Bytes"
}

// End of class Utf8Bytes

func Type_Utf8Bytes_() _dafny.TypeDescriptor {
	return type_Utf8Bytes_{}
}

type type_Utf8Bytes_ struct {
}

func (_this type_Utf8Bytes_) Default() interface{} {
	return UTF8.Companion_ValidUTF8Bytes_.Witness()
}

func (_this type_Utf8Bytes_) String() string {
	return "simpleconstraintsinternaldafnytypes.Utf8Bytes"
}
func (_this *CompanionStruct_Utf8Bytes_) Is_(__source _dafny.Sequence) bool {
	var _17_x _dafny.Sequence = (__source)
	_ = _17_x
	if UTF8.Companion_ValidUTF8Bytes_.Is_(_17_x) {
		return Companion_Default___.IsValid__Utf8Bytes(_17_x)
	}
	return false
}

// Definition of datatype Error
type Error struct {
	Data_Error_
}

func (_this Error) Get_() Data_Error_ {
	return _this.Data_Error_
}

type Data_Error_ interface {
	isError()
}

type CompanionStruct_Error_ struct {
}

var Companion_Error_ = CompanionStruct_Error_{}

type Error_SimpleConstraintsException struct {
	Message _dafny.Sequence
}

func (Error_SimpleConstraintsException) isError() {}

func (CompanionStruct_Error_) Create_SimpleConstraintsException_(Message _dafny.Sequence) Error {
	return Error{Error_SimpleConstraintsException{Message}}
}

func (_this Error) Is_SimpleConstraintsException() bool {
	_, ok := _this.Get_().(Error_SimpleConstraintsException)
	return ok
}

type Error_CollectionOfErrors struct {
	List    _dafny.Sequence
	Message _dafny.Sequence
}

func (Error_CollectionOfErrors) isError() {}

func (CompanionStruct_Error_) Create_CollectionOfErrors_(List _dafny.Sequence, Message _dafny.Sequence) Error {
	return Error{Error_CollectionOfErrors{List, Message}}
}

func (_this Error) Is_CollectionOfErrors() bool {
	_, ok := _this.Get_().(Error_CollectionOfErrors)
	return ok
}

type Error_Opaque struct {
	Obj interface{}
}

func (Error_Opaque) isError() {}

func (CompanionStruct_Error_) Create_Opaque_(Obj interface{}) Error {
	return Error{Error_Opaque{Obj}}
}

func (_this Error) Is_Opaque() bool {
	_, ok := _this.Get_().(Error_Opaque)
	return ok
}

func (CompanionStruct_Error_) Default() Error {
	return Companion_Error_.Create_SimpleConstraintsException_(_dafny.EmptySeq.SetString())
}

func (_this Error) Dtor_message() _dafny.Sequence {
	switch data := _this.Get_().(type) {
	case Error_SimpleConstraintsException:
		return data.Message
	default:
		return data.(Error_CollectionOfErrors).Message
	}
}

func (_this Error) Dtor_list() _dafny.Sequence {
	return _this.Get_().(Error_CollectionOfErrors).List
}

func (_this Error) Dtor_obj() interface{} {
	return _this.Get_().(Error_Opaque).Obj
}

func (_this Error) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case Error_SimpleConstraintsException:
		{
			return "SimpleConstraintsTypes.Error.SimpleConstraintsException" + "(" + _dafny.String(data.Message) + ")"
		}
	case Error_CollectionOfErrors:
		{
			return "SimpleConstraintsTypes.Error.CollectionOfErrors" + "(" + _dafny.String(data.List) + ", " + _dafny.String(data.Message) + ")"
		}
	case Error_Opaque:
		{
			return "SimpleConstraintsTypes.Error.Opaque" + "(" + _dafny.String(data.Obj) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this Error) Equals(other Error) bool {
	switch data1 := _this.Get_().(type) {
	case Error_SimpleConstraintsException:
		{
			data2, ok := other.Get_().(Error_SimpleConstraintsException)
			return ok && data1.Message.Equals(data2.Message)
		}
	case Error_CollectionOfErrors:
		{
			data2, ok := other.Get_().(Error_CollectionOfErrors)
			return ok && data1.List.Equals(data2.List) && data1.Message.Equals(data2.Message)
		}
	case Error_Opaque:
		{
			data2, ok := other.Get_().(Error_Opaque)
			return ok && _dafny.AreEqual(data1.Obj, data2.Obj)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this Error) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(Error)
	return ok && _this.Equals(typed)
}

func Type_Error_() _dafny.TypeDescriptor {
	return type_Error_{}
}

type type_Error_ struct {
}

func (_this type_Error_) Default() interface{} {
	return Companion_Error_.Default()
}

func (_this type_Error_) String() string {
	return "simpleconstraintsinternaldafnytypes.Error"
}
func (_this Error) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = Error{}

// End of datatype Error

// Definition of class OpaqueError
type OpaqueError struct {
}

func New_OpaqueError_() *OpaqueError {
	_this := OpaqueError{}

	return &_this
}

type CompanionStruct_OpaqueError_ struct {
}

var Companion_OpaqueError_ = CompanionStruct_OpaqueError_{}

func (*OpaqueError) String() string {
	return "simpleconstraintsinternaldafnytypes.OpaqueError"
}

// End of class OpaqueError

func Type_OpaqueError_() _dafny.TypeDescriptor {
	return type_OpaqueError_{}
}

type type_OpaqueError_ struct {
}

func (_this type_OpaqueError_) Default() interface{} {
	return Companion_Error_.Default()
}

func (_this type_OpaqueError_) String() string {
	return "simpleconstraintsinternaldafnytypes.OpaqueError"
}
func (_this *CompanionStruct_OpaqueError_) Is_(__source Error) bool {
	var _18_e Error = (__source)
	_ = _18_e
	return (_18_e).Is_Opaque()
}
