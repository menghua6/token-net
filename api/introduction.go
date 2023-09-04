package api

func Introduction() (string, error) {
	return "口令交流网 \n" + 
	"\n" +
	"创建需要的口令（示例）\n" +
	"创建限次口令 /token/10 \n" +
	"创建限时口令 /token/2023-09-21+00:00 \n" +
	"创建限时口令 /token/1h1m \n" +
	"\n" +
	"查看口令会话（示例） \n" +
	"/c31f9d6f28e243c1bef052ab46ea4b58a6d44d2505484a758a02f0dd2a63d82c \n" +
	"\n" +
	"发送口令会话（示例）\n" +
	"/c31f9d6f28e243c1bef052ab46ea4b58a6d44d2505484a758a02f0dd2a63d82c/测试 \n", nil
}
