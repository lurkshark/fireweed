//go:build !cgo
#include "textflag.h"
TEXT ·spawnAddress(SB),NOSPLIT,$0-8
 MOVD $fireweed_spawn_anchor(SB), R0
 MOVD R0, ret+0(FP)
 RET
DATA _cgo_init(SB)/8, $fireweed_early_entry(SB)
GLOBL _cgo_init(SB), NOPTR, $8
