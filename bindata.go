package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func bash_buildpack_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xac, 0x55,
		0x51, 0x6f, 0xe2, 0x46, 0x10, 0x7e, 0x36, 0xbf, 0x62, 0xe4, 0x43, 0xb9,
		0xa4, 0x92, 0x43, 0xe9, 0xf5, 0x29, 0x17, 0x4e, 0x4d, 0x02, 0xd7, 0xa2,
		0x26, 0x21, 0x25, 0xf0, 0x94, 0x46, 0x68, 0x59, 0x8f, 0x61, 0x95, 0x65,
		0xd7, 0xdd, 0x5d, 0x87, 0x43, 0x69, 0xfe, 0x7b, 0x67, 0x6d, 0x83, 0xed,
		0x40, 0x7a, 0x2f, 0x27, 0x21, 0x61, 0xef, 0xec, 0xcc, 0x7c, 0x33, 0xf3,
		0xcd, 0xe7, 0x56, 0x8c, 0x5c, 0x32, 0x83, 0x60, 0x51, 0x22, 0x77, 0x18,
		0xcf, 0x52, 0xe6, 0x96, 0xfb, 0xa7, 0x8a, 0xad, 0xb0, 0xd5, 0x9a, 0x67,
		0x42, 0xc6, 0x29, 0xe3, 0x4f, 0x51, 0xfe, 0x74, 0x7c, 0x02, 0x2f, 0xad,
		0x60, 0x7b, 0x37, 0x46, 0xcb, 0x7b, 0xe1, 0xa5, 0x37, 0x00, 0x53, 0xc0,
		0xd2, 0x54, 0x0a, 0xce, 0x9c, 0xd0, 0x0a, 0x32, 0x2b, 0xd4, 0x02, 0x84,
		0xb2, 0x8e, 0x49, 0x89, 0x31, 0xec, 0xe2, 0xd8, 0xb0, 0x15, 0xa0, 0xb2,
		0x99, 0xc1, 0xc8, 0xe7, 0xb5, 0xad, 0xa0, 0x4a, 0x61, 0xd1, 0x65, 0x29,
		0x7c, 0x81, 0x4e, 0x8c, 0xcf, 0x1d, 0x95, 0x49, 0xd9, 0x34, 0x7a, 0x60,
		0xf0, 0x2f, 0x05, 0x8d, 0x51, 0xb9, 0xba, 0x89, 0xeb, 0x55, 0x2a, 0x24,
		0xd6, 0x6c, 0xa9, 0xd1, 0x3c, 0xa1, 0xa3, 0xc8, 0x6d, 0x52, 0xb4, 0x95,
		0xe1, 0xb5, 0x5e, 0x51, 0x89, 0xee, 0x50, 0x4d, 0xc3, 0xc2, 0x54, 0xc1,
		0x86, 0xc4, 0xe8, 0x15, 0xfc, 0x2e, 0x1c, 0x4c, 0xc7, 0xd7, 0x54, 0x6c,
		0x0c, 0x3a, 0xf5, 0x85, 0x32, 0x09, 0x94, 0x7d, 0x25, 0x9c, 0x13, 0x76,
		0x19, 0x56, 0x61, 0x32, 0x23, 0x7b, 0x61, 0xbb, 0x1b, 0x96, 0x56, 0x7a,
		0xfe, 0x25, 0x04, 0xdf, 0x52, 0x7a, 0x7a, 0xf9, 0x74, 0x16, 0xb5, 0x8f,
		0xe7, 0xcc, 0xa2, 0x3f, 0x80, 0x76, 0xf7, 0xe4, 0x75, 0xaf, 0x29, 0x52,
		0x73, 0x0a, 0xed, 0x98, 0x59, 0xa0, 0xcb, 0x07, 0x44, 0x7e, 0x3b, 0x30,
		0xf9, 0x41, 0xa7, 0xed, 0xbd, 0xc9, 0x51, 0x24, 0xf0, 0xf0, 0x00, 0x91,
		0x82, 0xb0, 0x5d, 0x24, 0x0b, 0xe1, 0xf1, 0xf1, 0x33, 0xb8, 0x25, 0xaa,
		0x56, 0x10, 0x2c, 0x08, 0x32, 0x97, 0x5a, 0x21, 0x99, 0x09, 0x54, 0x48,
		0x7f, 0xb5, 0xa8, 0xe4, 0x1e, 0xf0, 0x78, 0xff, 0x2c, 0xf7, 0x5a, 0x22,
		0x7f, 0xd2, 0x99, 0x83, 0x28, 0xfa, 0x27, 0x13, 0xe8, 0xaa, 0xf8, 0x85,
		0x53, 0xd4, 0x9c, 0x14, 0x4a, 0x8b, 0x8d, 0x7c, 0x51, 0x14, 0x63, 0x4a,
		0xc0, 0xbb, 0xef, 0x65, 0x4e, 0x44, 0x2b, 0x30, 0x2b, 0x88, 0x4c, 0xd2,
		0x34, 0x75, 0x4e, 0x17, 0x3e, 0x49, 0x63, 0x56, 0x52, 0x58, 0x77, 0x68,
		0x50, 0xd7, 0x74, 0xfe, 0x1e, 0xcd, 0xa4, 0x85, 0xc8, 0x27, 0x6f, 0xf6,
		0xed, 0x4d, 0xe0, 0x9c, 0x73, 0x45, 0xe4, 0x0f, 0x70, 0xb9, 0x1b, 0x37,
		0x7e, 0x4b, 0x89, 0x6d, 0x39, 0x97, 0x69, 0x1a, 0xf4, 0xa6, 0x8d, 0x83,
		0x8b, 0xbb, 0xbb, 0x59, 0x7f, 0x38, 0xa6, 0x51, 0x10, 0xd5, 0xb7, 0x55,
		0x94, 0xb6, 0x3f, 0x46, 0x37, 0x83, 0x83, 0x86, 0xf1, 0xe0, 0xaf, 0xe9,
		0xe0, 0x7e, 0x32, 0x1b, 0xf6, 0x7b, 0x61, 0x9e, 0x36, 0x6a, 0x8f, 0x2f,
		0x6e, 0xfb, 0xa3, 0x9b, 0xea, 0xca, 0xfd, 0xe4, 0xe2, 0xea, 0xcf, 0x5e,
		0xc8, 0x31, 0x66, 0x26, 0xea, 0xfe, 0x4a, 0x06, 0x9e, 0x52, 0x5b, 0xa0,
		0x0a, 0xd7, 0x39, 0x0d, 0xb7, 0x85, 0x6c, 0xc3, 0x7b, 0xbc, 0x7d, 0xa3,
		0xd3, 0x94, 0xca, 0x4e, 0x8d, 0x78, 0x26, 0xb6, 0x2f, 0x90, 0xb0, 0x66,
		0x16, 0xcd, 0x4a, 0xd3, 0x78, 0xa2, 0xa5, 0xf6, 0xec, 0xf2, 0xc0, 0x40,
		0xe9, 0xb9, 0x8e, 0x37, 0x14, 0x77, 0xa9, 0xd7, 0x0a, 0xa2, 0x71, 0x79,
		0x70, 0xa6, 0xf4, 0xc2, 0x68, 0xda, 0xb9, 0xbf, 0x69, 0x72, 0x35, 0xf0,
		0xe5, 0x7b, 0x2d, 0x5f, 0x79, 0xc2, 0x19, 0xb1, 0xa2, 0x44, 0xe0, 0x01,
		0x4c, 0x2d, 0x26, 0x99, 0x24, 0xd1, 0x20, 0xfe, 0xab, 0x85, 0x85, 0x0e,
		0x24, 0xc8, 0x1c, 0xf1, 0xb8, 0x6a, 0xda, 0x15, 0x2d, 0xcc, 0xec, 0x6a,
		0x74, 0x7b, 0x3b, 0xb8, 0x9a, 0xcc, 0x26, 0xc3, 0x9b, 0xc1, 0x68, 0x3a,
		0xe9, 0x85, 0x9f, 0x7e, 0x2e, 0x02, 0xe4, 0x1d, 0xb7, 0x0e, 0x53, 0x98,
		0x53, 0xdb, 0xd7, 0xcc, 0xc4, 0xd6, 0xaf, 0x0c, 0x65, 0x10, 0x73, 0x21,
		0x85, 0xdb, 0xec, 0xf8, 0x9d, 0x34, 0xda, 0x81, 0xea, 0xb9, 0x41, 0x73,
		0xab, 0x33, 0xc3, 0x71, 0xef, 0x4a, 0x4e, 0xb2, 0x37, 0xf3, 0xf6, 0x32,
		0x52, 0x0c, 0xbc, 0xb6, 0x39, 0x97, 0xd3, 0xe1, 0x75, 0xff, 0x8e, 0xc6,
		0x30, 0x23, 0xb8, 0x8d, 0xc8, 0x4e, 0x38, 0x92, 0x96, 0xf0, 0x2b, 0x3a,
		0xbe, 0xf4, 0x9a, 0xc6, 0x33, 0xeb, 0x48, 0x0a, 0x76, 0x11, 0xfd, 0x36,
		0x78, 0x00, 0x75, 0x35, 0xdd, 0x5f, 0xd6, 0xc2, 0xcb, 0xdf, 0xdd, 0x31,
		0xbe, 0xe1, 0xe1, 0xbb, 0x11, 0x0c, 0xbf, 0xde, 0xf7, 0x3e, 0x7e, 0xf8,
		0x08, 0x06, 0x59, 0xec, 0x15, 0xa4, 0x14, 0x0f, 0x38, 0x3f, 0x3f, 0xdf,
		0x83, 0x48, 0xd7, 0xf7, 0x94, 0xac, 0x5a, 0xb4, 0xad, 0x10, 0x94, 0x60,
		0x8f, 0xea, 0x9b, 0x5a, 0x07, 0x5b, 0x2a, 0xd2, 0x71, 0xa6, 0x76, 0x14,
		0x8a, 0xa1, 0x89, 0xac, 0x33, 0x17, 0x8a, 0x9c, 0x9d, 0x17, 0xdf, 0x1a,
		0x23, 0x4e, 0xc2, 0xdd, 0xca, 0x17, 0x5a, 0x55, 0xad, 0x5e, 0xef, 0xf8,
		0x6d, 0xf5, 0x3f, 0x9d, 0xd0, 0xb5, 0x44, 0x9b, 0x9a, 0x9e, 0x0a, 0xdf,
		0xf5, 0x97, 0xca, 0xe9, 0xe1, 0xb7, 0xc7, 0xd7, 0xf0, 0x33, 0xc4, 0x9a,
		0x6e, 0x7e, 0x0f, 0xdf, 0xce, 0xeb, 0x5d, 0x6c, 0x39, 0x5d, 0x83, 0xe0,
		0xe8, 0x08, 0xde, 0x9d, 0x4b, 0xed, 0xce, 0x9c, 0x1a, 0xfe, 0x44, 0x2f,
		0x31, 0xc9, 0x56, 0x21, 0x4b, 0x35, 0x66, 0x34, 0xc7, 0x74, 0x80, 0x19,
		0xed, 0x06, 0x5a, 0xff, 0x15, 0x84, 0x02, 0x12, 0xc6, 0x55, 0x8f, 0xca,
		0xbb, 0x53, 0xc5, 0xe6, 0xf4, 0xef, 0x74, 0x89, 0x0b, 0x58, 0x93, 0x48,
		0xf8, 0x8d, 0xc6, 0xdd, 0xdd, 0x67, 0x6d, 0xf9, 0x85, 0xdb, 0xea, 0xd4,
		0x64, 0xd4, 0x1f, 0x9d, 0x81, 0x43, 0xaf, 0x7d, 0x09, 0xc1, 0x11, 0x16,
		0xe8, 0xa7, 0x90, 0xa3, 0xb5, 0xcc, 0x6c, 0x80, 0x3a, 0x4d, 0xe7, 0x6b,
		0x04, 0xee, 0x3f, 0xcb, 0x72, 0xcd, 0x36, 0x16, 0x52, 0x66, 0x2d, 0xb4,
		0x69, 0x29, 0xf2, 0x42, 0x60, 0x2d, 0xdc, 0xd2, 0x8b, 0xbb, 0xb0, 0x36,
		0xc3, 0x6d, 0xc1, 0xd4, 0x69, 0x2f, 0x9a, 0x17, 0xd5, 0xbd, 0x93, 0x46,
		0xc5, 0x8d, 0x31, 0x84, 0x07, 0x78, 0x52, 0xe2, 0x7c, 0x23, 0x56, 0xd0,
		0x10, 0x0e, 0x7a, 0xdb, 0x46, 0xaf, 0xfa, 0xf3, 0x43, 0x22, 0x17, 0xb3,
		0xfb, 0x6e, 0x28, 0x43, 0x27, 0xf4, 0xf1, 0xfd, 0x7f, 0x90, 0x5f, 0x1a,
		0xd6, 0xce, 0xe9, 0xd6, 0xa9, 0xf5, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff,
		0xff, 0x40, 0xb8, 0x14, 0xca, 0x41, 0x09, 0x00, 0x00,
	},
		"bash/buildpack.bash",
	)
}

func bash_cmd_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x9c, 0x54,
		0x41, 0x73, 0x9b, 0x3c, 0x10, 0x3d, 0xc3, 0xaf, 0xd8, 0x4f, 0x51, 0x32,
		0xc9, 0x81, 0xe1, 0xc3, 0xa7, 0x0e, 0x1e, 0x77, 0x9c, 0x69, 0x7b, 0x6b,
		0x7b, 0xc9, 0xd1, 0x38, 0x33, 0x2a, 0x88, 0xc0, 0x58, 0x16, 0x1e, 0x84,
		0xdd, 0x74, 0x30, 0xff, 0xbd, 0xbb, 0x42, 0x60, 0x88, 0x7d, 0x6a, 0x2e,
		0x46, 0xbb, 0xab, 0xdd, 0xf7, 0x9e, 0xde, 0xc6, 0xcf, 0x64, 0xaa, 0x44,
		0x2d, 0x21, 0x78, 0x86, 0x2f, 0x3f, 0xbe, 0xbe, 0xf8, 0x7e, 0xba, 0xcf,
		0x02, 0x55, 0x9a, 0xe6, 0xf1, 0x09, 0x5a, 0xdf, 0x1b, 0xd2, 0x99, 0x34,
		0xe9, 0x8a, 0x7d, 0xc7, 0xb8, 0x01, 0x71, 0x12, 0xa5, 0x12, 0xbf, 0x94,
		0x84, 0xb4, 0xda, 0xef, 0x85, 0xce, 0x0c, 0xbb, 0x14, 0x6a, 0xb3, 0x62,
		0x3c, 0xc2, 0xc0, 0xd0, 0x27, 0xd8, 0xc9, 0x3f, 0x06, 0x18, 0xd7, 0x86,
		0xc1, 0x19, 0x8c, 0xcc, 0x80, 0x99, 0x10, 0x4f, 0x71, 0x18, 0x32, 0xbf,
		0xbb, 0xcc, 0xb3, 0x75, 0xf3, 0xa1, 0x63, 0xaf, 0xbc, 0xaa, 0x61, 0x07,
		0xa5, 0xc6, 0x36, 0xed, 0x7f, 0x04, 0x73, 0xb3, 0xde, 0x76, 0x6c, 0x09,
		0x59, 0xe5, 0x7b, 0x9e, 0x4c, 0x8b, 0x0a, 0x13, 0x3b, 0x02, 0x51, 0x69,
		0x89, 0x43, 0xde, 0x6a, 0x79, 0x00, 0xf6, 0x4a, 0x43, 0xec, 0xcc, 0xaa,
		0x6e, 0x66, 0x93, 0xb4, 0x9b, 0xf3, 0x2f, 0x6d, 0x83, 0x13, 0xc4, 0x1f,
		0x7a, 0xca, 0xf7, 0x03, 0x9e, 0x6e, 0xe9, 0xf5, 0xcd, 0x66, 0x50, 0x31,
		0xc8, 0x8f, 0x3a, 0x6d, 0xca, 0x4a, 0x83, 0xa0, 0x93, 0xd3, 0x6d, 0x22,
		0x5b, 0xae, 0x2d, 0x55, 0x4c, 0xe3, 0x6f, 0xbb, 0x88, 0x03, 0x1e, 0x75,
		0x98, 0x56, 0x55, 0x2a, 0x94, 0xd5, 0xc1, 0xa9, 0xa0, 0x09, 0x2e, 0x7f,
		0x9c, 0x50, 0x79, 0x9a, 0xe3, 0xcd, 0x35, 0x9b, 0x2a, 0x10, 0x30, 0x78,
		0xf8, 0x0c, 0x61, 0x26, 0x4f, 0xa1, 0x3e, 0x2a, 0x05, 0x0f, 0x0f, 0xbd,
		0xaa, 0xda, 0xd1, 0xf2, 0x3d, 0xcb, 0x9b, 0x9e, 0x27, 0xe6, 0xad, 0x30,
		0xe1, 0x1d, 0x7e, 0x05, 0x61, 0xc7, 0xb6, 0x2b, 0xdb, 0x6b, 0xce, 0x71,
		0x54, 0xee, 0xc3, 0x0b, 0x39, 0xba, 0x7c, 0x81, 0x5d, 0xe5, 0x09, 0x11,
		0x63, 0x90, 0x0a, 0x01, 0x71, 0xcd, 0x14, 0x49, 0x18, 0xa7, 0xdf, 0x04,
		0xeb, 0xac, 0x41, 0xb4, 0x01, 0x1e, 0x41, 0xc2, 0x12, 0xbe, 0x4e, 0x50,
		0x78, 0xdf, 0xeb, 0x9c, 0x73, 0xfa, 0x79, 0xd0, 0x3f, 0xbf, 0x83, 0x18,
		0x59, 0x50, 0xd1, 0x88, 0x69, 0x00, 0x73, 0x11, 0x09, 0x93, 0x4b, 0x30,
		0x45, 0x99, 0x37, 0x30, 0x84, 0xb1, 0x70, 0x16, 0x3f, 0x9f, 0xa1, 0xa9,
		0x8f, 0x72, 0x48, 0x9b, 0x46, 0x34, 0x47, 0xb3, 0xfa, 0xdf, 0xf7, 0xca,
		0x1c, 0x06, 0x55, 0x47, 0xb7, 0x5a, 0x19, 0x5f, 0x39, 0xc6, 0x13, 0x3e,
		0xd3, 0x71, 0x09, 0x4d, 0x21, 0x35, 0x92, 0xe0, 0xed, 0x44, 0x3f, 0xac,
		0x63, 0xdb, 0x0e, 0x6f, 0xaf, 0x49, 0x07, 0x65, 0x70, 0x0a, 0xb5, 0xdd,
		0x6c, 0x30, 0x44, 0x39, 0xd8, 0x6e, 0xc7, 0x8b, 0xee, 0xbd, 0x7e, 0x56,
		0x60, 0x8e, 0x69, 0x31, 0x38, 0x22, 0x06, 0x5b, 0x48, 0x79, 0x87, 0x6c,
		0x41, 0x4f, 0xab, 0x86, 0x2e, 0x04, 0xeb, 0xba, 0x09, 0x7f, 0xcc, 0x75,
		0x40, 0xba, 0xf6, 0x15, 0x4f, 0x74, 0x3f, 0x2f, 0x9d, 0x27, 0x46, 0x6b,
		0x3c, 0x5f, 0xad, 0x6d, 0x6c, 0x0b, 0xd1, 0x55, 0x38, 0x73, 0xee, 0x2b,
		0xd7, 0xc8, 0x59, 0xcb, 0x3b, 0xd4, 0xa5, 0x6e, 0x72, 0x60, 0x00, 0xf7,
		0xc1, 0xe2, 0x93, 0x81, 0x7b, 0x93, 0xa0, 0xcd, 0x1c, 0xa9, 0xd9, 0xf8,
		0x6b, 0x39, 0x7a, 0x38, 0xbd, 0xd7, 0x46, 0x40, 0xef, 0x65, 0x03, 0xbc,
		0x67, 0xe8, 0x13, 0x54, 0xf7, 0xa2, 0x85, 0x54, 0x87, 0x5b, 0x7b, 0xf4,
		0x52, 0x54, 0xbf, 0x0d, 0x50, 0x16, 0x61, 0x22, 0xe0, 0xbd, 0xb0, 0xbb,
		0x44, 0xd0, 0x6f, 0x2d, 0x93, 0xa8, 0xdf, 0xc8, 0x0c, 0xf4, 0x0a, 0x83,
		0x70, 0x14, 0x9a, 0x48, 0x07, 0xf8, 0x37, 0x30, 0x27, 0x92, 0xd6, 0x13,
		0xf4, 0x81, 0x8b, 0x7e, 0x07, 0x4a, 0xa0, 0x02, 0x78, 0xa3, 0x2f, 0x9b,
		0xd8, 0xab, 0xa5, 0x36, 0xe1, 0x3d, 0x31, 0x0b, 0xe9, 0x1f, 0x45, 0x1f,
		0xd4, 0x18, 0x82, 0xf9, 0x19, 0xc2, 0x00, 0x9d, 0x3c, 0xb9, 0x6e, 0xf7,
		0xfb, 0x86, 0x38, 0x0e, 0x89, 0x0e, 0x88, 0x97, 0x5b, 0xdf, 0x68, 0x34,
		0x8f, 0x5b, 0x11, 0xbb, 0xfc, 0x24, 0xd2, 0x64, 0x33, 0x06, 0xb9, 0xac,
		0x2a, 0xfe, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x2e, 0xf9, 0x29,
		0xc7, 0x05, 0x00, 0x00,
	},
		"bash/cmd.bash",
	)
}

func bash_fn_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x8c, 0x51,
		0x41, 0x6e, 0xc2, 0x30, 0x10, 0x3c, 0xc7, 0xaf, 0x18, 0xad, 0x22, 0x01,
		0xaa, 0xa2, 0x08, 0xae, 0x34, 0x3d, 0x56, 0xea, 0x1b, 0x28, 0x07, 0xcb,
		0xac, 0x89, 0xd5, 0xd4, 0x89, 0x6c, 0x03, 0xaa, 0x28, 0x7f, 0xef, 0x1a,
		0xd2, 0x92, 0x43, 0x55, 0x55, 0xb9, 0x64, 0x77, 0x66, 0x67, 0x76, 0xd6,
		0xca, 0xfa, 0x4a, 0x87, 0x7d, 0x9c, 0x2f, 0x70, 0x56, 0xc5, 0x8e, 0x4d,
		0xa7, 0x03, 0x63, 0xc7, 0xd1, 0x34, 0xf4, 0xe2, 0xe3, 0xc0, 0x26, 0x41,
		0xc3, 0x1e, 0xbc, 0x49, 0xae, 0xf7, 0xb3, 0x08, 0x21, 0x1f, 0xde, 0xd9,
		0xa7, 0x48, 0xaa, 0xe8, 0x7a, 0xa3, 0xbb, 0xdc, 0xe9, 0x9c, 0xe7, 0xa6,
		0x9c, 0xa7, 0x8f, 0x81, 0x51, 0x2e, 0xf1, 0x89, 0x7d, 0xe0, 0x01, 0xdf,
		0x6a, 0x63, 0x59, 0x1d, 0x41, 0x53, 0x03, 0x12, 0xa0, 0x65, 0xbd, 0x43,
		0xb5, 0x5c, 0xa8, 0x82, 0x4d, 0xdb, 0xa3, 0x62, 0x50, 0x79, 0x1e, 0x05,
		0xeb, 0x1a, 0x35, 0xbd, 0x7a, 0xba, 0x64, 0xa2, 0x3e, 0xbd, 0xa1, 0x7a,
		0x6e, 0x30, 0xab, 0x9b, 0xfa, 0x3c, 0x04, 0xe7, 0x13, 0xe8, 0x91, 0xca,
		0x25, 0x3d, 0xd1, 0x65, 0x26, 0x78, 0x0a, 0xc8, 0x5c, 0xc8, 0xa7, 0x2e,
		0x2a, 0xa7, 0xca, 0x16, 0xff, 0x4e, 0x95, 0xa1, 0xe0, 0x86, 0x5c, 0x51,
		0x1e, 0xc8, 0x44, 0xf9, 0xe1, 0xa3, 0xe4, 0xa3, 0x5f, 0x82, 0x45, 0x33,
		0x59, 0x9e, 0xd6, 0xb8, 0x6e, 0x5f, 0xe6, 0xfe, 0xe8, 0xee, 0xbc, 0xed,
		0xff, 0x70, 0x8f, 0x13, 0x7b, 0xba, 0x73, 0xac, 0x6f, 0x72, 0x26, 0xc4,
		0xb6, 0x3f, 0xc5, 0xfe, 0x10, 0x0c, 0x4b, 0xbd, 0xa2, 0xf1, 0x3a, 0x54,
		0x5a, 0x8f, 0x72, 0x3e, 0xbe, 0x18, 0xa4, 0x5a, 0xfc, 0x40, 0xb8, 0x01,
		0xd7, 0xcd, 0x26, 0x80, 0x2a, 0x9c, 0xc5, 0x66, 0x23, 0xa3, 0x77, 0x49,
		0xc2, 0x76, 0xbb, 0x46, 0x6a, 0xd9, 0xab, 0xa2, 0xb8, 0x25, 0x13, 0x5d,
		0x39, 0xa1, 0x76, 0x1d, 0x2a, 0x8f, 0x87, 0x95, 0xf4, 0x6f, 0xc3, 0xd6,
		0x49, 0x9c, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xaf, 0x12, 0xe0,
		0x23, 0x02, 0x00, 0x00,
	},
		"bash/fn.bash",
	)
}

func bash_herokuish_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x54,
		0x5d, 0x4f, 0xeb, 0x38, 0x10, 0x7d, 0x4e, 0x7e, 0xc5, 0x60, 0x55, 0x74,
		0x41, 0xca, 0xa2, 0x5d, 0xf1, 0x44, 0x95, 0x15, 0xa5, 0x5b, 0x2e, 0xe8,
		0x72, 0xa1, 0xa2, 0x80, 0x74, 0x05, 0xa8, 0x32, 0xf1, 0x34, 0xb5, 0x70,
		0x6d, 0xcb, 0x76, 0x4a, 0x11, 0xe2, 0xbf, 0x5f, 0x3b, 0x6d, 0x3e, 0xda,
		0x42, 0x9e, 0x9c, 0x39, 0x67, 0x66, 0x8e, 0x67, 0xc6, 0x13, 0xf3, 0x29,
		0x3c, 0x3e, 0x02, 0xe9, 0x9c, 0xf5, 0xc7, 0x17, 0x93, 0x87, 0xe1, 0xed,
		0xf8, 0xf2, 0xfa, 0xfc, 0x86, 0x40, 0x22, 0x1c, 0x90, 0x63, 0x02, 0xcf,
		0xcf, 0x3d, 0x70, 0x33, 0x94, 0x71, 0x84, 0xd9, 0x4c, 0x01, 0xd9, 0xdb,
		0x83, 0xdf, 0xaa, 0x30, 0x60, 0xdf, 0xad, 0xc3, 0x39, 0x9c, 0x51, 0x3b,
		0x03, 0x6e, 0x41, 0x15, 0x0e, 0xd4, 0x14, 0x18, 0x75, 0x78, 0x02, 0xad,
		0x58, 0x37, 0xd7, 0xa4, 0xe5, 0x39, 0x12, 0x48, 0x2d, 0x42, 0xa1, 0x73,
		0x43, 0x19, 0x82, 0x53, 0x2b, 0xff, 0x63, 0x50, 0x06, 0x72, 0x83, 0xde,
		0xd9, 0xfc, 0x1d, 0xf8, 0x4b, 0xee, 0xe0, 0xdf, 0x78, 0xca, 0xe3, 0xd8,
		0x1b, 0x99, 0x92, 0xe2, 0x1d, 0xa8, 0xd6, 0x13, 0x4d, 0xdd, 0x2c, 0x25,
		0x9d, 0x8f, 0xfe, 0x68, 0x34, 0x19, 0xf5, 0xef, 0x2e, 0x4e, 0x92, 0x23,
		0x6f, 0xfe, 0x24, 0x0d, 0x0b, 0xe5, 0xa2, 0x66, 0x0d, 0xaf, 0x1f, 0x2a,
		0x96, 0x9b, 0xeb, 0x23, 0x0f, 0xb5, 0x99, 0x2f, 0x05, 0x17, 0xac, 0xe6,
		0x9e, 0xdd, 0x5f, 0x5e, 0xfd, 0xdf, 0x66, 0x97, 0x70, 0x9b, 0x9f, 0xd1,
		0x6c, 0x86, 0x35, 0x7f, 0xd0, 0x1f, 0x5c, 0x0c, 0xdb, 0xfc, 0x12, 0xde,
		0x89, 0xaf, 0x69, 0xf6, 0xba, 0x99, 0x63, 0xd4, 0x1f, 0xfc, 0xdc, 0xc9,
		0x13, 0x68, 0xd6, 0x3b, 0xc7, 0x0c, 0x33, 0x41, 0x8d, 0x2f, 0x90, 0xd4,
		0x86, 0x2f, 0xb8, 0xc0, 0x1c, 0xd9, 0xa4, 0xb0, 0x68, 0x52, 0x22, 0xd5,
		0x8b, 0x62, 0xef, 0xe4, 0x6b, 0x4a, 0x6e, 0x54, 0xa1, 0x03, 0xa7, 0x3c,
		0xf8, 0x40, 0x28, 0x6d, 0x61, 0x30, 0x09, 0xa9, 0xed, 0x5f, 0x07, 0xf0,
		0x11, 0x47, 0xf3, 0x57, 0xc6, 0x0d, 0x24, 0x1a, 0x9e, 0xe2, 0x28, 0x22,
		0x9d, 0xaa, 0x9c, 0x64, 0xfd, 0x5f, 0x15, 0xae, 0xfa, 0x6f, 0xca, 0x53,
		0x59, 0x9a, 0x02, 0x6c, 0x70, 0xea, 0x2b, 0x92, 0xf8, 0x33, 0x8e, 0x5b,
		0x09, 0x2b, 0xa1, 0x0c, 0x6d, 0x96, 0x92, 0xf1, 0x4c, 0xbd, 0x59, 0x08,
		0x30, 0x58, 0x74, 0x8e, 0xcb, 0xdc, 0xd6, 0x83, 0x51, 0xb5, 0x33, 0x6d,
		0x54, 0x55, 0x50, 0xd5, 0xc3, 0xb4, 0x11, 0x58, 0x41, 0x4d, 0xcb, 0xd2,
		0xb6, 0xda, 0x0a, 0x6e, 0x3a, 0x94, 0xb6, 0xa5, 0x6f, 0x78, 0xd7, 0xcd,
		0x48, 0xbf, 0xba, 0xcb, 0x02, 0x8d, 0xe5, 0x4a, 0x7e, 0x77, 0x1b, 0x58,
		0xe3, 0x40, 0x25, 0x03, 0x5b, 0x68, 0xad, 0x8c, 0x43, 0x56, 0x5b, 0xb9,
		0x9c, 0xaa, 0x3a, 0x1b, 0xc3, 0x45, 0x19, 0xd2, 0x71, 0x27, 0x70, 0x15,
		0xb0, 0x04, 0x3a, 0xdd, 0x27, 0x7c, 0xfc, 0xe7, 0x47, 0x12, 0xbe, 0xff,
		0xba, 0xd0, 0x39, 0x0c, 0x24, 0x2e, 0x19, 0x4a, 0xb7, 0x62, 0xbd, 0xcd,
		0x7c, 0x87, 0x21, 0x0c, 0x15, 0x08, 0x2e, 0xb1, 0x07, 0x4c, 0xf9, 0xd2,
		0x57, 0x4f, 0x36, 0x98, 0x08, 0xa4, 0x29, 0x24, 0xc9, 0x61, 0xeb, 0xa1,
		0x46, 0x1b, 0xc1, 0xbb, 0x25, 0xcd, 0x5b, 0x51, 0x58, 0xdc, 0x06, 0xa1,
		0xfc, 0xba, 0x55, 0x2c, 0x0f, 0xfb, 0x47, 0x17, 0xf9, 0x09, 0xc6, 0xa0,
		0xa4, 0x3d, 0x63, 0x2b, 0x3d, 0xbe, 0x77, 0x05, 0x67, 0x39, 0x67, 0xde,
		0x63, 0x67, 0x48, 0x09, 0x74, 0x4e, 0x83, 0xdb, 0x9c, 0x72, 0x59, 0xd3,
		0x21, 0x41, 0x05, 0x9a, 0x6b, 0x9c, 0x52, 0x2e, 0x7a, 0x2b, 0xdd, 0x77,
		0xb7, 0xfd, 0xc1, 0x30, 0xac, 0x16, 0xd8, 0xdf, 0x87, 0x92, 0xb3, 0x8c,
		0xe3, 0x28, 0x9b, 0xb3, 0x04, 0x97, 0xa1, 0x8c, 0xe5, 0x9c, 0xd8, 0x0d,
		0xcb, 0xba, 0xac, 0x1b, 0xb4, 0x44, 0xda, 0xe6, 0x99, 0x01, 0xb9, 0xf7,
		0x9b, 0x25, 0xf4, 0x82, 0x4b, 0xeb, 0xa8, 0x10, 0x0d, 0x14, 0x26, 0xad,
		0x15, 0xaa, 0xb6, 0x27, 0xe5, 0xe9, 0x1b, 0x6c, 0x1d, 0xe5, 0x1b, 0x54,
		0x70, 0xeb, 0x76, 0xb4, 0x58, 0x51, 0xe4, 0x40, 0x7e, 0x51, 0x49, 0x73,
		0x0c, 0xfb, 0x4a, 0xf0, 0x8c, 0xba, 0x30, 0x0b, 0x01, 0xd8, 0x12, 0x11,
		0x4c, 0x09, 0x9f, 0x87, 0xf3, 0xae, 0x3d, 0x47, 0x89, 0xc6, 0x2f, 0xc3,
		0x5d, 0x64, 0x75, 0xde, 0xc9, 0xac, 0x8d, 0xca, 0xa6, 0x61, 0x52, 0xc8,
		0x88, 0x9a, 0x75, 0x19, 0x70, 0x89, 0x59, 0xe1, 0x10, 0x46, 0x6b, 0x6c,
		0x4b, 0x40, 0xe5, 0x92, 0xf8, 0x6b, 0x6e, 0x69, 0xa8, 0xa1, 0x10, 0xe2,
		0x6b, 0x44, 0x87, 0x34, 0x71, 0xe4, 0xc1, 0xb0, 0xcf, 0x49, 0x67, 0x3c,
		0xbc, 0x3a, 0x27, 0xbe, 0xf2, 0xe0, 0x27, 0xe8, 0xa8, 0x8c, 0x78, 0x10,
		0x45, 0x9b, 0x29, 0x3c, 0xeb, 0x94, 0xf4, 0x7a, 0x81, 0x10, 0xe2, 0xb6,
		0xf1, 0xf0, 0xdf, 0xc0, 0x87, 0x1e, 0x2a, 0x93, 0xfa, 0x8b, 0x11, 0xb2,
		0xb2, 0xfb, 0xf7, 0x62, 0x69, 0x16, 0x7f, 0xfe, 0x09, 0x00, 0x00, 0xff,
		0xff, 0x3d, 0x9f, 0x61, 0x9d, 0xae, 0x06, 0x00, 0x00,
	},
		"bash/herokuish.bash",
	)
}

func bash_procfile_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x94,
		0xcd, 0x6e, 0xe3, 0x36, 0x10, 0x80, 0xcf, 0xd6, 0x53, 0x0c, 0x04, 0x61,
		0xd7, 0x6e, 0x97, 0x56, 0xf6, 0xda, 0x85, 0x8b, 0x16, 0x48, 0x81, 0x1e,
		0xba, 0xbb, 0x45, 0xae, 0x69, 0x60, 0xd0, 0xe2, 0x58, 0x22, 0x42, 0x89,
		0x2a, 0x49, 0xe5, 0xa7, 0x69, 0xde, 0x7d, 0x87, 0xa4, 0x64, 0x8b, 0xb1,
		0x72, 0x12, 0x35, 0x9c, 0x19, 0x7e, 0xf3, 0x9b, 0xf5, 0x46, 0x57, 0x47,
		0xa9, 0x90, 0xf5, 0xdc, 0x58, 0x5c, 0x6f, 0xe0, 0x25, 0x5b, 0x09, 0xac,
		0x14, 0x37, 0x08, 0x02, 0x6d, 0xb5, 0xcb, 0xff, 0xd2, 0xfa, 0x1e, 0x86,
		0x1e, 0x2a, 0xdd, 0xb6, 0xbc, 0x13, 0x60, 0x9d, 0x91, 0x5d, 0x0d, 0x47,
		0x6d, 0x80, 0x6c, 0x9c, 0xac, 0x06, 0x52, 0x06, 0xef, 0x07, 0xad, 0x05,
		0xf7, 0xdc, 0x23, 0x1c, 0x8d, 0x6e, 0xe1, 0xef, 0xd1, 0x73, 0x7e, 0x76,
		0xe8, 0x2f, 0x77, 0x79, 0xf1, 0x99, 0x44, 0x15, 0x77, 0x90, 0x17, 0xbc,
		0xef, 0xf7, 0x3d, 0x77, 0x4d, 0x79, 0x52, 0x86, 0xff, 0xe1, 0x99, 0xb7,
		0x8a, 0xd5, 0xe8, 0xef, 0xbd, 0x41, 0x9e, 0xbd, 0x66, 0x67, 0x4c, 0xeb,
		0xe8, 0xcd, 0x25, 0xcc, 0x9b, 0xa1, 0x4b, 0x21, 0x26, 0xde, 0x04, 0x06,
		0x5c, 0x63, 0xf4, 0x50, 0x37, 0x80, 0x4f, 0x58, 0x2d, 0x93, 0x9d, 0x9e,
		0xf2, 0x2a, 0x50, 0xac, 0xd3, 0x0c, 0x4d, 0x50, 0x9b, 0x84, 0xca, 0xab,
		0xbe, 0x07, 0x35, 0x71, 0x70, 0x0b, 0x86, 0xbe, 0xba, 0xfd, 0x04, 0x43,
		0xd7, 0x1b, 0xf9, 0x40, 0x86, 0x35, 0x0a, 0x18, 0x2c, 0x1a, 0x78, 0x94,
		0xae, 0x81, 0x2d, 0xf9, 0xf3, 0xee, 0xb6, 0x94, 0x64, 0x3d, 0x98, 0x0a,
		0xc5, 0x1c, 0x27, 0x1a, 0xcb, 0xff, 0x90, 0x79, 0x8b, 0xd9, 0x85, 0x45,
		0xc7, 0x1a, 0xdd, 0xe2, 0x4c, 0xa4, 0x34, 0x17, 0x0c, 0xbb, 0x87, 0xb7,
		0xa2, 0xf1, 0x81, 0x6c, 0x95, 0x10, 0x14, 0xbf, 0x25, 0xc1, 0xf8, 0xf8,
		0x6c, 0x8c, 0xc6, 0x49, 0x47, 0x39, 0xcb, 0xaf, 0xa5, 0xad, 0xf4, 0x03,
		0x86, 0xba, 0xcf, 0x73, 0x6c, 0x89, 0x4f, 0x1e, 0xe1, 0xf6, 0x16, 0xd8,
		0x91, 0x12, 0x73, 0x18, 0xa4, 0x12, 0x6f, 0xeb, 0x79, 0x77, 0xf7, 0x85,
		0x92, 0x8e, 0x5d, 0xb6, 0x5a, 0x29, 0x5d, 0x71, 0x15, 0x0d, 0xe9, 0x2f,
		0x7c, 0x29, 0xe7, 0x6b, 0xdf, 0x0a, 0x4b, 0xb6, 0x53, 0x2b, 0xdc, 0xe3,
		0xb3, 0xa5, 0xf3, 0x13, 0x37, 0xb5, 0x05, 0xac, 0x1a, 0xbd, 0xa1, 0x67,
		0x57, 0xfe, 0x00, 0xf9, 0x49, 0x77, 0xcc, 0xfb, 0xc8, 0x05, 0xec, 0x57,
		0x28, 0x5e, 0xc2, 0xb1, 0x2c, 0xa1, 0xfc, 0x04, 0xaf, 0xde, 0xc4, 0xa0,
		0x1b, 0x0c, 0x81, 0x1c, 0xe5, 0x09, 0xdb, 0xa6, 0xd8, 0x5b, 0x83, 0x0a,
		0xb9, 0x5d, 0xc2, 0x16, 0x78, 0xe4, 0x83, 0x72, 0xfb, 0x09, 0x3f, 0xf9,
		0x5f, 0x0a, 0x63, 0xf2, 0x95, 0x84, 0x31, 0x59, 0x8d, 0x59, 0x8c, 0xd6,
		0x97, 0xc1, 0x11, 0x5a, 0x5e, 0x24, 0x2f, 0x78, 0x22, 0xf8, 0xf0, 0x01,
		0xfe, 0xa1, 0xdb, 0x31, 0xf6, 0xeb, 0x78, 0x9f, 0x56, 0x24, 0xcc, 0x65,
		0x61, 0xe9, 0xe9, 0xca, 0xa1, 0xd8, 0x77, 0xbc, 0xc5, 0x98, 0x8c, 0xc4,
		0xdb, 0x3b, 0x49, 0x89, 0x7e, 0xbf, 0xe9, 0x0b, 0x97, 0x43, 0x27, 0xd2,
		0x31, 0x9c, 0x1a, 0x2c, 0xb6, 0xc9, 0x98, 0x4c, 0x41, 0xd0, 0x24, 0x0b,
		0xe1, 0x27, 0x19, 0xf4, 0x4c, 0x08, 0xb2, 0xa3, 0x69, 0x52, 0x16, 0x4e,
		0x3a, 0x9b, 0x2f, 0x20, 0x74, 0x88, 0xe7, 0xa9, 0xd7, 0xc6, 0xcf, 0x3b,
		0xee, 0xc6, 0x3c, 0x4e, 0x2a, 0x65, 0x81, 0x21, 0x21, 0x42, 0x77, 0x18,
		0x18, 0x2f, 0x20, 0xc6, 0x96, 0x8e, 0x20, 0xb6, 0xd1, 0xbd, 0xf3, 0x55,
		0xed, 0x06, 0xa5, 0x6a, 0xa5, 0x0f, 0xd9, 0xaa, 0xbd, 0x17, 0xd2, 0x00,
		0xeb, 0xe7, 0xdb, 0xe6, 0x3c, 0x68, 0xe4, 0xdb, 0xc3, 0x85, 0x16, 0xf2,
		0x7c, 0x0b, 0x2a, 0xe5, 0x4f, 0x5b, 0xdb, 0x8c, 0xa4, 0x71, 0x2c, 0xc9,
		0xd5, 0xb4, 0xd8, 0x02, 0x56, 0xc3, 0x6d, 0x03, 0xcc, 0xa4, 0x7b, 0x6a,
		0x1c, 0xca, 0xc8, 0x35, 0x06, 0xf8, 0xe7, 0xf7, 0xaf, 0x7f, 0xec, 0xce,
		0x1c, 0xe4, 0xc0, 0x8f, 0x72, 0xab, 0x05, 0xb0, 0xa0, 0x3c, 0x63, 0xcc,
		0xe9, 0x3c, 0x1f, 0xd2, 0xbd, 0xd7, 0xf4, 0x7b, 0xb3, 0xd1, 0x8f, 0x1d,
		0xb0, 0x9b, 0xa5, 0xeb, 0x5f, 0x52, 0x51, 0x4d, 0x9b, 0xae, 0xcf, 0xe7,
		0x3e, 0x13, 0xc2, 0x74, 0x9f, 0x44, 0xce, 0xd8, 0xea, 0xfe, 0x5f, 0x0a,
		0xdf, 0xd3, 0xeb, 0x9b, 0xdf, 0xbf, 0x5d, 0x7f, 0xff, 0xfa, 0xf3, 0xe7,
		0xab, 0xab, 0xab, 0x8d, 0x2f, 0xc4, 0x59, 0xc1, 0x77, 0xd6, 0x2e, 0x1f,
		0x8a, 0x97, 0xa8, 0x4d, 0xbd, 0x94, 0xad, 0xb8, 0x10, 0xe1, 0x51, 0x0a,
		0xe7, 0xdf, 0x41, 0xd2, 0x06, 0x67, 0xac, 0x96, 0xbe, 0x2b, 0xa2, 0x4e,
		0x3e, 0x9e, 0xbc, 0x69, 0x1e, 0xb4, 0xc3, 0xee, 0xf3, 0x2d, 0xcd, 0x98,
		0x6d, 0x50, 0x29, 0x28, 0x0f, 0xb2, 0x2b, 0x0f, 0x3e, 0x9f, 0x51, 0x2a,
		0xa4, 0xe5, 0x07, 0x85, 0x54, 0x66, 0x6e, 0xed, 0xa3, 0x36, 0x62, 0x94,
		0x53, 0xd1, 0x2a, 0x64, 0x07, 0x2e, 0x42, 0x83, 0x47, 0x59, 0xa7, 0x59,
		0x65, 0x90, 0x3b, 0x8c, 0xb9, 0x8c, 0xc2, 0x21, 0x79, 0x3f, 0xca, 0xea,
		0x25, 0x19, 0x56, 0xda, 0xc2, 0xc7, 0x8f, 0xe3, 0x6f, 0xc4, 0x8f, 0xe7,
		0x8b, 0xca, 0x78, 0x71, 0x12, 0x49, 0xba, 0x4f, 0x43, 0x2d, 0x76, 0x89,
		0xc2, 0x65, 0x5d, 0x92, 0xfb, 0xd7, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff,
		0xb1, 0x4f, 0xcb, 0xd0, 0x87, 0x07, 0x00, 0x00,
	},
		"bash/procfile.bash",
	)
}

func bash_slug_bash() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x7c, 0x53,
		0x51, 0x6f, 0xd3, 0x30, 0x10, 0x7e, 0x6e, 0x7e, 0xc5, 0x61, 0x55, 0x5a,
		0xfb, 0xe0, 0x96, 0xf1, 0x8a, 0x82, 0x84, 0x0a, 0x42, 0x93, 0x10, 0x20,
		0xd6, 0x4a, 0x48, 0x63, 0xaa, 0x3c, 0xe7, 0x9a, 0x58, 0x72, 0x62, 0xcb,
		0x76, 0x46, 0x17, 0xe0, 0xbf, 0x73, 0x4e, 0xd2, 0x36, 0x61, 0x6b, 0xfb,
		0xd4, 0x9c, 0xef, 0xbe, 0xfb, 0xee, 0xbb, 0xfb, 0x12, 0x87, 0x22, 0x33,
		0x95, 0x7e, 0x02, 0xaf, 0xeb, 0x7c, 0x6b, 0x45, 0x28, 0x52, 0xb6, 0x0c,
		0xa5, 0x5d, 0xc6, 0xef, 0x45, 0xc8, 0x1b, 0x96, 0x24, 0xf1, 0x2f, 0x57,
		0xa5, 0x35, 0x2e, 0xcc, 0xe6, 0xf0, 0x3b, 0x99, 0x64, 0x28, 0xb5, 0x70,
		0x08, 0x19, 0x7a, 0x99, 0xb2, 0x9b, 0xf6, 0x05, 0x04, 0xe4, 0x8d, 0xb2,
		0x16, 0xb3, 0x16, 0x0a, 0x82, 0x70, 0x0f, 0x42, 0x6b, 0xd8, 0x39, 0x53,
		0xc2, 0xe6, 0xfb, 0x67, 0x30, 0x0e, 0x6e, 0xd7, 0x1f, 0x6e, 0xbe, 0x00,
		0x3b, 0x01, 0xd4, 0x4e, 0xa7, 0x6c, 0x7a, 0x4d, 0x11, 0xb5, 0x83, 0xbb,
		0x3b, 0x60, 0xd3, 0x99, 0xf6, 0xc0, 0xdf, 0xc3, 0x54, 0x58, 0xdb, 0xb2,
		0x99, 0x33, 0xb8, 0xbf, 0x7f, 0x0b, 0xa1, 0xc0, 0x2a, 0x99, 0x4c, 0x1c,
		0x86, 0xda, 0x55, 0x70, 0x9d, 0x4c, 0x50, 0x1f, 0x2a, 0x08, 0x63, 0x94,
		0x23, 0x29, 0x00, 0x9c, 0x50, 0x0c, 0x2c, 0x33, 0x7c, 0x5c, 0x56, 0x35,
		0xb1, 0xe0, 0x9c, 0x4a, 0xdd, 0x13, 0xbc, 0xa1, 0x02, 0xef, 0x24, 0x83,
		0x3f, 0x91, 0x20, 0xf0, 0x7d, 0xb3, 0xa2, 0xc8, 0xa1, 0x19, 0x8b, 0xb8,
		0x1e, 0x23, 0x88, 0x08, 0x67, 0x53, 0x76, 0x2a, 0xf9, 0xdb, 0x8b, 0x92,
		0x63, 0x85, 0x4e, 0x04, 0x7c, 0x49, 0x96, 0x4f, 0xfd, 0xdb, 0x45, 0x61,
		0x88, 0x33, 0x10, 0x61, 0x87, 0x15, 0x09, 0x68, 0x2d, 0xa1, 0x6b, 0x23,
		0x85, 0x06, 0xab, 0xf2, 0x66, 0x6b, 0x6c, 0x50, 0xa6, 0x6a, 0xb5, 0xf9,
		0x55, 0x28, 0x59, 0xb4, 0x51, 0x78, 0x77, 0x1a, 0xeb, 0x38, 0xf3, 0x20,
		0x3d, 0x65, 0x9c, 0xd7, 0x1e, 0xb9, 0x34, 0xa5, 0x75, 0xe8, 0x3d, 0xb7,
		0xce, 0xe4, 0x4e, 0x94, 0x69, 0xcc, 0xe9, 0xd8, 0xf7, 0x3d, 0x22, 0x1b,
		0x95, 0x57, 0xc6, 0xe1, 0xb0, 0x13, 0x69, 0xca, 0x77, 0x34, 0xf0, 0x43,
		0xad, 0x74, 0xd6, 0x8e, 0xbc, 0x5c, 0x9c, 0x32, 0x47, 0x4a, 0x3f, 0x03,
		0xa0, 0xde, 0x3f, 0xe0, 0x5c, 0x65, 0xd7, 0x3a, 0x4a, 0x3a, 0x1d, 0xd0,
		0x85, 0xe9, 0x33, 0x14, 0xf8, 0x49, 0xd8, 0x9c, 0xe3, 0x5e, 0xea, 0x3a,
		0xc3, 0xf4, 0x6a, 0x91, 0xab, 0x70, 0xd5, 0x05, 0x57, 0x23, 0x62, 0xac,
		0x0b, 0xca, 0x48, 0xf7, 0x78, 0xbe, 0x31, 0x08, 0xf4, 0x9b, 0x2c, 0x86,
		0x73, 0x6e, 0xbd, 0x6a, 0x90, 0x4e, 0x6d, 0x96, 0xd5, 0xc0, 0x6f, 0x0b,
		0x38, 0xe5, 0xd3, 0x9a, 0x65, 0x1d, 0x68, 0xe6, 0xeb, 0x39, 0x51, 0x0c,
		0x2a, 0x68, 0x04, 0xb6, 0x22, 0xf1, 0x94, 0x3e, 0x6c, 0x2c, 0xd6, 0x82,
		0xf2, 0x7d, 0x51, 0xfc, 0x62, 0xc7, 0x0b, 0xc0, 0xfd, 0x39, 0x5b, 0x7c,
		0x6c, 0x5f, 0xe0, 0x70, 0x22, 0xff, 0x6d, 0x3f, 0x98, 0xd6, 0x14, 0xb3,
		0x47, 0x25, 0xe0, 0xdb, 0x66, 0x3d, 0xef, 0xed, 0xf1, 0x75, 0xb3, 0xbe,
		0x60, 0x8f, 0x57, 0xdd, 0x6a, 0x06, 0xb3, 0xbe, 0xe8, 0x8d, 0xa8, 0xf3,
		0x65, 0x77, 0xbc, 0xbe, 0x60, 0x10, 0x5a, 0x21, 0x11, 0x02, 0xbe, 0x1e,
		0x77, 0xea, 0xa0, 0x46, 0x06, 0x19, 0xbe, 0x77, 0xb6, 0xf8, 0x17, 0x00,
		0x00, 0xff, 0xff, 0xa6, 0x2f, 0xb4, 0x12, 0x4f, 0x04, 0x00, 0x00,
	},
		"bash/slug.bash",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"bash/buildpack.bash": bash_buildpack_bash,
	"bash/cmd.bash": bash_cmd_bash,
	"bash/fn.bash": bash_fn_bash,
	"bash/herokuish.bash": bash_herokuish_bash,
	"bash/procfile.bash": bash_procfile_bash,
	"bash/slug.bash": bash_slug_bash,
}
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"bash": &_bintree_t{nil, map[string]*_bintree_t{
		"buildpack.bash": &_bintree_t{bash_buildpack_bash, map[string]*_bintree_t{
		}},
		"cmd.bash": &_bintree_t{bash_cmd_bash, map[string]*_bintree_t{
		}},
		"fn.bash": &_bintree_t{bash_fn_bash, map[string]*_bintree_t{
		}},
		"herokuish.bash": &_bintree_t{bash_herokuish_bash, map[string]*_bintree_t{
		}},
		"procfile.bash": &_bintree_t{bash_procfile_bash, map[string]*_bintree_t{
		}},
		"slug.bash": &_bintree_t{bash_slug_bash, map[string]*_bintree_t{
		}},
	}},
}}