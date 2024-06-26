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

package security ;import (_f "bytes";_ca "crypto/aes";_gg "crypto/cipher";_ge "crypto/md5";_dg "crypto/rand";_eb "crypto/rc4";_c "crypto/sha256";_e "crypto/sha512";_d "encoding/binary";_ed "errors";_fa "fmt";_fe "github.com/unidoc/unipdf/v3/common";_ad "hash";
_a "io";_b "math";);

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_bbe stdHandlerR4 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){O ,_cga :=_bbe .alg3 (d .R ,upass ,opass );if _cga !=nil {_fe .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_cga );
return nil ,_cga ;};d .O =O ;_fe .Log .Trace ("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078",O );_dgbc :=_bbe .alg2 (d ,upass );U ,_cga :=_bbe .alg5 (_dgbc ,upass );if _cga !=nil {_fe .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_cga );
return nil ,_cga ;};d .U =U ;_fe .Log .Trace ("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078",U );return _dgbc ,nil ;};func _fb (_cb _gg .Block )*ecb {return &ecb {_faf :_cb ,_cc :_cb .BlockSize ()}};var _ StdHandler =stdHandlerR4 {};func (_db *ecbDecrypter )BlockSize ()int {return _db ._cc };
func _agd (_fcg []byte ,_abe int ){_fed :=_abe ;for _fed < len (_fcg ){copy (_fcg [_fed :],_fcg [:_fed ]);_fed *=2;};};const (PermOwner =Permissions (_b .MaxUint32 );PermPrinting =Permissions (1<<2);PermModify =Permissions (1<<3);PermExtractGraphics =Permissions (1<<4);
PermAnnotate =Permissions (1<<5);PermFillForms =Permissions (1<<8);PermDisabilityExtract =Permissions (1<<9);PermRotateInsert =Permissions (1<<10);PermFullPrintQuality =Permissions (1<<11););type ecbEncrypter ecb ;func _gc (_fbe _gg .Block )_gg .BlockMode {return (*ecbEncrypter )(_fb (_fbe ))};
func (_ag stdHandlerR4 )alg7 (_gec *StdEncryptDict ,_fcc []byte )([]byte ,error ){_ab :=_ag .alg3Key (_gec .R ,_fcc );_gcg :=make ([]byte ,len (_gec .O ));if _gec .R ==2{_abd ,_bfe :=_eb .NewCipher (_ab );if _bfe !=nil {return nil ,_ed .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");
};_abd .XORKeyStream (_gcg ,_gec .O );}else if _gec .R >=3{_acbe :=append ([]byte {},_gec .O ...);for _aea :=0;_aea < 20;_aea ++{_cgfc :=append ([]byte {},_ab ...);for _abdd :=0;_abdd < len (_ab );_abdd ++{_cgfc [_abdd ]^=byte (19-_aea );};_dbe ,_fbd :=_eb .NewCipher (_cgfc );
if _fbd !=nil {return nil ,_ed .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");};_dbe .XORKeyStream (_gcg ,_acbe );_acbe =append ([]byte {},_gcg ...);};}else {return nil ,_ed .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");
};_gag ,_bad :=_ag .alg6 (_gec ,_gcg );if _bad !=nil {return nil ,nil ;};return _gag ,nil ;};

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4 (id0 string ,length int )StdHandler {return stdHandlerR4 {ID0 :id0 ,Length :length }};func (_gcd errInvalidField )Error ()string {return _fa .Sprintf ("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064",_gcd .Func ,_gcd .Field ,_gcd .Exp ,_gcd .Got );
};func _de (_eab []byte )(_gg .Block ,error ){_aeeg ,_dgd :=_ca .NewCipher (_eab );if _dgd !=nil {_fe .Log .Error ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076",_dgd );
return nil ,_dgd ;};return _aeeg ,nil ;};func (_feg stdHandlerR4 )alg4 (_bbg []byte ,_fda []byte )([]byte ,error ){_fdd ,_ga :=_eb .NewCipher (_bbg );if _ga !=nil {return nil ,_ed .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_dcdg :=[]byte (_fgb );_edd :=make ([]byte ,len (_dcdg ));_fdd .XORKeyStream (_edd ,_dcdg );return _edd ,nil ;};

// StdEncryptDict is a set of additional fields used in standard encryption dictionary.
type StdEncryptDict struct{R int ;P Permissions ;EncryptMetadata bool ;O ,U []byte ;OE ,UE []byte ;Perms []byte ;};func (stdHandlerR4 )paddedPass (_fc []byte )[]byte {_ae :=make ([]byte ,32);_fd :=copy (_ae ,_fc );for ;_fd < 32;_fd ++{_ae [_fd ]=_fgb [_fd -len (_fc )];
};return _ae ;};func (_cg *ecbDecrypter )CryptBlocks (dst ,src []byte ){if len (src )%_cg ._cc !=0{_fe .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_fe .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_cg ._faf .Decrypt (dst ,src [:_cg ._cc ]);src =src [_cg ._cc :];dst =dst [_cg ._cc :];};};func (_ddfe stdHandlerR6 )alg11 (_bdga *StdEncryptDict ,_ddb []byte )([]byte ,error ){if _beec :=_fg ("\u0061\u006c\u00671\u0031","\u0055",48,_bdga .U );
_beec !=nil {return nil ,_beec ;};_fac :=make ([]byte ,len (_ddb )+8);_gdg :=copy (_fac ,_ddb );_gdg +=copy (_fac [_gdg :],_bdga .U [32:40]);_ecgbc ,_bbd :=_ddfe .alg2b (_bdga .R ,_fac ,_ddb ,nil );if _bbd !=nil {return nil ,_bbd ;};_ecgbc =_ecgbc [:32];
if !_f .Equal (_ecgbc ,_bdga .U [:32]){return nil ,nil ;};return _ecgbc ,nil ;};

// Authenticate implements StdHandler interface.
func (_ccf stdHandlerR4 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){_fe .Log .Trace ("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073");
_fddb ,_gf :=_ccf .alg7 (d ,pass );if _gf !=nil {return nil ,0,_gf ;};if _fddb !=nil {_fe .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _fddb ,PermOwner ,nil ;
};_fe .Log .Trace ("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073");_fddb ,_gf =_ccf .alg6 (d ,pass );if _gf !=nil {return nil ,0,_gf ;};
if _fddb !=nil {_fe .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _fddb ,d .P ,nil ;};return nil ,0,nil ;};type ecb struct{_faf _gg .Block ;_cc int ;};

// AuthEvent is an event type that triggers authentication.
type AuthEvent string ;func (_ecgb stdHandlerR6 )alg8 (_dbge *StdEncryptDict ,_egb []byte ,_egd []byte )error {if _ege :=_fg ("\u0061\u006c\u0067\u0038","\u004b\u0065\u0079",32,_egb );_ege !=nil {return _ege ;};var _ebg [16]byte ;if _ ,_afba :=_a .ReadFull (_dg .Reader ,_ebg [:]);
_afba !=nil {return _afba ;};_dcae :=_ebg [0:8];_gba :=_ebg [8:16];_bgd :=make ([]byte ,len (_egd )+len (_dcae ));_bgb :=copy (_bgd ,_egd );copy (_bgd [_bgb :],_dcae );_bcb ,_afd :=_ecgb .alg2b (_dbge .R ,_bgd ,_egd ,nil );if _afd !=nil {return _afd ;};
U :=make ([]byte ,len (_bcb )+len (_dcae )+len (_gba ));_bgb =copy (U ,_bcb [:32]);_bgb +=copy (U [_bgb :],_dcae );copy (U [_bgb :],_gba );_dbge .U =U ;_bgb =len (_egd );copy (_bgd [_bgb :],_gba );_bcb ,_afd =_ecgb .alg2b (_dbge .R ,_bgd ,_egd ,nil );if _afd !=nil {return _afd ;
};_gbb ,_afd :=_de (_bcb [:32]);if _afd !=nil {return _afd ;};_egca :=make ([]byte ,_ca .BlockSize );_aeee :=_gg .NewCBCEncrypter (_gbb ,_egca );UE :=make ([]byte ,32);_aeee .CryptBlocks (UE ,_egb [:32]);_dbge .UE =UE ;return nil ;};

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_fcf stdHandlerR6 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){_facc :=make ([]byte ,32);if _ ,_fafa :=_a .ReadFull (_dg .Reader ,_facc );_fafa !=nil {return nil ,_fafa ;};d .U =nil ;d .O =nil ;d .UE =nil ;d .OE =nil ;
d .Perms =nil ;if len (upass )> 127{upass =upass [:127];};if len (opass )> 127{opass =opass [:127];};if _gaga :=_fcf .alg8 (d ,_facc ,upass );_gaga !=nil {return nil ,_gaga ;};if _cac :=_fcf .alg9 (d ,_facc ,opass );_cac !=nil {return nil ,_cac ;};if d .R ==5{return _facc ,nil ;
};if _ggcf :=_fcf .alg10 (d ,_facc );_ggcf !=nil {return nil ,_ggcf ;};return _facc ,nil ;};func (_bag stdHandlerR6 )alg10 (_agg *StdEncryptDict ,_ffbg []byte )error {if _bec :=_fg ("\u0061\u006c\u00671\u0030","\u004b\u0065\u0079",32,_ffbg );_bec !=nil {return _bec ;
};_fge :=uint64 (uint32 (_agg .P ))|(_b .MaxUint32 <<32);Perms :=make ([]byte ,16);_d .LittleEndian .PutUint64 (Perms [:8],_fge );if _agg .EncryptMetadata {Perms [8]='T';}else {Perms [8]='F';};copy (Perms [9:12],"\u0061\u0064\u0062");if _ ,_acf :=_a .ReadFull (_dg .Reader ,Perms [12:16]);
_acf !=nil {return _acf ;};_fbg ,_cbff :=_de (_ffbg [:32]);if _cbff !=nil {return _cbff ;};_eeb :=_gc (_fbg );_eeb .CryptBlocks (Perms ,Perms );_agg .Perms =Perms [:16];return nil ;};type stdHandlerR4 struct{Length int ;ID0 string ;};func (_bfb stdHandlerR4 )alg3 (R int ,_cgb ,_faff []byte )([]byte ,error ){var _geb []byte ;
if len (_faff )> 0{_geb =_bfb .alg3Key (R ,_faff );}else {_geb =_bfb .alg3Key (R ,_cgb );};_af ,_caa :=_eb .NewCipher (_geb );if _caa !=nil {return nil ,_ed .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_gb :=_bfb .paddedPass (_cgb );
_gee :=make ([]byte ,len (_gb ));_af .XORKeyStream (_gee ,_gb );if R >=3{_dgb :=make ([]byte ,len (_geb ));for _gcff :=0;_gcff < 19;_gcff ++{for _da :=0;_da < len (_geb );_da ++{_dgb [_da ]=_geb [_da ]^byte (_gcff +1);};_bdg ,_cgf :=_eb .NewCipher (_dgb );
if _cgf !=nil {return nil ,_ed .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_bdg .XORKeyStream (_gee ,_gee );};};return _gee ,nil ;};type stdHandlerR6 struct{};func (_ddd stdHandlerR6 )alg12 (_dagb *StdEncryptDict ,_fbec []byte )([]byte ,error ){if _ecc :=_fg ("\u0061\u006c\u00671\u0032","\u0055",48,_dagb .U );
_ecc !=nil {return nil ,_ecc ;};if _eca :=_fg ("\u0061\u006c\u00671\u0032","\u004f",48,_dagb .O );_eca !=nil {return nil ,_eca ;};_fecg :=make ([]byte ,len (_fbec )+8+48);_ddbb :=copy (_fecg ,_fbec );_ddbb +=copy (_fecg [_ddbb :],_dagb .O [32:40]);_ddbb +=copy (_fecg [_ddbb :],_dagb .U [0:48]);
_aff ,_bdf :=_ddd .alg2b (_dagb .R ,_fecg ,_fbec ,_dagb .U [0:48]);if _bdf !=nil {return nil ,_bdf ;};_aff =_aff [:32];if !_f .Equal (_aff ,_dagb .O [:32]){return nil ,nil ;};return _aff ,nil ;};func (_ebe stdHandlerR4 )alg5 (_fdad []byte ,_aee []byte )([]byte ,error ){_fbf :=_ge .New ();
_fbf .Write ([]byte (_fgb ));_fbf .Write ([]byte (_ebe .ID0 ));_ea :=_fbf .Sum (nil );_fe .Log .Trace ("\u0061\u006c\u0067\u0035");_fe .Log .Trace ("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x",_fdad );_fe .Log .Trace ("\u0049D\u003a\u0020\u0025\u0020\u0078",_ebe .ID0 );
if len (_ea )!=16{return nil ,_ed .New ("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073");};_bee ,_dce :=_eb .NewCipher (_fdad );if _dce !=nil {return nil ,_ed .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_afb :=make ([]byte ,16);_bee .XORKeyStream (_afb ,_ea );_gbc :=make ([]byte ,len (_fdad ));for _eaa :=0;_eaa < 19;_eaa ++{for _eef :=0;_eef < len (_fdad );_eef ++{_gbc [_eef ]=_fdad [_eef ]^byte (_eaa +1);};_bee ,_dce =_eb .NewCipher (_gbc );if _dce !=nil {return nil ,_ed .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_bee .XORKeyStream (_afb ,_afb );_fe .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078",_eaa ,_gbc );_fe .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078",_eaa ,_afb );
};_fff :=make ([]byte ,32);for _cec :=0;_cec < 16;_cec ++{_fff [_cec ]=_afb [_cec ];};_ ,_dce =_dg .Read (_fff [16:32]);if _dce !=nil {return nil ,_ed .New ("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r");
};return _fff ,nil ;};func (_eed stdHandlerR4 )alg3Key (R int ,_ba []byte )[]byte {_aef :=_ge .New ();_cce :=_eed .paddedPass (_ba );_aef .Write (_cce );if R >=3{for _dcd :=0;_dcd < 50;_dcd ++{_gga :=_aef .Sum (nil );_aef =_ge .New ();_aef .Write (_gga );
};};_fgg :=_aef .Sum (nil );if R ==2{_fgg =_fgg [0:5];}else {_fgg =_fgg [0:_eed .Length /8];};return _fgg ;};

// Allowed checks if a set of permissions can be granted.
func (_ce Permissions )Allowed (p2 Permissions )bool {return _ce &p2 ==p2 };var _ StdHandler =stdHandlerR6 {};

// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32 ;func _ffed (_fad []byte )([]byte ,error ){_gded :=_c .New ();_gded .Write (_fad );return _gded .Sum (nil ),nil ;};func _bb (_ccb _gg .Block )_gg .BlockMode {return (*ecbDecrypter )(_fb (_ccb ))};func (_cbb stdHandlerR6 )alg9 (_beaa *StdEncryptDict ,_babb []byte ,_gcb []byte )error {if _fde :=_fg ("\u0061\u006c\u0067\u0039","\u004b\u0065\u0079",32,_babb );
_fde !=nil {return _fde ;};if _fdb :=_fg ("\u0061\u006c\u0067\u0039","\u0055",48,_beaa .U );_fdb !=nil {return _fdb ;};var _fbc [16]byte ;if _ ,_cge :=_a .ReadFull (_dg .Reader ,_fbc [:]);_cge !=nil {return _cge ;};_gfd :=_fbc [0:8];_daf :=_fbc [8:16];
_cd :=_beaa .U [:48];_edb :=make ([]byte ,len (_gcb )+len (_gfd )+len (_cd ));_gdf :=copy (_edb ,_gcb );_gdf +=copy (_edb [_gdf :],_gfd );_gdf +=copy (_edb [_gdf :],_cd );_abdg ,_ded :=_cbb .alg2b (_beaa .R ,_edb ,_gcb ,_cd );if _ded !=nil {return _ded ;
};O :=make ([]byte ,len (_abdg )+len (_gfd )+len (_daf ));_gdf =copy (O ,_abdg [:32]);_gdf +=copy (O [_gdf :],_gfd );_gdf +=copy (O [_gdf :],_daf );_beaa .O =O ;_gdf =len (_gcb );_gdf +=copy (_edb [_gdf :],_daf );_abdg ,_ded =_cbb .alg2b (_beaa .R ,_edb ,_gcb ,_cd );
if _ded !=nil {return _ded ;};_ffb ,_ded :=_de (_abdg [:32]);if _ded !=nil {return _ded ;};_bbc :=make ([]byte ,_ca .BlockSize );_gbe :=_gg .NewCBCEncrypter (_ffb ,_bbc );OE :=make ([]byte ,32);_gbe .CryptBlocks (OE ,_babb [:32]);_beaa .OE =OE ;return nil ;
};

// Authenticate implements StdHandler interface.
func (_gdfg stdHandlerR6 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){return _gdfg .alg2a (d ,pass );};

// StdHandler is an interface for standard security handlers.
type StdHandler interface{

// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
// It assumes that R, P and EncryptMetadata are already set.
GenerateParams (_be *StdEncryptDict ,_bc ,_ade []byte )([]byte ,error );

// Authenticate uses encryption dictionary parameters and the password to calculate
// the document encryption key. It also returns permissions that should be granted to a user.
// In case of failed authentication, it returns empty key and zero permissions with no error.
Authenticate (_gd *StdEncryptDict ,_ef []byte )([]byte ,Permissions ,error );};func (_ffc stdHandlerR6 )alg13 (_ggc *StdEncryptDict ,_bfee []byte )error {if _aag :=_fg ("\u0061\u006c\u00671\u0033","\u004b\u0065\u0079",32,_bfee );_aag !=nil {return _aag ;
};if _edca :=_fg ("\u0061\u006c\u00671\u0033","\u0050\u0065\u0072m\u0073",16,_ggc .Perms );_edca !=nil {return _edca ;};_aafc :=make ([]byte ,16);copy (_aafc ,_ggc .Perms [:16]);_fee ,_cdb :=_ca .NewCipher (_bfee [:32]);if _cdb !=nil {return _cdb ;};_cfg :=_bb (_fee );
_cfg .CryptBlocks (_aafc ,_aafc );if !_f .Equal (_aafc [9:12],[]byte ("\u0061\u0064\u0062")){return _ed .New ("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};_aeg :=Permissions (_d .LittleEndian .Uint32 (_aafc [0:4]));if _aeg !=_ggc .P {return _ed .New ("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064");
};var _gda bool ;if _aafc [8]=='T'{_gda =true ;}else if _aafc [8]=='F'{_gda =false ;}else {return _ed .New ("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};if _gda !=_ggc .EncryptMetadata {return _ed .New ("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064");
};return nil ;};func (_dc *ecbEncrypter )CryptBlocks (dst ,src []byte ){if len (src )%_dc ._cc !=0{_fe .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_fe .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_dc ._faf .Encrypt (dst ,src [:_dc ._cc ]);src =src [_dc ._cc :];dst =dst [_dc ._cc :];};};func (_ccbf stdHandlerR6 )alg2b (R int ,_dae ,_cedg ,_gdd []byte )([]byte ,error ){if R ==5{return _ffed (_dae );};return _cf (_dae ,_cedg ,_gdd );
};func _cf (_ecg ,_fec ,_cae []byte )([]byte ,error ){var (_bdb ,_agde ,_ggg _ad .Hash ;);_bdb =_c .New ();_bafc :=make ([]byte ,64);_ged :=_bdb ;_ged .Write (_ecg );K :=_ged .Sum (_bafc [:0]);_ddf :=make ([]byte ,64*(127+64+48));_aaf :=func (_bea int )([]byte ,error ){_bcf :=len (_fec )+len (K )+len (_cae );
_dbc :=_ddf [:_bcf ];_fffb :=copy (_dbc ,_fec );_fffb +=copy (_dbc [_fffb :],K [:]);_fffb +=copy (_dbc [_fffb :],_cae );if _fffb !=_bcf {_fe .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e");
return nil ,_ed .New ("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065");};K1 :=_ddf [:_bcf *64];_agd (K1 ,_bcf );_bfa ,_fgga :=_de (K [0:16]);if _fgga !=nil {return nil ,_fgga ;};_fcge :=_gg .NewCBCEncrypter (_bfa ,K [16:32]);_fcge .CryptBlocks (K1 ,K1 );
E :=K1 ;_bae :=0;for _ccg :=0;_ccg < 16;_ccg ++{_bae +=int (E [_ccg ]%3);};var _dag _ad .Hash ;switch _bae %3{case 0:_dag =_bdb ;case 1:if _agde ==nil {_agde =_e .New384 ();};_dag =_agde ;case 2:if _ggg ==nil {_ggg =_e .New ();};_dag =_ggg ;};_dag .Reset ();
_dag .Write (E );K =_dag .Sum (_bafc [:0]);return E ,nil ;};for _egc :=0;;{E ,_beg :=_aaf (_egc );if _beg !=nil {return nil ,_beg ;};_gebg :=E [len (E )-1];_egc ++;if _egc >=64&&_gebg <=uint8 (_egc -32){break ;};};return K [:32],nil ;};func (_fca stdHandlerR4 )alg2 (_bff *StdEncryptDict ,_acb []byte )[]byte {_fe .Log .Trace ("\u0061\u006c\u0067\u0032");
_ee :=_fca .paddedPass (_acb );_bd :=_ge .New ();_bd .Write (_ee );_bd .Write (_bff .O );var _dd [4]byte ;_d .LittleEndian .PutUint32 (_dd [:],uint32 (_bff .P ));_bd .Write (_dd [:]);_fe .Log .Trace ("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x",_dd );
_bd .Write ([]byte (_fca .ID0 ));_fe .Log .Trace ("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076",_bff .R ,_bff .EncryptMetadata );
if (_bff .R >=4)&&!_bff .EncryptMetadata {_bd .Write ([]byte {0xff,0xff,0xff,0xff});};_cca :=_bd .Sum (nil );if _bff .R >=3{_bd =_ge .New ();for _dca :=0;_dca < 50;_dca ++{_bd .Reset ();_bd .Write (_cca [0:_fca .Length /8]);_cca =_bd .Sum (nil );};};if _bff .R >=3{return _cca [0:_fca .Length /8];
};return _cca [0:5];};func (_ced stdHandlerR4 )alg6 (_gbg *StdEncryptDict ,_dgcc []byte )([]byte ,error ){var (_ffe []byte ;_dbg error ;);_edc :=_ced .alg2 (_gbg ,_dgcc );if _gbg .R ==2{_ffe ,_dbg =_ced .alg4 (_edc ,_dgcc );}else if _gbg .R >=3{_ffe ,_dbg =_ced .alg5 (_edc ,_dgcc );
}else {return nil ,_ed .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");};if _dbg !=nil {return nil ,_dbg ;};_fe .Log .Trace ("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f",string (_ffe ),string (_gbg .U ));
_dgg :=_ffe ;_ddc :=_gbg .U ;if _gbg .R >=3{if len (_dgg )> 16{_dgg =_dgg [0:16];};if len (_ddc )> 16{_ddc =_ddc [0:16];};};if !_f .Equal (_dgg ,_ddc ){return nil ,nil ;};return _edc ,nil ;};func (_fgd stdHandlerR6 )alg2a (_ada *StdEncryptDict ,_eg []byte )([]byte ,Permissions ,error ){if _cbf :=_fg ("\u0061\u006c\u00672\u0061","\u004f",48,_ada .O );
_cbf !=nil {return nil ,0,_cbf ;};if _cab :=_fg ("\u0061\u006c\u00672\u0061","\u0055",48,_ada .U );_cab !=nil {return nil ,0,_cab ;};if len (_eg )> 127{_eg =_eg [:127];};_dea ,_fgf :=_fgd .alg12 (_ada ,_eg );if _fgf !=nil {return nil ,0,_fgf ;};var (_bdc []byte ;
_bg []byte ;_gde []byte ;);var _ec Permissions ;if len (_dea )!=0{_ec =PermOwner ;_cgfd :=make ([]byte ,len (_eg )+8+48);_dab :=copy (_cgfd ,_eg );_dab +=copy (_cgfd [_dab :],_ada .O [40:48]);copy (_cgfd [_dab :],_ada .U [0:48]);_bdc =_cgfd ;_bg =_ada .OE ;
_gde =_ada .U [0:48];}else {_dea ,_fgf =_fgd .alg11 (_ada ,_eg );if _fgf ==nil &&len (_dea )==0{_dea ,_fgf =_fgd .alg11 (_ada ,[]byte (""));};if _fgf !=nil {return nil ,0,_fgf ;}else if len (_dea )==0{return nil ,0,nil ;};_ec =_ada .P ;_dda :=make ([]byte ,len (_eg )+8);
_gae :=copy (_dda ,_eg );copy (_dda [_gae :],_ada .U [40:48]);_bdc =_dda ;_bg =_ada .UE ;_gde =nil ;};if _cbfc :=_fg ("\u0061\u006c\u00672\u0061","\u004b\u0065\u0079",32,_bg );_cbfc !=nil {return nil ,0,_cbfc ;};_bg =_bg [:32];_baf ,_fgf :=_fgd .alg2b (_ada .R ,_bdc ,_eg ,_gde );
if _fgf !=nil {return nil ,0,_fgf ;};_adad ,_fgf :=_ca .NewCipher (_baf [:32]);if _fgf !=nil {return nil ,0,_fgf ;};_bba :=make ([]byte ,_ca .BlockSize );_afbd :=_gg .NewCBCDecrypter (_adad ,_bba );_gfc :=make ([]byte ,32);_afbd .CryptBlocks (_gfc ,_bg );
if _ada .R ==5{return _gfc ,_ec ,nil ;};_fgf =_fgd .alg13 (_ada ,_gfc );if _fgf !=nil {return nil ,0,_fgf ;};return _gfc ,_ec ,nil ;};func _fg (_ff ,_cba string ,_ac int ,_bf []byte )error {if len (_bf )< _ac {return errInvalidField {Func :_ff ,Field :_cba ,Exp :_ac ,Got :len (_bf )};
};return nil ;};func (_gcf *ecbEncrypter )BlockSize ()int {return _gcf ._cc };type ecbDecrypter ecb ;type errInvalidField struct{Func string ;Field string ;Exp int ;Got int ;};const (EventDocOpen =AuthEvent ("\u0044o\u0063\u004f\u0070\u0065\u006e");EventEFOpen =AuthEvent ("\u0045\u0046\u004f\u0070\u0065\u006e");
);

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6 ()StdHandler {return stdHandlerR6 {}};const _fgb ="\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377"+"\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014"+"\251\xfe\x64\x53\x69\172";