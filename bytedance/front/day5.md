## DAY5
### 实现浏览器存储数据
#### Cookie
一小段存储在客户端浏览器的文本信息，用于辨认用户状态
#### Web存储API(JS对象)
- localStorage：用于长久保存整个网站的数据，保存的数据没有过期时间，直到手动删除
- sessionStorage：用于临时保存同一窗口或标签页的数据，在关闭窗口或标签页后会自动删除这些数据

### 代码执行结果
```js
let arr = [1,2,3,4,5];
let arr2 = [1,,3];

// 若执行arr.length = 3, arr变为[1,2,3]
// 若执行arr[10] = 11, arr.length变为11
// 若执行delete arr[2], arr.length不变
// arr2.length为3
```