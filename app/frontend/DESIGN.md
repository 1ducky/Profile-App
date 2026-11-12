# Design Spec — Website Profil Puguh Pamungkas
 
> Kombinasi Template A & B
> Tampilan profesional, hangat, dan humanis dengan kesan rendah hati tetapi tetap berwibawa. Tidak terlalu premium, lebih dekat dengan masyarakat.
 
**Tagline:** *"Teknologi untuk manusia, bukan sebaliknya."*
 
---
 
## 1. Prinsip Desain
 
- **Profesional tapi membumi**: rapi dan kredibel, tidak mewah.
- **Humanis**: foto orang nyata, kegiatan sosial, bahasa yang hangat.
- **Fokus pada dampak**: setiap section menjawab "apa manfaatnya bagi masyarakat".
- **Bersih dan lega**: whitespace besar, hierarki jelas, mudah dipindai.
---
 
## 2. Struktur Halaman
 
| No | Section | Isi |
|----|---------|-----|
| 01 | **Hero** | Headline, subheadline, foto, CTA utama + sekunder |
| 02 | **Produk & Karya** | Highlight produk (Algebra), proyek lain, pencapaian |
| 03 | **Dampak Sosial** | Dokumentasi kegiatan, cerita singkat, dampak yang dicapai |
| 04 | **Tentang Saya** | Profil, nilai, keahlian, pengalaman singkat |
| 05 | **Artikel / Insights** | Tulisan, opini, berita, kegiatan terbaru |
| 06 | **Kontak** | Form, informasi kontak, media sosial |
 
Navigasi: `Beranda · Tentang · Produk · Dampak Sosial · Artikel · Kontak` + tombol **Hubungi Saya**.
 
---
 
## 3. Arah Visual & Gaya Desain
 
### Foto Hero
- Latar belakang **grayscale** dengan foto subjek berwarna.
- Subjek sedikit **offset** (kanan), dengan bentuk organik berwarna biru muda/teal di belakangnya.
- Memberi kesan fokus pada sosok, tetap profesional dan hangat.
### Tipografi
- Sans-serif modern, bersih, mudah dibaca, hierarki jelas.
- Pilihan font: **Inter**, **Plus Jakarta Sans**, atau sejenisnya.
- Heading: semibold/bold, tracking rapat. Body: regular, line-height longgar.
- Eyebrow/label section: huruf kapital kecil dengan letter-spacing (mis. `FOKUS UTAMA`, `PENGABDIAN`, `SIAPA SAYA`).
### Warna
- Palet terang dan natural, aksen biru/teal untuk teknologi dan hijau/earth tone untuk sosial.
- Section Dampak Sosial memakai latar krem/pasir lembut untuk nuansa hangat.
---
 
## 4. Palet Warna
 
| Peran | Nama | Hex | Penggunaan |
|-------|------|-----|------------|
| Primary | Tech / Trust | `#0B63CE` | CTA utama, link, ikon, aksen teknologi |
| Secondary | Growth / Social | `#16A34A` | Aksen sosial, indikator positif |
| Neutral | Background | `#F8FAFC` | Latar halaman |
| Text | Primary | `#1F2937` | Heading dan teks utama |
| Text Muted | Secondary | `#6B7280` | Deskripsi, caption, teks pendukung |
 
**Aksen tambahan** (dari swatch): teal `#1AA6B7`, hijau muda `#86C98B`, pasir `#E6CFA8`, krem `#F5EFE3`. Warna navbar CTA dan footer CTA memakai navy gelap (`#1F2937`).
 
```css
:root {
  --color-primary: #0B63CE;
  --color-secondary: #16A34A;
  --color-bg: #F8FAFC;
  --color-text: #1F2937;
  --color-text-muted: #6B7280;
  --color-accent-teal: #1AA6B7;
  --color-accent-green-soft: #86C98B;
  --color-accent-sand: #E6CFA8;
  --color-section-warm: #F5EFE3;
  --radius-card: 16px;
  --radius-pill: 999px;
  --shadow-card: 0 4px 16px rgba(31, 41, 55, 0.06);
}
```
 
---
 
## 5. Detail Per Section
 
### 01. Hero
- Eyebrow: `TEKNOLOGI UNTUK PENDIDIKAN & MASYARAKAT`
- Headline: **Membangun Solusi Teknologi, Menguatkan Manusia.**
- Subheadline: "Saya percaya bahwa teknologi akan bermakna lebih besar ketika bisa diakses oleh lebih banyak orang dan memberi dampak nyata bagi masyarakat."
- CTA utama: **Lihat Karya & Dampak** (pill biru)
- CTA sekunder: **Tentang Saya →** (teks + ikon panah dalam lingkaran)
- Layout: teks kiri, foto kanan, foto grayscale menyatu ke latar putih (gradient fade).
### 02. Produk & Karya
- Eyebrow: `FOKUS UTAMA` · Judul: **Produk & Karya** · Link: *Lihat semua →*
- Deskripsi: "Menghadirkan solusi teknologi yang berdampak, mulai dari digitalisasi pendidikan hingga berbagai inisiatif inovasi lainnya."
- **Kartu unggulan (Algebra)**: logo, tagline "Digitalisasi Pendidikan", deskripsi, tombol outline **Pelajari lebih lanjut →**, mockup laptop + tablet di kanan.
- **Statistik**: `+10K` Pengguna aktif · `+500` Sekolah mitra · `+20` Provinsi
- **3 kartu kecil** (ikon + judul + deskripsi + panah):
  1. Proyek Teknologi Lainnya
  2. Kolaborasi & Kemitraan
  3. Inovasi Berkelanjutan
### 03. Dampak Sosial
- Eyebrow: `PENGABDIAN` · Judul: **Dampak Sosial**
- Deskripsi: "Turun langsung ke masyarakat, mendengarkan, belajar, dan berkontribusi bersama untuk perubahan yang lebih baik."
- Link: *Lihat semua →*
- Galeri foto (1 foto besar + 2 foto kecil bertumpuk), carousel dengan **dot indicator** dan tombol panah kiri/kanan.
- Item cerita dengan ikon: **Pendidikan & Literasi Digital**.
- Latar section: krem/pasir lembut.
### 04. Tentang Saya
- Eyebrow: `SIAPA SAYA` · Judul: **Tentang Saya**
- Deskripsi: "Seorang teknolog, pembelajar, dan pegiat sosial yang percaya bahwa kolaborasi adalah kunci untuk menciptakan dampak yang lebih besar."
- CTA: **Selengkapnya tentang saya →** (pill biru)
- Daftar nilai (ikon + judul + deskripsi):
  - **IT Expert & Product Builder**
  - **Pembelajar Sepanjang Hayat**
  - **Aktivis Sosial**
### 05. Artikel / Insights
- Tulisan, opini, berita, dan kegiatan terbaru (grid kartu artikel dengan thumbnail, judul, tanggal, ringkasan).
### 06. Kontak (CTA Footer)
- Banner dengan latar foto pegunungan + overlay gelap biru.
- Judul: **Mari Terhubung**
- Teks: "Terbuka untuk kolaborasi, diskusi, dan kesempatan berdampak bersama."
- Tombol: **Hubungi Saya** (primary) + **Ikuti di Media Sosial** (outline putih).
- Tambahan: form kontak, info kontak, ikon media sosial.
---
 
## 6. Komponen Utama
 
### Navbar
- Logo + nama di kiri, menu di tengah, tombol **Hubungi Saya** (pill, biru/navy) di kanan.
- Latar putih/transparan, sticky saat scroll.
- Teks menu kecil, warna `--color-text`; item aktif diberi penekanan.
### Tombol & Link
| Varian | Gaya |
|--------|------|
| **Button Primary** | Pill, latar `#0B63CE`, teks putih |
| **Button Secondary** | Pill, outline tipis abu, teks gelap, latar putih |
| **Link Text** | Teks biru + ikon panah → |
 
### Card Produk
- Latar putih, radius besar, shadow lembut, border tipis.
- Isi: logo/ikon, judul, deskripsi singkat, link **Pelajari lebih lanjut →**.
### Card Dampak Sosial
- Thumbnail foto di kiri, judul + deskripsi di kanan, dot indicator dan tombol panah di bawah.
### Card Info Kecil
- Ikon di kiri-atas, panah di kanan-atas, judul tebal, deskripsi muted.
---
 
## 7. Responsive Behavior
 
| Perangkat | Lebar | Perilaku |
|-----------|-------|----------|
| **Desktop** | ≥ 1024px | Layout 2 kolom di Hero, kartu 3 kolom, navbar penuh |
| **Tablet** | 768 – 1023px | Grid 2 kolom, foto hero mengecil, galeri tetap carousel |
| **Mobile** | ≤ 767px | 1 kolom, navbar jadi hamburger, foto hero di atas teks, CTA full-width, kartu bertumpuk |
 
---
 
## 8. Catatan Implementasi
 
- Gunakan **grid 12 kolom**, lebar konten maksimum ± 1200px.
- Jarak antar section besar (± 80–120px desktop, 56–72px mobile).
- Foto hero dioptimalkan (WebP/AVIF), gunakan `loading="lazy"` untuk galeri.
- Pastikan kontras teks memenuhi **WCAG AA**.
- Berikan `alt` deskriptif untuk semua foto kegiatan.
- Animasi halus saja (fade/slide ringan), hormati `prefers-reduced-motion`.
 