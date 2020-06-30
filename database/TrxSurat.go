package dto

import (
	"database/sql"
)

type TrxSurat struct {
	SuratID                     int                    `db:"surat_id"`
	SuratNumber                 sql.NullString         `db:"surat_number"`
	ArsipNumber                 sql.NullString         `db:"arsip_number"`
	SuratDate                   []uint8                `db:"surat_date"`
	ReceivedDate                sql.NullTime           `db:"received_date"`
	Classification              sql.NullString         `db:"classification"`
	CategoryID                  sql.NullString         `db:"category_id"`
	PriorityID                  sql.NullString         `db:"priority_id"`
	StatusID                    sql.NullInt64          `db:"status_id"`
	Subject                     sql.NullString         `db:"subject"`
	Sender                      map[string]interface{} `db:"sender"`
	Receiver                    map[string]interface{} `db:"receiver"`
	TypeID                      sql.NullString         `db:"type_id"`
	TemplateID                  sql.NullString         `db:"template_id"`
	TemplateContent             sql.NullString         `db:"template_content"`
	SuratNotes                  sql.NullString         `db:"surat_notes"`
	IsDispositionMasuk          sql.NullInt64          `db:"is_disposition_masuk"`
	DispositionMasukNotes       sql.NullString         `db:"disposition_masuk_notes"`
	ApprovedDispositionMasuk    sql.NullInt64          `db:"approved_disposition_masuk"`
	ApprovedDispositionMasukBy  sql.NullString         `db:"approved_disposition_masuk_by"`
	ApprovedDispositionMasukAt  sql.NullString         `db:"approved_disposition_masuk_at"`
	IsDispositionKeluar         sql.NullInt64          `db:"is_disposition_keluar"`
	DispositionKeluar           sql.NullInt64          `db:"disposition_keluar"`
	ApprovedDispositionKeluar   sql.NullInt64          `db:"approved_disposition_keluar"`
	ApprovedDispositionKeluarBy sql.NullString         `db:"approved_disposition_keluar_by"`
	ApprovedDispositionKeluarAt sql.NullString         `db:"approved_disposition_keluar_at"`
	SuratAttachment             sql.NullString         `db:"surat_attachment"`
	SuratRecipient              []RelSuratRecipient    `db:"surat_recipient"`
	CreatedBy                   sql.NullString         `db:"created_by"`
	UpdatedBy                   sql.NullString         `db:"updated_by"`
	CreatedAt                   []uint8                `db:"created_at"`
	UpdatedAt                   []uint8                `db:"updated_at"`
	SuratAttachmentReply        sql.NullString         `db:"surat_attachment_reply"`
}
