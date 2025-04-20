package goja

var GT Tracer

type Tracer interface {
	String(tag TraceTag, vs ...any)
	Array(tag TraceTag, vs ...any)
	Operation(tag TraceTag, vs ...any)
}

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
	if GT != nil {
		GT.Operation(tag, vs...)
	}
}

func traceString(tag TraceTag, vs ...any) {
	if GT != nil {
		GT.String(tag, vs...)
	}
}

func traceArray(tag TraceTag, vs ...any) {
	if GT != nil {
		GT.Array(tag, vs...)
	}
}
