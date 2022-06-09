package model

type BlogStatus string

const (
	BlogActive   BlogStatus = "active"
	BlogInactive BlogStatus = "inactive"
)

func (m BlogStatus) String() string {
	return string(m)
}

func (m BlogStatus) IsActive() bool {
	return m == BlogActive
}

func (m BlogStatus) IsValid() bool {
	return m == BlogActive || m == BlogInactive
}
