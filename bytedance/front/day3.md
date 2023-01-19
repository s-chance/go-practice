## DAY3
### 关于事件冒泡
在一个对象上触发事件（比如单击事件），如果此对象定义了此事件的处理程序，那么此事件就会调用这个处理程序，如果没有定义此事件的处理程序或事件返回true，那么这个事件会向这个对象的父级对象传播，直到被处理或者到达最顶层对象，document对象

### 使脚本在DOMContentLoaded事件之后加载的script标签属性
```js
<script async>
<script type="module" async>
```
