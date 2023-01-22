## DAY8
### 下列说法正确的有哪些
A. visibility:hidden 表示所占据的空间位置仍然存在，仅为视觉上的完全透明  
B. display:none 不为被隐藏的对象保留其物理空间  
C. visibility:hidden 与display:none 两者没有本质上的区别  
D. visibility:hidden 回流与重绘  

答案：AD

### 最长时间计算
若主机甲与主机已已建立一条 TCP 链接，最大段长（MSS）为 1KB，往返时间（RTT）为 2 ms，则在不出现拥塞的前提下，拥塞窗口从 8KB 增长到 32KB 所需的最长时间是48ms  

每个RTT会增加1个MSS，这里是1KB，从8KB到32KB有24KB的调整空间，需要24个RTT，于是总时长为24个RTT，即24*2ms=48ms
