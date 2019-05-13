// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/file.proto

package resultstore // import "google.golang.org/genproto/googleapis/devtools/resultstore/v2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// If known, the hash function used to compute this digest.
type File_HashType int32

const (
	// Unknown
	File_HASH_TYPE_UNSPECIFIED File_HashType = 0
	// MD5
	File_MD5 File_HashType = 1
	// SHA-1
	File_SHA1 File_HashType = 2
	// SHA-256
	File_SHA256 File_HashType = 3
)

var File_HashType_name = map[int32]string{
	0: "HASH_TYPE_UNSPECIFIED",
	1: "MD5",
	2: "SHA1",
	3: "SHA256",
}
var File_HashType_value = map[string]int32{
	"HASH_TYPE_UNSPECIFIED": 0,
	"MD5":                   1,
	"SHA1":                  2,
	"SHA256":                3,
}

func (x File_HashType) String() string {
	return proto.EnumName(File_HashType_name, int32(x))
}
func (File_HashType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_file_ce5c35947c0b9efa, []int{0, 0}
}

// The metadata for a file or an archive file entry.
type File struct {
	// The identifier of the file or archive entry.
	// User-provided, must be unique for the repeated field it is in. When an
	// Append RPC is called with a Files field populated, if a File already exists
	// with this ID, that File will be overwritten with the new File proto.
	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// The URI of a file.
	// This could also be the URI of an entire archive.
	// Most log data doesn't need to be stored forever, so a ttl is suggested.
	// Note that if you ever move or delete the file at this URI, the link from
	// the server will be broken.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// (Optional) The length of the file in bytes.  Allows the filesize to be
	// shown in the UI.  Omit if file is still being written or length is
	// not known.  This could also be the length of an entire archive.
	Length *wrappers.Int64Value `protobuf:"bytes,3,opt,name=length,proto3" json:"length,omitempty"`
	// (Optional) The content-type (aka MIME-type) of the file.  This is sent to
	// the web browser so it knows how to handle the file. (e.g. text/plain,
	// image/jpeg, text/html, etc). For zip archives, use "application/zip".
	ContentType string `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// (Optional) If the above path, length, and content_type are referring to an
	// archive, and you wish to refer to a particular entry within that archive,
	// put the particular archive entry data here.
	ArchiveEntry *ArchiveEntry `protobuf:"bytes,5,opt,name=archive_entry,json=archiveEntry,proto3" json:"archive_entry,omitempty"`
	// (Optional) A url to a content display app/site for this file or archive
	// entry.
	ContentViewer string `protobuf:"bytes,6,opt,name=content_viewer,json=contentViewer,proto3" json:"content_viewer,omitempty"`
	// (Optional) Whether to hide this file or archive entry in the UI.  Defaults
	// to false. A checkbox lets users see hidden files, but they're hidden by
	// default.
	Hidden bool `protobuf:"varint,7,opt,name=hidden,proto3" json:"hidden,omitempty"`
	// (Optional) A short description of what this file or archive entry
	// contains. This description should help someone viewing the list of these
	// files to understand the purpose of this file and what they would want to
	// view it for.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// (Optional) digest of this file in hexadecimal-like string if known.
	Digest string `protobuf:"bytes,9,opt,name=digest,proto3" json:"digest,omitempty"`
	// (Optional) The algorithm corresponding to the digest if known.
	HashType             File_HashType `protobuf:"varint,10,opt,name=hash_type,json=hashType,proto3,enum=google.devtools.resultstore.v2.File_HashType" json:"hash_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_ce5c35947c0b9efa, []int{0}
}
func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (dst *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(dst, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *File) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *File) GetLength() *wrappers.Int64Value {
	if m != nil {
		return m.Length
	}
	return nil
}

func (m *File) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *File) GetArchiveEntry() *ArchiveEntry {
	if m != nil {
		return m.ArchiveEntry
	}
	return nil
}

func (m *File) GetContentViewer() string {
	if m != nil {
		return m.ContentViewer
	}
	return ""
}

func (m *File) GetHidden() bool {
	if m != nil {
		return m.Hidden
	}
	return false
}

func (m *File) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *File) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *File) GetHashType() File_HashType {
	if m != nil {
		return m.HashType
	}
	return File_HASH_TYPE_UNSPECIFIED
}

// Information specific to an entry in an archive.
type ArchiveEntry struct {
	// The relative path of the entry within the archive.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// (Optional) The uncompressed length of the archive entry in bytes.  Allows
	// the entry size to be shown in the UI.  Omit if the length is not known.
	Length *wrappers.Int64Value `protobuf:"bytes,2,opt,name=length,proto3" json:"length,omitempty"`
	// (Optional) The content-type (aka MIME-type) of the archive entry. (e.g.
	// text/plain, image/jpeg, text/html, etc). This is sent to the web browser
	// so it knows how to handle the entry.
	ContentType          string   `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArchiveEntry) Reset()         { *m = ArchiveEntry{} }
func (m *ArchiveEntry) String() string { return proto.CompactTextString(m) }
func (*ArchiveEntry) ProtoMessage()    {}
func (*ArchiveEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_file_ce5c35947c0b9efa, []int{1}
}
func (m *ArchiveEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchiveEntry.Unmarshal(m, b)
}
func (m *ArchiveEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchiveEntry.Marshal(b, m, deterministic)
}
func (dst *ArchiveEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchiveEntry.Merge(dst, src)
}
func (m *ArchiveEntry) XXX_Size() int {
	return xxx_messageInfo_ArchiveEntry.Size(m)
}
func (m *ArchiveEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchiveEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ArchiveEntry proto.InternalMessageInfo

func (m *ArchiveEntry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ArchiveEntry) GetLength() *wrappers.Int64Value {
	if m != nil {
		return m.Length
	}
	return nil
}

func (m *ArchiveEntry) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func init() {
	proto.RegisterType((*File)(nil), "google.devtools.resultstore.v2.File")
	proto.RegisterType((*ArchiveEntry)(nil), "google.devtools.resultstore.v2.ArchiveEntry")
	proto.RegisterEnum("google.devtools.resultstore.v2.File_HashType", File_HashType_name, File_HashType_value)
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/file.proto", fileDescriptor_file_ce5c35947c0b9efa)
}

var fileDescriptor_file_ce5c35947c0b9efa = []byte{
	// 453 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x61, 0x8b, 0xd3, 0x40,
	0x10, 0x35, 0x4d, 0xed, 0xa5, 0xd3, 0xde, 0x11, 0x16, 0x94, 0xa8, 0x20, 0xb1, 0x20, 0x54, 0xd0,
	0x0d, 0xe6, 0xbc, 0xfb, 0xe2, 0xa7, 0x6a, 0x7b, 0xa4, 0x82, 0x52, 0xd3, 0xf3, 0x40, 0xbf, 0x94,
	0x5c, 0x33, 0x97, 0x2c, 0xc4, 0x6c, 0xdc, 0xdd, 0xe6, 0xa8, 0xbf, 0xd6, 0x9f, 0x22, 0xd9, 0x6c,
	0x21, 0x20, 0x7a, 0xe0, 0xb7, 0x99, 0x37, 0xef, 0xcd, 0xe4, 0xed, 0x23, 0xf0, 0x22, 0xe3, 0x3c,
	0x2b, 0x30, 0x48, 0xb1, 0x56, 0x9c, 0x17, 0x32, 0x10, 0x28, 0x77, 0x85, 0x92, 0x8a, 0x0b, 0x0c,
	0xea, 0x30, 0xb8, 0x61, 0x05, 0xd2, 0x4a, 0x70, 0xc5, 0xc9, 0xd3, 0x96, 0x4a, 0x0f, 0x54, 0xda,
	0xa1, 0xd2, 0x3a, 0x7c, 0x6c, 0xe6, 0x81, 0x66, 0x5f, 0xef, 0x6e, 0x82, 0x5b, 0x91, 0x54, 0x15,
	0x0a, 0xd9, 0xea, 0x27, 0xbf, 0x6c, 0xe8, 0x5f, 0xb0, 0x02, 0x89, 0x0b, 0xf6, 0x8e, 0xa5, 0x9e,
	0xe5, 0x5b, 0xd3, 0x61, 0xdc, 0x94, 0x1a, 0x11, 0xcc, 0xeb, 0x19, 0x44, 0x30, 0x72, 0x0a, 0x83,
	0x02, 0xcb, 0x4c, 0xe5, 0x9e, 0xed, 0x5b, 0xd3, 0x51, 0xf8, 0x84, 0x9a, 0xeb, 0x87, 0xed, 0x74,
	0x59, 0xaa, 0xf3, 0x37, 0x57, 0x49, 0xb1, 0xc3, 0xd8, 0x50, 0xc9, 0x33, 0x18, 0x6f, 0x79, 0xa9,
	0xb0, 0x54, 0x1b, 0xb5, 0xaf, 0xd0, 0xeb, 0xeb, 0x7d, 0x23, 0x83, 0x5d, 0xee, 0x2b, 0x24, 0x9f,
	0xe1, 0x38, 0x11, 0xdb, 0x9c, 0xd5, 0xb8, 0xc1, 0x52, 0x89, 0xbd, 0x77, 0x5f, 0xaf, 0x7f, 0x49,
	0xff, 0x6d, 0x8e, 0xce, 0x5a, 0xd1, 0xa2, 0xd1, 0xc4, 0xe3, 0xa4, 0xd3, 0x91, 0xe7, 0x70, 0x72,
	0xb8, 0x5a, 0x33, 0xbc, 0x45, 0xe1, 0x0d, 0xf4, 0xdd, 0x63, 0x83, 0x5e, 0x69, 0x90, 0x3c, 0x84,
	0x41, 0xce, 0xd2, 0x14, 0x4b, 0xef, 0xc8, 0xb7, 0xa6, 0x4e, 0x6c, 0x3a, 0xe2, 0xc3, 0x28, 0x45,
	0xb9, 0x15, 0xac, 0x52, 0x8c, 0x97, 0x9e, 0xd3, 0x7e, 0x73, 0x07, 0x6a, 0x94, 0x29, 0xcb, 0x50,
	0x2a, 0x6f, 0xa8, 0x87, 0xa6, 0x23, 0x1f, 0x60, 0x98, 0x27, 0x32, 0x6f, 0xbd, 0x82, 0x6f, 0x4d,
	0x4f, 0xc2, 0x57, 0x77, 0xf9, 0x68, 0x02, 0xa0, 0x51, 0x22, 0xf3, 0xe6, 0x35, 0x62, 0x27, 0x37,
	0xd5, 0x64, 0x0e, 0xce, 0x01, 0x25, 0x8f, 0xe0, 0x41, 0x34, 0x5b, 0x47, 0x9b, 0xcb, 0xaf, 0xab,
	0xc5, 0xe6, 0xcb, 0xa7, 0xf5, 0x6a, 0xf1, 0x7e, 0x79, 0xb1, 0x5c, 0xcc, 0xdd, 0x7b, 0xe4, 0x08,
	0xec, 0x8f, 0xf3, 0x33, 0xd7, 0x22, 0x0e, 0xf4, 0xd7, 0xd1, 0xec, 0xb5, 0xdb, 0x23, 0x00, 0x83,
	0x75, 0x34, 0x0b, 0xcf, 0xce, 0x5d, 0x7b, 0xf2, 0x13, 0xc6, 0xdd, 0x87, 0x22, 0x04, 0xfa, 0x55,
	0xa2, 0x72, 0x13, 0xb5, 0xae, 0x3b, 0xc9, 0xf6, 0xfe, 0x3f, 0x59, 0xfb, 0x8f, 0x64, 0xdf, 0xfd,
	0x80, 0xc9, 0x96, 0x7f, 0xbf, 0xc3, 0xff, 0xca, 0xfa, 0xb6, 0x34, 0x8c, 0x8c, 0x17, 0x49, 0x99,
	0x51, 0x2e, 0xb2, 0x20, 0xc3, 0x52, 0x1f, 0x0f, 0xda, 0x51, 0x52, 0x31, 0xf9, 0xb7, 0x1f, 0xe2,
	0x6d, 0xa7, 0xbd, 0x1e, 0x68, 0xd5, 0xe9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0xc7, 0x53,
	0x67, 0x45, 0x03, 0x00, 0x00,
}
