# Technical_Test
Golang Developer Technical Test

Soal 1:

Soal 2:

Soal 3:

Soal 4:
untuk sender
 >Untuk Rerty Mekanisme koneksi ke rabbitMQ saya set untuk 5 kali mencoba menghubungkan dan jika gagal akan mengulangi lagi tapi menunggu dulu selama 2 detik
 >> 🔄 Retry 1/5 - Gagal koneksi: dial tcp [::1]:5678: connectex: No connection could be made because the target machine actively refused it. (pesan error jika gagal dan akan mengulangi sebanyak 5 kali)
 >> 2025/02/04 16:42:30 ❌ Error koneksi ke RabbitMQ: gagal koneksi ke RabbitMQ setelah 5 percobaan (jika sudah 5 kali percobaan tetap gagal akan ada pesan seperti ini)
 >Untuk Rerty Mekanisme mengirim pesan akan melakukan 5 kali percobaan juga
 >>⚠️ Gagal mengirim pesan. Retry 1/5 dalam 1s... (pesan error jika gagal dan akan mengulangi sebanyak 5 kali)
 >> gagal mengirim pesan setelah 5 percobaan (jika sudah 5 kali percobaan tetap gagal akan ada pesan seperti ini)
untuk receiver
 > Untuk Rerty Mekanisme koneksi ke rabbitMQ saya set untuk 5 kali mencoba menghubungkan dan jika gagal akan mengulangi lagi tapi menunggu dulu selama 2 detik
 >>🔄 Retry 1/5 - Gagal koneksi: dial tcp [::1]:5678: connectex: No connection could be made because the target machine actively refused it. (pesan error jika gagal dan akan mengulangi sebanyak 5 kali)
 >>2025/02/04 16:42:30 ❌ Error koneksi ke RabbitMQ: gagal koneksi ke RabbitMQ setelah 5 percobaan (jika sudah 5 kali percobaan tetap gagal akan ada pesan seperti ini)

 Soal 5: