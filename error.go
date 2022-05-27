package gos7

// Copyright 2018 Trung Hieu Le. All rights reserved.
// This software may be modified and distributed under the terms
// of the BSD license. See the LICENSE file for details.
import (
	"fmt"
)

const (
	ErrTCPSocketCreation    CommError = 1
	ErrTCPConnectionTimeout CommError = 2
	ErrTCPConnectionFailed  CommError = 3
	ErrTCPReceiveTimeout    CommError = 4
	ErrTCPDataReceive       CommError = -5
	ErrTCPSendTimeout       CommError = 0x00000006
	ErrTCPDataSend          CommError = 0x00000007
	ErrTCPConnectionReset   CommError = 0x00000008
	ErrTCPNotConnected      CommError = 0x00000009
	ErrTCPUnreachableHost   CommError = 0x00002751

	ErrIsoConnect         CommError = 0x00010000 // Connection error
	ErrIsoInvalidPDU      CommError = 0x00030000 // Bad format
	ErrIsoInvalidDataSize CommError = 0x00040000 // Bad Datasize passed to send/recv : buffer is invalid

	ErrCliNegotiatingPDU         CommError = 0x00100000
	ErrCliInvalidParams          CommError = 0x00200000
	ErrCliJobPending             CommError = 0x00300000
	ErrCliTooManyItems           CommError = 0x00400000
	ErrCliInvalidWordLen         CommError = 0x00500000
	ErrCliPartialDataWritten     CommError = 0x00600000
	ErrCliSizeOverPDU            CommError = 0x00700000
	ErrCliInvalidPlcAnswer       CommError = 0x00800000
	ErrCliAddressOutOfRange      CommError = 0x00900000
	ErrCliInvalidTransportSize   CommError = 0x00A00000
	ErrCliWriteDataSizeMismatch  CommError = 0x00B00000
	ErrCliItemNotAvailable       CommError = 0x00C00000
	ErrCliInvalidValue           CommError = 0x00D00000
	ErrCliCannotStartPLC         CommError = 0x00E00000
	ErrCliAlreadyRun             CommError = 0x00F00000
	ErrCliCannotStopPLC          CommError = 0x01000000
	ErrCliCannotCopyRAMToRom     CommError = 0x01100000
	ErrCliCannotCompress         CommError = 0x01200000
	ErrCliAlreadyStop            CommError = 0x01300000
	ErrCliFunNotAvailable        CommError = 0x01400000
	ErrCliUploadSequenceFailed   CommError = 0x01500000
	ErrCliInvalidDataSizeRecvd   CommError = 0x01600000
	ErrCliInvalidBlockType       CommError = 0x01700000
	ErrCliInvalidBlockNumber     CommError = 0x01800000
	ErrCliInvalidBlockSize       CommError = 0x01900000
	ErrCliNeedPassword           CommError = 0x01D00000
	ErrCliInvalidPassword        CommError = 0x01E00000
	ErrCliNoPasswordToSetOrClear CommError = 0x01F00000
	ErrCliJobTimeout             CommError = 0x02000000
	ErrCliPartialDataRead        CommError = 0x02100000
	ErrCliBufferTooSmall         CommError = 0x02200000
	ErrCliFunctionRefused        CommError = 0x02300000
	ErrCliDestroying             CommError = 0x02400000
	ErrCliInvalidParamNumber     CommError = 0x02500000
	ErrCliCannotChangeParam      CommError = 0x02600000
	ErrCliFunctionNotImplemented CommError = 0x02700000

	code7Ok                    = 0
	code7AddressOutOfRange     = 5
	code7InvalidTransportSize  = 6
	code7WriteDataSizeMismatch = 7
	code7ResItemNotAvailable   = 10
	code7ResItemNotAvailable1  = 53769
	code7InvalidValue          = 56321
	code7NeedPassword          = 53825
	code7InvalidPassword       = 54786
	code7NoPasswordToClear     = 54788
	code7NoPasswordToSet       = 54789
	code7FunNotAvailable       = 33028
	code7DataOverPDU           = 34048
)

type CommError int
func (err CommError) Error() string {
	return ErrorText(err)
}

//ErrorText return a string error text from error code integer
func ErrorText(err CommError) string {
	switch err {
	case 0:
		return "OK"
	case ErrTCPSocketCreation:
		return "SYS : Error creating the Socket"
	case ErrTCPConnectionTimeout:
		return "TCP : Connection Timeout"
	case ErrTCPConnectionFailed:
		return "TCP : Connection Error"
	case ErrTCPReceiveTimeout:
		return "TCP : Data receive Timeout"
	case ErrTCPDataReceive:
		return "TCP : Error receiving Data"
	case ErrTCPSendTimeout:
		return "TCP : Data send Timeout"
	case ErrTCPDataSend:
		return "TCP : Error sending Data"
	case ErrTCPConnectionReset:
		return "TCP : Connection reset by the Peer"
	case ErrTCPNotConnected:
		return "CLI : Client not connected"
	case ErrTCPUnreachableHost:
		return "TCP : Unreachable host"
	case ErrIsoConnect:
		return "ISO : Connection Error"
	case ErrIsoInvalidPDU:
		return "ISO : Invalid PDU received"
	case ErrIsoInvalidDataSize:
		return "ISO : Invalid Buffer passed to Send/Receive"
	case ErrCliNegotiatingPDU:
		return "CLI : Error in PDU negotiation"
	case ErrCliInvalidParams:
		return "CLI : invalid param(s) supplied"
	case ErrCliJobPending:
		return "CLI : Job pending"
	case ErrCliTooManyItems:
		return "CLI : too may items (>20) in multi read/write"
	case ErrCliInvalidWordLen:
		return "CLI : invalid WordLength"
	case ErrCliPartialDataWritten:
		return "CLI : Partial data written"
	case ErrCliSizeOverPDU:
		return "CPU : total data exceeds the PDU size"
	case ErrCliInvalidPlcAnswer:
		return "CLI : invalid CPU answer"
	case ErrCliAddressOutOfRange:
		return "CPU : Address out of range"
	case ErrCliInvalidTransportSize:
		return "CPU : Invalid Transport size"
	case ErrCliWriteDataSizeMismatch:
		return "CPU : Data size mismatch"
	case ErrCliItemNotAvailable:
		return "CPU : Item not available"
	case ErrCliInvalidValue:
		return "CPU : Invalid value supplied"
	case ErrCliCannotStartPLC:
		return "CPU : Cannot start PLC"
	case ErrCliAlreadyRun:
		return "CPU : PLC already RUN"
	case ErrCliCannotStopPLC:
		return "CPU : Cannot stop PLC"
	case ErrCliCannotCopyRAMToRom:
		return "CPU : Cannot copy RAM to ROM"
	case ErrCliCannotCompress:
		return "CPU : Cannot compress"
	case ErrCliAlreadyStop:
		return "CPU : PLC already STOP"
	case ErrCliFunNotAvailable:
		return "CPU : Function not available"
	case ErrCliUploadSequenceFailed:
		return "CPU : Upload sequence failed"
	case ErrCliInvalidDataSizeRecvd:
		return "CLI : Invalid data size received"
	case ErrCliInvalidBlockType:
		return "CLI : Invalid block type"
	case ErrCliInvalidBlockNumber:
		return "CLI : Invalid block number"
	case ErrCliInvalidBlockSize:
		return "CLI : Invalid block size"
	case ErrCliNeedPassword:
		return "CPU : Function not authorized for current protection level"
	case ErrCliInvalidPassword:
		return "CPU : Invalid password"
	case ErrCliNoPasswordToSetOrClear:
		return "CPU : No password to set or clear"
	case ErrCliJobTimeout:
		return "CLI : Job Timeout"
	case ErrCliFunctionRefused:
		return "CLI : function refused by CPU (Unknown error)"
	case ErrCliPartialDataRead:
		return "CLI : Partial data read"
	case ErrCliBufferTooSmall:
		return "CLI : The buffer supplied is too small to accomplish the operation"
	case ErrCliDestroying:
		return "CLI : Cannot perform (destroying)"
	case ErrCliInvalidParamNumber:
		return "CLI : Invalid Param Number"
	case ErrCliCannotChangeParam:
		return "CLI : Cannot change this param now"
	case ErrCliFunctionNotImplemented:
		return "CLI : Function not implemented"
	default:
		return fmt.Sprintf("CLI : Unknown error (%v)", err)
	}
}

//CPUError specific CPU error after response
func CPUError(err uint) CommError {
	switch err {
	case 0:
		return 0
	case code7AddressOutOfRange:
		return ErrCliAddressOutOfRange
	case code7InvalidTransportSize:
		return ErrCliInvalidTransportSize
	case code7WriteDataSizeMismatch:
		return ErrCliWriteDataSizeMismatch
	case code7ResItemNotAvailable, code7ResItemNotAvailable1:
		return ErrCliItemNotAvailable
	case code7DataOverPDU:
		return ErrCliSizeOverPDU
	case code7InvalidValue:
		return ErrCliInvalidValue
	case code7FunNotAvailable:
		return ErrCliFunNotAvailable
	case code7NeedPassword:
		return ErrCliNeedPassword
	case code7InvalidPassword:
		return ErrCliInvalidPassword
	case code7NoPasswordToSet, code7NoPasswordToClear:
		return ErrCliNoPasswordToSetOrClear
	default:
		return ErrCliFunctionRefused
	}
	return 0
}
