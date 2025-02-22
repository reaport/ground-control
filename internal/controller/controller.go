package controller

type MapService interface{}

type Controller struct {
	mapService MapService
}

func New(mapService MapService) *Controller {
	return &Controller{
		mapService: mapService,
	}
}
