package data

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	*gorm.Model
	Name        string    gorm:"column:name;type:varchar(255)"           // Nama event
	Description string    gorm:"column:description;type:text"            // Deskripsi event
	StartDate   time.Time gorm:"column:start_date;type:timestamp"        // Tanggal mulai event
	EndDate     time.Time gorm:"column:end_date;type:timestamp"          // Tanggal berakhir event
	IsActive    bool      gorm:"column:is_active;type:bool"              // Status aktif event
}

type Voucher struct {
	*gorm.Model
	Code        string    gorm:"column:code;type:varchar(255);unique"    // Kode voucher
	Description string    gorm:"column:description;type:text"            // Deskripsi voucher
	IsActive    bool      gorm:"column:is_active;type:bool"              // Status aktif voucher
	ExpiredAt   time.Time gorm:"column:expired_at;type:timestamp"        // Tanggal kadaluarsa voucher
	EventID     uint      gorm:"column:event_id"                         // Foreign key ke tabel event
	Event       Event     gorm:"foreignKey:EventID"                      // Relasi ke tabel event
}