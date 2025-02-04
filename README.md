# Technical_Test
Golang Developer Technical Test

## Soal 1:

## Soal 2:

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
