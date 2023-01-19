## DAY5
### RDBMS跟常见的对象存储系统的不同
RDBMS是Relational Database Management System的缩写，即关系数据库管理系统。是SQL语言和所有现代数据库系统的基础  
OODBMS是Object-oriented Database Management System的缩写，即面向对象数据库管理系统。将数据以对象形式存储  
**与常见对象存储系统的不同点**
- RDBMS中数据库支持事务
- OODBMS中数据库会向用户暴露put/get接口
- RDBMS一般存储结构化数据

### 常见的存储系统IO性能优化的方式
- 预读
- 减少IO路径上的内存拷贝
- batch写入