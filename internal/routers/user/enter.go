package user

type UserRouterGroup struct {
	CustomerRouter
	OrderRouter
	OrderDetailRouter
}
