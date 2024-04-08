
########################****************************#######################
    Buat koneksi dari psql ke apps go dengan orm sqlc
########################****************************#######################

1. Install sqlc di wsl2 setelah install postgresql di wsl2 juga
2. Buat direktori baru seetelah itu : touch main.go -> go mod init // bisa nama github.com/ardileon/nama folder atau project.

3. craete sqlc.yaml : 
version: "2"
sql:
  - engine: "postgresql" // nama engienya
    queries: "query.sql" // nanti kita buat nama filenya sama 
    schema: "schema.sql" // nanti kita buat nama filenya sama
    gen:
      go:
        package: "tutorial" // jadi setelah nanti di sqlc generate nanti turunannya pake dari nama package : tutorial
        out: "tutorial"
        sql_package: "pgx/v5" // ini kalau di hapus maka driver untuk psql ke golang juga beda

4. touch schema.sql dan query.sql
5. schema.sql  :

CREATE TABLE IF NOT EXISTS product (
  id   BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price NUMERIC(6,2) NOT NULL,
  available BOOLEAN, 
  created TIMESTAMP DEFAULT NOW()
);  

6. query.sql :

-- name: GetProduct :one // ini nantinya sebagai nama func jadi harus tepat penamaannya
SELECT * FROM product
WHERE id = $1 LIMIT 1;

-- name: AmbilProduct :many // ini nantinya sebagai nama func jadi harus tepat penamaannya
SELECT * FROM product
ORDER BY name;

7. jalankan sqlc generate
8. setelah itu kita tambahkan driver untuk menghubungkan psql ke go apps kita dengan:  
go get github.com/jackc/pgx/v5 untuk yang pakai sql_package : "pgx/v5" di yaml file kalau kg pakai go get github.com/lib/pq
9. Jangan lupa kita buat database di psql dulu dengan login dulu ke psqlnya sudo -u postgres psql -W  
10. Create database : create database namadatabase;
11. Untuk connection ke database : \c namadatabase;
12. Kita buat table di databasenya dengan copykan text yang ada di schema.sql terus enter 
13. Untuk memastikan bahwa database kita sudah memilih tabelnya tuliskan : \dt
14. Masukkin data ke table: INSERT INTO namatabelnya (nama atrribut) VALUES (nilai yang akan kita input ke attribut); contoh:
INSERT INTO product (name, price, available) VALUES ('BOOK', 10.99, true); // untuk input value type string pakai ' value stringnya ' kalau berhasil nanti ada INSERT 0 1
15. Kita bisa verfikasi value yang kita masukin ke table tadi dengan : SELECT * FROM namatablenya;
16. Untuk keluar dari terminal psql bisa di ketik : \q
17. buat koneksi dari main.go ke postgresql
18. tuliskan di main.go nya 

    import _ "github.com/lib/pq" // 

    // buat koneksi ke postgresql dengan sqlc
	connStr := "postgresql://postgresql:12345@localhost:5432/sqlctest?sslmode=disable" // jadi instance variable untuk connection ke db psql"
	// postgresql://username db postgresql : password db postresql @localhost:5432 itu entry point dari db psqlnya  /sqlctest adalah nama dari databasenya
	// sslmode=disable berati nggak ada enkripsi antara aplikasi dan databasenya
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

    