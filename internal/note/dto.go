package note

type Dto struct {
	Text   string `form:"text"`
	Topic  string	`form:"topic"`
}


func (dto Dto)Parse() *Note {
	return &Note{
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