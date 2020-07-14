package sample

import (
	"pcbook/pb"

	"github.com/golang/protobuf/ptypes"
)

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

	minGz := randomFloat64(2.0, 3.5)
	maxGz := randomFloat64(minGz, 5.0)

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

	minGz := randomFloat64(1.0, 1.5)
	maxGz := randomFloat64(minGz, 2.0)

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

//NewRAM returns a random amount of ram
func NewRAM() *pb.Memory {
	ram := &pb.Memory{
		Unit:  pb.Memory_GIGABYTE,
		Value: uint64(randomInt(4, 64)),
	}

	return ram
}

//NewSSD returns a ssd disk with random capacity
func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Unit:  pb.Memory_GIGABYTE,
			Value: uint64(randomInt(128, 1024)),
		},
	}

	return ssd
}

//NewHDD returns a hdd disk with random capcacity
func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Unit:  pb.Memory_TERABYTE,
			Value: uint64(randomInt(1, 6)),
		},
	}

	return hdd
}

//NewScreen returns a random screen
func NewScreen() *pb.Screen {
	heigh := randomInt(1080, 4320)
	width := heigh * 16 / 10

	screen := &pb.Screen{
		Multitouch: randomBool(),
		Panel:      randomScreenPanel(),
		Resolution: &pb.Screen_Resolution{
			Width: uint32(width),
			Heigh: uint32(heigh),
		},
		SizeInch: randomFloat32(13, 17),
	}

	return screen
}

//NewLaptop returns a generated laptop :D
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weigh: &pb.Laptop_WeighKg{
			WeighKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   ptypes.TimestampNow(),
	}

	return laptop
}
