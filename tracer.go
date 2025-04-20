package goja

import "fmt"

type TraceTag string

const (
	TraceTagADD  = "add"
	TraceTagSUB  = "sub"
	TraceTagXOR  = "xor"
	TraceTagOR   = "or"
	TraceTagMUL  = "mul"
	TraceTagEXP  = "exp"
	TraceTagDIV  = "div"
	TraceTagNEG  = "neg"
	TraceTagINC  = "inc"
	TraceTagDEC  = "dec"
	TraceTagAND  = "and"
	TraceTagBNOT = "bnot"
	TraceTagSAL  = "sal"
	TraceTagSAR  = "sar"
	TraceTagSHR  = "shr"

	TraceTagStringConcat    = "string_Concat"
	TraceTagStringCharAt    = "string_CharAt"
	TraceTagStringLength    = "string_Length"
	TraceTagStringSubstring = "string_Substring"

	TraceTagArraySet       = "array_set"
	TraceTagArraySetOwnIdx = "array_setOwnIdx"
	TraceTagArraySetOwnStr = "array_setOwnStr"
)

func traceOperation(tag TraceTag, vs ...any) {

}

func traceString(tag TraceTag, vs ...any) {
}

func traceArray(tag TraceTag, vs ...any) {
	fmt.Println(tag, vs)
}
