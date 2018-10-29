# 匿名Struct的大用途

## 适用场景
经常我们使用Struct结构只是临时一用，所以不想利用type进行定义，这时候就可以使用匿名struct来解决了。

## 代码说明
* StudentSimpleData接口：用于返回学生基本信息（ID+Name），采用匿名struct实现
* StudentFullData接口：用于返回学生所有信息，不采用匿名struct实现