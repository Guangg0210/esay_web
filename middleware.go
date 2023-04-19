package web

/*
• 抽象出来一个 QueryContext，代表查询上下文
• 抽象出来一个 QueryResult，代表查询结果
• 抽象出来 Handler，代表在这个上下文里面做点什么事情
• 抽象出来 Middleware，连接不同的 Handler
*/

type Middleware func(next HandleFunc) HandleFunc
