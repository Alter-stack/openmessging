package message

type Message struct {
	A int64 //随机整数,可能重复
	T int64 //输入时间戳,基本是升序的,可能重复

	Body []byte //消息体, 随机byte数组, 大小固定34字节

	HashCode int
}
