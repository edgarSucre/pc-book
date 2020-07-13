package sample

import "pcbook/pb"

//NewKeyboard returns a random keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

//NewCPU returns a random CPU
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGz := randomFloat(2.0, 3.5)
	maxGz := randomFloat(minGz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGz:         minGz,
		MaxGz:         maxGz,
	}
	return cpu
}

//NewGPU returns a random GPU
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUBName(brand)

	minGz := randomFloat(1.0, 1.5)
	maxGz := randomFloat(minGz, 2.0)

	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGz:  minGz,
		MaxGz:  maxGz,
		Memory: memory,
	}
	return gpu
}
