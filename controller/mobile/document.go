package mobile

import (
	dto "PNM/database"
	"PNM/model"
	"github.com/Masterminds/squirrel"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetAllPendingList(c echo.Context) (err error) {
	id := model.GetJWTId(c)

	type ModelData struct {
		SuratID      string `json:"surat_id"`
		SuratNumber  string `json:"surat_number"`
		SuratDate    string `json:"surat_date"`
		ReceivedDate string `json:"received_date"`
		CreatedBy    string `json:"created_by"`
		DivisionName string `json:"division_name"`
		SuratNotes   string `json:"surat_notes"`
	}

	var data []ModelData

	con := model.InitDB()
	defer con.Close()

	rows, err := squirrel.Select(
		"trx_surat.surat_id",
		"surat_number",
		"surat_date",
		"received_date",
		"rel_surat_recipient.created_by",
		"rel_employee_division.division_name",
		"COALESCE(surat_notes,'')",
	).From("trx_surat").
		LeftJoin("rel_surat_recipient on rel_surat_recipient.surat_id = trx_surat.surat_id").
		LeftJoin("rel_employee_division on rel_employee_division.division_id = trx_surat.sender").
		Where(squirrel.Eq{"rel_surat_recipient.user_id": id}).
		RunWith(con).
		Query()

	if err != nil {
		return err
	}

	for rows.Next() {
		var tempdata ModelData

		if err := rows.Scan(&tempdata.SuratID,
			&tempdata.SuratNumber,
			&tempdata.ReceivedDate,
			&tempdata.CreatedBy,
			&tempdata.DivisionName,
			&tempdata.SuratNotes); err != nil {
			return err
		}

		data = append(data, tempdata)
	}

	return c.JSON(http.StatusOK, data)
}

func GetHistory(c echo.Context) (err error) {
	id := model.GetJWTId(c)

	type ModelData struct {
		SuratID      string `json:"surat_id"`
		SuratNumber  string `json:"surat_number"`
		SuratDate    string `json:"surat_date"`
		ReceivedDate string `json:"received_date"`
		CreatedBy    string `json:"created_by"`
		DivisionName string `json:"division_name"`
		SuratNotes   string `json:"surat_notes"`
	}

	var data []ModelData

	con := model.InitDB()
	defer con.Close()

	rows, err := squirrel.Select(
		"trx_surat.surat_id",
		"surat_number",
		"surat_date",
		"received_date",
		"rel_surat_recipient.created_by",
		"rel_employee_division.division_name",
		"COALESCE(surat_notes,'')",
	).From("trx_surat").
		LeftJoin("rel_surat_recipient on rel_surat_recipient.surat_id = trx_surat.surat_id").
		LeftJoin("rel_employee_division on rel_employee_division.division_id = trx_surat.sender").
		Where(squirrel.Eq{"rel_surat_recipient.user_id": id}).
		//Where(squirrel.Eq{"rel_surat_recipient.sign_status": 0}).
		RunWith(con).
		Query()

	if err != nil {
		return err
	}

	for rows.Next() {
		var tempdata ModelData

		if err := rows.Scan(
			&tempdata.SuratID,
			&tempdata.SuratNumber,
			&tempdata.ReceivedDate,
			&tempdata.CreatedBy,
			&tempdata.DivisionName,
			&tempdata.SuratNotes); err != nil {
			return err
		}

		data = append(data, tempdata)
	}

	return c.JSON(http.StatusOK, data)
}

func GetDocsById(c echo.Context) (err error) {
	request := new(model.FindId)
	type SuratData struct {
		SuratId              uint64                  `json:"surat_id"`
		SuratNumber          string                  `json:"surat_number"`
		ArsipNumber          string                  `json:"arsip_number"`
		SuratDate            string                  `json:"surat_date"`
		ReceivedDate         string                  `json:"received_date"`
		Classification       string                  `json:"classification"`
		Category             string                  `json:"category"`
		CategoryId           string                  `json:"category_id"`
		Priority             string                  `json:"priority"`
		PriorityId           string                  `json:"priority_id"`
		Status               string                  `json:"status"`
		StatusId             string                  `json:"status_id"`
		Subject              string                  `json:"subject"`
		SenderName           string                  `json:"sender_name"`
		SenderId             string                  `json:"sender_id"`
		SenderDivision       string                  `json:"sender_division"`
		SenderDivisionId     string                  `json:"sender_division_id"`
		SenderPosition       string                  `json:"sender_position"`
		SenderPositionId     string                  `json:"sender_position_id"`
		ReceiverName         string                  `json:"receiver_name"`
		ReceiverId           string                  `json:"receiver_id"`
		ReceiverDivision     string                  `json:"receiver_division"`
		ReceiverDivisionId   string                  `json:"receiver_division_id"`
		ReceiverPosition     string                  `json:"receiver_position"`
		ReceiverPositionId   string                  `json:"receiver_position_id"`
		SuratType            string                  `json:"surat_type"`
		SuratTypeId          string                  `json:"surat_type_id"`
		TemplateName         string                  `json:"template_name"`
		TemplateId           string                  `json:"template_name_id"`
		TemplateContent      string                  `json:"template_content"`
		TemplateHeader       string                  `json:"template_header"`
		TemplateFooter       string                  `json:"template_footer"`
		SuratNotes           string                  `json:"surat_notes"`
		SuratAttachment      string                  `json:"surat_attachment"`
		SuratAttachmentReply string                  `json:"surat_attachment_reply"`
		SuratRecipient       []dto.RelSuratRecipient `json:"surat_recipient"`
	}
	var (
		surat SuratData
	)
	if err := c.Bind(request); err != nil {
		return err
	}
	con := model.InitDB()
	defer con.Close()

	if err := squirrel.Select(
		"COALESCE(trx_surat.surat_id,'')",
		"COALESCE(trx_surat.surat_number,'')",
		"COALESCE(trx_surat.arsip_number,'')",
		"COALESCE(trx_surat.surat_date,'')",
		"COALESCE(trx_surat.received_date,'')",
		"COALESCE(trx_surat.classification,'')",
		"COALESCE(c.category_name,'')",
		"COALESCE(c.category_id,'')",
		"COALESCE(p.priority_name,'')",
		"COALESCE(p.priority_id,'')",
		"COALESCE(st.status_alias,'')",
		"COALESCE(st.status_id,'')",
		"COALESCE(trx_surat.subject,'')",
		"COALESCE(s.username as sender_name,'')",
		"COALESCE(s.user_id as sender_id,'')",
		"COALESCE(sd.division_name as sender_division,'')",
		"COALESCE(sd.division_id as sender_division_id,'')",
		"COALESCE(sp.position_name as sender_position,'')",
		"COALESCE(sp.position_id as sender_position_id,'')",
		"COALESCE(r.username as receiver_name,'')",
		"COALESCE(r.user_id as receiver_id,'')",
		"COALESCE(rd.division_name as receiver_division,'')",
		"COALESCE(rd.division_id as receiver_division_id,'')",
		"COALESCE(rp.position_name as receiver_position,'')",
		"COALESCE(rp.position_id as receiver_position_id,'')",
		"COALESCE(ty.type_name as surat_type,'')",
		"COALESCE(trx_surat.type_id as surat_type_id,'')",
		"COALESCE(tt.memo_name as template_name,'')",
		"COALESCE(tt.template_id as template_id,'')",
		"COALESCE(trx_surat.template_content,'')",
		"COALESCE(tt.memo_header as template_header,'')",
		"COALESCE(tt.memo_footer as template_footer,'')",
		"COALESCE(trx_surat.surat_notes,'')",
		"COALESCE(trx_surat.surat_attachment,'')",
		"COALESCE(trx_surat.surat_attachment_reply,'')",
	).
		From("trx_surat").
		Where(squirrel.Eq{"trx_surat.surat_id": request.Id}).
		LeftJoin("mst_surat_category c on c.category_id = trx_surat.category_id").
		LeftJoin("mst_surat_priority p on p.priority_id = trx_surat.priority_id").
		LeftJoin("mst_surat_status st on st.status_id = trx_surat.status_id").
		LeftJoin("mst_surat_type ty on ty.type_id = trx_surat.type_id").
		LeftJoin("mst_memo_template tt on tt.template_id = trx_surat.template_id").
		LeftJoin("mst_user s on s.user_id = trx_surat.sender").
		LeftJoin("ref_employee_division sd on sd.division_id = s.division_id").
		LeftJoin("ref_employee_position sp on sp.position_id = s.position_id").
		LeftJoin("mst_user r on r.user_id = trx_surat.receiver").
		LeftJoin("ref_employee_division rd on r.division_id = rd.division_id").
		LeftJoin("ref_employee_position rp on r.position_id = rp.position_id").
		RunWith(con).
		QueryRow().
		Scan(
			&surat.SuratId,
			&surat.SuratNumber,
			&surat.ArsipNumber,
			&surat.SuratDate,
			&surat.ReceivedDate,
			&surat.Classification,
			&surat.Category,
			&surat.CategoryId,
			&surat.Priority,
			&surat.PriorityId,
			&surat.Status,
			&surat.StatusId,
			&surat.Subject,
			&surat.SenderName,
			&surat.SenderId,
			&surat.SenderDivision,
			&surat.SenderDivisionId,
			&surat.SenderPosition,
			&surat.SenderPositionId,
			&surat.ReceiverName,
			&surat.ReceiverId,
			&surat.ReceiverDivision,
			&surat.ReceiverDivisionId,
			&surat.ReceiverPosition,
			&surat.ReceiverPositionId,
			&surat.SuratType,
			&surat.SuratTypeId,
			&surat.TemplateName,
			&surat.TemplateId,
			&surat.TemplateContent,
			&surat.TemplateHeader,
			&surat.TemplateFooter,
			&surat.SuratNotes,
			&surat.SuratAttachment,
			&surat.SuratAttachmentReply,
		); err != nil {
		return err
	}

	if surat.SuratAttachment != "" {
		surat.SuratAttachment = "uploads/surat_attachment/" + surat.SuratAttachment
	}

	if surat.SuratAttachmentReply != "" {
		surat.SuratAttachmentReply = "uploads/surat_attachment/" + surat.SuratAttachmentReply
	}

	recRows, err := squirrel.Select("recipient_name").
		From("rel_surat_recipient").
		Where(squirrel.Eq{"surat_id": surat.SuratId}).RunWith(con).Query()
	if err != nil {
		return err
	}
	defer recRows.Close()

	var recipientNameList []dto.RelSuratRecipient

	for recRows.Next() {
		var recipientName dto.RelSuratRecipient
		if err := recRows.Scan(&recipientName.RecipientName); err != nil {
			return err
		}
		recipientNameList = append(recipientNameList, recipientName)
	}

	surat.SuratRecipient = recipientNameList

	return c.JSON(http.StatusOK, surat)
}
func SubmitSignature(c echo.Context) (err error) {
	suratId := c.FormValue("surat_id")
	notes := c.FormValue("notes")
	gmbr, err := c.FormFile("image")
	userId := model.GetJWTId(c)
	username := model.GetJWTName(c)

	if suratId == "" {
		return echo.ErrBadRequest
	}

	namaFile, err := model.InsertImage(gmbr, "uploads/signature",suratId + "_" + userId + "_" + time.Now().Format("20060102150405"))
	if err != nil {
		return err
	}
	query := squirrel.Update("rel_surat_recipient").
		Where(squirrel.Eq{
			"surat_id": suratId,
			"user_id": userId,
		}).
		Set("sign_status", 1).
		Set("updatedBy", username).
		Set("signature_image", namaFile).
		Set("updated_at", time.Now())

	if notes != "" {
		query.Set("notes", notes)
	}

	return c.JSON(http.StatusOK, model.Response{Message: "Updated"})

}
