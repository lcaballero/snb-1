package anchored

type Map map[string]interface{}

type Anchor interface {
	GetAnchor() Map
	SetAnchor(Map)
}
