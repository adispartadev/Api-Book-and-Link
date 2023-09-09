Project ini merupakan `api` dari [Book and Link](https://github.com/adispartadev/Front-Book-and-Link), yang menggunakan framework [Go fiber](https://docs.gofiber.io/).

Sebelum menjalankan project ini, silahkan sesuaikan beberapa variable berikut pada file `.env`

```
PORT=8002  

DB_HOST=
DB_PORT=
DB_NAME=
DB_USERNAME=
DB_PASSWORD=  
  
JWT_SECRET=0c5aH3zEubh032jTTTfq7XBEQqmC967EdqneuOt1BZ1jfLI2wKrON1QsEEtzWJKm  
JWT_ALGO=HS256  
  
MAIL_SMTP_HOST=smtp.gmail.com  
MAIL_SMTP_PORT=587  
MAIL_SENDER_NAME=""  
MAIL_AUTH_EMAIL=""  
MAIL_AUTH_PASSWORD=""  
  
FRONT_PAGE_BASE_URL=http://localhost:3000
```
Setelah konfigurasi disesuaikan, anda dapat menjalankan service dari aplikasi ini sebelum menjalankan versi [front end](https://github.com/adispartadev/Front-Book-and-Link) dari project ini.
