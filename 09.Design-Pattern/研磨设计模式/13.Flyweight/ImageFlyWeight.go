package Flyweight

import "fmt"

type ImageFlyweight struct {
	data string
}

// NewImageFlyweight: 初始化
func NewImageFlyweight(filename string) *ImageFlyweight {
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

// Data: 获取数据
func (i *ImageFlyweight) Data() string {
	return i.data
}
