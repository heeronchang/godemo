package functional_options

type User struct {
	Name    string
	Age     int64
	Hobbies []string
}

type UserFuncOption func(user *User)

func NewUser(userOptions ...UserFuncOption) *User {
	u := &User{}
	for _, option := range userOptions {
		option(u)
	}

	return u
}

func WithName(name string) UserFuncOption {
	return func(u *User) {
		u.Name = name
	}
}

func WithAge(age int64) UserFuncOption {
	return func(u *User) {
		u.Age = age
	}
}

func WithHobbies(hobbies []string) UserFuncOption {
	return func(u *User) {
		if u.Hobbies == nil {
			u.Hobbies = make([]string, 0)
		}
		u.Hobbies = append(u.Hobbies, hobbies...)
	}
}

type Configurable interface {
	Apply(u *User)
}

func NewUser2(options ...Configurable) *User {
	u := &User{}
	for _, option := range options {
		option.Apply(u)
	}

	return u
}

type NameAgeOption struct {
	Name string
	Age  int64
}

func (n *NameAgeOption) Apply(u *User) {
	u.Name = n.Name
	u.Age = n.Age
}

type HobbiesOption struct {
	Hobbies []string
}

func (h *HobbiesOption) Apply(u *User) {
	if u.Hobbies == nil {
		u.Hobbies = make([]string, 0)
	}
	u.Hobbies = append(u.Hobbies, h.Hobbies...)
}
