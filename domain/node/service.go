package node

type Repository interface {
	SetPort(port string)
	GetPort() string
	SetBlock(block Block)
	GetBlockHeight() uint64
	GetNodeInfo() *Node
}

type Service interface {
	SetPort(port string)
	GetPort() string
	SetBlock(block Block)
	GetBlockHeight() uint64
	GetNodeInfo() *Node
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetBlockHeight() uint64 {
	return s.repo.GetBlockHeight()
}

func (s *service) SetPort(port string) {
	s.repo.SetPort(port)
}

func (s *service) GetPort() string {
	return s.repo.GetPort()
}

func (s *service) SetBlock(block Block) {
	s.repo.SetBlock(block)
}

func (s *service) GetNodeInfo() *Node {
	return s.repo.GetNodeInfo()
}
