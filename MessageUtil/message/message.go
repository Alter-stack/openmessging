package message

type Message struct {
	a int64 //随机整数,可能重复
	t int64 //输入时间戳,基本是升序的,可能重复

	body []byte //消息体, 随机byte数组, 大小固定34字节

	hashCode int
}

