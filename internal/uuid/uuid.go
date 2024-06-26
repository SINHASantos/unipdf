//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package uuid ;import (_ee "crypto/rand";_af "encoding/hex";_a "io";);var Nil =_eee ;var _c =_ee .Reader ;type UUID [16]byte ;func MustUUID ()UUID {uuid ,_cb :=NewUUID ();if _cb !=nil {panic (_cb );};return uuid ;};func (_afd UUID )String ()string {var _ec [36]byte ;
_g (_ec [:],_afd );return string (_ec [:])};func NewUUID ()(UUID ,error ){var uuid UUID ;_ ,_b :=_a .ReadFull (_c ,uuid [:]);if _b !=nil {return _eee ,_b ;};uuid [6]=(uuid [6]&0x0f)|0x40;uuid [8]=(uuid [8]&0x3f)|0x80;return uuid ,nil ;};func _g (_gc []byte ,_eb UUID ){_af .Encode (_gc ,_eb [:4]);
_gc [8]='-';_af .Encode (_gc [9:13],_eb [4:6]);_gc [13]='-';_af .Encode (_gc [14:18],_eb [6:8]);_gc [18]='-';_af .Encode (_gc [19:23],_eb [8:10]);_gc [23]='-';_af .Encode (_gc [24:],_eb [10:]);};var _eee UUID ;