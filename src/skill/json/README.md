# Json解析技巧

## 适用场景
由于Golang是静态语言，所以在解析json数据的时候必须定义清楚结构。
但有时候我们对接的接口，正确返回和错误返回是不一样的接口，从而造成解析出错。
为了解决这个问题，我们可以采用实现Json的解析接口来进行特殊判断