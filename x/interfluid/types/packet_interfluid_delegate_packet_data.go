package types

// ValidateBasic is used for validating the packet
func (p InterfluidDelegatePacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p InterfluidDelegatePacketData) GetBytes() ([]byte, error) {
	var modulePacket InterfluidPacketData

	modulePacket.Packet = &InterfluidPacketData_InterfluidDelegatePacket{&p}

	return modulePacket.Marshal()
}
