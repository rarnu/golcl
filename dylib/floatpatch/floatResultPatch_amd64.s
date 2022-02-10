#include "textflag.h"

// float32 Getfloat32();
TEXT ·Getfloat32(SB),NOSPLIT|NOFRAME,$0
    MOVL X0, ret(FP)
    RET

// float64 Getfloat64();
TEXT ·Getfloat64(SB),NOSPLIT|NOFRAME,$0
    MOVQ  X0, ret(FP)
    RET
