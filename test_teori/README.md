# Jawaban Soal Teori

## Database Diagram

![Database Diagram]

## Tech Stack

1. Bahasa Pemograman: Karena aplikasi ini harus bisa menyelesaikan banyak request dengan cepat, maka untuk bahasa pemograman akan menggunakan Go yang mempunyai Concurrency.
2. Database: Karna data yang akan dibuat sudah terstruktur,  butuh aggregate data dari beberapa table untuk fitur laporan maka membutukan database yang mendukung pembuatan relasi antar table. maka akan menggunakan Relational Database (MySql/Postgresql)
3. Cache: Untuk mengurangi latency dan mengatasi request data yang berulang maka akan digunakan cahce pada data menu menggunakan nosql seperti redis, cache akan diperbahuri ketika ada perubahan data di menu.
4. Application Monitoring: Karena applikasi yang sudah live di production membutuhkan tools untuk memonitoring baik logs, data tracing, query analysis, dll, maka akan menggunakan ElasticSearch yang mempunyai fitur - fitur tersebut.
5. Metric Server Monitoring: Metric dari server juga merupakan salah satu hal penting guna memastikan aplikasi berjalan baik,disini pengimplentasian grafana perlu diterapkan.
6. Containerization:  Dalam hal development seringkali terjadi perbedaan disisi environment local & server, untuk mengatasi hal tersebut, pengimplementasian docker perlu diterapkan.
7. CI/CD:  Untuk memudahkan  proses development & otomasi, penerapan jenkins perlu dilakukan