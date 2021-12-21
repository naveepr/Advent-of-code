package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func parseInput(scanner *bufio.Scanner) []string {
	var matrix []string
	for scanner.Scan() {
		s := scanner.Text()
		matrix = append(matrix, s)
	}
	return matrix
}

func hexToBits(hexastr string) string {
	var binaryStr string
	var BCD = [16]string{
		"0000", "0001", "0010", "0011", "0100", "0101", "0110", "0111",
		"1000", "1001", "1010", "1011", "1100", "1101", "1110", "1111"}

	for i := 0; i < len(hexastr); i++ {
		index, err := strconv.ParseUint(hexastr[i:i+1], 16, 32)
		check(err)
		binaryStr += BCD[index]
	}
	return binaryStr
}

func readBits(data []byte, startpos, count int64) int64 {
	var out int64
	if startpos+count > int64(len(data)) {
		log.Fatal("Data out of bound")
	}
	for _, bit := range data[startpos : startpos+count] {
		out <<= 1
		a := bit - '0'
		out |= int64(a)
	}

	return out
}

func readNumber(data []byte, startpos int64) (out int64, count int64) {
	out = 0
	for {
		part := readBits(data, startpos, 5)
		out <<= 4
		out |= int64(part & 0x0f)
		startpos += 5
		count += 5
		if (part & 0x10) == 0 {
			break
		}
	}

	return out, count
}

type Literal struct {
	Version int64
	Typeid  int64
	Value   int64
}

type Operand struct {
	Version  int64
	Typeid   int64
	Lengthid int64
	Length   int64
	Packets  []interface{}
}

func readPacket(bytes []byte, startpos int64) (l interface{}, count int64) {
	var op Operand
	n := startpos
	version := readBits(bytes, n, 3)
	n += 3

	typeid := readBits(bytes, n, 3)
	n += 3

	switch typeid {
	case 4:
		/* Literal */
		value, count := readNumber(bytes, n)
		n += count
		return Literal{
			Version: version,
			Typeid:  typeid,
			Value:   value,
		}, n - startpos
	default:
		lengthid := readBits(bytes, n, 1)
		n += 1
		op = Operand{
			Version:  version,
			Typeid:   typeid,
			Lengthid: lengthid,
			Length:   0,
			Packets:  nil,
		}
		if lengthid == 0 {
			length := readBits(bytes, n, 15)
			n += 15
			op.Length = length
			for packetlen := int64(0); packetlen < length; {
				packet, plen := readPacket(bytes, n)
				packetlen += plen
				n += plen
				op.Packets = append(op.Packets, packet)
			}

		} else {
			length := readBits(bytes, n, 11)
			n += 11
			op.Length = length

			for numSubpackets := int64(0); numSubpackets < length; numSubpackets++ {
				packet, plen := readPacket(bytes, n)
				n += plen
				op.Packets = append(op.Packets, packet)
			}
		}

		return op, n - startpos
	}

}

func getVersionSumPacket(packet interface{}) int64 {
	var version int64

	switch packets := packet.(type) {
	case Literal:
		version += packets.Version
	case Operand:
		var res int64
		version += packets.Version

		for _, p := range packets.Packets {
			res += getVersionSumPacket(p)
		}

		version += res
	}

	return version
}

func eval(packet interface{}) int64 {

	switch p := packet.(type) {
	case Literal:
		return p.Value
	case Operand:
		switch p.Typeid {
		case 7:
			p1 := eval(p.Packets[0])
			p2 := eval(p.Packets[1])
			if p1 == p2 {
				return 1
			} else {
				return 0
			}
		case 6:
			p1 := eval(p.Packets[0])
			p2 := eval(p.Packets[1])
			if p1 < p2 {
				return 1
			} else {
				return 0
			}
		case 5:
			p1 := eval(p.Packets[0])
			p2 := eval(p.Packets[1])
			if p1 > p2 {
				return 1
			} else {
				return 0
			}
		case 3:
			maxVal := math.MinInt64
			for _, subpacket := range p.Packets {
				maxVal = max(maxVal, int(eval(subpacket)))
			}
			return int64(maxVal)
		case 2:
			minVal := math.MaxInt64
			for _, subpacket := range p.Packets {
				minVal = min(minVal, int(eval(subpacket)))
			}
			return int64(minVal)
		case 1:
			prodVal := int64(1)
			for _, subpacket := range p.Packets {
				prodVal *= eval(subpacket)
			}
			return prodVal

		case 0:
			var sumVal int64
			for _, subpacket := range p.Packets {
				sumVal += eval(subpacket)
			}
			return sumVal

		default:
			break
		}
	default:
		break
	}

	return 0
}

func Solve(hexastr string) (int64, int64) {
	var res1, res2 int64

	binaryStr := hexToBits(hexastr)
	fmt.Println(binaryStr)

	packet, plen := readPacket([]byte(binaryStr), 0)

	fmt.Println(packet, plen)

	res1 = getVersionSumPacket(packet)
	res2 = eval(packet)
	return res1, res2
}

func main() {
	arg := os.Args
	f, err := os.Open(arg[1])
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	matrix := parseInput(scanner)
	for _, s := range matrix {
		fmt.Println(s)
		ans1, ans2 := Solve(s)
		fmt.Println("Answer is ", ans1, ans2)
	}
}
