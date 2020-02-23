package Flyweight

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

var imageFactory *ImageFlyweightFactory

// GetImageFlyweightFactory: 对象仅初始化依次
func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}

func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename] // 取出，map去重
	if image == nil {
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}
	return image
}
