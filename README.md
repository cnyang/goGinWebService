# 建議作法

## 把專案clone回來後，建議作法如下:

1. 先把專案/bkd-apiService/go.mod, go.sum刪除

2. 然後在cmd裡/bkd-apiService目錄下指令

   - go mod init iBP  ->會生出go.mod (大小寫要一樣，go.mod的第一行會是module iBP)

   - go mod tidy ->會生出go.sum
 
   - go mod vendor -> 會把要用到的resource抓進vendor資料夾裡

3. 最後可以用go run main.go 或是直接下air (可以hot reload，但有時候會卡住)

## Testing

在cmd裡/bkd-apiService目錄下指令

go test iBP -v -> 會跑main_test.go ( -v 是verbose會把過程印出來)
或是在vs code右側工具列的圖示裡按測試