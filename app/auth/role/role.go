package role

type Role byte

// 권한의 오름 차순으로 작성
const (
	Member Role = iota
	Admin
)

var (
	roleList = []Role{Member, Admin}
)

func (r Role) GenerateAuthority() map[Role]bool {
	authority := make(map[Role]bool)
	for _, value := range roleList {
		if r >= value {
			authority[value] = true
		}
	}

	return authority
}

func (r Role) Valid(authority map[Role]bool) bool {
	return authority[r]
}
