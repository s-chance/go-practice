## DAY7
### 清除浮动带来的影响的方法
- 在浮动元素末尾添加一个空的标签，例如
  ```html
  <div style="clear: both"></div>
  ```
- 设置父元素overflow值为hidden
- 给父元素添加clearfix类

### 代码运行结果分析
```js
var f = function () {
    console.log("1");
}
function f() {
    console.log("2");
}

f();

// 输出结果是1
```
原理：[变量提升](https://juejin.cn/post/7007224479218663455)

