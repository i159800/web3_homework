#### 注册请求
curl -i  -X POST  -d '{"Username":"zxh1","Password":"zxh1"}' http://127.0.0.1:8080/register;echo

#### 登录请求
curl -i  -X POST  -d '{"Username":"zxh1","Password":"zxh1"}' http://127.0.0.1:8080/login;echo

#### 文章创建请求
curl -i  -X POST -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsImV4cCI6MTc1OTg1NjQ0MCwicm9sZXMiOlsiYWRtaW4iXX0.072kUMxKDiRULd47dBidRY_KZNWXlqOs3URaVf3f-A0"  -d '{"Content":"content12","Title":"title12"}' http://127.0.0.1:8080/posts;echo

#### 根据用户名字获取某个用户下的所有文章列表
curl -i -X GET http://127.0.0.1:8080/posts/zxh1
