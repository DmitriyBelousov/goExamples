package main

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64

byte

rune

float32 float64

complex64 complex128

x := 1
x :=3

const limit = 512
const top uint16 = 1421

const (
	Cyan = 0
	Magenta = 1
	Yellow = 2
)

const (
	Cyan = iota // 0
	Magenta // 1
	Yellow // 2
)

type BitFlag int
const (
	Active BitFlag = 1 << iota // 1 << 0 == 1
	Send // Неявно BitFlag = 1 << iota // 1 << 1 == 2
	Receive // Неявно BitFlag = 1 << iota // 1 << 2 == 4
)
flag := Active | Send