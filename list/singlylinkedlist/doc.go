package singlylinkedlist

//head 元素 value 一直为nil      即 list.head.pNext 为第一个节点

//nil
//3->2->1->nil
// 头插速度快
//尾部查入 要遍历到最后

//链表 适用于 频繁到 删除和插入，查询少       不需要知道长度，遍历到nil 跳出
//数组 适用于 频繁到查找 ，删除和插入比较少，  要知道长度
//进程列表  双链表
