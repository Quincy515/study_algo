package Facade

import "fmt"

// API is facade interface of facade package
type API interface {
	Test() string // 测试
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// apiImpl: facade implement
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (api *apiImpl) Test() string {
	aRet := api.a.TestA()
	bRet := api.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

//AModuleAPI: A 测试
type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct{}

// NewAModuleAPI return new AModuleAPI
func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func (a *aModuleImpl) TestA() string {
	return "A module running"
}

// BModuleAPI: B 测试
type BModuleAPI interface {
	TestB() string
}

type bModuleAPI struct{}

// NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleAPI{}
}

func (b *bModuleAPI) TestB() string {
	return "B module running"
}
