package binaryx

import (
	"bytes"
	"encoding/binary"
	fbinary "github.com/funny/binary"
)

/*   htonl()--"Host to Network Long"  	小端输入, 大端输出   	uint32  */
func HToNL(in uint32) uint32 {
	data := make([]byte, 4)
	fbinary.PutUint32LE(data, in)
	return fbinary.GetUint32BE(data)
}

/*    ntohl()--"Network to Host Long"  	大端输入, 小端输出   	uint32  */
func NToHL(in uint32) uint32 {
	data := make([]byte, 4)
	fbinary.PutUint32BE(data, in)
	return fbinary.GetUint32LE(data)
}

/*    htons()--"Host to Network Short" 	小端输入, 大端输出	uint16  */
func HToNS(in uint16) uint16 {
	data := make([]byte, 2)
	fbinary.PutUint16LE(data, in)
	return fbinary.GetUint16BE(data)
}

/*    ntohs()--"Network to Host Short" 	大端输入, 小端输出	uint16  */
func NToHS(in uint16) uint16 {
	data := make([]byte, 2)
	fbinary.PutUint16BE(data, in)
	return fbinary.GetUint16LE(data)
}

func ByteSliceToInt16Slice(data []byte) ([]int16, error) {
	var utt16 []int16 = make([]int16, len(data)/2)
	br := bytes.NewReader(data)
	//err := binaryx.Read(br, binaryx.BigEndian, &utt32)
	err := binary.Read(br, binary.LittleEndian, &utt16)
	if err != nil {
		return nil, err
	}

	return utt16, nil
}

func Int16SliceToByteSlice(data []int16) []byte {
	lenOfWav16 := len(data)
	bdata := make([]byte, 0)
	for i := 0; i < lenOfWav16; i++ {
		bdata = append(bdata, byte(data[i]&0xff))
		bdata = append(bdata, byte((data[i]>>8)&0xff))
	}
	return bdata
}
