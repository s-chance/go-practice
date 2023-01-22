## DAY6
### css选择器优先级
id选择器 > 类选择器 > 标签选择器

### 分析代码结果
```js
document.body.addEventListener('click', () => {
    Promise.resolve().then(() => console.log(1));
    console.log(2);
    }, false);
document.body.addEventListener('click', () => {
    Promise.resolve().then(() => console.log(3));
    console.log(4);
}, false);

// 输出结果为2 4 1 3
```

**原理：js事件循环机制**  
js是一门单线程语言（html5中的Web-Worker并没有改变js是单线程），但是浏览器又能够处理异步请求  
> 浏览器执行线程  
> 浏览器是多进程的，每一个tab代表一个独立的进程。  
> 其中浏览器渲染进程（浏览器内核）属于浏览器多进程中的一种，主要负责页面渲染，脚本执行，事件处理等。  
> 其包含的线程有：GUI渲染线程（负责渲染页面，解析HTML，CSS构成DOM树）、JS引擎线程、事件触发线程、定时器触发线程、http请求线程等主要线程
>
> 执行中的线程  
> 主线程：即js引擎执行的线程，这个线程只有一个，页面渲染、函数处理都在这个主线程上执行。  
> 工作线程：也称为幕后线程，这个线程可能存在于浏览器或js引擎内，与主线程分开，负责处理文件读取、网络请求等异步事件。  
> 
> 任务队列  
> 所有的任务可分为同步任务和异步任务  
> 同步任务是立即执行的任务，一般会直接进入到主线程中执行  
> 异步任务是异步执行的任务，比如ajax请求，setTimeout定时函数等不会立即执行的任务  
> 异步任务会通过任务队列的机制（先进先出）来进行协调
> 
> 宏任务（Macro Task）和微任务（Micro Task）  
> 宏任务主要包含：script整体、setTimeout、setInterval、I/O、UI交互事件、setImmediate  
> 微任务主要包含：Promise、MutationObserver、process.nextTick  
> 在一次事件循环中，宏任务优先于微任务  
> 以此分析以上代码：整个script作为第一个宏任务，之后遇到第一个Promise微任务放入任务队列，  
> 之后遇到第二个宏任务，执行任务输出2，遇到第二个微任务放入任务队列，遇到第三个宏任务执行，输出4。  
> 最后将任务队列中的微任务依次执行，输出1 3



