package messageStore

import "openmessging/MessageUtil/message"

type DefaultMessageStore struct {
	NavigableMap map[int64][]message.Message
}

func (d *DefaultMessageStore) Put(message message.Message) {
	panic("implement me")
}

func (d DefaultMessageStore) GetMessage(aMin, aMax, tMin, tMax int64) []message.Message {
	panic("implement me")
}

func (d DefaultMessageStore) GetAvgValue(aMin, aMax, tMin, tMax int64) int64 {
	panic("implement me")
}


///**
// * 这是一个简单的基于内存的实现，以方便选手理解题意；
// * 实际提交时，请维持包名和类名不变，把方法实现修改为自己的内容；
// */
//public class DefaultMessageStoreImpl extends MessageStore {
//
//    private NavigableMap<Long, List<Message>> msgMap = new TreeMap<Long, List<Message>>();
//
//    @Override
//    public synchronized void put(Message message) {
//        if (!msgMap.containsKey(message.getT())) {
//            msgMap.put(message.getT(), new ArrayList<Message>());
//        }
//
//        msgMap.get(message.getT()).add(message);
//    }
//
//
//    @Override
//    public synchronized List<Message> getMessage(long aMin, long aMax, long tMin, long tMax) {
//        ArrayList<Message> res = new ArrayList<Message>();
//        NavigableMap<Long, List<Message>> subMap = msgMap.subMap(tMin, true, tMax, true);
//        for (Map.Entry<Long, List<Message>> mapEntry : subMap.entrySet()) {
//            List<Message> msgQueue = mapEntry.getValue();
//            for (Message msg : msgQueue) {
//                if (msg.getA() >= aMin && msg.getA() <= aMax) {
//                    res.add(msg);
//                }
//            }
//        }
//
//        return res;
//    }
//
//
//    @Override
//    public long getAvgValue(long aMin, long aMax, long tMin, long tMax) {
//        long sum = 0;
//        long count = 0;
//        NavigableMap<Long, List<Message>> subMap = msgMap.subMap(tMin, true, tMax, true);
//        for (Map.Entry<Long, List<Message>> mapEntry : subMap.entrySet()) {
//            List<Message> msgQueue = mapEntry.getValue();
//            for (Message msg : msgQueue) {
//                if (msg.getA() >= aMin && msg.getA() <= aMax) {
//                    sum += msg.getA();
//                    count++;
//                }
//            }
//        }
//
//        return count == 0 ? 0 : sum / count;
//    }
//
//}