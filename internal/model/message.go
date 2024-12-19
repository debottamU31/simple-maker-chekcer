package model

type MessageStatus string

const (
    StatusPending  MessageStatus = "PENDING"
    StatusApproved MessageStatus = "APPROVED"
    StatusRejected MessageStatus = "REJECTED"
    StatusSent     MessageStatus = "SENT"
)