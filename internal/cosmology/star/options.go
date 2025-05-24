package star

type StarOption func(*Star)

func WithPosition(p SystemPosition) StarOption {
	return func(s *Star) {
		s.Position = p
	}
}
