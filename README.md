# vote
一个简单的投票练手小项目

## V0版本
1. 实现了用户登录和投票的最基本的功能。
2. 使用了Go的前端模板TMPL，写了一些简单的页面。
3. 使用了Cookie。
4. 在投票的方法中，使用了MySQL的事务功能。

## V1版本
1. 将绝大数方法改为了前后端交互，使用了Jq的Ajax。
2. 将Cookie改为了Session。

## V2版本
1. 将所有接口改造为前后端交互，引入RestFul的概念。
2. 改造了验签中间件，使其能够更好的支持接口。
3. 将MySQL操作改造为原生SQL语句。
4. 增加了一些增删改查的操作。
5. 引入swagger生成接口文档 [swaggo/swag](https://github.com/swaggo/swag/blob/master/README_zh-CN.md)。
6. 将Gorm升级至新版本。
7. 引入JWT。
8. 接口测试使用APIFOX，不要用POSTMAN

