package gockground

type Gallery struct {

  Id      string  `json:"id"`
  Title   string  `json:"title"`
  Ups     int     `json:"ups"`
  Downs   int     `json:"downs"`
  Images  []Image `json:"images"`
}
