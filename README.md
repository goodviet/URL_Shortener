# URL_Shortener
URL_Shortener_Service


[![My Skills](https://skillicons.dev/icons?i=golang,mongodb&theme=light)](https://skillicons.dev)


## Mô tả bài toán

Thiết kế một service cho phép thực thi việc làm ngắn url từ người dùng:

- Chuyển một URL dài thành URL ngắn (ví dụ: https://example.com/... → http://localhost:8080/abc123)

- Khi người dùng truy cập link ngắn, tự động redirect về URL gốc

- Theo dõi số lượt click

- Xem thông tin link, danh sách các link đã tạo

## Cách chạy Project Url_Shortener

Yêu cầu:

- Golang >= 1.20

- MongoDB (server đã cài hoặc Mongo Atlas)

Hướng dẫn:

1. Clone Project:

    ```sh
    git clone <link-repo>
    cd url-shortener
    ```
2. Cài dependencies:

    ```sh
    go mod tidy
    ```

* tạo file .env và sử dụng DB: 
# .env
Điền MongoDB URI của bạn :

    ```sh
    MONGO_URI="mongodb+srv://<username>:<password>@<cluster>/url_shortener"
    ```

3. Chạy server

   ```sh
   go run cmd/server/main.go
   ```
4. Test API bằng Postman hoặc curl:

    ```sh
    curl -X POST http://localhost:8080/api/shorten
    -H "Content-Type: application/json"
    -d '{"url":"https://example.com"}'
    ```

## Thiết kế & Quyết định kỹ thuật:

1. Tại sao lại chon Database này:

- Trong project này đang dùng là NoSQL (MongoDB) vì dữ liệu url dạng này không quá phức tạp, cứ mỗi lần ghi là một document riêng biệt.
Không có sự liên hệ giữa các bảng hay ràng buộc nhiều. Và đặc biệt sử dụng NoSQL(MongoDB) cũng giúp việc dễ thay đổi và cập nhật dữ liệu

2. Tại sao thiết kế API kiểu này:
REST API
User => Handler → Service → DB
Tách riêng ra giữa model, handler và service để có sự quản lí dễ dàng.
- model: chứa cấu trúc của schem, định nghĩa kiểu dữ liệu để ghi vào DB
- service: chứa các logic thao tacs với DB
- handler: nhận request, validate input, trả response.
=> dễ đọc và mở rộng.

## API

| Method | Endpoint              | chức năng |
|------|----------------------|------|
| POST | /api/shorten         | Tạo short URL |
| GET  | /:codeURL              | Redirect & tăng click |
| GET  | /api/links/:codeURL     | Lấy thông tin link |
| GET  | /api/links           | Lấy danh sách link |


3. Thuật toán để dùng cho việc rút gọn link và trả về một new link ngắn hơn:

- Lấy input(url) của user sau đó sẽ tạo ra output một đoạn (random a-zA-Z0-9) mã làm route để gán cho url gốc và lưu vài DB.
- Khi click vào new URL thì sẽ được redirect về link gốc.

    url original: https://example.com/some/long/path

    mã random để gán cho link gốc: abc123

    url short: http://localhost:8080/abc123

* check sự trùng lập : Kiểm tra xem đã tồn tại trong DB chưa => nếu có => tạo lại.

4. Xử lý conflict/duplicate như thế nào?

- Khi nào thì duplicate?

Nếu như lượng url lớn thì việc trùng mã random từ (random a-zA-Z0-9) vẫn có thể xảy ra.

- Em đang dùng cách là sinh ra mã random 6 ký tự nên là sẽ sinh lại mã mới.
- Giới hạn số lần và có thể output ra đoạn mã code dài hơn.

## Trade - offs:

Có 2 vấn đề để nói ở đây là việc sử dụng dùng cái nào:

1. Thuật toán generate random string

# Em chọn thuật toán random string 6 ký tự thay vì hash hoặc incremental ID theo tìm hiểu thì:

- Độ dài ngắn (6 ký tự) vừa đủ cho số lượng link nhỏ => URL ngắn gọn.

- Generate random ra nhanh

# Cách này có nhược điểm:

- Có xác suất trùng code (duplicate), cần kiểm tra DB.

- Cùng URL dài có thể tạo ra nhiều code khác nhau.

# Nhưng phù hợp vì:

- Giúp nhanh chóng có URL ngắn gọn mà vẫn đảm bảo unique.

2. Chọn MongoDB (NoSQL) thay vì SQL

# Em chọn MongoDB thay vì các DB quan hệ (MySQL, PostgreSQL) vì lý do:

Dữ liệu đơn giản, không cần nhiều tính chất quan hệ => MongoDB có sự linh hoạt schema.

Dễ mở rộng khi scale số lượng URL lớn. Dễ cập nhật dữ liệu thay đổi.

# Cách này có nhược điểm:

Không phù hợp với dữ liệu phức tạp, nhiều bảng liên quan.

Việc join và query phức tạp có thể khó khăn hơn SQL.

# Nhưng phù hợp vì:

URL shortener chủ yếu lưu từng record độc lập => MongoDB đủ đáp ứng.

## Challenge

* Duplicate code ( Trùng code. khi generate): Khi sinh random string 6 ký tự, đôi khi trùng với code đã tồn tại trong DB.
- Sinh lại random string nếu trùng, tối đa 5 lần, báo lỗi nếu vẫn trùng.
=> Xử lý duplicate là cần thiết trong các hệ thống generate code.

* Validation URL: URL không có scheme (http/https), ký tự đặc biệt, hoặc rỗng.
- Dùng regex để kiểm tra URL hợp lệ trước khi lưu.
- Loại bỏ URL rỗng và không đúng định dạng HTTP/HTTPS.
=> Validate input kỹ càng trước khi lưu vào DB, xử lý các edge case để dự án ổn định hơn.

## Limitations & Improvements:

# Code hiệnt tại thiếu:
- Custom alias: người dùng chưa thể đặt short code theo ý muốn.
- Expiration date / link hạn sử dụng: chưa hỗ trợ tự động hết hạn link.

# Nếu có thêm thời gian:
- Thêm custom alias để người dùng tự đặt short URL.
- Thêm expiration date và QR code cho mỗi link (*).

# Production-ready cần thêm gì?
- HTTPS / SSL: bảo mật khi truy cập link.
- Phục vụ lượng truy cập lớn.
- tối ưu truy vấn.

--------------------------------
