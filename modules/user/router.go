package user

import "github.com/dibimbing-satkom-indo/onion-architecture-go/dto"

type Router struct {
	rq RequestHandlerInterface
}


func (r Router) Route(request dto.Request) {

}
