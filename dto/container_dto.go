package dto

type ContainerDto struct {
	ID     uint
	PID    uint
	CID    string
	Image  string
	Status uint
}

type ContainerProtocolDto struct {
	ID            uint
	CID           uint
	Protocol      string
	HostIP        string
	HostPort      string
	ContainerPort string
}

type ContainerAuthDto struct {
	ID       uint
	CID      uint
	Username string
	Password string
}
