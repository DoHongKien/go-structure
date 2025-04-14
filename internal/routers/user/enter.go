package user

type UserRouterGroup struct {
	AuthRouter
	CustomerRouter
	OrderRouter
	OrderDetailRouter
}
