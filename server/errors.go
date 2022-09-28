package server

import "errors"

var (
	ErrInvalidNick = errors.New("a same nick already exoists")
	ErrEmptyNick   = errors.New("an empty nick disallowed")
	ErrReadInput   = errors.New("unable to read msg from the client")
	ErrWrongNick   = errors.New("choose correct name with alpha numeric symbols")
)
