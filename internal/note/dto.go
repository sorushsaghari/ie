package note

type Dto struct {
	Text   string `form:"text"`
	Topic  string	`form:"topic"`
}
type ReadDto struct {
	Dto
	ID uint`form:"id"`
}
type CreateDto struct {
	Dto
	UserID uint
}
func (dto Dto)Parse() *Note {
	return &Note{
		Text: dto.Text,
		Topic: dto.Topic,
	}
}

func (dto CreateDto)Parse() *Note {
	return &Note{
		UserID: dto.UserID,
		Text: dto.Text,
		Topic: dto.Topic,
	}
}

func NewDto(note Note) *Dto{
	return &Dto{
		Text: note.Text,
		Topic: note.Topic,
	}
}



func NewReadDto(note Note) *ReadDto{
	return &ReadDto{
		Dto: *NewDto(note),
		ID:  note.ID,
	}
}