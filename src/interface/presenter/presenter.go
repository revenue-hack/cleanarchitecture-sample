package presenter

type Presenter interface {
	WriteJson(v interface{}) error
}
