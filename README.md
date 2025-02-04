# Technical_Test
Golang Developer Technical Test

## Soal 1:
 - saya disini menggunakan gin dengan arsitecture yang saya kuasai, dimana terdapat Direktori config, controller, manager, middleware, models, service, dan utils
1. **config/**  
   Berisi file konfigurasi yang digunakan oleh aplikasi, seperti pengaturan koneksi database, environment variables, atau konfigurasi lain yang diperlukan oleh sistem.

2. **controller/**  
   Menangani request dari klien dan memberikan response. Controller ini bertanggung jawab untuk memanggil service dan manager untuk memproses logika lebih lanjut serta memberikan hasil kepada klien.

3. **manager/**  
   Biasanya berisi logika bisnis yang lebih kompleks dan berhubungan dengan pengelolaan data atau model. Manager menangani operasi yang melibatkan lebih dari satu entitas atau operasi yang memerlukan aturan bisnis tertentu.

4. **middleware/**  
   Berfungsi untuk memfilter request sebelum diteruskan ke controller. Middleware bisa digunakan untuk berbagai tujuan seperti autentikasi, logging, validasi input, atau pengelolaan session.

5. **models/**  
   Berisi representasi data atau struktur yang digunakan dalam aplikasi, sering kali berhubungan langsung dengan database atau API eksternal. Di sini juga bisa ditemukan definisi model untuk ORM (seperti GORM) yang memetakan data ke dalam bentuk yang dapat dikelola aplikasi.

6. **service/**  
   Berisi logika yang lebih spesifik yang menangani pengolahan data atau komunikasi dengan layanan eksternal, seperti interaksi dengan database, API eksternal, atau pengolahan data yang lebih rumit.

7. **utils/**  
   Menyimpan berbagai fungsi utilitas yang dapat digunakan di berbagai bagian aplikasi, seperti helper functions, validasi umum, atau tools lainnya yang mendukung operasi aplikasi.

## Soal 2:
 - Penjelasan Pentingnya **Unit Testing** dalam Pengembangan Software:
1. Menjamin Keandalan Kode: Unit testing memastikan bahwa setiap bagian kecil dari aplikasi (fungsi atau metode) bekerja seperti yang diharapkan. Ini membantu menemukan dan memperbaiki bug pada tahap awal sebelum kode digabungkan dengan bagian lain dari sistem.

2. Meningkatkan Kepercayaan Developer: 
   Dengan memiliki unit tests yang baik, pengembang dapat lebih percaya diri dalam perubahan atau refactor kode. Hal ini juga memudahkan untuk memahami kode yang lebih kompleks, karena pengujian membuktikan bahwa fungsionalitas tertentu berjalan sesuai yang diinginkan.

3. Mempermudah Perubahan dan Refactoring: 
   Dengan adanya unit test, perubahan kode (refactoring) dapat dilakukan dengan lebih mudah. Jika ada perubahan yang tidak diinginkan, unit test akan segera memberi tahu jika ada fungsionalitas yang rusak, menghindari regresi.

4. Dokumentasi Kode: 
   Unit test berfungsi sebagai dokumentasi untuk kode. Pengembang baru dapat melihat bagaimana fungsi diharapkan bekerja hanya dengan melihat unit test yang ada.

5. Pengurangan Biaya Pemeliharaan: 
   Unit testing mengurangi biaya pemeliharaan jangka panjang, karena deteksi masalah pada tahap awal dapat mengurangi waktu yang dibutuhkan untuk mengidentifikasi dan memperbaiki bug saat aplikasi sudah berjalan.

## Soal 3:

## Soal 4:

### Sender:
- Untuk **Retry Mekanisme koneksi ke RabbitMQ**, saya set untuk 5 kali mencoba menghubungkan, dan jika gagal, akan mengulangi lagi dengan jeda 2 detik.
    - Contoh pesan error jika gagal:
    ```
    🔄 Retry 1/5 - Gagal koneksi: dial tcp [::1]:5678: connectex: No connection could be made because the target machine actively refused it.
    ```
    - Jika setelah 5 percobaan gagal, pesan berikut akan muncul:
    ```
    2025/02/04 16:42:30 ❌ Error koneksi ke RabbitMQ: gagal koneksi ke RabbitMQ setelah 5 percobaan.
    ```

- Untuk **Retry Mekanisme mengirim pesan**, proses ini juga melakukan 5 kali percobaan.
    - Contoh pesan error jika gagal:
    ```
    ⚠️ Gagal mengirim pesan. Retry 1/5 dalam 1s...
    ```
    - Jika setelah 5 percobaan tetap gagal, pesan berikut akan muncul:
    ```
    Gagal mengirim pesan setelah 5 percobaan.
    ```

### Receiver:
- Untuk **Retry Mekanisme koneksi ke RabbitMQ**, saya set untuk 5 kali mencoba menghubungkan, dan jika gagal, akan mengulangi lagi dengan jeda 2 detik.
    - Contoh pesan error jika gagal:
    ```
    🔄 Retry 1/5 - Gagal koneksi: dial tcp [::1]:5678: connectex: No connection could be made because the target machine actively refused it.
    ```
    - Jika setelah 5 percobaan gagal, pesan berikut akan muncul:
    ```
    2025/02/04 16:42:30 ❌ Error koneksi ke RabbitMQ: gagal koneksi ke RabbitMQ setelah 5 percobaan.
    ```

## Soal 5:
