// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package waltuhimpl

import (
	"context"
	"example/api/waltuh"
	"sync"
)

// AnotherRepositoryMock is a mock implementation of waltuh.AnotherRepository.
//
//	func TestSomethingThatUsesAnotherRepository(t *testing.T) {
//
//		// make and configure a mocked waltuh.AnotherRepository
//		mockedAnotherRepository := &AnotherRepositoryMock{
//			AFunc: func() (string, error) {
//				panic("mock out the A method")
//			},
//			BFunc: func(ctx context.Context) error {
//				panic("mock out the B method")
//			},
//			CFunc: func() (string, error) {
//				panic("mock out the C method")
//			},
//		}
//
//		// use mockedAnotherRepository in code that requires waltuh.AnotherRepository
//		// and then make assertions.
//
//	}
type AnotherRepositoryMock struct {
	// AFunc mocks the A method.
	AFunc func() (string, error)

	// BFunc mocks the B method.
	BFunc func(ctx context.Context) error

	// CFunc mocks the C method.
	CFunc func() (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// A holds details about calls to the A method.
		A []struct {
		}
		// B holds details about calls to the B method.
		B []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// C holds details about calls to the C method.
		C []struct {
		}
	}
	lockA sync.RWMutex
	lockB sync.RWMutex
	lockC sync.RWMutex
}

// A calls AFunc.
func (mock *AnotherRepositoryMock) A() (string, error) {
	if mock.AFunc == nil {
		panic("AnotherRepositoryMock.AFunc: method is nil but AnotherRepository.A was just called")
	}
	callInfo := struct {
	}{}
	mock.lockA.Lock()
	mock.calls.A = append(mock.calls.A, callInfo)
	mock.lockA.Unlock()
	return mock.AFunc()
}

// ACalls gets all the calls that were made to A.
// Check the length with:
//
//	len(mockedAnotherRepository.ACalls())
func (mock *AnotherRepositoryMock) ACalls() []struct {
} {
	var calls []struct {
	}
	mock.lockA.RLock()
	calls = mock.calls.A
	mock.lockA.RUnlock()
	return calls
}

// B calls BFunc.
func (mock *AnotherRepositoryMock) B(ctx context.Context) error {
	if mock.BFunc == nil {
		panic("AnotherRepositoryMock.BFunc: method is nil but AnotherRepository.B was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockB.Lock()
	mock.calls.B = append(mock.calls.B, callInfo)
	mock.lockB.Unlock()
	return mock.BFunc(ctx)
}

// BCalls gets all the calls that were made to B.
// Check the length with:
//
//	len(mockedAnotherRepository.BCalls())
func (mock *AnotherRepositoryMock) BCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockB.RLock()
	calls = mock.calls.B
	mock.lockB.RUnlock()
	return calls
}

// C calls CFunc.
func (mock *AnotherRepositoryMock) C() (string, error) {
	if mock.CFunc == nil {
		panic("AnotherRepositoryMock.CFunc: method is nil but AnotherRepository.C was just called")
	}
	callInfo := struct {
	}{}
	mock.lockC.Lock()
	mock.calls.C = append(mock.calls.C, callInfo)
	mock.lockC.Unlock()
	return mock.CFunc()
}

// CCalls gets all the calls that were made to C.
// Check the length with:
//
//	len(mockedAnotherRepository.CCalls())
func (mock *AnotherRepositoryMock) CCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockC.RLock()
	calls = mock.calls.C
	mock.lockC.RUnlock()
	return calls
}

// BRepositoryMock is a mock implementation of waltuh.BRepository.
//
//	func TestSomethingThatUsesBRepository(t *testing.T) {
//
//		// make and configure a mocked waltuh.BRepository
//		mockedBRepository := &BRepositoryMock{
//			YepFunc: func(ctx context.Context, id string) (string, error) {
//				panic("mock out the Yep method")
//			},
//			YopeFunc: func() (string, error) {
//				panic("mock out the Yope method")
//			},
//		}
//
//		// use mockedBRepository in code that requires waltuh.BRepository
//		// and then make assertions.
//
//	}
type BRepositoryMock struct {
	// YepFunc mocks the Yep method.
	YepFunc func(ctx context.Context, id string) (string, error)

	// YopeFunc mocks the Yope method.
	YopeFunc func() (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Yep holds details about calls to the Yep method.
		Yep []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// Yope holds details about calls to the Yope method.
		Yope []struct {
		}
	}
	lockYep  sync.RWMutex
	lockYope sync.RWMutex
}

// Yep calls YepFunc.
func (mock *BRepositoryMock) Yep(ctx context.Context, id string) (string, error) {
	if mock.YepFunc == nil {
		panic("BRepositoryMock.YepFunc: method is nil but BRepository.Yep was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockYep.Lock()
	mock.calls.Yep = append(mock.calls.Yep, callInfo)
	mock.lockYep.Unlock()
	return mock.YepFunc(ctx, id)
}

// YepCalls gets all the calls that were made to Yep.
// Check the length with:
//
//	len(mockedBRepository.YepCalls())
func (mock *BRepositoryMock) YepCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockYep.RLock()
	calls = mock.calls.Yep
	mock.lockYep.RUnlock()
	return calls
}

// Yope calls YopeFunc.
func (mock *BRepositoryMock) Yope() (string, error) {
	if mock.YopeFunc == nil {
		panic("BRepositoryMock.YopeFunc: method is nil but BRepository.Yope was just called")
	}
	callInfo := struct {
	}{}
	mock.lockYope.Lock()
	mock.calls.Yope = append(mock.calls.Yope, callInfo)
	mock.lockYope.Unlock()
	return mock.YopeFunc()
}

// YopeCalls gets all the calls that were made to Yope.
// Check the length with:
//
//	len(mockedBRepository.YopeCalls())
func (mock *BRepositoryMock) YopeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockYope.RLock()
	calls = mock.calls.Yope
	mock.lockYope.RUnlock()
	return calls
}

// RepositoryMock is a mock implementation of waltuh.Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked waltuh.Repository
//		mockedRepository := &RepositoryMock{
//			DropWaltJrOffAtSchoolFunc: func(ctx context.Context) (bool, error) {
//				panic("mock out the DropWaltJrOffAtSchool method")
//			},
//			GetFunc: func(ctx context.Context, id string) (string, error) {
//				panic("mock out the Get method")
//			},
//			KillKrazy8Func: func(ctx context.Context, missingPlateShards int) (string, error) {
//				panic("mock out the KillKrazy8 method")
//			},
//			MakeBreakfastFunc: func(birthday int, kilograms int) waltuh.Waltuh {
//				panic("mock out the MakeBreakfast method")
//			},
//			MakeMoneyFunc: func(ctx context.Context, poundsOfMeth int) (int, error) {
//				panic("mock out the MakeMoney method")
//			},
//			NopeFunc: func()  {
//				panic("mock out the Nope method")
//			},
//			SynthesizeMethFunc: func(ctx context.Context, flyPresent bool, withJesse bool) int {
//				panic("mock out the SynthesizeMeth method")
//			},
//		}
//
//		// use mockedRepository in code that requires waltuh.Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// DropWaltJrOffAtSchoolFunc mocks the DropWaltJrOffAtSchool method.
	DropWaltJrOffAtSchoolFunc func(ctx context.Context) (bool, error)

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, id string) (string, error)

	// KillKrazy8Func mocks the KillKrazy8 method.
	KillKrazy8Func func(ctx context.Context, missingPlateShards int) (string, error)

	// MakeBreakfastFunc mocks the MakeBreakfast method.
	MakeBreakfastFunc func(birthday int, kilograms int) waltuh.Waltuh

	// MakeMoneyFunc mocks the MakeMoney method.
	MakeMoneyFunc func(ctx context.Context, poundsOfMeth int) (int, error)

	// NopeFunc mocks the Nope method.
	NopeFunc func()

	// SynthesizeMethFunc mocks the SynthesizeMeth method.
	SynthesizeMethFunc func(ctx context.Context, flyPresent bool, withJesse bool) int

	// calls tracks calls to the methods.
	calls struct {
		// DropWaltJrOffAtSchool holds details about calls to the DropWaltJrOffAtSchool method.
		DropWaltJrOffAtSchool []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID string
		}
		// KillKrazy8 holds details about calls to the KillKrazy8 method.
		KillKrazy8 []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// MissingPlateShards is the missingPlateShards argument value.
			MissingPlateShards int
		}
		// MakeBreakfast holds details about calls to the MakeBreakfast method.
		MakeBreakfast []struct {
			// Birthday is the birthday argument value.
			Birthday int
			// Kilograms is the kilograms argument value.
			Kilograms int
		}
		// MakeMoney holds details about calls to the MakeMoney method.
		MakeMoney []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// PoundsOfMeth is the poundsOfMeth argument value.
			PoundsOfMeth int
		}
		// Nope holds details about calls to the Nope method.
		Nope []struct {
		}
		// SynthesizeMeth holds details about calls to the SynthesizeMeth method.
		SynthesizeMeth []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// FlyPresent is the flyPresent argument value.
			FlyPresent bool
			// WithJesse is the withJesse argument value.
			WithJesse bool
		}
	}
	lockDropWaltJrOffAtSchool sync.RWMutex
	lockGet                   sync.RWMutex
	lockKillKrazy8            sync.RWMutex
	lockMakeBreakfast         sync.RWMutex
	lockMakeMoney             sync.RWMutex
	lockNope                  sync.RWMutex
	lockSynthesizeMeth        sync.RWMutex
}

// DropWaltJrOffAtSchool calls DropWaltJrOffAtSchoolFunc.
func (mock *RepositoryMock) DropWaltJrOffAtSchool(ctx context.Context) (bool, error) {
	if mock.DropWaltJrOffAtSchoolFunc == nil {
		panic("RepositoryMock.DropWaltJrOffAtSchoolFunc: method is nil but Repository.DropWaltJrOffAtSchool was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockDropWaltJrOffAtSchool.Lock()
	mock.calls.DropWaltJrOffAtSchool = append(mock.calls.DropWaltJrOffAtSchool, callInfo)
	mock.lockDropWaltJrOffAtSchool.Unlock()
	return mock.DropWaltJrOffAtSchoolFunc(ctx)
}

// DropWaltJrOffAtSchoolCalls gets all the calls that were made to DropWaltJrOffAtSchool.
// Check the length with:
//
//	len(mockedRepository.DropWaltJrOffAtSchoolCalls())
func (mock *RepositoryMock) DropWaltJrOffAtSchoolCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockDropWaltJrOffAtSchool.RLock()
	calls = mock.calls.DropWaltJrOffAtSchool
	mock.lockDropWaltJrOffAtSchool.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *RepositoryMock) Get(ctx context.Context, id string) (string, error) {
	if mock.GetFunc == nil {
		panic("RepositoryMock.GetFunc: method is nil but Repository.Get was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  string
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedRepository.GetCalls())
func (mock *RepositoryMock) GetCalls() []struct {
	Ctx context.Context
	ID  string
} {
	var calls []struct {
		Ctx context.Context
		ID  string
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// KillKrazy8 calls KillKrazy8Func.
func (mock *RepositoryMock) KillKrazy8(ctx context.Context, missingPlateShards int) (string, error) {
	if mock.KillKrazy8Func == nil {
		panic("RepositoryMock.KillKrazy8Func: method is nil but Repository.KillKrazy8 was just called")
	}
	callInfo := struct {
		Ctx                context.Context
		MissingPlateShards int
	}{
		Ctx:                ctx,
		MissingPlateShards: missingPlateShards,
	}
	mock.lockKillKrazy8.Lock()
	mock.calls.KillKrazy8 = append(mock.calls.KillKrazy8, callInfo)
	mock.lockKillKrazy8.Unlock()
	return mock.KillKrazy8Func(ctx, missingPlateShards)
}

// KillKrazy8Calls gets all the calls that were made to KillKrazy8.
// Check the length with:
//
//	len(mockedRepository.KillKrazy8Calls())
func (mock *RepositoryMock) KillKrazy8Calls() []struct {
	Ctx                context.Context
	MissingPlateShards int
} {
	var calls []struct {
		Ctx                context.Context
		MissingPlateShards int
	}
	mock.lockKillKrazy8.RLock()
	calls = mock.calls.KillKrazy8
	mock.lockKillKrazy8.RUnlock()
	return calls
}

// MakeBreakfast calls MakeBreakfastFunc.
func (mock *RepositoryMock) MakeBreakfast(birthday int, kilograms int) waltuh.Waltuh {
	if mock.MakeBreakfastFunc == nil {
		panic("RepositoryMock.MakeBreakfastFunc: method is nil but Repository.MakeBreakfast was just called")
	}
	callInfo := struct {
		Birthday  int
		Kilograms int
	}{
		Birthday:  birthday,
		Kilograms: kilograms,
	}
	mock.lockMakeBreakfast.Lock()
	mock.calls.MakeBreakfast = append(mock.calls.MakeBreakfast, callInfo)
	mock.lockMakeBreakfast.Unlock()
	return mock.MakeBreakfastFunc(birthday, kilograms)
}

// MakeBreakfastCalls gets all the calls that were made to MakeBreakfast.
// Check the length with:
//
//	len(mockedRepository.MakeBreakfastCalls())
func (mock *RepositoryMock) MakeBreakfastCalls() []struct {
	Birthday  int
	Kilograms int
} {
	var calls []struct {
		Birthday  int
		Kilograms int
	}
	mock.lockMakeBreakfast.RLock()
	calls = mock.calls.MakeBreakfast
	mock.lockMakeBreakfast.RUnlock()
	return calls
}

// MakeMoney calls MakeMoneyFunc.
func (mock *RepositoryMock) MakeMoney(ctx context.Context, poundsOfMeth int) (int, error) {
	if mock.MakeMoneyFunc == nil {
		panic("RepositoryMock.MakeMoneyFunc: method is nil but Repository.MakeMoney was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		PoundsOfMeth int
	}{
		Ctx:          ctx,
		PoundsOfMeth: poundsOfMeth,
	}
	mock.lockMakeMoney.Lock()
	mock.calls.MakeMoney = append(mock.calls.MakeMoney, callInfo)
	mock.lockMakeMoney.Unlock()
	return mock.MakeMoneyFunc(ctx, poundsOfMeth)
}

// MakeMoneyCalls gets all the calls that were made to MakeMoney.
// Check the length with:
//
//	len(mockedRepository.MakeMoneyCalls())
func (mock *RepositoryMock) MakeMoneyCalls() []struct {
	Ctx          context.Context
	PoundsOfMeth int
} {
	var calls []struct {
		Ctx          context.Context
		PoundsOfMeth int
	}
	mock.lockMakeMoney.RLock()
	calls = mock.calls.MakeMoney
	mock.lockMakeMoney.RUnlock()
	return calls
}

// Nope calls NopeFunc.
func (mock *RepositoryMock) Nope() {
	if mock.NopeFunc == nil {
		panic("RepositoryMock.NopeFunc: method is nil but Repository.Nope was just called")
	}
	callInfo := struct {
	}{}
	mock.lockNope.Lock()
	mock.calls.Nope = append(mock.calls.Nope, callInfo)
	mock.lockNope.Unlock()
	mock.NopeFunc()
}

// NopeCalls gets all the calls that were made to Nope.
// Check the length with:
//
//	len(mockedRepository.NopeCalls())
func (mock *RepositoryMock) NopeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockNope.RLock()
	calls = mock.calls.Nope
	mock.lockNope.RUnlock()
	return calls
}

// SynthesizeMeth calls SynthesizeMethFunc.
func (mock *RepositoryMock) SynthesizeMeth(ctx context.Context, flyPresent bool, withJesse bool) int {
	if mock.SynthesizeMethFunc == nil {
		panic("RepositoryMock.SynthesizeMethFunc: method is nil but Repository.SynthesizeMeth was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		FlyPresent bool
		WithJesse  bool
	}{
		Ctx:        ctx,
		FlyPresent: flyPresent,
		WithJesse:  withJesse,
	}
	mock.lockSynthesizeMeth.Lock()
	mock.calls.SynthesizeMeth = append(mock.calls.SynthesizeMeth, callInfo)
	mock.lockSynthesizeMeth.Unlock()
	return mock.SynthesizeMethFunc(ctx, flyPresent, withJesse)
}

// SynthesizeMethCalls gets all the calls that were made to SynthesizeMeth.
// Check the length with:
//
//	len(mockedRepository.SynthesizeMethCalls())
func (mock *RepositoryMock) SynthesizeMethCalls() []struct {
	Ctx        context.Context
	FlyPresent bool
	WithJesse  bool
} {
	var calls []struct {
		Ctx        context.Context
		FlyPresent bool
		WithJesse  bool
	}
	mock.lockSynthesizeMeth.RLock()
	calls = mock.calls.SynthesizeMeth
	mock.lockSynthesizeMeth.RUnlock()
	return calls
}
