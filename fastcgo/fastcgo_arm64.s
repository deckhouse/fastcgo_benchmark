#include "go_asm.h"
#include "textflag.h"

TEXT ·EmptyAsmFunction(SB), NOSPLIT, $0-0
    RET

TEXT ·SimpleCall0(SB), NOSPLIT, $0-0
    MOVD    fn+0(FP), R0
    CALL    R0
    RET

TEXT ·Call0(SB), NOSPLIT, $0-0
    MOVD    fn+0(FP), R0

    MOVD    RSP, R19

    MOVD    g_m(g), R9
    MOVD    m_g0(R9), R10
    MOVD    (g_sched+gobuf_sp)(R10), R9
    AND     $~15, R9

    MOVD    R9, RSP
    CALL    R0
    MOVD    R19, RSP

    RET

TEXT ·Call2(SB), NOSPLIT, $0-0
    MOVD    fn+0(FP), R2
    MOVD    fn+8(FP), R0
    MOVD    fn+16(FP), R1

    MOVD    RSP, R19

    MOVD    g_m(g), R9
    MOVD    m_g0(R9), R10
    MOVD    (g_sched+gobuf_sp)(R10), R9
    AND     $~15, R9

    MOVD    R9, RSP
    CALL    R2
    MOVD    R19, RSP

    RET
