package universal

type HttpMethod int

const (
	GET HttpMethod = iota
	HEAD
	POST
	PUT
	DELETE
)

type ContentType int

const (
	Default ContentType = iota
	FormUrlencoded
	ApplicationJSON
)
