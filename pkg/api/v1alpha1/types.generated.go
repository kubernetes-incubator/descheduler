/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by codecgen - DO NOT EDIT.

package v1alpha1

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg2_v1 "k8s.io/api/core/v1"
	pkg1_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
	"runtime"
)

const (
	// ----- content types ----
	codecSelferCcUTF81234 = 1
	codecSelferCcRAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234  = 10
	codecSelferValueTypeMap1234    = 9
	codecSelferValueTypeString1234 = 6
	codecSelferValueTypeInt1234    = 2
	codecSelferValueTypeUint1234   = 3
	codecSelferValueTypeFloat1234  = 4
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	errCodecSelferOnlyMapOrArrayEncodeToStruct1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 8 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			8, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg2_v1.ResourceName
		var v1 pkg1_v1.TypeMeta
		_, _ = v0, v1
	}
}

func (x *DeschedulerPolicy) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [3]bool
			_ = yyq2
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = len(x.Strategies) != 0
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(3)
			} else {
				var yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.WriteMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferCcUTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferCcUTF81234, "")
				}
			} else {
				if yyq2[0] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `kind`)
					r.WriteMapElemValue()
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferCcUTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferCcUTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferCcUTF81234, "")
				}
			} else {
				if yyq2[1] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `apiVersion`)
					r.WriteMapElemValue()
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferCcUTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[2] {
					if x.Strategies == nil {
						r.EncodeNil()
					} else {
						x.Strategies.CodecEncodeSelf(e)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `strategies`)
					r.WriteMapElemValue()
					if x.Strategies == nil {
						r.EncodeNil()
					} else {
						x.Strategies.CodecEncodeSelf(e)
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}

func (x *DeschedulerPolicy) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct1234)
		}
	}
}

func (x *DeschedulerPolicy) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3 := z.StringView(r.DecStructFieldKey(codecSelferValueTypeString1234, z.DecScratchArrayBuffer()))
		r.ReadMapElemValue()
		switch yys3 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				yyv4 := &x.Kind
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				yyv6 := &x.APIVersion
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "strategies":
			if r.TryDecodeAsNil() {
				x.Strategies = nil
			} else {
				yyv8 := &x.Strategies
				yyv8.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	r.ReadMapEnd()
}

func (x *DeschedulerPolicy) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj9 int
	var yyb9 bool
	var yyhl9 bool = l >= 0
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		yyv10 := &x.Kind
		yym11 := z.DecBinary()
		_ = yym11
		if false {
		} else {
			*((*string)(yyv10)) = r.DecodeString()
		}
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		yyv12 := &x.APIVersion
		yym13 := z.DecBinary()
		_ = yym13
		if false {
		} else {
			*((*string)(yyv12)) = r.DecodeString()
		}
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Strategies = nil
	} else {
		yyv14 := &x.Strategies
		yyv14.CodecDecodeSelf(d)
	}
	for {
		yyj9++
		if yyhl9 {
			yyb9 = yyj9 > l
		} else {
			yyb9 = r.CheckBreak()
		}
		if yyb9 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj9-1, "")
	}
	r.ReadArrayEnd()
}

func (x StrategyName) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	yym1 := z.EncBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.EncExtension(x, yyxt1)
	} else {
		r.EncodeString(codecSelferCcUTF81234, string(x))
	}
}

func (x *StrategyName) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		*((*string)(x)) = r.DecodeString()
	}
}

func (x StrategyList) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			h.encStrategyList((StrategyList)(x), e)
		}
	}
}

func (x *StrategyList) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		h.decStrategyList((*StrategyList)(x), d)
	}
}

func (x *DeschedulerStrategy) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [3]bool
			_ = yyq2
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Enabled != false
			yyq2[1] = x.Weight != 0
			yyq2[2] = true
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(3)
			} else {
				var yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.WriteMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeBool(bool(x.Enabled))
					}
				} else {
					r.EncodeBool(false)
				}
			} else {
				if yyq2[0] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `enabled`)
					r.WriteMapElemValue()
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeBool(bool(x.Enabled))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeInt(int64(x.Weight))
					}
				} else {
					r.EncodeInt(0)
				}
			} else {
				if yyq2[1] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `weight`)
					r.WriteMapElemValue()
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeInt(int64(x.Weight))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[2] {
					yy10 := &x.Params
					yy10.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `params`)
					r.WriteMapElemValue()
					yy12 := &x.Params
					yy12.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}

func (x *DeschedulerStrategy) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct1234)
		}
	}
}

func (x *DeschedulerStrategy) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3 := z.StringView(r.DecStructFieldKey(codecSelferValueTypeString1234, z.DecScratchArrayBuffer()))
		r.ReadMapElemValue()
		switch yys3 {
		case "enabled":
			if r.TryDecodeAsNil() {
				x.Enabled = false
			} else {
				yyv4 := &x.Enabled
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*bool)(yyv4)) = r.DecodeBool()
				}
			}
		case "weight":
			if r.TryDecodeAsNil() {
				x.Weight = 0
			} else {
				yyv6 := &x.Weight
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*int)(yyv6)) = int(r.DecodeInt(codecSelferBitsize1234))
				}
			}
		case "params":
			if r.TryDecodeAsNil() {
				x.Params = StrategyParameters{}
			} else {
				yyv8 := &x.Params
				yyv8.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	r.ReadMapEnd()
}

func (x *DeschedulerStrategy) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj9 int
	var yyb9 bool
	var yyhl9 bool = l >= 0
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Enabled = false
	} else {
		yyv10 := &x.Enabled
		yym11 := z.DecBinary()
		_ = yym11
		if false {
		} else {
			*((*bool)(yyv10)) = r.DecodeBool()
		}
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Weight = 0
	} else {
		yyv12 := &x.Weight
		yym13 := z.DecBinary()
		_ = yym13
		if false {
		} else {
			*((*int)(yyv12)) = int(r.DecodeInt(codecSelferBitsize1234))
		}
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Params = StrategyParameters{}
	} else {
		yyv14 := &x.Params
		yyv14.CodecDecodeSelf(d)
	}
	for {
		yyj9++
		if yyhl9 {
			yyb9 = yyj9 > l
		} else {
			yyb9 = r.CheckBreak()
		}
		if yyb9 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj9-1, "")
	}
	r.ReadArrayEnd()
}

func (x *StrategyParameters) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [3]bool
			_ = yyq2
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = true
			yyq2[1] = len(x.NodeAffinityType) != 0
			yyq2[2] = true
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(3)
			} else {
				var yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.WriteMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[0] {
					yy4 := &x.NodeResourceUtilizationThresholds
					yy4.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[0] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `nodeResourceUtilizationThresholds`)
					r.WriteMapElemValue()
					yy6 := &x.NodeResourceUtilizationThresholds
					yy6.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[1] {
					if x.NodeAffinityType == nil {
						r.EncodeNil()
					} else {
						yym9 := z.EncBinary()
						_ = yym9
						if false {
						} else {
							z.F.EncSliceStringV(x.NodeAffinityType, e)
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[1] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `nodeAffinityType`)
					r.WriteMapElemValue()
					if x.NodeAffinityType == nil {
						r.EncodeNil()
					} else {
						yym10 := z.EncBinary()
						_ = yym10
						if false {
						} else {
							z.F.EncSliceStringV(x.NodeAffinityType, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[2] {
					yy12 := &x.PodsHavingTooManyRestarts
					yy12.CodecEncodeSelf(e)
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `podsHavingTooManyRestarts`)
					r.WriteMapElemValue()
					yy14 := &x.PodsHavingTooManyRestarts
					yy14.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}

func (x *StrategyParameters) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct1234)
		}
	}
}

func (x *StrategyParameters) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3 := z.StringView(r.DecStructFieldKey(codecSelferValueTypeString1234, z.DecScratchArrayBuffer()))
		r.ReadMapElemValue()
		switch yys3 {
		case "nodeResourceUtilizationThresholds":
			if r.TryDecodeAsNil() {
				x.NodeResourceUtilizationThresholds = NodeResourceUtilizationThresholds{}
			} else {
				yyv4 := &x.NodeResourceUtilizationThresholds
				yyv4.CodecDecodeSelf(d)
			}
		case "nodeAffinityType":
			if r.TryDecodeAsNil() {
				x.NodeAffinityType = nil
			} else {
				yyv5 := &x.NodeAffinityType
				yym6 := z.DecBinary()
				_ = yym6
				if false {
				} else {
					z.F.DecSliceStringX(yyv5, d)
				}
			}
		case "podsHavingTooManyRestarts":
			if r.TryDecodeAsNil() {
				x.PodsHavingTooManyRestarts = PodsHavingTooManyRestarts{}
			} else {
				yyv7 := &x.PodsHavingTooManyRestarts
				yyv7.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	r.ReadMapEnd()
}

func (x *StrategyParameters) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj8 int
	var yyb8 bool
	var yyhl8 bool = l >= 0
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.NodeResourceUtilizationThresholds = NodeResourceUtilizationThresholds{}
	} else {
		yyv9 := &x.NodeResourceUtilizationThresholds
		yyv9.CodecDecodeSelf(d)
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.NodeAffinityType = nil
	} else {
		yyv10 := &x.NodeAffinityType
		yym11 := z.DecBinary()
		_ = yym11
		if false {
		} else {
			z.F.DecSliceStringX(yyv10, d)
		}
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PodsHavingTooManyRestarts = PodsHavingTooManyRestarts{}
	} else {
		yyv12 := &x.PodsHavingTooManyRestarts
		yyv12.CodecDecodeSelf(d)
	}
	for {
		yyj8++
		if yyhl8 {
			yyb8 = yyj8 > l
		} else {
			yyb8 = r.CheckBreak()
		}
		if yyb8 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj8-1, "")
	}
	r.ReadArrayEnd()
}

func (x Percentage) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	yym1 := z.EncBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.EncExtension(x, yyxt1)
	} else {
		r.EncodeFloat64(float64(x))
	}
}

func (x *Percentage) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		*((*float64)(x)) = r.DecodeFloat64()
	}
}

func (x ResourceThresholds) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			h.encResourceThresholds((ResourceThresholds)(x), e)
		}
	}
}

func (x *ResourceThresholds) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		h.decResourceThresholds((*ResourceThresholds)(x), d)
	}
}

func (x *NodeResourceUtilizationThresholds) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [3]bool
			_ = yyq2
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = len(x.Thresholds) != 0
			yyq2[1] = len(x.TargetThresholds) != 0
			yyq2[2] = x.NumberOfNodes != 0
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(3)
			} else {
				var yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.WriteMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[0] {
					if x.Thresholds == nil {
						r.EncodeNil()
					} else {
						x.Thresholds.CodecEncodeSelf(e)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[0] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `thresholds`)
					r.WriteMapElemValue()
					if x.Thresholds == nil {
						r.EncodeNil()
					} else {
						x.Thresholds.CodecEncodeSelf(e)
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[1] {
					if x.TargetThresholds == nil {
						r.EncodeNil()
					} else {
						x.TargetThresholds.CodecEncodeSelf(e)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[1] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `targetThresholds`)
					r.WriteMapElemValue()
					if x.TargetThresholds == nil {
						r.EncodeNil()
					} else {
						x.TargetThresholds.CodecEncodeSelf(e)
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[2] {
					yym10 := z.EncBinary()
					_ = yym10
					if false {
					} else {
						r.EncodeInt(int64(x.NumberOfNodes))
					}
				} else {
					r.EncodeInt(0)
				}
			} else {
				if yyq2[2] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `numberOfNodes`)
					r.WriteMapElemValue()
					yym11 := z.EncBinary()
					_ = yym11
					if false {
					} else {
						r.EncodeInt(int64(x.NumberOfNodes))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}

func (x *NodeResourceUtilizationThresholds) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct1234)
		}
	}
}

func (x *NodeResourceUtilizationThresholds) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3 := z.StringView(r.DecStructFieldKey(codecSelferValueTypeString1234, z.DecScratchArrayBuffer()))
		r.ReadMapElemValue()
		switch yys3 {
		case "thresholds":
			if r.TryDecodeAsNil() {
				x.Thresholds = nil
			} else {
				yyv4 := &x.Thresholds
				yyv4.CodecDecodeSelf(d)
			}
		case "targetThresholds":
			if r.TryDecodeAsNil() {
				x.TargetThresholds = nil
			} else {
				yyv5 := &x.TargetThresholds
				yyv5.CodecDecodeSelf(d)
			}
		case "numberOfNodes":
			if r.TryDecodeAsNil() {
				x.NumberOfNodes = 0
			} else {
				yyv6 := &x.NumberOfNodes
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*int)(yyv6)) = int(r.DecodeInt(codecSelferBitsize1234))
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	r.ReadMapEnd()
}

func (x *NodeResourceUtilizationThresholds) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj8 int
	var yyb8 bool
	var yyhl8 bool = l >= 0
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Thresholds = nil
	} else {
		yyv9 := &x.Thresholds
		yyv9.CodecDecodeSelf(d)
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.TargetThresholds = nil
	} else {
		yyv10 := &x.TargetThresholds
		yyv10.CodecDecodeSelf(d)
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.NumberOfNodes = 0
	} else {
		yyv11 := &x.NumberOfNodes
		yym12 := z.DecBinary()
		_ = yym12
		if false {
		} else {
			*((*int)(yyv11)) = int(r.DecodeInt(codecSelferBitsize1234))
		}
	}
	for {
		yyj8++
		if yyhl8 {
			yyb8 = yyj8 > l
		} else {
			yyb8 = r.CheckBreak()
		}
		if yyb8 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj8-1, "")
	}
	r.ReadArrayEnd()
}

func (x *PodsHavingTooManyRestarts) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [2]bool
			_ = yyq2
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.PodeRestartThresholds != 0
			yyq2[1] = x.IncludingInitContainers != false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(2)
			} else {
				var yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.WriteMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeInt(int64(x.PodeRestartThresholds))
					}
				} else {
					r.EncodeInt(0)
				}
			} else {
				if yyq2[0] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `podeRestartThresholds`)
					r.WriteMapElemValue()
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeInt(int64(x.PodeRestartThresholds))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeBool(bool(x.IncludingInitContainers))
					}
				} else {
					r.EncodeBool(false)
				}
			} else {
				if yyq2[1] {
					r.WriteMapElemKey()
					r.EncStructFieldKey(codecSelferValueTypeString1234, `includingInitContainers`)
					r.WriteMapElemValue()
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeBool(bool(x.IncludingInitContainers))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}

func (x *PodsHavingTooManyRestarts) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct1234)
		}
	}
}

func (x *PodsHavingTooManyRestarts) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3 := z.StringView(r.DecStructFieldKey(codecSelferValueTypeString1234, z.DecScratchArrayBuffer()))
		r.ReadMapElemValue()
		switch yys3 {
		case "podeRestartThresholds":
			if r.TryDecodeAsNil() {
				x.PodeRestartThresholds = 0
			} else {
				yyv4 := &x.PodeRestartThresholds
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*int32)(yyv4)) = int32(r.DecodeInt(32))
				}
			}
		case "includingInitContainers":
			if r.TryDecodeAsNil() {
				x.IncludingInitContainers = false
			} else {
				yyv6 := &x.IncludingInitContainers
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*bool)(yyv6)) = r.DecodeBool()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	r.ReadMapEnd()
}

func (x *PodsHavingTooManyRestarts) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj8 int
	var yyb8 bool
	var yyhl8 bool = l >= 0
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PodeRestartThresholds = 0
	} else {
		yyv9 := &x.PodeRestartThresholds
		yym10 := z.DecBinary()
		_ = yym10
		if false {
		} else {
			*((*int32)(yyv9)) = int32(r.DecodeInt(32))
		}
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.IncludingInitContainers = false
	} else {
		yyv11 := &x.IncludingInitContainers
		yym12 := z.DecBinary()
		_ = yym12
		if false {
		} else {
			*((*bool)(yyv11)) = r.DecodeBool()
		}
	}
	for {
		yyj8++
		if yyhl8 {
			yyb8 = yyj8 > l
		} else {
			yyb8 = r.CheckBreak()
		}
		if yyb8 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj8-1, "")
	}
	r.ReadArrayEnd()
}

func (x codecSelfer1234) encStrategyList(v StrategyList, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.WriteMapStart(len(v))
	for yyk1, yyv1 := range v {
		r.WriteMapElemKey()
		yyk1.CodecEncodeSelf(e)
		r.WriteMapElemValue()
		yy3 := &yyv1
		yy3.CodecEncodeSelf(e)
	}
	r.WriteMapEnd()
}

func (x codecSelfer1234) decStrategyList(v *StrategyList, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyl1 := r.ReadMapStart()
	yybh1 := z.DecBasicHandle()
	if yyv1 == nil {
		yyrl1 := z.DecInferLen(yyl1, yybh1.MaxInitLen, 88)
		yyv1 = make(map[StrategyName]DeschedulerStrategy, yyrl1)
		*v = yyv1
	}
	var yymk1 StrategyName
	var yymv1 DeschedulerStrategy
	var yymg1, yymdn1 bool
	if yybh1.MapValueReset {
		yymg1 = true
	}
	if yyl1 != 0 {
		yyhl1 := yyl1 > 0
		for yyj1 := 0; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || r.CheckBreak()); yyj1++ {
			r.ReadMapElemKey()
			if r.TryDecodeAsNil() {
				yymk1 = ""
			} else {
				yyv2 := &yymk1
				yyv2.CodecDecodeSelf(d)
			}

			if yymg1 {
				yymv1 = yyv1[yymk1]
			} else {
				yymv1 = DeschedulerStrategy{}
			}
			r.ReadMapElemValue()
			yymdn1 = false
			if r.TryDecodeAsNil() {
				yymdn1 = true
			} else {
				yyv3 := &yymv1
				yyv3.CodecDecodeSelf(d)
			}

			if yymdn1 {
				if yybh1.DeleteOnNilMapValue {
					delete(yyv1, yymk1)
				} else {
					yyv1[yymk1] = DeschedulerStrategy{}
				}
			} else if yyv1 != nil {
				yyv1[yymk1] = yymv1
			}
		}
	} // else len==0: TODO: Should we clear map entries?
	r.ReadMapEnd()
}

func (x codecSelfer1234) encResourceThresholds(v ResourceThresholds, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.WriteMapStart(len(v))
	for yyk1, yyv1 := range v {
		r.WriteMapElemKey()
		yym2 := z.EncBinary()
		_ = yym2
		if false {
		} else if yyxt2 := z.Extension(z.I2Rtid(yyk1)); yyxt2 != nil {
			z.EncExtension(yyk1, yyxt2)
		} else {
			r.EncodeString(codecSelferCcUTF81234, string(yyk1))
		}
		r.WriteMapElemValue()
		yyv1.CodecEncodeSelf(e)
	}
	r.WriteMapEnd()
}

func (x codecSelfer1234) decResourceThresholds(v *ResourceThresholds, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyl1 := r.ReadMapStart()
	yybh1 := z.DecBasicHandle()
	if yyv1 == nil {
		yyrl1 := z.DecInferLen(yyl1, yybh1.MaxInitLen, 24)
		yyv1 = make(map[pkg2_v1.ResourceName]Percentage, yyrl1)
		*v = yyv1
	}
	var yymk1 pkg2_v1.ResourceName
	var yymv1 Percentage
	var yymg1, yymdn1 bool
	if yybh1.MapValueReset {
	}
	if yyl1 != 0 {
		yyhl1 := yyl1 > 0
		for yyj1 := 0; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || r.CheckBreak()); yyj1++ {
			r.ReadMapElemKey()
			if r.TryDecodeAsNil() {
				yymk1 = ""
			} else {
				yyv2 := &yymk1
				yym3 := z.DecBinary()
				_ = yym3
				if false {
				} else if yyxt3 := z.Extension(z.I2Rtid(yyv2)); yyxt3 != nil {
					z.DecExtension(yyv2, yyxt3)
				} else {
					*((*string)(yyv2)) = r.DecodeString()
				}
			}

			if yymg1 {
				yymv1 = yyv1[yymk1]
			}
			r.ReadMapElemValue()
			yymdn1 = false
			if r.TryDecodeAsNil() {
				yymdn1 = true
			} else {
				yyv4 := &yymv1
				yyv4.CodecDecodeSelf(d)
			}

			if yymdn1 {
				if yybh1.DeleteOnNilMapValue {
					delete(yyv1, yymk1)
				} else {
					yyv1[yymk1] = 0
				}
			} else if yyv1 != nil {
				yyv1[yymk1] = yymv1
			}
		}
	} // else len==0: TODO: Should we clear map entries?
	r.ReadMapEnd()
}
