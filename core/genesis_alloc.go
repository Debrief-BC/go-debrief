// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xc0"
const testnetAllocData = "\xf9\x03\xa4\u0080\x01\xc2\x01\x01\xc2\x02\x01\xc2\x03\x01\xc2\x04\x01\xc2\x05\x01\xc2\x06\x01\xc2\a\x01\xc2\b\x01\xc2\t\x01\xc2\n\x80\xc2\v\x80\xc2\f\x80\xc2\r\x80\xc2\x0e\x80\xc2\x0f\x80\xc2\x10\x80\xc2\x11\x80\xc2\x12\x80\xc2\x13\x80\xc2\x14\x80\xc2\x15\x80\xc2\x16\x80\xc2\x17\x80\xc2\x18\x80\xc2\x19\x80\xc2\x1a\x80\xc2\x1b\x80\xc2\x1c\x80\xc2\x1d\x80\xc2\x1e\x80\xc2\x1f\x80\xc2 \x80\xc2!\x80\xc2\"\x80\xc2#\x80\xc2$\x80\xc2%\x80\xc2&\x80\xc2'\x80\xc2(\x80\xc2)\x80\xc2*\x80\xc2+\x80\xc2,\x80\xc2-\x80\xc2.\x80\xc2/\x80\xc20\x80\xc21\x80\xc22\x80\xc23\x80\xc24\x80\xc25\x80\xc26\x80\xc27\x80\xc28\x80\xc29\x80\xc2:\x80\xc2;\x80\xc2<\x80\xc2=\x80\xc2>\x80\xc2?\x80\xc2@\x80\xc2A\x80\xc2B\x80\xc2C\x80\xc2D\x80\xc2E\x80\xc2F\x80\xc2G\x80\xc2H\x80\xc2I\x80\xc2J\x80\xc2K\x80\xc2L\x80\xc2M\x80\xc2N\x80\xc2O\x80\xc2P\x80\xc2Q\x80\xc2R\x80\xc2S\x80\xc2T\x80\xc2U\x80\xc2V\x80\xc2W\x80\xc2X\x80\xc2Y\x80\xc2Z\x80\xc2[\x80\xc2\\\x80\xc2]\x80\xc2^\x80\xc2_\x80\xc2`\x80\xc2a\x80\xc2b\x80\xc2c\x80\xc2d\x80\xc2e\x80\xc2f\x80\xc2g\x80\xc2h\x80\xc2i\x80\xc2j\x80\xc2k\x80\xc2l\x80\xc2m\x80\xc2n\x80\xc2o\x80\xc2p\x80\xc2q\x80\xc2r\x80\xc2s\x80\xc2t\x80\xc2u\x80\xc2v\x80\xc2w\x80\xc2x\x80\xc2y\x80\xc2z\x80\xc2{\x80\xc2|\x80\xc2}\x80\xc2~\x80\xc2\u007f\x80\u00c1\x80\x80\u00c1\x81\x80\u00c1\x82\x80\u00c1\x83\x80\u00c1\x84\x80\u00c1\x85\x80\u00c1\x86\x80\u00c1\x87\x80\u00c1\x88\x80\u00c1\x89\x80\u00c1\x8a\x80\u00c1\x8b\x80\u00c1\x8c\x80\u00c1\x8d\x80\u00c1\x8e\x80\u00c1\x8f\x80\u00c1\x90\x80\u00c1\x91\x80\u00c1\x92\x80\u00c1\x93\x80\u00c1\x94\x80\u00c1\x95\x80\u00c1\x96\x80\u00c1\x97\x80\u00c1\x98\x80\u00c1\x99\x80\u00c1\x9a\x80\u00c1\x9b\x80\u00c1\x9c\x80\u00c1\x9d\x80\u00c1\x9e\x80\u00c1\x9f\x80\u00c1\xa0\x80\u00c1\xa1\x80\u00c1\xa2\x80\u00c1\xa3\x80\u00c1\xa4\x80\u00c1\xa5\x80\u00c1\xa6\x80\u00c1\xa7\x80\u00c1\xa8\x80\u00c1\xa9\x80\u00c1\xaa\x80\u00c1\xab\x80\u00c1\xac\x80\u00c1\xad\x80\u00c1\xae\x80\u00c1\xaf\x80\u00c1\xb0\x80\u00c1\xb1\x80\u00c1\xb2\x80\u00c1\xb3\x80\u00c1\xb4\x80\u00c1\xb5\x80\u00c1\xb6\x80\u00c1\xb7\x80\u00c1\xb8\x80\u00c1\xb9\x80\u00c1\xba\x80\u00c1\xbb\x80\u00c1\xbc\x80\u00c1\xbd\x80\u00c1\xbe\x80\u00c1\xbf\x80\u00c1\xc0\x80\u00c1\xc1\x80\u00c1\u0080\u00c1\u00c0\u00c1\u0100\u00c1\u0140\u00c1\u0180\u00c1\u01c0\u00c1\u0200\u00c1\u0240\u00c1\u0280\u00c1\u02c0\u00c1\u0300\u00c1\u0340\u00c1\u0380\u00c1\u03c0\u00c1\u0400\u00c1\u0440\u00c1\u0480\u00c1\u04c0\u00c1\u0500\u00c1\u0540\u00c1\u0580\u00c1\u05c0\u00c1\u0600\u00c1\u0640\u00c1\u0680\u00c1\u06c0\u00c1\u0700\u00c1\u0740\u00c1\u0780\u00c1\u07c0\u00c1\xe0\x80\u00c1\xe1\x80\u00c1\xe2\x80\u00c1\xe3\x80\u00c1\xe4\x80\u00c1\xe5\x80\u00c1\xe6\x80\u00c1\xe7\x80\u00c1\xe8\x80\u00c1\xe9\x80\u00c1\xea\x80\u00c1\xeb\x80\u00c1\xec\x80\u00c1\xed\x80\u00c1\xee\x80\u00c1\xef\x80\u00c1\xf0\x80\u00c1\xf1\x80\u00c1\xf2\x80\u00c1\xf3\x80\u00c1\xf4\x80\u00c1\xf5\x80\u00c1\xf6\x80\u00c1\xf7\x80\u00c1\xf8\x80\u00c1\xf9\x80\u00c1\xfa\x80\u00c1\xfb\x80\u00c1\xfc\x80\u00c1\xfd\x80\u00c1\xfe\x80\u00c1\xff\x80\u3507KT\xa8\xbd\x15)f\xd6?pk\xae\x1f\xfe\xb0A\x19!\xe5\x8d\f\x9f,\x9c\xd0Ft\xed\xea@\x00\x00\x00"
const rinkebyAllocData = "\xf9\x03\xb7\u0080\x01\xc2\x01\x01\xc2\x02\x01\xc2\x03\x01\xc2\x04\x01\xc2\x05\x01\xc2\x06\x01\xc2\a\x01\xc2\b\x01\xc2\t\x01\xc2\n\x01\xc2\v\x01\xc2\f\x01\xc2\r\x01\xc2\x0e\x01\xc2\x0f\x01\xc2\x10\x01\xc2\x11\x01\xc2\x12\x01\xc2\x13\x01\xc2\x14\x01\xc2\x15\x01\xc2\x16\x01\xc2\x17\x01\xc2\x18\x01\xc2\x19\x01\xc2\x1a\x01\xc2\x1b\x01\xc2\x1c\x01\xc2\x1d\x01\xc2\x1e\x01\xc2\x1f\x01\xc2 \x01\xc2!\x01\xc2\"\x01\xc2#\x01\xc2$\x01\xc2%\x01\xc2&\x01\xc2'\x01\xc2(\x01\xc2)\x01\xc2*\x01\xc2+\x01\xc2,\x01\xc2-\x01\xc2.\x01\xc2/\x01\xc20\x01\xc21\x01\xc22\x01\xc23\x01\xc24\x01\xc25\x01\xc26\x01\xc27\x01\xc28\x01\xc29\x01\xc2:\x01\xc2;\x01\xc2<\x01\xc2=\x01\xc2>\x01\xc2?\x01\xc2@\x01\xc2A\x01\xc2B\x01\xc2C\x01\xc2D\x01\xc2E\x01\xc2F\x01\xc2G\x01\xc2H\x01\xc2I\x01\xc2J\x01\xc2K\x01\xc2L\x01\xc2M\x01\xc2N\x01\xc2O\x01\xc2P\x01\xc2Q\x01\xc2R\x01\xc2S\x01\xc2T\x01\xc2U\x01\xc2V\x01\xc2W\x01\xc2X\x01\xc2Y\x01\xc2Z\x01\xc2[\x01\xc2\\\x01\xc2]\x01\xc2^\x01\xc2_\x01\xc2`\x01\xc2a\x01\xc2b\x01\xc2c\x01\xc2d\x01\xc2e\x01\xc2f\x01\xc2g\x01\xc2h\x01\xc2i\x01\xc2j\x01\xc2k\x01\xc2l\x01\xc2m\x01\xc2n\x01\xc2o\x01\xc2p\x01\xc2q\x01\xc2r\x01\xc2s\x01\xc2t\x01\xc2u\x01\xc2v\x01\xc2w\x01\xc2x\x01\xc2y\x01\xc2z\x01\xc2{\x01\xc2|\x01\xc2}\x01\xc2~\x01\xc2\u007f\x01\u00c1\x80\x01\u00c1\x81\x01\u00c1\x82\x01\u00c1\x83\x01\u00c1\x84\x01\u00c1\x85\x01\u00c1\x86\x01\u00c1\x87\x01\u00c1\x88\x01\u00c1\x89\x01\u00c1\x8a\x01\u00c1\x8b\x01\u00c1\x8c\x01\u00c1\x8d\x01\u00c1\x8e\x01\u00c1\x8f\x01\u00c1\x90\x01\u00c1\x91\x01\u00c1\x92\x01\u00c1\x93\x01\u00c1\x94\x01\u00c1\x95\x01\u00c1\x96\x01\u00c1\x97\x01\u00c1\x98\x01\u00c1\x99\x01\u00c1\x9a\x01\u00c1\x9b\x01\u00c1\x9c\x01\u00c1\x9d\x01\u00c1\x9e\x01\u00c1\x9f\x01\u00c1\xa0\x01\u00c1\xa1\x01\u00c1\xa2\x01\u00c1\xa3\x01\u00c1\xa4\x01\u00c1\xa5\x01\u00c1\xa6\x01\u00c1\xa7\x01\u00c1\xa8\x01\u00c1\xa9\x01\u00c1\xaa\x01\u00c1\xab\x01\u00c1\xac\x01\u00c1\xad\x01\u00c1\xae\x01\u00c1\xaf\x01\u00c1\xb0\x01\u00c1\xb1\x01\u00c1\xb2\x01\u00c1\xb3\x01\u00c1\xb4\x01\u00c1\xb5\x01\u00c1\xb6\x01\u00c1\xb7\x01\u00c1\xb8\x01\u00c1\xb9\x01\u00c1\xba\x01\u00c1\xbb\x01\u00c1\xbc\x01\u00c1\xbd\x01\u00c1\xbe\x01\u00c1\xbf\x01\u00c1\xc0\x01\u00c1\xc1\x01\u00c1\xc2\x01\u00c1\xc3\x01\u00c1\xc4\x01\u00c1\xc5\x01\u00c1\xc6\x01\u00c1\xc7\x01\u00c1\xc8\x01\u00c1\xc9\x01\u00c1\xca\x01\u00c1\xcb\x01\u00c1\xcc\x01\u00c1\xcd\x01\u00c1\xce\x01\u00c1\xcf\x01\u00c1\xd0\x01\u00c1\xd1\x01\u00c1\xd2\x01\u00c1\xd3\x01\u00c1\xd4\x01\u00c1\xd5\x01\u00c1\xd6\x01\u00c1\xd7\x01\u00c1\xd8\x01\u00c1\xd9\x01\u00c1\xda\x01\u00c1\xdb\x01\u00c1\xdc\x01\u00c1\xdd\x01\u00c1\xde\x01\u00c1\xdf\x01\u00c1\xe0\x01\u00c1\xe1\x01\u00c1\xe2\x01\u00c1\xe3\x01\u00c1\xe4\x01\u00c1\xe5\x01\u00c1\xe6\x01\u00c1\xe7\x01\u00c1\xe8\x01\u00c1\xe9\x01\u00c1\xea\x01\u00c1\xeb\x01\u00c1\xec\x01\u00c1\xed\x01\u00c1\xee\x01\u00c1\xef\x01\u00c1\xf0\x01\u00c1\xf1\x01\u00c1\xf2\x01\u00c1\xf3\x01\u00c1\xf4\x01\u00c1\xf5\x01\u00c1\xf6\x01\u00c1\xf7\x01\u00c1\xf8\x01\u00c1\xf9\x01\u00c1\xfa\x01\u00c1\xfb\x01\u00c1\xfc\x01\u00c1\xfd\x01\u00c1\xfe\x01\u00c1\xff\x01\xf6\x941\xb9\x8d\x14\x00{\xde\xe67)\x80\x86\x98\x8a\v\xbd1\x18E#\xa0\x02\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
const goerliAllocData = "\xf9\x04\x06\u0080\x01\xc2\x01\x01\xc2\x02\x01\xc2\x03\x01\xc2\x04\x01\xc2\x05\x01\xc2\x06\x01\xc2\a\x01\xc2\b\x01\xc2\t\x01\xc2\n\x01\xc2\v\x01\xc2\f\x01\xc2\r\x01\xc2\x0e\x01\xc2\x0f\x01\xc2\x10\x01\xc2\x11\x01\xc2\x12\x01\xc2\x13\x01\xc2\x14\x01\xc2\x15\x01\xc2\x16\x01\xc2\x17\x01\xc2\x18\x01\xc2\x19\x01\xc2\x1a\x01\xc2\x1b\x01\xc2\x1c\x01\xc2\x1d\x01\xc2\x1e\x01\xc2\x1f\x01\xc2 \x01\xc2!\x01\xc2\"\x01\xc2#\x01\xc2$\x01\xc2%\x01\xc2&\x01\xc2'\x01\xc2(\x01\xc2)\x01\xc2*\x01\xc2+\x01\xc2,\x01\xc2-\x01\xc2.\x01\xc2/\x01\xc20\x01\xc21\x01\xc22\x01\xc23\x01\xc24\x01\xc25\x01\xc26\x01\xc27\x01\xc28\x01\xc29\x01\xc2:\x01\xc2;\x01\xc2<\x01\xc2=\x01\xc2>\x01\xc2?\x01\xc2@\x01\xc2A\x01\xc2B\x01\xc2C\x01\xc2D\x01\xc2E\x01\xc2F\x01\xc2G\x01\xc2H\x01\xc2I\x01\xc2J\x01\xc2K\x01\xc2L\x01\xc2M\x01\xc2N\x01\xc2O\x01\xc2P\x01\xc2Q\x01\xc2R\x01\xc2S\x01\xc2T\x01\xc2U\x01\xc2V\x01\xc2W\x01\xc2X\x01\xc2Y\x01\xc2Z\x01\xc2[\x01\xc2\\\x01\xc2]\x01\xc2^\x01\xc2_\x01\xc2`\x01\xc2a\x01\xc2b\x01\xc2c\x01\xc2d\x01\xc2e\x01\xc2f\x01\xc2g\x01\xc2h\x01\xc2i\x01\xc2j\x01\xc2k\x01\xc2l\x01\xc2m\x01\xc2n\x01\xc2o\x01\xc2p\x01\xc2q\x01\xc2r\x01\xc2s\x01\xc2t\x01\xc2u\x01\xc2v\x01\xc2w\x01\xc2x\x01\xc2y\x01\xc2z\x01\xc2{\x01\xc2|\x01\xc2}\x01\xc2~\x01\xc2\u007f\x01\u00c1\x80\x01\u00c1\x81\x01\u00c1\x82\x01\u00c1\x83\x01\u00c1\x84\x01\u00c1\x85\x01\u00c1\x86\x01\u00c1\x87\x01\u00c1\x88\x01\u00c1\x89\x01\u00c1\x8a\x01\u00c1\x8b\x01\u00c1\x8c\x01\u00c1\x8d\x01\u00c1\x8e\x01\u00c1\x8f\x01\u00c1\x90\x01\u00c1\x91\x01\u00c1\x92\x01\u00c1\x93\x01\u00c1\x94\x01\u00c1\x95\x01\u00c1\x96\x01\u00c1\x97\x01\u00c1\x98\x01\u00c1\x99\x01\u00c1\x9a\x01\u00c1\x9b\x01\u00c1\x9c\x01\u00c1\x9d\x01\u00c1\x9e\x01\u00c1\x9f\x01\u00c1\xa0\x01\u00c1\xa1\x01\u00c1\xa2\x01\u00c1\xa3\x01\u00c1\xa4\x01\u00c1\xa5\x01\u00c1\xa6\x01\u00c1\xa7\x01\u00c1\xa8\x01\u00c1\xa9\x01\u00c1\xaa\x01\u00c1\xab\x01\u00c1\xac\x01\u00c1\xad\x01\u00c1\xae\x01\u00c1\xaf\x01\u00c1\xb0\x01\u00c1\xb1\x01\u00c1\xb2\x01\u00c1\xb3\x01\u00c1\xb4\x01\u00c1\xb5\x01\u00c1\xb6\x01\u00c1\xb7\x01\u00c1\xb8\x01\u00c1\xb9\x01\u00c1\xba\x01\u00c1\xbb\x01\u00c1\xbc\x01\u00c1\xbd\x01\u00c1\xbe\x01\u00c1\xbf\x01\u00c1\xc0\x01\u00c1\xc1\x01\u00c1\xc2\x01\u00c1\xc3\x01\u00c1\xc4\x01\u00c1\xc5\x01\u00c1\xc6\x01\u00c1\xc7\x01\u00c1\xc8\x01\u00c1\xc9\x01\u00c1\xca\x01\u00c1\xcb\x01\u00c1\xcc\x01\u00c1\xcd\x01\u00c1\xce\x01\u00c1\xcf\x01\u00c1\xd0\x01\u00c1\xd1\x01\u00c1\xd2\x01\u00c1\xd3\x01\u00c1\xd4\x01\u00c1\xd5\x01\u00c1\xd6\x01\u00c1\xd7\x01\u00c1\xd8\x01\u00c1\xd9\x01\u00c1\xda\x01\u00c1\xdb\x01\u00c1\xdc\x01\u00c1\xdd\x01\u00c1\xde\x01\u00c1\xdf\x01\u00c1\xe0\x01\u00c1\xe1\x01\u00c1\xe2\x01\u00c1\xe3\x01\u00c1\xe4\x01\u00c1\xe5\x01\u00c1\xe6\x01\u00c1\xe7\x01\u00c1\xe8\x01\u00c1\xe9\x01\u00c1\xea\x01\u00c1\xeb\x01\u00c1\xec\x01\u00c1\xed\x01\u00c1\xee\x01\u00c1\xef\x01\u00c1\xf0\x01\u00c1\xf1\x01\u00c1\xf2\x01\u00c1\xf3\x01\u00c1\xf4\x01\u00c1\xf5\x01\u00c1\xf6\x01\u00c1\xf7\x01\u00c1\xf8\x01\u00c1\xf9\x01\u00c1\xfa\x01\u00c1\xfb\x01\u00c1\xfc\x01\u00c1\xfd\x01\u00c1\xfe\x01\u00c1\xff\x01\xe0\x94L*\xe4\x82Y5\x05\xf0\x16<\xde\xfc\a>\x81\xc6<\xdaA\a\x8a\x15-\x02\xc7\xe1J\xf6\x80\x00\x00\xe0\x94\xa8\xe8\xf1G2e\x8eKQ\xe8q\x191\x05:\x8ai\xba\xf2\xb1\x8a\x15-\x02\xc7\xe1J\xf6\x80\x00\x00\xe1\x94\u0665\x17\x9f\t\x1d\x85\x05\x1d<\x98'\x85\xef\xd1E\\\uc199\x8b\bE\x95\x16\x14\x01HJ\x00\x00\x00\xe1\x94\u08bdBX\xd2v\x887\xba\xa2j(\xfeq\xdc\a\x9f\x84\u01cbJG\xe3\xc1$H\xf4\xad\x00\x00\x00"
