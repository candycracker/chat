package types

const (
	// ModuleName defines the module name
	ModuleName = "chat"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_chat"

	// MsgGroupKey defines the group messages store key
	MsgGroupKey = "Chat/group"

	// MsgSenderKey defines the Sender's messages store key
	MsgSenderKey = "Chat/sender"

	// MsgRecipientKey defines the Recipient's messages store key
	MsgRecipientKey = "Chat/recipient"

	// MsgRecipientKey defines the Recipient's messages store key
	MsgCountKey = "Chat/count"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
