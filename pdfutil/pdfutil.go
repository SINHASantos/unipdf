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

package pdfutil ;import (_b "github.com/unidoc/unipdf/v3/common";_dc "github.com/unidoc/unipdf/v3/contentstream";_bc "github.com/unidoc/unipdf/v3/contentstream/draw";_f "github.com/unidoc/unipdf/v3/core";_a "github.com/unidoc/unipdf/v3/model";);

// NormalizePage performs the following operations on the passed in page:
//   - Normalize the page rotation.
//     Rotates the contents of the page according to the Rotate entry, thus
//     flattening the rotation. The Rotate entry of the page is set to nil.
//   - Normalize the media box.
//     If the media box of the page is offsetted (Llx != 0 or Lly != 0),
//     the contents of the page are translated to (-Llx, -Lly). After
//     normalization, the media box is updated (Llx and Lly are set to 0 and
//     Urx and Ury are updated accordingly).
//   - Normalize the crop box.
//     The crop box of the page is updated based on the previous operations.
//
// After normalization, the page should look the same if openend using a
// PDF viewer.
// NOTE: This function does not normalize annotations, outlines other parts
// that are not part of the basic geometry and page content streams.
func NormalizePage (page *_a .PdfPage )error {_dcd ,_ag :=page .GetMediaBox ();if _ag !=nil {return _ag ;};_fd ,_ag :=page .GetRotate ();if _ag !=nil {_b .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a",_ag .Error ());
};_df :=_fd %360!=0&&_fd %90==0;_dcd .Normalize ();_db ,_aa ,_fda ,_g :=_dcd .Llx ,_dcd .Lly ,_dcd .Width (),_dcd .Height ();_dg :=_db !=0||_aa !=0;if !_df &&!_dg {return nil ;};_e :=func (_de ,_fg ,_bg float64 )_bc .BoundingBox {return _bc .Path {Points :[]_bc .Point {_bc .NewPoint (0,0).Rotate (_bg ),_bc .NewPoint (_de ,0).Rotate (_bg ),_bc .NewPoint (0,_fg ).Rotate (_bg ),_bc .NewPoint (_de ,_fg ).Rotate (_bg )}}.GetBoundingBox ();
};_bf :=_dc .NewContentCreator ();var _eg float64 ;if _df {_eg =-float64 (_fd );_dfa :=_e (_fda ,_g ,_eg );_bf .Translate ((_dfa .Width -_fda )/2+_fda /2,(_dfa .Height -_g )/2+_g /2);_bf .RotateDeg (_eg );_bf .Translate (-_fda /2,-_g /2);_fda ,_g =_dfa .Width ,_dfa .Height ;
};if _dg {_bf .Translate (-_db ,-_aa );};_bd :=_bf .Operations ();_ec ,_ag :=_f .MakeStream (_bd .Bytes (),_f .NewFlateEncoder ());if _ag !=nil {return _ag ;};_ad :=_f .MakeArray (_ec );_ad .Append (page .GetContentStreamObjs ()...);*_dcd =_a .PdfRectangle {Urx :_fda ,Ury :_g };
if _c :=page .CropBox ;_c !=nil {_c .Normalize ();_ecg ,_ab ,_ae ,_ff :=_c .Llx -_db ,_c .Lly -_aa ,_c .Width (),_c .Height ();if _df {_bgg :=_e (_ae ,_ff ,_eg );_ae ,_ff =_bgg .Width ,_bgg .Height ;};*_c =_a .PdfRectangle {Llx :_ecg ,Lly :_ab ,Urx :_ecg +_ae ,Ury :_ab +_ff };
};_b .Log .Debug ("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066",_eg ,_bd ,_dcd );page .Contents =_ad ;page .Rotate =nil ;
return nil ;};