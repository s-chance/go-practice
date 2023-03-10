## DAY4
### MySQL数据库中实现MVCC机制的日志
MVCC就是多版本并发控制，是为了在读取数据时不加锁来提高读取效率和并发性的一种手段。
#### 日志
MySQL日志主要包括错误日志、查询日志、慢查询日志、事务日志、二进制日志几大类  
事务日志
- undo Log：事务中每次修改都会生成undo log记录，主要用于事务回滚
- redo Log：只记录事务对数据具体进行了哪些修改，主要用于解决性能问题  

二进制日志
- Binary Log：记录了所有的DDL和DML语句  

慢查询日志
- Slow Log：记录存在效率问题的查询语句
### 关于排序算法
- 归并排序任何情况下都能保持时间复杂度为O(n*log n)
- 插入排序时间复杂度为O(n*n)，但不表示在任何情况下都慢于快速排序
- 快速排序的最坏情况是初始序列已经有序，时间复杂度为O(n*n)
- 希尔排序最坏情况时间复杂度为O(n*n)
- 比较算法时，需要根据实际情况考虑，不能以最好情况去比较最坏情况，同时也不能盲目追求效率而忽视稳定性
