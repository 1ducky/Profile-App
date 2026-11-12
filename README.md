# Profile Application — Docker Setup

Panduan instalasi dan menjalankan aplikasi menggunakan Docker Compose.

## 1. Architecture

Aplikasi terdiri dari beberapa service:

| Service | Technology | Purpose |
|---|---|---|
| `frontend` | Astro SSR + Node.js | Web application dan reverse proxy |
| `api` | Go | Backend REST API |
| `mysql` | MySQL | Relational database |
| `s3Service` | RustFS | S3-compatible object storage |
| `s3-init` | AWS CLI | Automatic bucket initialization |
| `migration` | MySQL CLI | Database schema migration |

### Request Flow

```text
Browser
   |
   | HTTP / HTTPS
   v
Astro Frontend (:4321)
   |
   | Internal Docker Network
   v
Go Backend (:8080)
   |
   +------ MySQL (:3306)
   |
   +------ RustFS (:9000)
```

Browser hanya mengakses Astro. Seluruh request API browser diarahkan melalui Astro reverse proxy menggunakan `/api/*`.

Backend dan service infrastruktur berkomunikasi melalui internal Docker network.

---

## 2. Requirements

Pastikan perangkat sudah memiliki:

- Docker Engine / Docker Desktop
- Docker Compose v2
- Git
- Project source code

Periksa instalasi:

```bash
docker --version
docker compose version
```

## 3. Project Structure

Struktur direktori yang digunakan:

```text
project-root/
├── app/
│   ├── backend/
│   │   ├── Dockerfile
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── cmd/
│   │   ├── internal/
│   │   └── infra/
│   │       └── db/
│   │           └── schema/
│   │               └── 000_initial.sql
│   │
│   └── frontend/
│       ├── Dockerfile
│       ├── package.json
│       ├── pnpm-lock.yaml
│       ├── astro.config.mjs
│       └── src/
│
├── docker-compose.yml
├── .env
└── README.md
```

## 4. Environment Configuration

Buat file `.env` di root project, sejajar dengan `docker-compose.yml`.

Contoh konfigurasi:

```env
# =========================
# Backend HTTP
# =========================
HTTP_HOST=0.0.0.0
HTTP_PORT=8080

# =========================
# Database
# =========================
DB_HOST=mysql
DB_PORT=3306
DB_NAME=profile_db
DB_USER=profile_app
DB_PASSWORD=your_app_password
DB_ROOT_PASSWORD=your_root_password

# =========================
# JWT
# =========================
SECRET_KEY_JWT=replace_with_secure_random_secret

# =========================
# Storage
# =========================
STORAGE_ENDPOINT=http://s3Service:9000
STORAGE_REGION=us-east-1
STORAGE_ACCESS_KEY=your_storage_access_key
STORAGE_SECRET_KEY=your_storage_secret_key
STORAGE_BUCKET=profile-app

STORAGE_ROOT=storage
STORAGE_PATH_UPLOAD=uploads
STORAGE_PATH_TEMP=tmp
STORAGE_PATH_AVATAR=avatars
STORAGE_PATH_DOCUMENT=documents

# =========================
# Frontend
# =========================
API_URL=http://api:8080
APP_ORIGIN=http://localhost:4321
```

### Important Environment Rules

Environment variables memiliki lifecycle berbeda:

| Variable | Consumer | Lifecycle |
|---|---|---|
| `DB_*` | Go Backend / MySQL | Runtime |
| `SECRET_KEY_JWT` | Go Backend | Runtime |
| `STORAGE_*` | Go Backend / RustFS initialization | Runtime |
| `API_URL` | Astro SSR | Runtime |
| `APP_ORIGIN` | Astro SSR | Runtime |

**API_URL**

Gunakan hostname service Docker:

```env
API_URL=http://api:8080
```

Jangan gunakan `localhost:8080` dari dalam container frontend karena `localhost` mengarah ke container frontend itu sendiri.

**APP_ORIGIN**

Gunakan origin yang benar-benar diakses browser:

```env
APP_ORIGIN=http://localhost:4321
```

Untuk production:

```env
APP_ORIGIN=https://your-domain.com
```

Jangan commit `.env` ke Git. Gunakan `.env.example` tanpa credential asli sebagai template.

---

## 5. Persistent Volumes

Compose menggunakan external volumes:

```yaml
volumes:
  profile_db:
    external: true
    name: profile_db

  s3_storage:
    external: true
    name: s3_storage
```

Buat volume sebelum menjalankan aplikasi jika belum tersedia:

```bash
docker volume create profile_db
docker volume create s3_storage
```

Volume digunakan untuk mempertahankan data database dan object storage meskipun container dihapus atau dibuat ulang.

**Warning:** Jangan menggunakan `docker compose down -v` untuk restart aplikasi karena dapat menghapus volume yang dikelola Compose.

---

## 6. Build Docker Images

Jalankan seluruh perintah dari root project.

### Validate Compose Configuration

```bash
docker compose --env-file .env config
```

Periksa apakah seluruh interpolasi environment berhasil.

Jangan membagikan hasil konfigurasi lengkap jika mengandung credential.

### Build Backend

```bash
docker compose --env-file .env build api
```

### Build Frontend

```bash
docker compose --env-file .env build frontend
```

### Build Both Images

```bash
docker compose --env-file .env build api frontend
```

Untuk rebuild tanpa menggunakan Docker build cache:

```bash
docker compose --env-file .env build --no-cache api frontend
```

### Environment During Build

Perintah:

```bash
docker compose --env-file .env build frontend
```

membaca `.env` untuk interpolasi variabel Compose.

Namun, variabel tersebut tidak otomatis menjadi build arguments.

Pada arsitektur ini, `API_URL` dan `APP_ORIGIN` merupakan runtime environment. Karena itu, keduanya tidak perlu dimasukkan sebagai Docker build arguments.

---

## 7. Start Application

Jalankan seluruh service:

```bash
docker compose --env-file .env up -d
```

Compose akan menjalankan infrastruktur, melakukan bucket initialization, kemudian menjalankan aplikasi sesuai dependency yang didefinisikan.

Untuk melihat status:

```bash
docker compose ps
```

Untuk melihat log seluruh service:

```bash
docker compose logs -f
```

Untuk melihat log aplikasi saja:

```bash
docker compose logs -f api frontend
```

### Recreate Application Containers

Setelah konfigurasi runtime berubah:

```bash
docker compose --env-file .env up -d --force-recreate api frontend
```

Perubahan source code yang memengaruhi image membutuhkan rebuild terlebih dahulu.

---

## 8. Database Migration

Migration menggunakan service dengan profile `tools`.

Untuk menjalankan initial schema migration:

```bash
docker compose --env-file .env --profile tools run --rm migration
```

Migration membaca file:

```text
app/backend/infra/db/schema/000_initial.sql
```

Pastikan SQL migration aman dijalankan sebelum mengulangnya pada database yang sudah memiliki schema.

Jangan mengasumsikan migration otomatis idempotent.

---

## 9. Verify Runtime Environment

Pastikan environment diteruskan dengan benar ke container.

### Frontend

```bash
docker exec frontend-profile sh -c 'echo "$API_URL"'
```

Expected:

```text
http://api:8080
```

### Backend

```bash
docker exec backend-profile sh -c 'echo "$DB_HOST"'
```

Expected:

```text
mysql
```

Periksa endpoint storage:

```bash
docker exec backend-profile sh -c 'echo "$STORAGE_ENDPOINT"'
```

Expected:

```text
http://s3Service:9000
```

Hindari mencetak seluruh environment backend karena terdapat credential dan secret.

### Verify Backend Connectivity

```bash
docker exec frontend-profile node -e \
"fetch(process.env.API_URL).then(r => console.log(r.status)).catch(console.error)"
```

HTTP 404 pada root endpoint masih dapat menunjukkan bahwa koneksi HTTP berhasil. Gunakan endpoint API yang benar untuk pengujian fungsional.

---

## 10. Troubleshooting

### API_URL is not configured

Pastikan:

1. `.env` berada di root project.
2. `API_URL` sudah didefinisikan.
3. Compose meneruskan variabel melalui `environment`.
4. Astro menggunakan `process.env.API_URL` untuk SSR.
5. Container frontend sudah direcreate.

### Connection Refused

Periksa:

```bash
docker compose ps
docker compose logs api
docker compose logs frontend
```

Pastikan backend listening pada:

```env
HTTP_HOST=0.0.0.0
HTTP_PORT=8080
```

Jangan menggunakan `127.0.0.1` untuk HTTP server yang perlu diakses container lain.

### Database Connection Failed

Pastikan:

- MySQL sudah healthy.
- `DB_HOST=mysql`.
- `DB_PORT=3306`.
- Database dan user sudah tersedia.
- Credential cocok dengan database yang sudah tersimpan di volume.

Perubahan `MYSQL_PASSWORD` pada `.env` tidak otomatis mengubah password user yang sudah ada pada persistent volume.

### Storage Connection Failed

Pastikan:

- RustFS sudah healthy.
- Endpoint menggunakan hostname `s3Service`.
- Access key dan secret key sesuai.
- Bucket sudah berhasil dibuat oleh `s3-init`.

### Frontend Cannot Reach Backend

Pastikan frontend dan backend berada dalam network yang sama:

```yaml
networks:
  - profile-app-network
```

Pastikan alamat internal:

```env
API_URL=http://api:8080
```

Jangan gunakan alamat localhost untuk komunikasi antar-container.

---

## 11. Useful Docker Commands

| Task | Command |
|---|---|
| List containers | `docker compose ps` |
| View all logs | `docker compose logs -f` |
| View API logs | `docker compose logs -f api` |
| View frontend logs | `docker compose logs -f frontend` |
| Stop services | `docker compose stop` |
| Start services | `docker compose start` |
| Rebuild images | `docker compose build api frontend` |
| Recreate containers | `docker compose up -d --force-recreate api frontend` |
| Stop and remove containers | `docker compose down` |

---

## 12. Production Notes

Sebelum deployment production:

- Replace all default credentials and secrets.
- Use HTTPS through a reverse proxy or ingress.
- Configure the correct `APP_ORIGIN`.
- Do not expose MySQL or RustFS administrative interfaces publicly unless required.
- Keep `.env` outside version control.
- Maintain backups for MySQL and object storage.
- Use pinned Docker image versions instead of relying indefinitely on `latest`.
- Add backend healthcheck for reliable service readiness.
- Implement a versioned database migration strategy.

---

## 13. Git Ignore

Pastikan file berikut tidak masuk repository:

```gitignore
.env
.env.*
!.env.example

**/node_modules/
**/dist/
**/storage/
```

Sesuaikan pengecualian jika ada file environment non-secret yang memang perlu di-version-control.

---

**End of Setup Guide**