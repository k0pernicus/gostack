package errors

import "errors"

var MaximumStackCapacityExceededError = errors.New("Maximum stack capacity has been exceeded")
var NoMoreContentError = errors.New("No more content - the current length is 0")
var MinCapacityError = errors.New("Canno't change current capacity - the length of the stack is greater than the capacity")
