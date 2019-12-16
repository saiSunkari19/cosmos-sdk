package channel

// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/04-channel/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"
)

const (
	NONE                           = types.NONE
	UNORDERED                      = types.UNORDERED
	ORDERED                        = types.ORDERED
	OrderNone                      = types.OrderNone
	OrderUnordered                 = types.OrderUnordered
	OrderOrdered                   = types.OrderOrdered
	CLOSED                         = types.CLOSED
	INIT                           = types.INIT
	OPENTRY                        = types.OPENTRY
	OPEN                           = types.OPEN
	StateClosed                    = types.StateClosed
	StateInit                      = types.StateInit
	StateOpenTry                   = types.StateOpenTry
	StateOpen                      = types.StateOpen
	DefaultCodespace               = types.DefaultCodespace
	CodeChannelExists              = types.CodeChannelExists
	CodeChannelNotFound            = types.CodeChannelNotFound
	CodeInvalidCounterpartyChannel = types.CodeInvalidCounterpartyChannel
	CodeChannelCapabilityNotFound  = types.CodeChannelCapabilityNotFound
	CodeInvalidPacket              = types.CodeInvalidPacket
	CodeSequenceNotFound           = types.CodeSequenceNotFound
	CodePacketTimeout              = types.CodePacketTimeout
	CodeInvalidChannel             = types.CodeInvalidChannel
	CodeInvalidChannelState        = types.CodeInvalidChannelState
	AttributeKeySenderPort         = types.AttributeKeySenderPort
	AttributeKeyReceiverPort       = types.AttributeKeyReceiverPort
	AttributeKeyChannelID          = types.AttributeKeyChannelID
	AttributeKeySequence           = types.AttributeKeySequence
	SubModuleName                  = types.SubModuleName
	StoreKey                       = types.StoreKey
	RouterKey                      = types.RouterKey
	QuerierRoute                   = types.QuerierRoute
	QueryAllChannels               = types.QueryAllChannels
	QueryChannel                   = types.QueryChannel
)

var (
	// functions aliases
	NewKeeper                     = keeper.NewKeeper
	QuerierChannels               = keeper.QuerierChannels
	QuerierChannel                = keeper.QuerierChannel
	NewChannel                    = types.NewChannel
	NewCounterparty               = types.NewCounterparty
	OrderFromString               = types.OrderFromString
	StateFromString               = types.StateFromString
	RegisterCodec                 = types.RegisterCodec
	ErrChannelExists              = types.ErrChannelExists
	ErrChannelNotFound            = types.ErrChannelNotFound
	ErrInvalidCounterpartyChannel = types.ErrInvalidCounterpartyChannel
	ErrChannelCapabilityNotFound  = types.ErrChannelCapabilityNotFound
	ErrInvalidPacket              = types.ErrInvalidPacket
	ErrSequenceNotFound           = types.ErrSequenceNotFound
	ErrPacketTimeout              = types.ErrPacketTimeout
	ErrInvalidChannel             = types.ErrInvalidChannel
	ErrInvalidChannelState        = types.ErrInvalidChannelState
	ChannelPath                   = types.ChannelPath
	ChannelCapabilityPath         = types.ChannelCapabilityPath
	NextSequenceSendPath          = types.NextSequenceSendPath
	NextSequenceRecvPath          = types.NextSequenceRecvPath
	PacketCommitmentPath          = types.PacketCommitmentPath
	PacketAcknowledgementPath     = types.PacketAcknowledgementPath
	KeyChannel                    = types.KeyChannel
	KeyChannelCapabilityPath      = types.KeyChannelCapabilityPath
	KeyNextSequenceSend           = types.KeyNextSequenceSend
	KeyNextSequenceRecv           = types.KeyNextSequenceRecv
	KeyPacketCommitment           = types.KeyPacketCommitment
	KeyPacketAcknowledgement      = types.KeyPacketAcknowledgement
	NewMsgChannelOpenInit         = types.NewMsgChannelOpenInit
	NewMsgChannelOpenTry          = types.NewMsgChannelOpenTry
	NewMsgChannelOpenAck          = types.NewMsgChannelOpenAck
	NewMsgChannelOpenConfirm      = types.NewMsgChannelOpenConfirm
	NewMsgChannelCloseInit        = types.NewMsgChannelCloseInit
	NewMsgChannelCloseConfirm     = types.NewMsgChannelCloseConfirm
	NewPacket                     = types.NewPacket
	NewOpaquePacket               = types.NewOpaquePacket
	NewChannelResponse            = types.NewChannelResponse
	NewQueryChannelParams         = types.NewQueryChannelParams

	// variable aliases
	SubModuleCdc                 = types.SubModuleCdc
	EventTypeChannelOpenInit     = types.EventTypeChannelOpenInit
	EventTypeChannelOpenTry      = types.EventTypeChannelOpenTry
	EventTypeChannelOpenAck      = types.EventTypeChannelOpenAck
	EventTypeChannelOpenConfirm  = types.EventTypeChannelOpenConfirm
	EventTypeChannelCloseInit    = types.EventTypeChannelCloseInit
	EventTypeChannelCloseConfirm = types.EventTypeChannelCloseConfirm
	AttributeValueCategory       = types.AttributeValueCategory
)

type (
	Keeper                 = keeper.Keeper
	Channel                = types.Channel
	Counterparty           = types.Counterparty
	Order                  = types.Order
	State                  = types.State
	ClientKeeper           = types.ClientKeeper
	ConnectionKeeper       = types.ConnectionKeeper
	PortKeeper             = types.PortKeeper
	MsgChannelOpenInit     = types.MsgChannelOpenInit
	MsgChannelOpenTry      = types.MsgChannelOpenTry
	MsgChannelOpenAck      = types.MsgChannelOpenAck
	MsgChannelOpenConfirm  = types.MsgChannelOpenConfirm
	MsgChannelCloseInit    = types.MsgChannelCloseInit
	MsgChannelCloseConfirm = types.MsgChannelCloseConfirm
	Packet                 = types.Packet
	OpaquePacket           = types.OpaquePacket
	ChannelResponse        = types.ChannelResponse
	QueryChannelParams     = types.QueryChannelParams
)
