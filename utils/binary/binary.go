package binary

import (
	"github.com/funny/binary"
)

/*   htonl()--"Host to Network Long"  	小端输入, 大端输出   	uint32  */
func HToNL(in uint32) uint32 {
	data := make([]byte, 4)
	binary.PutUint32LE(data, in)
	return binary.GetUint32BE(data)
}

/*    ntohl()--"Network to Host Long"  	大端输入, 小端输出   	uint32  */
func NToHL(in uint32) uint32 {
	data := make([]byte, 4)
	binary.PutUint32BE(data, in)
	return binary.GetUint32LE(data)
}

/*    htons()--"Host to Network Short" 	小端输入, 大端输出	uint16  */
func HToNS(in uint16) uint16 {
	data := make([]byte, 2)
	binary.PutUint16LE(data, in)
	return binary.GetUint16BE(data)
}

/*    ntohs()--"Network to Host Short" 	大端输入, 小端输出	uint16  */
func NToHS(in uint16) uint16 {
	data := make([]byte, 2)
	binary.PutUint16BE(data, in)
	return binary.GetUint16LE(data)
}
