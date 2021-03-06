package opcode

var opcodeStrings = [...]string{
	NOP:              "nop",
	ESWAP:            "eswap",
	WIDEN:            "widen",
	VERSION:          "version",
	VLIST:            "vlist",
	BLOCK:            "block",
	COLLECT:          "collect",
	DROP:             "drop",
	MARK:             "mark",
	SCAN:             "scan",
	SMATCH:           "smatch",
	UTF8:             "utf8",
	BITS:             "bits",
	NIBBLES:          "nibbles",
	OCTETS:           "octets",
	NATLONG:          "natlong",
	IEEE754DP:        "ieee754dp",
	ISTR:             "istr",
	FSTR:             "fstr",
	NAME:             "name",
	APPLY:            "apply",
	OPNULL:           "null",
	MTC:              "mtc",
	SELF:             "self",
	SYSTEM:           "system",
	SKIP:             "skip",
	OMIT:             "omit",
	INFINITYN:        "infinityn",
	INFINITYP:        "infinityp",
	FALSE:            "false",
	TRUE:             "true",
	NONE:             "none",
	PASS:             "pass",
	INCONC:           "inconc",
	FAIL:             "fail",
	ERROR:            "error",
	IF:               "if",
	IFELSE:           "ifelse",
	RETURN:           "return",
	STOPI:            "stopi",
	FOR:              "for",
	WHILE:            "while",
	DOWHILE:          "dowhile",
	EXEC:             "exec",
	ALT:              "alt",
	INTERLEAVE:       "interleave",
	ELSE:             "else",
	REPEAT:           "repeat",
	BREAK:            "break",
	CONTINUE:         "continue",
	STEP:             "step",
	ACACHE:           "acache",
	POS:              "pos",
	ADD:              "add",
	NEG:              "neg",
	SUB:              "sub",
	MUL:              "mul",
	DIV:              "div",
	MOD:              "mod",
	REM:              "rem",
	AND:              "and",
	NOT:              "not",
	OR:               "or",
	XOR:              "xor",
	SHL:              "shl",
	SHR:              "shr",
	ROL:              "rol",
	ROR:              "ror",
	EQ:               "eq",
	GE:               "ge",
	GT:               "gt",
	LE:               "le",
	LT:               "lt",
	NE:               "ne",
	CAT:              "cat",
	GET:              "get",
	ASSIGN:           "assign",
	ASSIGND:          "assignd",
	IGET:             "iget",
	DEF:              "def",
	IDEF:             "idef",
	IFIELD:           "ifield",
	MOVE:             "move",
	ANY:              "any",
	ANYN:             "anyn",
	COMPLEMENT:       "complement",
	IFPRESENT:        "ifpresent",
	LENGTH:           "length",
	PATTERN:          "pattern",
	PERMUTATION:      "permutation",
	RANGE:            "range",
	SUBSET:           "subset",
	SUPERSET:         "superset",
	ALLFROM:          "allfrom",
	ALLFROMP:         "allfromp",
	VARDUP:           "vardup",
	VAR:              "var",
	TERM:             "term",
	FIELD:            "field",
	FIELDO:           "fieldo",
	IN:               "in",
	INOUT:            "inout",
	OUT:              "out",
	PERMITO:          "permito",
	PERMITT:          "permitt",
	PERMITP:          "permitp",
	ALTSTEP:          "altstep",
	ALTSTEPB:         "altstepb",
	ALTSTEPW:         "altstepw",
	ALTSTEPBW:        "altstepbw",
	CONST:            "const",
	CONSTW:           "constw",
	CONTROL:          "control",
	CONTROLW:         "controlw",
	GROUP:            "group",
	GROUPW:           "groupw",
	IMPORT:           "import",
	IMPORTW:          "importw",
	EXTERN:           "extern",
	SOURCE:           "source",
	MODULE:           "module",
	MODULEW:          "modulew",
	FUNCTION:         "function",
	FUNCTIONB:        "functionb",
	FUNCTIONW:        "functionw",
	FUNCTIONBW:       "functionbw",
	FUNCTIONV:        "functionv",
	FUNCTIONVB:       "functionvb",
	FUNCTIONVW:       "functionvw",
	FUNCTIONVBW:      "functionvbw",
	FUNCTIONX:        "functionx",
	FUNCTIONXW:       "functionxw",
	FUNCTIONXV:       "functionxv",
	FUNCTIONXVW:      "functionxvw",
	MPAR:             "mpar",
	MPARD:            "mpard",
	MPARW:            "mparw",
	MPARDW:           "mpardw",
	TEMPLATE:         "template",
	TEMPLATEW:        "templatew",
	TYPE:             "type",
	TYPEW:            "typew",
	TESTCASE:         "testcase",
	TESTCASES:        "testcases",
	TESTCASEW:        "testcasew",
	TESTCASESW:       "testcasesw",
	SIG:              "sig",
	SIGA:             "siga",
	SIGW:             "sigw",
	SIGAW:            "sigaw",
	SIGV:             "sigv",
	SIGVW:            "sigvw",
	SIGX:             "sigx",
	SIGXA:            "sigxa",
	SIGXW:            "sigxw",
	SIGXAW:           "sigxaw",
	SIGXV:            "sigxv",
	SIGXVW:           "sigxvw",
	DISPLAY:          "display",
	DISPLAYQ:         "displayq",
	DISPLAYO:         "displayo",
	DISPLAYQO:        "displayqo",
	ENCODE:           "encode",
	ENCODEQ:          "encodeq",
	ENCODEO:          "encodeo",
	ENCODEQO:         "encodeqo",
	EXTENSION:        "extension",
	EXTENSIONQ:       "extensionq",
	EXTENSIONO:       "extensiono",
	EXTENSIONQO:      "extensionqo",
	VARIANT:          "variant",
	VARIANTQ:         "variantq",
	VARIANTO:         "varianto",
	VARIANTQO:        "variantqo",
	ACTIVATE:         "activate",
	DEACTIVATE:       "deactivate",
	DEACTIVATEA:      "deactivatea",
	SPECPLC:          "specplc",
	BITSTRING:        "bitstring",
	BOOLEAN:          "boolean",
	CHARSTRING:       "charstring",
	CHARSTRINGU:      "charstringu",
	FLOAT:            "float",
	HEXSTRING:        "hexstring",
	INTEGER:          "integer",
	OCTETSTRING:      "octetstring",
	VERDICTTYPE:      "verdicttype",
	CLOSURETYPE:      "closuretype",
	ADDRESS:          "address",
	DEFAULT:          "default",
	TIMER:            "timer",
	ARRAY:            "array",
	SUBTYPE:          "subtype",
	ENUMERATED:       "enumerated",
	COMPONENT:        "component",
	COMPONENTX:       "componentx",
	PORTM:            "portm",
	PORTP:            "portp",
	PORTMA:           "portma",
	PORTPA:           "portpa",
	RECORD:           "record",
	SET:              "set",
	RECORDOF:         "recordof",
	SETOF:            "setof",
	UNION:            "union",
	ACTION:           "action",
	LOG:              "log",
	MATCH:            "match",
	VALUEOF:          "valueof",
	GETVERDICT:       "getverdict",
	SETVERDICT:       "setverdict",
	EXECUTE:          "execute",
	EXECUTEL:         "executel",
	CONNECT:          "connect",
	DISCONNECT:       "disconnect",
	DISCONNECTA:      "disconnecta",
	DISCONNECTAA:     "disconnectaa",
	MAP:              "map",
	UNMAP:            "unmap",
	UNMAPA:           "unmapa",
	UNMAPAA:          "unmapaa",
	ALIVE:            "alive",
	ALIVE1:           "alive1",
	ALIVEA:           "alivea",
	DONE:             "done",
	DONE1:            "done1",
	DONEA:            "donea",
	CREATE:           "create",
	CREATEA:          "createa",
	CREATEN:          "createn",
	CREATEAN:         "createan",
	EXECUTELD:        "executeld",
	EXECUTED:         "executed",
	TCSTOP:           "tcstop",
	CLOSURE:          "closure",
	RUNNING:          "running",
	RUNNING1C:        "running1c",
	RUNNINGAC:        "runningac",
	KILLED:           "killed",
	KILLED1:          "killed1",
	KILLEDA:          "killeda",
	START:            "start",
	STARTC:           "startc",
	STARTD:           "startd",
	STARTAP:          "startap",
	STOP:             "stop",
	STOPAC:           "stopac",
	STOPAP:           "stopap",
	STOPAT:           "stopat",
	READ:             "read",
	RUNNING1T:        "running1t",
	TIMEOUT:          "timeout",
	TIMEOUT1:         "timeout1",
	NOW:              "now",
	WAIT:             "wait",
	CLEAR:            "clear",
	CLEARA:           "cleara",
	HALT:             "halt",
	HALTA:            "halta",
	KILL:             "kill",
	KILLA:            "killa",
	SEND:             "send",
	SEND1:            "send1",
	SENDN:            "sendn",
	SENDA:            "senda",
	RECEIVE:          "receive",
	RECEIVEC:         "receivec",
	RECEIVE1:         "receive1",
	RECEIVEC1:        "receivec1",
	TRIGGER:          "trigger",
	TRIGGER1:         "trigger1",
	CHECK:            "check",
	CHECK1:           "check1",
	MAP3:             "map3",
	MAP4:             "map4",
	CALL:             "call",
	CALL1:            "call1",
	CALLN:            "calln",
	CALLA:            "calla",
	CALLB:            "callb",
	CALLB1:           "callb1",
	CALLBN:           "callbn",
	CALLBA:           "callba",
	CALLW:            "callw",
	CALLW1:           "callw1",
	CALLWN:           "callwn",
	CALLWA:           "callwa",
	GETCALL:          "getcall",
	GETCALLC:         "getcallc",
	GETCALL1:         "getcall1",
	GETCALLC1:        "getcallc1",
	REPLY:            "reply",
	REPLY1:           "reply1",
	REPLYN:           "replyn",
	REPLYA:           "replya",
	REPLYV:           "replyv",
	REPLYV1:          "replyv1",
	REPLYVN:          "replyvn",
	REPLYVA:          "replyva",
	GETREPLY:         "getreply",
	GETREPLYC:        "getreplyc",
	GETREPLY1:        "getreply1",
	GETREPLYC1:       "getreplyc1",
	RAISE:            "raise",
	RAISE1:           "raise1",
	RAISEN:           "raisen",
	RAISEA:           "raisea",
	CATCH:            "catch",
	CATCHC:           "catchc",
	CATCH1:           "catch1",
	CATCHC1:          "catchc1",
	CHECKSTATE:       "checkstate",
	CHECKSTATEAL:     "checkstateal",
	CHECKSTATEAN:     "checkstatean",
	VALUE:            "value",
	PARAM:            "param",
	SENDER:           "sender",
	TIMESTAMP:        "timestamp",
	INT2CHAR:         "int2char",
	INT2UNICHAR:      "int2unichar",
	INT2BIT:          "int2bit",
	INT2HEX:          "int2hex",
	INT2OCT:          "int2oct",
	INT2STR:          "int2str",
	INT2FLOAT:        "int2float",
	FLOAT2INT:        "float2int",
	CHAR2INT:         "char2int",
	CHAR2OCT:         "char2oct",
	UNICHAR2INT:      "unichar2int",
	BIT2INT:          "bit2int",
	BIT2HEX:          "bit2hex",
	BIT2OCT:          "bit2oct",
	BIT2STR:          "bit2str",
	HEX2INT:          "hex2int",
	HEX2BIT:          "hex2bit",
	HEX2OCT:          "hex2oct",
	HEX2STR:          "hex2str",
	OCT2INT:          "oct2int",
	OCT2BIT:          "oct2bit",
	OCT2HEX:          "oct2hex",
	OCT2STR:          "oct2str",
	OCT2CHR:          "oct2chr",
	STR2INT:          "str2int",
	STR2OCT:          "str2oct",
	STR2FLOAT:        "str2float",
	LENGTHOF:         "lengthof",
	SIZEOF:           "sizeof",
	ISPRESENT:        "ispresent",
	ISCHOSEN:         "ischosen",
	REGEXP:           "regexp",
	SUBSTR:           "substr",
	REPLACE:          "replace",
	RND:              "rnd",
	RNDS:             "rnds",
	ENUM2INT:         "enum2int",
	ISVALUE:          "isvalue",
	ENCVALUE:         "encvalue",
	DECVALUE:         "decvalue",
	TESTCASENAME:     "testcasename",
	INT2ENUM:         "int2enum",
	XINT2ENUM:        "xint2enum",
	STR2HEX:          "str2hex",
	ISBOUND:          "isbound",
	VAL2STR:          "val2str",
	REF_INT2CHAR:     "Rint2char",
	REF_INT2UNICHAR:  "Rint2unichar",
	REF_INT2BIT:      "Rint2bit",
	REF_INT2HEX:      "Rint2hex",
	REF_INT2OCT:      "Rint2oct",
	REF_INT2STR:      "Rint2str",
	REF_INT2FLOAT:    "Rint2float",
	REF_FLOAT2INT:    "Rfloat2int",
	REF_CHAR2INT:     "Rchar2int",
	REF_CHAR2OCT:     "Rchar2oct",
	REF_UNICHAR2INT:  "Runichar2int",
	REF_BIT2INT:      "Rbit2int",
	REF_BIT2HEX:      "Rbit2hex",
	REF_BIT2OCT:      "Rbit2oct",
	REF_BIT2STR:      "Rbit2str",
	REF_HEX2INT:      "Rhex2int",
	REF_HEX2BIT:      "Rhex2bit",
	REF_HEX2OCT:      "Rhex2oct",
	REF_HEX2STR:      "Rhex2str",
	REF_OCT2INT:      "Roct2int",
	REF_OCT2BIT:      "Roct2bit",
	REF_OCT2HEX:      "Roct2hex",
	REF_OCT2STR:      "Roct2str",
	REF_OCT2CHR:      "Roct2chr",
	REF_STR2INT:      "Rstr2int",
	REF_STR2OCT:      "Rstr2oct",
	REF_STR2FLOAT:    "Rstr2float",
	REF_LENGTHOF:     "Rlengthof",
	REF_SIZEOF:       "Rsizeof",
	REF_VAL2STR:      "Rval2str",
	REF_ISPRESENT:    "Rispresent",
	REF_ISCHOSEN:     "Rischosen",
	REF_REGEXP:       "Rregexp",
	REF_SUBSTR:       "Rsubstr",
	REF_REPLACE:      "Rreplace",
	REF_RND:          "Rrnd",
	REF_RNDS:         "Rrnds",
	REF_ENUM2INT:     "Renum2int",
	REF_ISVALUE:      "Risvalue",
	REF_ENCVALUE:     "Rencvalue",
	REF_DECVALUE:     "Rdecvalue",
	REF_TESTCASENAME: "Rtestcasename",
	REF_INT2ENUM:     "Rint2enum",
	REF_XINT2ENUM:    "Rxint2enum",
	REF_STR2HEX:      "Rstr2hex",
	REF_ISBOUND:      "Risbound",
	PROFILE_TIME:     "profile_time",
	AT_DEFAULT:       "at_default",
}
