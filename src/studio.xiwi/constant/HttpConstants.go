package constant

const (
	HTTP_200 int = 200 //ok  - 成功返回状态，对应，GET,PUT,PATCH,DELETE.
	HTTP_201 int = 201 //created  - 成功创建。
	HTTP_304 int = 304 //not modified   - HTTP缓存有效。
	HTTP_400 int = 400 //bad request   - 请求格式错误。
	HTTP_401 int = 401 //unauthorized   - 未授权。
	HTTP_403 int = 403 //forbidden   - 鉴权成功，但是该用户没有权限。
	HTTP_404 int = 404 //not found - 请求的资源不存在
	HTTP_405 int = 405 //method not allowed - 该http方法不被允许。
	HTTP_410 int = 410 //gone - 这个url对应的资源现在不可用。
	HTTP_415 int = 415 //unsupported media type - 请求类型错误。
	HTTP_422 int = 422 //unprocessable entity - 校验错误时用。
	HTTP_429 int = 429 //too many request - 请求过多。
)
