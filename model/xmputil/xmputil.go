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

// Package xmputil provides abstraction used by the pdf document XMP Metadata.
package xmputil ;import (_ec "errors";_bc "fmt";_eg "github.com/trimmer-io/go-xmp/models/pdf";_d "github.com/trimmer-io/go-xmp/models/xmp_base";_ca "github.com/trimmer-io/go-xmp/models/xmp_mm";_e "github.com/trimmer-io/go-xmp/xmp";_ga "github.com/unidoc/unipdf/v3/core";
_eb "github.com/unidoc/unipdf/v3/internal/timeutils";_gd "github.com/unidoc/unipdf/v3/internal/uuid";_ea "github.com/unidoc/unipdf/v3/model/xmputil/pdfaextension";_b "github.com/unidoc/unipdf/v3/model/xmputil/pdfaid";_ee "strconv";_c "time";);

// NewDocument creates a new document without any previous xmp information.
func NewDocument ()*Document {_dc :=_e .NewDocument ();return &Document {_egd :_dc }};

// MediaManagementDerivedFrom is a structure that contains references of identifiers and versions
// from which given document was derived.
type MediaManagementDerivedFrom struct{OriginalDocumentID GUID ;DocumentID GUID ;InstanceID GUID ;VersionID string ;};

// MediaManagement are the values from the document media management metadata.
type MediaManagement struct{

// OriginalDocumentID  as media is imported and projects is started, an original-document ID
// must be created to identify a new document. This identifies a document as a conceptual entity.
OriginalDocumentID GUID ;

// DocumentID when a document is copied to a new file path or converted to a new format with
// Save As, another new document ID should usually be assigned. This identifies a general version or
// branch of a document. You can use it to track different versions or extracted portions of a document
// with the same original-document ID.
DocumentID GUID ;

// InstanceID to track a document’s editing history, you must assign a new instance ID
// whenever a document is saved after any changes. This uniquely identifies an exact version of a
// document. It is used in resource references (to identify both the document or part itself and the
// referenced or referencing documents), and in document-history resource events (to identify the
// document instance that resulted from the change).
InstanceID GUID ;

// DerivedFrom references the source document from which this one is derived,
// typically through a Save As operation that changes the file name or format. It is a minimal reference;
// missing components can be assumed to be unchanged. For example, a new version might only need
// to specify the instance ID and version number of the previous version, or a rendition might only need
// to specify the instance ID and rendition class of the original.
DerivedFrom *MediaManagementDerivedFrom ;

// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
// document history, but should not be used to identify a document uniquely in any context.
// Usually it simply works by incrementing integers 1,2,3...
VersionID string ;

// Versions is the history of the document versions along with the comments, timestamps and issuers.
Versions []MediaManagementVersion ;};

// GetGoXmpDocument gets direct access to the go-xmp.Document.
// All changes done to specified document would result in change of this document 'd'.
func (_dcb *Document )GetGoXmpDocument ()*_e .Document {return _dcb ._egd };

// PdfInfoOptions are the options used for setting pdf info.
type PdfInfoOptions struct{InfoDict _ga .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;

// Overwrite if set to true, overwrites all values found in the current pdf info xmp model to the ones provided.
Overwrite bool ;};

// SetPdfAID sets up pdfaid xmp metadata.
// In example: Part: '1' Conformance: 'B' states for PDF/A 1B.
func (_ag *Document )SetPdfAID (part int ,conformance string )error {_cde ,_adf :=_b .MakeModel (_ag ._egd );if _adf !=nil {return _adf ;};_cde .Part =part ;_cde .Conformance =conformance ;if _acb :=_cde .SyncToXMP (_ag ._egd );_acb !=nil {return _acb ;
};return nil ;};

// PdfAID is the result of the XMP pdfaid metadata.
type PdfAID struct{Part int ;Conformance string ;};

// MediaManagementVersion is the version of the media management xmp metadata.
type MediaManagementVersion struct{VersionID string ;ModifyDate _c .Time ;Comments string ;Modifier string ;};

// GetMediaManagement gets the media management metadata from provided xmp document.
func (_ebe *Document )GetMediaManagement ()(*MediaManagement ,bool ){_dcc :=_ca .FindModel (_ebe ._egd );if _dcc ==nil {return nil ,false ;};_ffc :=make ([]MediaManagementVersion ,len (_dcc .Versions ));for _fedb ,_ceb :=range _dcc .Versions {_ffc [_fedb ]=MediaManagementVersion {VersionID :_ceb .Version ,ModifyDate :_ceb .ModifyDate .Value (),Comments :_ceb .Comments ,Modifier :_ceb .Modifier };
};_efb :=&MediaManagement {OriginalDocumentID :GUID (_dcc .OriginalDocumentID .Value ()),DocumentID :GUID (_dcc .DocumentID .Value ()),InstanceID :GUID (_dcc .InstanceID .Value ()),VersionID :_dcc .VersionID ,Versions :_ffc };if _dcc .DerivedFrom !=nil {_efb .DerivedFrom =&MediaManagementDerivedFrom {OriginalDocumentID :GUID (_dcc .DerivedFrom .OriginalDocumentID ),DocumentID :GUID (_dcc .DerivedFrom .DocumentID ),InstanceID :GUID (_dcc .DerivedFrom .InstanceID ),VersionID :_dcc .DerivedFrom .VersionID };
};return _efb ,true ;};

// SetPdfInfo sets the pdf info into selected document.
func (_ba *Document )SetPdfInfo (options *PdfInfoOptions )error {if options ==nil {return _ec .New ("\u006ei\u006c\u0020\u0070\u0064\u0066\u0020\u006f\u0070\u0074\u0069\u006fn\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};_ad ,_cc :=_eg .MakeModel (_ba ._egd );
if _cc !=nil {return _cc ;};if options .Overwrite {*_ad =_eg .PDFInfo {};};if options .InfoDict !=nil {_ccd ,_ff :=_ga .GetDict (options .InfoDict );if !_ff {return _bc .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",options .InfoDict );
};var _dcg *_ga .PdfObjectString ;for _ ,_bcb :=range _ccd .Keys (){switch _bcb {case "\u0054\u0069\u0074l\u0065":_dcg ,_ff =_ga .GetString (_ccd .Get ("\u0054\u0069\u0074l\u0065"));if _ff {_ad .Title =_e .NewAltString (_dcg );};case "\u0041\u0075\u0074\u0068\u006f\u0072":_dcg ,_ff =_ga .GetString (_ccd .Get ("\u0041\u0075\u0074\u0068\u006f\u0072"));
if _ff {_ad .Author =_e .NewStringList (_dcg .String ());};case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":_dcg ,_ff =_ga .GetString (_ccd .Get ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"));if _ff {_ad .Keywords =_dcg .String ();};case "\u0043r\u0065\u0061\u0074\u006f\u0072":_dcg ,_ff =_ga .GetString (_ccd .Get ("\u0043r\u0065\u0061\u0074\u006f\u0072"));
if _ff {_ad .Creator =_e .AgentName (_dcg .String ());};case "\u0053u\u0062\u006a\u0065\u0063\u0074":_dcg ,_ff =_ga .GetString (_ccd .Get ("\u0053u\u0062\u006a\u0065\u0063\u0074"));if _ff {_ad .Subject =_e .NewAltString (_dcg .String ());};case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":_dcg ,_ff =_ga .GetString (_ccd .Get ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"));
if _ff {_ad .Producer =_e .AgentName (_dcg .String ());};case "\u0054r\u0061\u0070\u0070\u0065\u0064":_egf ,_cad :=_ga .GetName (_ccd .Get ("\u0054r\u0061\u0070\u0070\u0065\u0064"));if _cad {switch _egf .String (){case "\u0054\u0072\u0075\u0065":_ad .Trapped =true ;
case "\u0046\u0061\u006cs\u0065":_ad .Trapped =false ;default:_ad .Trapped =true ;};};case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":if _ge ,_fe :=_ga .GetString (_ccd .Get ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065"));
_fe &&_ge .String ()!=""{_gab ,_fed :=_eb .ParsePdfTime (_ge .String ());if _fed !=nil {return _bc .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077",_fed );
};_ad .CreationDate =_e .NewDate (_gab );};case "\u004do\u0064\u0044\u0061\u0074\u0065":if _db ,_egg :=_ga .GetString (_ccd .Get ("\u004do\u0064\u0044\u0061\u0074\u0065"));_egg &&_db .String ()!=""{_ecf ,_cca :=_eb .ParsePdfTime (_db .String ());if _cca !=nil {return _bc .Errorf ("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w",_cca );
};_ad .ModifyDate =_e .NewDate (_ecf );};};};};if options .PdfVersion !=""{_ad .PDFVersion =options .PdfVersion ;};if options .Marked {_ad .Marked =_e .Bool (options .Marked );};if options .Copyright !=""{_ad .Copyright =options .Copyright ;};if _cc =_ad .SyncToXMP (_ba ._egd );
_cc !=nil {return _cc ;};return nil ;};

// SetMediaManagement sets up XMP media management metadata: namespace xmpMM.
func (_ed *Document )SetMediaManagement (options *MediaManagementOptions )error {_gc ,_ege :=_ca .MakeModel (_ed ._egd );if _ege !=nil {return _ege ;};if options ==nil {options =new (MediaManagementOptions );};_fc :=_ca .ResourceRef {};switch {case options .DocumentID !="":_gc .DocumentID =_e .GUID (options .DocumentID );
case options .NewDocumentID ||_gc .DocumentID .IsZero ():if !_gc .DocumentID .IsZero (){_fc .DocumentID =_gc .DocumentID ;};_cd ,_dce :=_gd .NewUUID ();if _dce !=nil {return _dce ;};_gc .DocumentID =_e .GUID (_cd .String ());};if !_gc .InstanceID .IsZero (){_fc .InstanceID =_gc .InstanceID ;
};_gc .InstanceID =_e .GUID (options .InstanceID );if _gc .InstanceID ==""{_cfc ,_cef :=_gd .NewUUID ();if _cef !=nil {return _cef ;};_gc .InstanceID =_e .GUID (_cfc .String ());};if !_fc .IsZero (){_gc .DerivedFrom =&_fc ;};_cec :=options .VersionID ;
if _gc .VersionID !=""{_be ,_eggc :=_ee .Atoi (_gc .VersionID );if _eggc !=nil {_cec =_ee .Itoa (len (_gc .Versions )+1);}else {_cec =_ee .Itoa (_be +1);};};if _cec ==""{_cec ="\u0031";};_gc .VersionID =_cec ;if _ege =_gc .SyncToXMP (_ed ._egd );_ege !=nil {return _ege ;
};return nil ;};

// GetPdfInfo gets the document pdf info.
func (_feg *Document )GetPdfInfo ()(*PdfInfo ,bool ){_fef :=PdfInfo {};var _gdd *_ga .PdfObjectDictionary ;_ae :=func (_efg string ,_fegb _ga .PdfObject ){if _gdd ==nil {_gdd =_ga .MakeDict ();};_gdd .Set (_ga .PdfObjectName (_efg ),_fegb );};_cf ,_gbc :=_feg ._egd .FindModel (_eg .NsPDF ).(*_eg .PDFInfo );
if !_gbc {_gbf ,_aac :=_feg ._egd .FindModel (_d .NsXmp ).(*_d .XmpBase );if !_aac {return nil ,false ;};if _gbf .CreatorTool !=""{_ae ("\u0043r\u0065\u0061\u0074\u006f\u0072",_ga .MakeString (string (_gbf .CreatorTool )));};if !_gbf .CreateDate .IsZero (){_ae ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065",_ga .MakeString (_eb .FormatPdfTime (_gbf .CreateDate .Value ())));
};if !_gbf .ModifyDate .IsZero (){_ae ("\u004do\u0064\u0044\u0061\u0074\u0065",_ga .MakeString (_eb .FormatPdfTime (_gbf .ModifyDate .Value ())));};_fef .InfoDict =_gdd ;return &_fef ,true ;};_fef .Copyright =_cf .Copyright ;_fef .PdfVersion =_cf .PDFVersion ;
_fef .Marked =bool (_cf .Marked );if len (_cf .Title )> 0{_ae ("\u0054\u0069\u0074l\u0065",_ga .MakeString (_cf .Title .Default ()));};if len (_cf .Author )> 0{_ae ("\u0041\u0075\u0074\u0068\u006f\u0072",_ga .MakeString (_cf .Author [0]));};if _cf .Keywords !=""{_ae ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",_ga .MakeString (_cf .Keywords ));
};if len (_cf .Subject )> 0{_ae ("\u0053u\u0062\u006a\u0065\u0063\u0074",_ga .MakeString (_cf .Subject .Default ()));};if _cf .Creator !=""{_ae ("\u0043r\u0065\u0061\u0074\u006f\u0072",_ga .MakeString (string (_cf .Creator )));};if _cf .Producer !=""{_ae ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072",_ga .MakeString (string (_cf .Producer )));
};if _cf .Trapped {_ae ("\u0054r\u0061\u0070\u0070\u0065\u0064",_ga .MakeName ("\u0054\u0072\u0075\u0065"));};if !_cf .CreationDate .IsZero (){_ae ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065",_ga .MakeString (_eb .FormatPdfTime (_cf .CreationDate .Value ())));
};if !_cf .ModifyDate .IsZero (){_ae ("\u004do\u0064\u0044\u0061\u0074\u0065",_ga .MakeString (_eb .FormatPdfTime (_cf .ModifyDate .Value ())));};_fef .InfoDict =_gdd ;return &_fef ,true ;};

// MarshalIndent the document into xml byte stream with predefined prefix and indent.
func (_dab *Document )MarshalIndent (prefix ,indent string )([]byte ,error ){if _dab ._egd .IsDirty (){if _bb :=_dab ._egd .SyncModels ();_bb !=nil {return nil ,_bb ;};};return _e .MarshalIndent (_dab ._egd ,prefix ,indent );};

// GetPdfaExtensionSchemas gets a pdfa extension schemas.
func (_daa *Document )GetPdfaExtensionSchemas ()([]_ea .Schema ,error ){_aa :=_daa ._egd .FindModel (_ea .Namespace );if _aa ==nil {return nil ,nil ;};_ef ,_ac :=_aa .(*_ea .Model );if !_ac {return nil ,_bc .Errorf ("\u0069\u006eva\u006c\u0069\u0064 \u006d\u006f\u0064\u0065l f\u006fr \u0070\u0064\u0066\u0061\u0045\u0078\u0074en\u0073\u0069\u006f\u006e\u0073\u003a\u0020%\u0054",_aa );
};return _ef .Schemas ,nil ;};

// LoadDocument loads up the xmp document from provided input stream.
func LoadDocument (stream []byte )(*Document ,error ){_cg :=_e .NewDocument ();if _a :=_e .Unmarshal (stream ,_cg );_a !=nil {return nil ,_a ;};return &Document {_egd :_cg },nil ;};

// Marshal the document into xml byte stream.
func (_da *Document )Marshal ()([]byte ,error ){if _da ._egd .IsDirty (){if _f :=_da ._egd .SyncModels ();_f !=nil {return nil ,_f ;};};return _e .Marshal (_da ._egd );};

// Document is an implementation of the xmp document.
// It is a wrapper over go-xmp/xmp.Document that provides some Pdf predefined functionality.
type Document struct{_egd *_e .Document };

// PdfInfo is the xmp document pdf info.
type PdfInfo struct{InfoDict _ga .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;};

// GetPdfAID gets the pdfaid xmp metadata model.
func (_fd *Document )GetPdfAID ()(*PdfAID ,bool ){_ccdc ,_cff :=_fd ._egd .FindModel (_b .Namespace ).(*_b .Model );if !_cff {return nil ,false ;};return &PdfAID {Part :_ccdc .Part ,Conformance :_ccdc .Conformance },true ;};

// GUID is a string representing a globally unique identifier.
type GUID string ;

// SetPdfAExtension sets the pdfaExtension XMP metadata.
func (_bf *Document )SetPdfAExtension ()error {_cgf ,_gb :=_ea .MakeModel (_bf ._egd );if _gb !=nil {return _gb ;};if _gb =_ea .FillModel (_bf ._egd ,_cgf );_gb !=nil {return _gb ;};if _gb =_cgf .SyncToXMP (_bf ._egd );_gb !=nil {return _gb ;};return nil ;
};

// MediaManagementOptions are the options for the Media management xmp metadata.
type MediaManagementOptions struct{

// OriginalDocumentID  as media is imported and projects is started, an original-document ID
// must be created to identify a new document. This identifies a document as a conceptual entity.
// By default, this value is generated.
OriginalDocumentID string ;

// NewDocumentID is a flag which generates a new Document identifier while setting media management.
// This value should be set to true only if the document is stored and saved as new document.
// Otherwise, if the document is modified and overwrites previous file, it should be set to false.
NewDocumentID bool ;

// DocumentID when a document is copied to a new file path or converted to a new format with
// Save As, another new document ID should usually be assigned. This identifies a general version or
// branch of a document. You can use it to track different versions or extracted portions of a document
// with the same original-document ID.
// By default, this value is generated if NewDocumentID is true or previous doesn't exist.
DocumentID string ;

// InstanceID to track a document’s editing history, you must assign a new instance ID
// whenever a document is saved after any changes. This uniquely identifies an exact version of a
// document. It is used in resource references (to identify both the document or part itself and the
// referenced or referencing documents), and in document-history resource events (to identify the
// document instance that resulted from the change).
// By default, this value is generated.
InstanceID string ;

// DerivedFrom references the source document from which this one is derived,
// typically through a Save As operation that changes the file name or format. It is a minimal reference;
// missing components can be assumed to be unchanged. For example, a new version might only need
// to specify the instance ID and version number of the previous version, or a rendition might only need
// to specify the instance ID and rendition class of the original.
// By default, the derived from structure is filled from previous XMP metadata (if exists).
DerivedFrom string ;

// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
// document history, but should not be used to identify a document uniquely in any context.
// Usually it simply works by incrementing integers 1,2,3...
// By default, this values is incremented or set to the next version number.
VersionID string ;

// ModifyComment is a comment to given modification
ModifyComment string ;

// ModifyDate is a custom modification date for the versions.
// By default, this would be set to time.Now().
ModifyDate _c .Time ;

// Modifier is a person who did the modification.
Modifier string ;};