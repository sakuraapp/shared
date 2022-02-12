package resource

type UserFormatter func(user *User) *User

type Builder struct {
	userFormatter UserFormatter
}

func (b *Builder) SetUserFormatter(formatter UserFormatter) {
	b.userFormatter = formatter
}

func NewBuilder() *Builder {
	return &Builder{}
}