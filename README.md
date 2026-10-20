# Smart Runway

**Go, PostgreSQL ve Valkey teknolojileriyle geliştirilen finansal yönetim ve analiz uygulaması.**

Smart Runway; kullanıcıların gelirlerini, giderlerini, finansal işlemlerini ve nakit akışlarını tek bir platform üzerinden yönetmelerini sağlamak amacıyla geliştirilen bir yazılım projesidir.

Projenin temel hedefi; modern bir kullanıcı arayüzü, güvenilir veri yönetimi, etkileşimli grafikler ve gelecekte eklenecek yapay zekâ destekli tahmin sistemlerini bir araya getirmektir.

## Kullanılan Teknolojiler

| Teknoloji | Kullanım Amacı |
|---|---|
| Go | Backend geliştirme ve HTTP API |
| PostgreSQL 18 | İlişkisel veritabanı ve kalıcı veri saklama |
| Valkey 9 | Önbellekleme ve performans optimizasyonu |
| Docker | Konteyner yönetimi |
| Git & GitHub | Sürüm kontrolü |
| Web GUI | Planlanan kullanıcı arayüzü |

## Mevcut Geliştirme Durumu

**Aşama 1 – Backend Altyapısı**

- [x] Go geliştirme ortamı kuruldu.
- [x] PostgreSQL 18 kuruldu ve çalıştırıldı.
- [x] Docker Desktop yapılandırıldı.
- [x] Valkey 9.1.2 Docker üzerinde çalıştırıldı.
- [x] Go HTTP sunucusu oluşturuldu.
- [x] PostgreSQL ve Valkey bağlantı kontrol kodları eklendi.
- [x] GitHub deposu oluşturuldu.
- [ ] PostgreSQL ve Valkey bağlantılarının Go üzerinden doğrulanması.
- [ ] Veritabanı tablolarının tamamlanması.

## Planlanan Özellikler

### 1. Finansal Kontrol Paneli

- Toplam gelir ve gider görüntüleme
- Güncel finansal durum analizi
- Aylık ve yıllık finansal özetler
- Etkileşimli grafikler
- Kategori bazlı harcama analizi

### 2. Gelir ve Gider Yönetimi

- Gelir ve gider ekleme
- Mevcut kayıtları düzenleme ve silme
- İşlemleri kategorilere ayırma
- Tarih aralığına göre filtreleme
- PostgreSQL üzerinden güvenilir veri yönetimi

### 3. Yapay Zekâ Destekli Tahmin Sistemi

- Gelecek dönem gelir-gider tahminleri
- Nakit akışı öngörüleri
- Harcama eğilimlerinin analizi
- Finansal risk göstergeleri
- Veriye dayalı finansal öneriler

### 4. Performans ve Ölçeklenebilirlik

- Valkey ile önbellekleme
- REST API mimarisi
- Docker tabanlı servis yönetimi
- Modüler backend yapısı
- Hata yönetimi ve otomatik testler

## Kurulum

### Gereksinimler

- Go
- PostgreSQL 18
- Docker Desktop
- Valkey

### Projeyi Çalıştırma

Projeyi GitHub üzerinden klonlayın:

```bash
git clone https://github.com/Elifksaglam/smart-runway-go.git
cd smart-runway-go
```

Gerekli Go paketlerini yükleyin:

```bash
go mod download
```

Proje dizininde `.env` dosyası oluşturun:

```env
DATABASE_URL=postgres://postgres:YOUR_PASSWORD@localhost:5432/smart_runway?sslmode=disable
VALKEY_ADDR=localhost:6379
```

`YOUR_PASSWORD` kısmını kendi PostgreSQL parolanızla değiştirin. `.env` dosyasını GitHub'a yüklemeyin.

Valkey'i Docker üzerinden çalıştırın:

```bash
docker run -d --name runway-valkey -p 127.0.0.1:6379:6379 --restart unless-stopped valkey/valkey:9.1.2-alpine
```

PostgreSQL'de `smart_runway` veritabanının oluşturulduğundan emin olun.

Go sunucusunu başlatın:

```bash
go run main.go
```

Uygulama adresi:

http://localhost:8080

Veritabanı bağlantı kontrolü:

http://localhost:8080/health

**Not:** PostgreSQL ve Valkey bağlantılarının Go üzerinden başarılı şekilde kurulduğu henüz doğrulanmamıştır.

## Geliştirme Yol Haritası

| Aşama | Hedef | Durum |
|---|---|---|
| 1 | Geliştirme ortamı ve temel backend | Devam ediyor |
| 2 | PostgreSQL ve Valkey entegrasyonu | Devam ediyor |
| 3 | Finansal işlem API'leri | Planlandı |
| 4 | Modern Dashboard GUI | Planlandı |
| 5 | Finansal analiz ve AI tahminleri | Planlandı |
| 6 | Testler ve dağıtım | Planlandı |

## Proje Durumu

**Aktif Geliştirme Aşamasında**

Smart Runway henüz geliştirme sürecinin başlangıcındadır. Yeni özellikler, kullanıcı arayüzü ve sistem mimarisi aşamalı olarak geliştirilecektir.

## Lisans

Bu projeye henüz açık kaynak lisansı atanmadı. Tüm hakları saklıdır.


# Smart Runway

**A financial management and analytics application built with Go, PostgreSQL, and Valkey.**

Smart Runway is an ongoing software development project designed to help users manage income, expenses, financial records, and cash flow through an intuitive dashboard.

The project aims to combine financial data management, interactive visualizations, and future AI-powered forecasting capabilities.

## Tech Stack

| Technology | Purpose |
|---|---|
| Go | Backend development and HTTP API |
| PostgreSQL 18 | Relational database and persistent storage |
| Valkey 9 | In-memory caching |
| Docker | Containerized services |
| Git & GitHub | Version control |
| Web GUI | Planned dashboard interface |

## Current Development Status

**Phase 1 – Initial Backend Setup**

- [x] Go development environment configured
- [x] PostgreSQL installed and running
- [x] Valkey configured using Docker
- [x] Initial Go HTTP server implemented
- [x] Database connectivity health-check code added
- [x] GitHub repository initialized
- [ ] Verify PostgreSQL and Valkey connections through the API

## Planned Features

### Financial Dashboard
- Overview of income, expenses, and cash flow
- Interactive financial charts
- Monthly and annual summaries

### Transaction Management
- Add, edit, and delete transactions
- Organize transactions by category
- Store financial records in PostgreSQL

### AI-Powered Forecasting
- Future cash-flow predictions
- Expense trend analysis
- Financial risk indicators
- Data-driven financial insights

### Performance and Scalability
- Valkey caching
- REST API development
- Docker-based deployment

## Getting Started

### Prerequisites

- Go
- PostgreSQL
- Docker Desktop
- Valkey

### Run Locally

Clone the repository:

```bash
git clone https://github.com/YOUR_USERNAME/smart-runway-go.git
cd smart-runway-go
```

Install dependencies:

```bash
go mod download
```

Create a `.env` file in the project directory:

```env
DATABASE_URL=postgres://postgres:YOUR_PASSWORD@localhost:5432/smart_runway?sslmode=disable
VALKEY_ADDR=localhost:6379
```

Start Valkey:

```bash
docker run -d --name runway-valkey -p 127.0.0.1:6379:6379 valkey/valkey:9.1.2-alpine
```

Ensure the PostgreSQL database `smart_runway` exists.

Run the application:

```bash
go run main.go
```

Open:

- Application: http://localhost:8080
- Connection health check: http://localhost:8080/health

**Note:** Database connectivity must be verified locally. Do not commit `.env` files or database passwords.

## Development Roadmap

| Phase | Goal | Status |
|---|---|---|
| 1 | Initial backend infrastructure | In progress |
| 2 | PostgreSQL and Valkey integration | In progress |
| 3 | Financial transaction REST API | Planned |
| 4 | Dashboard GUI | Planned |
| 5 | Analytics and forecasting | Planned |
| 6 | Testing and deployment | Planned |

## Project Status

**Under Active Development**

Smart Runway is currently in its early development stage. Features, architecture, and documentation will evolve as development progresses.

## License

No open-source license has been assigned. All rights reserved.
