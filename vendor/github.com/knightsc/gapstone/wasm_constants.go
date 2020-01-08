/*
Gapstone is a Go binding for the Capstone disassembly library. For examples,
try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Scott Knight
	License: BSD style - see LICENSE file for details
    (c) 2019 Scott Knight

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
	Command: ./genconst ../capstone/bindings/python/capstone
	2019-12-30T10:12:58-05:00

*/

package gapstone

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

// For Capstone Engine. AUTO-GENERATED FILE, DO NOT EDIT [wasm_const.py]
const (
	WASM_OP_INVALID              = C.WASM_OP_INVALID
	WASM_OP_NONE                 = C.WASM_OP_NONE
	WASM_OP_INT7                 = C.WASM_OP_INT7
	WASM_OP_VARUINT32            = C.WASM_OP_VARUINT32
	WASM_OP_VARUINT64            = C.WASM_OP_VARUINT64
	WASM_OP_UINT32               = C.WASM_OP_UINT32
	WASM_OP_UINT64               = C.WASM_OP_UINT64
	WASM_OP_IMM                  = C.WASM_OP_IMM
	WASM_OP_BRTABLE              = C.WASM_OP_BRTABLE
	WASM_INS_UNREACHABLE         = C.WASM_INS_UNREACHABLE
	WASM_INS_NOP                 = C.WASM_INS_NOP
	WASM_INS_BLOCK               = C.WASM_INS_BLOCK
	WASM_INS_LOOP                = C.WASM_INS_LOOP
	WASM_INS_IF                  = C.WASM_INS_IF
	WASM_INS_ELSE                = C.WASM_INS_ELSE
	WASM_INS_END                 = C.WASM_INS_END
	WASM_INS_BR                  = C.WASM_INS_BR
	WASM_INS_BR_IF               = C.WASM_INS_BR_IF
	WASM_INS_BR_TABLE            = C.WASM_INS_BR_TABLE
	WASM_INS_RETURN              = C.WASM_INS_RETURN
	WASM_INS_CALL                = C.WASM_INS_CALL
	WASM_INS_CALL_INDIRECT       = C.WASM_INS_CALL_INDIRECT
	WASM_INS_DROP                = C.WASM_INS_DROP
	WASM_INS_SELECT              = C.WASM_INS_SELECT
	WASM_INS_GET_LOCAL           = C.WASM_INS_GET_LOCAL
	WASM_INS_SET_LOCAL           = C.WASM_INS_SET_LOCAL
	WASM_INS_TEE_LOCAL           = C.WASM_INS_TEE_LOCAL
	WASM_INS_GET_GLOBAL          = C.WASM_INS_GET_GLOBAL
	WASM_INS_SET_GLOBAL          = C.WASM_INS_SET_GLOBAL
	WASM_INS_I32_LOAD            = C.WASM_INS_I32_LOAD
	WASM_INS_I64_LOAD            = C.WASM_INS_I64_LOAD
	WASM_INS_F32_LOAD            = C.WASM_INS_F32_LOAD
	WASM_INS_F64_LOAD            = C.WASM_INS_F64_LOAD
	WASM_INS_I32_LOAD8_S         = C.WASM_INS_I32_LOAD8_S
	WASM_INS_I32_LOAD8_U         = C.WASM_INS_I32_LOAD8_U
	WASM_INS_I32_LOAD16_S        = C.WASM_INS_I32_LOAD16_S
	WASM_INS_I32_LOAD16_U        = C.WASM_INS_I32_LOAD16_U
	WASM_INS_I64_LOAD8_S         = C.WASM_INS_I64_LOAD8_S
	WASM_INS_I64_LOAD8_U         = C.WASM_INS_I64_LOAD8_U
	WASM_INS_I64_LOAD16_S        = C.WASM_INS_I64_LOAD16_S
	WASM_INS_I64_LOAD16_U        = C.WASM_INS_I64_LOAD16_U
	WASM_INS_I64_LOAD32_S        = C.WASM_INS_I64_LOAD32_S
	WASM_INS_I64_LOAD32_U        = C.WASM_INS_I64_LOAD32_U
	WASM_INS_I32_STORE           = C.WASM_INS_I32_STORE
	WASM_INS_I64_STORE           = C.WASM_INS_I64_STORE
	WASM_INS_F32_STORE           = C.WASM_INS_F32_STORE
	WASM_INS_F64_STORE           = C.WASM_INS_F64_STORE
	WASM_INS_I32_STORE8          = C.WASM_INS_I32_STORE8
	WASM_INS_I32_STORE16         = C.WASM_INS_I32_STORE16
	WASM_INS_I64_STORE8          = C.WASM_INS_I64_STORE8
	WASM_INS_I64_STORE16         = C.WASM_INS_I64_STORE16
	WASM_INS_I64_STORE32         = C.WASM_INS_I64_STORE32
	WASM_INS_CURRENT_MEMORY      = C.WASM_INS_CURRENT_MEMORY
	WASM_INS_GROW_MEMORY         = C.WASM_INS_GROW_MEMORY
	WASM_INS_I32_CONST           = C.WASM_INS_I32_CONST
	WASM_INS_I64_CONST           = C.WASM_INS_I64_CONST
	WASM_INS_F32_CONST           = C.WASM_INS_F32_CONST
	WASM_INS_F64_CONST           = C.WASM_INS_F64_CONST
	WASM_INS_I32_EQZ             = C.WASM_INS_I32_EQZ
	WASM_INS_I32_EQ              = C.WASM_INS_I32_EQ
	WASM_INS_I32_NE              = C.WASM_INS_I32_NE
	WASM_INS_I32_LT_S            = C.WASM_INS_I32_LT_S
	WASM_INS_I32_LT_U            = C.WASM_INS_I32_LT_U
	WASM_INS_I32_GT_S            = C.WASM_INS_I32_GT_S
	WASM_INS_I32_GT_U            = C.WASM_INS_I32_GT_U
	WASM_INS_I32_LE_S            = C.WASM_INS_I32_LE_S
	WASM_INS_I32_LE_U            = C.WASM_INS_I32_LE_U
	WASM_INS_I32_GE_S            = C.WASM_INS_I32_GE_S
	WASM_INS_I32_GE_U            = C.WASM_INS_I32_GE_U
	WASM_INS_I64_EQZ             = C.WASM_INS_I64_EQZ
	WASM_INS_I64_EQ              = C.WASM_INS_I64_EQ
	WASM_INS_I64_NE              = C.WASM_INS_I64_NE
	WASM_INS_I64_LT_S            = C.WASM_INS_I64_LT_S
	WASM_INS_I64_LT_U            = C.WASM_INS_I64_LT_U
	WASM_INS_I64_GT_U            = C.WASM_INS_I64_GT_U
	WASM_INS_I64_LE_S            = C.WASM_INS_I64_LE_S
	WASM_INS_I64_LE_U            = C.WASM_INS_I64_LE_U
	WASM_INS_I64_GE_S            = C.WASM_INS_I64_GE_S
	WASM_INS_I64_GE_U            = C.WASM_INS_I64_GE_U
	WASM_INS_F32_EQ              = C.WASM_INS_F32_EQ
	WASM_INS_F32_NE              = C.WASM_INS_F32_NE
	WASM_INS_F32_LT              = C.WASM_INS_F32_LT
	WASM_INS_F32_GT              = C.WASM_INS_F32_GT
	WASM_INS_F32_LE              = C.WASM_INS_F32_LE
	WASM_INS_F32_GE              = C.WASM_INS_F32_GE
	WASM_INS_F64_EQ              = C.WASM_INS_F64_EQ
	WASM_INS_F64_NE              = C.WASM_INS_F64_NE
	WASM_INS_F64_LT              = C.WASM_INS_F64_LT
	WASM_INS_F64_GT              = C.WASM_INS_F64_GT
	WASM_INS_F64_LE              = C.WASM_INS_F64_LE
	WASM_INS_F64_GE              = C.WASM_INS_F64_GE
	WASM_INS_I32_CLZ             = C.WASM_INS_I32_CLZ
	WASM_INS_I32_CTZ             = C.WASM_INS_I32_CTZ
	WASM_INS_I32_POPCNT          = C.WASM_INS_I32_POPCNT
	WASM_INS_I32_ADD             = C.WASM_INS_I32_ADD
	WASM_INS_I32_SUB             = C.WASM_INS_I32_SUB
	WASM_INS_I32_MUL             = C.WASM_INS_I32_MUL
	WASM_INS_I32_DIV_S           = C.WASM_INS_I32_DIV_S
	WASM_INS_I32_DIV_U           = C.WASM_INS_I32_DIV_U
	WASM_INS_I32_REM_S           = C.WASM_INS_I32_REM_S
	WASM_INS_I32_REM_U           = C.WASM_INS_I32_REM_U
	WASM_INS_I32_AND             = C.WASM_INS_I32_AND
	WASM_INS_I32_OR              = C.WASM_INS_I32_OR
	WASM_INS_I32_XOR             = C.WASM_INS_I32_XOR
	WASM_INS_I32_SHL             = C.WASM_INS_I32_SHL
	WASM_INS_I32_SHR_S           = C.WASM_INS_I32_SHR_S
	WASM_INS_I32_SHR_U           = C.WASM_INS_I32_SHR_U
	WASM_INS_I32_ROTL            = C.WASM_INS_I32_ROTL
	WASM_INS_I32_ROTR            = C.WASM_INS_I32_ROTR
	WASM_INS_I64_CLZ             = C.WASM_INS_I64_CLZ
	WASM_INS_I64_CTZ             = C.WASM_INS_I64_CTZ
	WASM_INS_I64_POPCNT          = C.WASM_INS_I64_POPCNT
	WASM_INS_I64_ADD             = C.WASM_INS_I64_ADD
	WASM_INS_I64_SUB             = C.WASM_INS_I64_SUB
	WASM_INS_I64_MUL             = C.WASM_INS_I64_MUL
	WASM_INS_I64_DIV_S           = C.WASM_INS_I64_DIV_S
	WASM_INS_I64_DIV_U           = C.WASM_INS_I64_DIV_U
	WASM_INS_I64_REM_S           = C.WASM_INS_I64_REM_S
	WASM_INS_I64_REM_U           = C.WASM_INS_I64_REM_U
	WASM_INS_I64_AND             = C.WASM_INS_I64_AND
	WASM_INS_I64_OR              = C.WASM_INS_I64_OR
	WASM_INS_I64_XOR             = C.WASM_INS_I64_XOR
	WASM_INS_I64_SHL             = C.WASM_INS_I64_SHL
	WASM_INS_I64_SHR_S           = C.WASM_INS_I64_SHR_S
	WASM_INS_I64_SHR_U           = C.WASM_INS_I64_SHR_U
	WASM_INS_I64_ROTL            = C.WASM_INS_I64_ROTL
	WASM_INS_I64_ROTR            = C.WASM_INS_I64_ROTR
	WASM_INS_F32_ABS             = C.WASM_INS_F32_ABS
	WASM_INS_F32_NEG             = C.WASM_INS_F32_NEG
	WASM_INS_F32_CEIL            = C.WASM_INS_F32_CEIL
	WASM_INS_F32_FLOOR           = C.WASM_INS_F32_FLOOR
	WASM_INS_F32_TRUNC           = C.WASM_INS_F32_TRUNC
	WASM_INS_F32_NEAREST         = C.WASM_INS_F32_NEAREST
	WASM_INS_F32_SQRT            = C.WASM_INS_F32_SQRT
	WASM_INS_F32_ADD             = C.WASM_INS_F32_ADD
	WASM_INS_F32_SUB             = C.WASM_INS_F32_SUB
	WASM_INS_F32_MUL             = C.WASM_INS_F32_MUL
	WASM_INS_F32_DIV             = C.WASM_INS_F32_DIV
	WASM_INS_F32_MIN             = C.WASM_INS_F32_MIN
	WASM_INS_F32_MAX             = C.WASM_INS_F32_MAX
	WASM_INS_F32_COPYSIGN        = C.WASM_INS_F32_COPYSIGN
	WASM_INS_F64_ABS             = C.WASM_INS_F64_ABS
	WASM_INS_F64_NEG             = C.WASM_INS_F64_NEG
	WASM_INS_F64_CEIL            = C.WASM_INS_F64_CEIL
	WASM_INS_F64_FLOOR           = C.WASM_INS_F64_FLOOR
	WASM_INS_F64_TRUNC           = C.WASM_INS_F64_TRUNC
	WASM_INS_F64_NEAREST         = C.WASM_INS_F64_NEAREST
	WASM_INS_F64_SQRT            = C.WASM_INS_F64_SQRT
	WASM_INS_F64_ADD             = C.WASM_INS_F64_ADD
	WASM_INS_F64_SUB             = C.WASM_INS_F64_SUB
	WASM_INS_F64_MUL             = C.WASM_INS_F64_MUL
	WASM_INS_F64_DIV             = C.WASM_INS_F64_DIV
	WASM_INS_F64_MIN             = C.WASM_INS_F64_MIN
	WASM_INS_F64_MAX             = C.WASM_INS_F64_MAX
	WASM_INS_F64_COPYSIGN        = C.WASM_INS_F64_COPYSIGN
	WASM_INS_I32_WARP_I64        = C.WASM_INS_I32_WARP_I64
	WASM_INS_I32_TRUNC_U_F32     = C.WASM_INS_I32_TRUNC_U_F32
	WASM_INS_I32_TRUNC_S_F64     = C.WASM_INS_I32_TRUNC_S_F64
	WASM_INS_I32_TRUNC_U_F64     = C.WASM_INS_I32_TRUNC_U_F64
	WASM_INS_I64_EXTEND_S_I32    = C.WASM_INS_I64_EXTEND_S_I32
	WASM_INS_I64_EXTEND_U_I32    = C.WASM_INS_I64_EXTEND_U_I32
	WASM_INS_I64_TRUNC_S_F32     = C.WASM_INS_I64_TRUNC_S_F32
	WASM_INS_I64_TRUNC_U_F32     = C.WASM_INS_I64_TRUNC_U_F32
	WASM_INS_I64_TRUNC_S_F64     = C.WASM_INS_I64_TRUNC_S_F64
	WASM_INS_I64_TRUNC_U_F64     = C.WASM_INS_I64_TRUNC_U_F64
	WASM_INS_F32_CONVERT_S_I32   = C.WASM_INS_F32_CONVERT_S_I32
	WASM_INS_F32_CONVERT_U_I32   = C.WASM_INS_F32_CONVERT_U_I32
	WASM_INS_F32_CONVERT_S_I64   = C.WASM_INS_F32_CONVERT_S_I64
	WASM_INS_F32_CONVERT_U_I64   = C.WASM_INS_F32_CONVERT_U_I64
	WASM_INS_F32_DEMOTE_F64      = C.WASM_INS_F32_DEMOTE_F64
	WASM_INS_F64_CONVERT_S_I32   = C.WASM_INS_F64_CONVERT_S_I32
	WASM_INS_F64_CONVERT_U_I32   = C.WASM_INS_F64_CONVERT_U_I32
	WASM_INS_F64_CONVERT_S_I64   = C.WASM_INS_F64_CONVERT_S_I64
	WASM_INS_F64_CONVERT_U_I64   = C.WASM_INS_F64_CONVERT_U_I64
	WASM_INS_F64_PROMOTE_F32     = C.WASM_INS_F64_PROMOTE_F32
	WASM_INS_I32_REINTERPRET_F32 = C.WASM_INS_I32_REINTERPRET_F32
	WASM_INS_I64_REINTERPRET_F64 = C.WASM_INS_I64_REINTERPRET_F64
	WASM_INS_F32_REINTERPRET_I32 = C.WASM_INS_F32_REINTERPRET_I32
	WASM_INS_F64_REINTERPRET_I64 = C.WASM_INS_F64_REINTERPRET_I64
	WASM_INS_INVALID             = C.WASM_INS_INVALID
	WASM_INS_ENDING              = C.WASM_INS_ENDING
)

const (
	WASM_GRP_INVALID    = C.WASM_GRP_INVALID
	WASM_GRP_NUMBERIC   = C.WASM_GRP_NUMBERIC
	WASM_GRP_PARAMETRIC = C.WASM_GRP_PARAMETRIC
	WASM_GRP_VARIABLE   = C.WASM_GRP_VARIABLE
	WASM_GRP_MEMORY     = C.WASM_GRP_MEMORY
	WASM_GRP_CONTROL    = C.WASM_GRP_CONTROL
	WASM_GRP_ENDING     = C.WASM_GRP_ENDING
)
