package io

import ()

type PlayerReader interface {
	ReadInt() int
	ReadBool() bool
}

type PlayerWriter interface {
	Write(interface{})
}
