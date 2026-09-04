package domain

type Theatre struct {
	ID string
	Name string
	Location string
	Screens []*Screen
}
	
func NewTheatre(id, name , location string ) *Theatre {
	return &Theatre{
		ID:id, 
		Name:name,
		Location:location,
	}
}

func (t *Theatre) AddScreen(s *Screen){
	t.Screens = append(t.Screens,s)
}

