# api生成代码
goctl api go -api book.api  -dir .

# 数据库model生成
goctl model mysql ddl -src book.sql -dir .
