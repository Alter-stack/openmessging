package messageStore

import (
	m "openmessging/MessageUtil/message"
)


type MessageStore interface {
	/**
	* 根据a和t的条件,返回符合条件的消息的集合. t是输入时间戳模拟值，和实际时间戳没有关系, 线程内升序
	* 这个接口需要是线程安全的，也即评测程序会并发调用该接口
	* 返回的List需要按照t升序排列 (a不要求排序). 如果没有符合的消息, 返回大小为0的List. 如果List里有null元素, 会当结果失败处理
	* 单条线程最大返回消息数量不会超过8万
	* @param aMin 代表a的最小值(包含此值)
	* @param aMax 代表a的最大值(包含此值)
	* @param tMin 代表t的最小值(包含此值)
	* @param tMax 代表t的最大值(包含此值)
	*/
	Put(message m.Message)

	/**
	* 根据a和t的条件,返回符合条件消息的a值的求平均结果. t是输入时间戳模拟值，和实际时间戳没有关系, 线程内升序
	* 这个接口需要是线程安全的，也即评测程序会并发调用该接口
	* 结果忽略小数位,返回整数位即可. 如果没有符合的消息, 返回0
	* 单次查询求和最大值不会超过Long.MAX_VALUE
	* @param aMin 代表a的最小值(包含此值)
	* @param aMax 代表a的最大值(包含此值)
	* @param tMin 代表t的最小值(包含此值)
	* @param tMax 代表t的最大值(包含此值)
	*/
	GetMessage(aMin, aMax, tMin, tMax int64) []m.Message

	/**
	* 根据a和t的条件,返回符合条件消息的a值的求平均结果. t是输入时间戳模拟值，和实际时间戳没有关系, 线程内升序
	* 这个接口需要是线程安全的，也即评测程序会并发调用该接口
	* 结果忽略小数位,返回整数位即可. 如果没有符合的消息, 返回0
	* 单次查询求和最大值不会超过Long.MAX_VALUE
	* @param aMin 代表a的最小值(包含此值)
	* @param aMax 代表a的最大值(包含此值)
	* @param tMin 代表t的最小值(包含此值)
	* @param tMax 代表t的最大值(包含此值)
	*/
	GetAvgValue(aMin, aMax, tMin, tMax int64) int64
}

