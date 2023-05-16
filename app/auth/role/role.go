package role

type Role int

// 권한의 오름 차순으로 작성
const (
	Member Role = iota
	Admin
)

var (
	roleList = []Role{Member, Admin}
)

func (r Role) GenerateAuthority() map[string]bool {
	authority := make(map[string]bool)
	for _, value := range roleList {
		if r >= value {
			authority[value.String()] = true
		} else {
			authority[value.String()] = false
		}
	}

	return authority
}

func (r Role) Valid(authority map[string]interface{}) bool {
	return authority[r.String()].(bool)
}
