package io

type validator func(int) bool

type PlayerReader interface {
	ReadInt(validator) int
	ReadBool() bool
}

type PlayerWriter interface {
	Write(interface{})
}
