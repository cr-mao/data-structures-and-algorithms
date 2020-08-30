package queue

import "errors"

//循环队列
const QueueSize = 5 //最多存储QueueSize-1个数据
//空一个空格标识满格
type CricleQueue struct {
	data  [QueueSize]interface{} //存储数据的结构
	front int                    //头部的位置
	rear  int                    //尾部的位置
}

func InitQueue(q *CricleQueue) { //初始化，头部尾部重合，为空
	q.front = 0
	q.rear = 0
}
func Queuelength(q *CricleQueue) int { //队列长度
	return (q.rear - q.front + QueueSize) % QueueSize
}
func EnQueue(q *CricleQueue, data interface{}) (err error) {
	if (q.rear+1)%QueueSize == q.front%QueueSize {
		return errors.New("队列已经满了")
	}
	q.data[q.rear] = data             //入队
	q.rear = (q.rear + 1) % QueueSize //循环 6，1
	return nil
}

func DeQueue(q *CricleQueue) (data interface{}, err error) {
	if q.rear == q.front {
		return nil, errors.New("队列为空")
	}
	res := q.data[q.front]              //取出第一个数据
	q.data[q.front] = 0                 //清空数据
	q.front = (q.front + 1) % QueueSize //小于5  + 1，6回到1的位置
	return res, nil
}

