package coroutine

import "time"

var resumeWaitDuration time.Duration = 4 * 1e9

func SetResumeWaitSecond(d time.Duration) { resumeWaitDuration = d }

func NewCoroutine() *Coroutine {
	return &Coroutine{
		mYield:  make(chan struct{}),
		mResume: make(chan struct{}),
		mDone:   make(chan struct{}),

		mRunDone:            make(chan error),
		mResumeWaitDuration: resumeWaitDuration,
	}
}

type Coroutine struct {
	mYield              chan struct{}
	mResume             chan struct{}
	mDone               chan struct{}
	mRunDone            chan error
	mResumeWaitDuration time.Duration
}

func (co *Coroutine) SetResumeWaitDuration(d time.Duration) {
	co.mResumeWaitDuration = d
}

func (co *Coroutine) Run(exe func() error) error {
	go func() {
		err := exe()
		if co.mRunDone != nil {
			co.mRunDone <- err
		}
	}()
	var err error
	select {
	case err = <-co.mRunDone:
	case <-co.mYield:
		co.mYield = nil
	}
	return err
}

func (co *Coroutine) Yield() {
	if co.mYield != nil {
		co.mYield <- struct{}{}
	}
	select {
	case <-co.mResume:
	}
}

func (co *Coroutine) Resume() {
	co.mResume <- struct{}{}
	select {
	case <-co.mDone:
	case <-time.Tick(co.mResumeWaitDuration):
	}
}

func (co *Coroutine) Done() {
	co.mDone <- struct{}{}
}
