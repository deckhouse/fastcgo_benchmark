package fastcgo

type stack struct {
	lo, hi uintptr
}

type gobuf struct {
	sp, pc, g, ctxt, ret, lr, bp uintptr
}

type m struct {
	g0 *g
}

type g struct {
	stack                    stack
	stackguard0, stackguard1 uintptr

	_panic, _defer uintptr
	m              *m
	sched          gobuf
}
