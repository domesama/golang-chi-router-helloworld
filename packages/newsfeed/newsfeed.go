package newsfeed

//Interfaces for loosely coupled interactions between the given methods
//Thus we can have other way to interacts with the DTO using the GetAll() functions that returns item
type Getter interface {
	GetAll() []Item
}

type Adder interface {
	Add(item Item)
}

//Provides what ever suits your data
type Item struct{
	Title string `json:"title"`
	Post string `json:post`
}

//Constructs a DTO for that given Data interface
type Repo struct{
	Items []Item
}

//Creates a referenced custom call to that DTO object
func New() *Repo{
	return &Repo{
		Items: []Item{},
	}
}


//Add a referenced way to interacts with our DTO object
func (r *Repo) Add(item Item){
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll()  []Item{
	return r.Items
}


