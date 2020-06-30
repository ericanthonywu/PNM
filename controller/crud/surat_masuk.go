package crud

import (
	dto "PNM/database"
	"PNM/model"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/labstack/echo/v4"
	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"net/http"
	"strings"
	"strconv"
)

func CreateSuratMasuk(c echo.Context) (err error) {
	request := new(model.SuratMasukReq)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	if _, err := squirrel.Insert("trx_surat").
		Columns(
			"classification",
			"category_id",
			"priority_id",
			"type_id",
			"subject",
			"sender", // dr session login, ambil id di concat bosQ
			"status_id",
			"template_id",
		).
		Values(
			request.Classification,
			request.CategoryID,
			request.PriorityID,
			request.TypeID,
			request.Subject,
			2,
			1,
			1,
		).RunWith(con).Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, model.Response{Message: "created"})
}

func EditSuratMasuk(c echo.Context) (err error) {
	var (
		suratAttachmentCheckFile      = true
		suratAttachmentReplyCheckFile = true
	)

	con := model.InitDB()
	defer con.Close()

	Query := squirrel.Update("trx_surat")

	if c.FormValue("surat_number") != "" {
		Query = Query.Set("surat_number", c.FormValue("surat_number"))
	}
	if c.FormValue("arsip_number") != "" {
		Query = Query.Set("arsip_number", c.FormValue("arsip_number"))
	}
	if c.FormValue("received_date") != "" {
		Query = Query.Set("received_date", c.FormValue("received_date"))
	}

	suratAttachment, err := c.FormFile("surat_attachment")
	if err != nil {
		if err.Error() == "http: no such file" {
			suratAttachmentCheckFile = false
		} else {
			return err
		}
	}
	suratAttachmentReply, err := c.FormFile("surat_attachment_reply")
	if err != nil {
		if err.Error() == "http: no such file" {
			suratAttachmentReplyCheckFile = false
		} else {
			return err
		}
	}
	if suratAttachmentCheckFile {
		fmt.Println("Masuk2")
		suratAttachmentFilename, err := model.InsertImage(suratAttachment, "uploads/surat_attachment/")
		if err != nil {
			return err
		}
		Query = Query.Set("surat_attachment", suratAttachmentFilename)
	}
	if suratAttachmentReplyCheckFile {
		suratAttachmentReplyFilename, err := model.InsertImage(suratAttachmentReply, "uploads/surat_attachment_reply/")
		if err != nil {
			return err
		}
		Query = Query.Set("surat_attachment_reply", suratAttachmentReplyFilename)
	}

	if c.FormValue("classification") != "" {
		Query = Query.Set("classification", c.FormValue("classification"))
	}
	if c.FormValue("category_id") != "" {
		//n := new(big.Int)
		//categoryId, ok := n.SetString(c.FormValue("category_id"), 10)
		//if !ok {
		////	return echo.ErrBadRequest
		//}
		Query = Query.Set("category_id", c.FormValue("category_id"))
	}
	if c.FormValue("status_id") != "" {
		//n := new(big.Int)
		//statusId, ok := n.SetString(c.FormValue("status_id"), 10)
		//if !ok {
		////	return echo.ErrBadRequest
		//}
		Query = Query.Set("status_id", c.FormValue("status_id"))
	}
	if c.FormValue("priority_id") != "" {
		//n := new(big.Int)
		//priorityId, ok := n.SetString(c.FormValue("priority_id"), 10)
		//if !ok {
		////	return echo.ErrBadRequest
		//}
		Query = Query.Set("priority_id", c.FormValue("priority_id"))
	}

	if c.FormValue("type_id") != "" {
		//n := new(big.Int)
		//typeId, ok := n.SetString(c.FormValue("type_id"), 10)
		//if !ok {
		////	return echo.ErrBadRequest
		//}
		Query = Query.Set("type_id", c.FormValue("type_id"))
	}

	if c.FormValue("template_id") != "" {
		//n := new(big.Int)
		//templateId, ok := n.SetString(c.FormValue("template_id"), 10)
		//if !ok {
		////	return echo.ErrBadRequest
		//}
		Query = Query.Set("template_id", c.FormValue("template_id"))
	}
	if c.FormValue("is_disposition_masuk") != "" {
		//n := new(big.Int)
		//isDispositionMasuk, ok := n.SetString(c.FormValue("is_disposition_masuk"), 10)
		//if !ok {
		////	return echo.ErrBadRequest
		//}
		Query = Query.Set("is_disposition_masuk", c.FormValue("is_disposition_masuk"))
	}
	if c.FormValue("approved_disposition_masuk") != "" {
		//n := new(big.Int)
		//approvedDispositionMasuk, ok := n.SetString(c.FormValue("approved_disposition_masuk"), 10)
		//if !ok {
		////	return echo.ErrBadRequest
		//}
		Query = Query.Set("approved_disposition_masuk", c.FormValue("approved_disposition_masuk"))
	}
	if c.FormValue("is_disposition_keluar") != "" {
		//n := new(big.Int)
		//var isDispotitionKeluar, ok = n.SetString(c.FormValue("is_dispotition_keluar"), 10)
		//if !ok {
		////	return echo.ErrBadRequest
		//}
		Query = Query.Set("is_disposition_keluar", c.FormValue("is_dispotition_keluar"))
	}
	if c.FormValue("approved_disposition_keluar") != "" {
		//n := new(big.Int)
		//approvedDispotitionKeluar, ok := n.SetString(c.FormValue("approved_dispotition_keluar"), 10)
		//if !ok {
		////	return echo.ErrBadRequest
		//}
		Query = Query.Set("approved_disposition_keluar", c.FormValue("approved_dispotition_keluar"))
	}
	if c.FormValue("disposition_keluar") != "" {
		//n := new(big.Int)
		//dispositionKeluar, ok := n.SetString(c.FormValue("disposition_keluar"), 10)
		//if !ok {
		////	return echo.ErrBadRequest
		//}
		Query = Query.Set("disposition_keluar", c.FormValue("disposition_keluar"))
	}

	if c.FormValue("subject") != "" {
		Query = Query.Set("subject", c.FormValue("subject"))
	}

	if c.FormValue("sender") != "" {
		Query = Query.Set("sender", c.FormValue("sender"))
	}
	if c.FormValue("receiver") != "" {
		Query = Query.Set("receiver", c.FormValue("receiver"))
	}

	if c.FormValue("template_content") != "" {
		Query = Query.Set("template_content", c.FormValue("template_content"))
	}

	if c.FormValue("surat_notes") != "" {
		Query = Query.Set("surat_notes", c.FormValue("surat_notes"))
	}

	if c.FormValue("disposition_masuk_notes") != "" {
		Query = Query.Set("disposition_masuk_notes", c.FormValue("disposition_masuk_notes"))
	}

	if c.FormValue("approved_disposition_masuk_by") != "" {
		Query = Query.Set("approved_disposition_masuk_by", c.FormValue("approved_disposition_masuk_by"))
	}

	if _, err := Query.
		Where(squirrel.Eq{"surat_id": c.FormValue("surat_id")}).
		RunWith(con).Exec(); err != nil {
		fmt.Println(Query.ToSql())
		return err
	}

	return c.JSON(http.StatusOK, model.Response{Message: "Updated"})
}

func EditStatus(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)

	if err := c.Bind(request); err != nil {
		return err
	}

	con := model.InitDB()
	defer con.Close()

	//if err := squirrel.
	//	Select("status_id").
	//	From("trx_surat").
	//	Where(squirrel.Eq{"surat_id": request.Id}).
	//	RunWith(con).QueryRow().Scan(&request.StatusId); err != nil {
	//	return err
	//}

	//if surat.StatusID != (request.StatusID + 1))  {
	//	return c.JSON(http.StatusOK, model.Response{Message: "Ente mau ngehack?"})
	//}

	//mssql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.AtP)

	if _, err := squirrel.
		Update("trx_surat").
		Set("status_id", request.StatusId).
		Where(squirrel.Eq{"surat_id": request.Id}).
		RunWith(con).
		Exec(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, model.Response{Message: "updated"})
}

func ViewByStatus(c echo.Context) (err error) {
	request := new(model.MasterCrudDefault)
	type SuratData struct {
		SuratId              uint64                  `json:"surat_id"`
		SuratNumber          sql.NullString          `json:"surat_number"`
		ArsipNumber          sql.NullString          `json:"arsip_number"`
		SuratDate            sql.NullTime            `json:"surat_date"`
		ReceivedDate         sql.NullTime            `json:"received_date"`
		Classification       sql.NullString          `json:"classification"`
		Category             sql.NullString          `json:"category"`
		CategoryId           sql.NullInt64           `json:"category_id"`
		Priority             sql.NullString          `json:"priority"`
		PriorityId           sql.NullString          `json:"priority_id"`
		Status               sql.NullString          `json:"status"`
		StatusId             sql.NullString          `json:"status_id"`
		Subject              sql.NullString          `json:"subject"`
		SenderName           sql.NullString          `json:"sender_name"`
		SenderId             sql.NullString          `json:"sender_id"`
		SenderDivision       sql.NullString          `json:"sender_division"`
		SenderDivisionId     sql.NullString          `json:"sender_division_id"`
		SenderPosition       sql.NullString          `json:"sender_position"`
		SenderPositionId     sql.NullString          `json:"sender_position_id"`
		ReceiverName         sql.NullString          `json:"receiver_name"`
		ReceiverId           sql.NullString          `json:"receiver_id"`
		ReceiverDivision     sql.NullString          `json:"receiver_division"`
		ReceiverDivisionId   sql.NullString          `json:"receiver_division_id"`
		ReceiverPosition     sql.NullString          `json:"receiver_position"`
		ReceiverPositionId   sql.NullString          `json:"receiver_position_id"`
		SuratType            sql.NullString          `json:"surat_type"`
		SuratTypeId          sql.NullString          `json:"surat_type_id"`
		TemplateName         sql.NullString          `json:"template_name"`
		TemplateId           sql.NullString          `json:"template_name_id"`
		TemplateContent      sql.NullString          `json:"template_content"`
		TemplateHeader 		 sql.NullString          `json:"template_header"`
		TemplateFooter       sql.NullString          `json:"template_footer"`
		SuratNotes           sql.NullString          `json:"surat_notes"`
		SuratAttachment      sql.NullString          `json:"surat_attachment"`
		SuratAttachmentReply sql.NullString          `json:"surat_attachment_reply"`
		SuratRecipient       []dto.RelSuratRecipient `json:"surat_recipient"`
	}
	var (
		suratArray []SuratData
		statusList []int
	)
	if err := c.Bind(request); err != nil {
		return err
	}
	con := model.InitDB()
	defer con.Close()

	switch request.StatusName {
	case "srtmsk":
		statusList = []int{1, 5, 6, 9}
		break
	case "dispmsk":
		statusList = []int{2}
		break
	case "srtklr":
		statusList = []int{3, 7}
		break
	case "dispklr":
		statusList = []int{4, 5, 6, 7, 8, 9}
		break
	default:
		return echo.ErrBadRequest
	}

	rows, err := squirrel.Select(
		"trx_surat.surat_id",
		"trx_surat.surat_number",
		"trx_surat.arsip_number",
		"trx_surat.surat_date",
		"trx_surat.received_date",
		"trx_surat.classification",
		"c.category_name",
		"c.category_id",
		"p.priority_name",
		"p.priority_id",
		"st.status_alias",
		"st.status_id",
		"trx_surat.subject",
		"s.username as sender_name",
		"s.user_id as sender_id",
		"sd.division_name as sender_division",
		"sd.division_id as sender_division_id",
		"sp.position_name as sender_position",
		"sp.position_id as sender_position_id",
		"r.username as receiver_name",
		"r.user_id as receiver_id",
		"rd.division_name as receiver_division",
		"rd.division_id as receiver_division_id",
		"rp.position_name as receiver_position",
		"rp.position_id as receiver_position_id",
		"ty.type_name as surat_type",
		"trx_surat.type_id as surat_type_id",
		"tt.memo_name as template_name",
		"tt.template_id as template_id",
		"trx_surat.template_content",
		"tt.memo_header as template_header",
		"tt.memo_footer as template_footer",
		"trx_surat.surat_notes",
		"trx_surat.surat_attachment",
		"trx_surat.surat_attachment_reply",
	).
		From("trx_surat").
		Where(squirrel.Eq{"trx_surat.status_id": statusList}).
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
		Limit(request.Limit).
		Offset(request.Offset).
		RunWith(con).
		Query()

	var count int

	if err := squirrel.Select("count(*)").From("trx_surat").RunWith(con).QueryRow().Scan(&count); err != nil {
		return err
	}

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var surat SuratData
		if err := rows.Scan(
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
		if surat.SuratAttachment.Valid {
			surat.SuratAttachment.String = "uploads/surat_attachment/" + surat.SuratAttachment.String
		}

		if surat.SuratAttachmentReply.Valid {
			surat.SuratAttachmentReply.String = "uploads/surat_attachment_reply/" + surat.SuratAttachmentReply.String
		}

		recRows, err := squirrel.Select("recipient_name").
			From("rel_surat_recipient").
			Where(squirrel.Eq{"surat_id": surat.SuratId}).RunWith(con).Query()
		if err != nil {
			return err
		}

		var recipientNameList []dto.RelSuratRecipient

		for recRows.Next() {
			var recipientName dto.RelSuratRecipient
			if err := recRows.Scan(&recipientName.RecipientName); err != nil {
				return err
			}
			recipientNameList = append(recipientNameList, recipientName)
		}

		surat.SuratRecipient = recipientNameList
		suratArray = append(suratArray, surat)
	}

	return c.JSON(http.StatusOK, model.CustomResponse{
		"data": suratArray,
		"count": count,
	})
}

func ViewById(c echo.Context) (err error) {
	request := new(model.FindId)
	type SuratData struct {
		SuratId              uint64                  `json:"surat_id"`
		SuratNumber          sql.NullString          `json:"surat_number"`
		ArsipNumber          sql.NullString          `json:"arsip_number"`
		SuratDate            sql.NullTime            `json:"surat_date"`
		ReceivedDate         sql.NullTime            `json:"received_date"`
		Classification       sql.NullString          `json:"classification"`
		Category             sql.NullString          `json:"category"`
		CategoryId           sql.NullInt64           `json:"category_id"`
		Priority             sql.NullString          `json:"priority"`
		PriorityId           sql.NullString          `json:"priority_id"`
		Status               sql.NullString          `json:"status"`
		StatusId             sql.NullString          `json:"status_id"`
		Subject              sql.NullString          `json:"subject"`
		SenderName           sql.NullString          `json:"sender_name"`
		SenderId             sql.NullString          `json:"sender_id"`
		SenderDivision       sql.NullString          `json:"sender_division"`
		SenderDivisionId     sql.NullString          `json:"sender_division_id"`
		SenderPosition       sql.NullString          `json:"sender_position"`
		SenderPositionId     sql.NullString          `json:"sender_position_id"`
		ReceiverName         sql.NullString          `json:"receiver_name"`
		ReceiverId           sql.NullString          `json:"receiver_id"`
		ReceiverDivision     sql.NullString          `json:"receiver_division"`
		ReceiverDivisionId   sql.NullString          `json:"receiver_division_id"`
		ReceiverPosition     sql.NullString          `json:"receiver_position"`
		ReceiverPositionId   sql.NullString          `json:"receiver_position_id"`
		SuratType            sql.NullString          `json:"surat_type"`
		SuratTypeId          sql.NullString          `json:"surat_type_id"`
		TemplateName         sql.NullString          `json:"template_name"`
		TemplateId           sql.NullString          `json:"template_name_id"`
		TemplateContent      sql.NullString          `json:"template_content"`
		TemplateHeader 		 sql.NullString          `json:"template_header"`
		TemplateFooter       sql.NullString          `json:"template_footer"`
		SuratNotes           sql.NullString          `json:"surat_notes"`
		SuratAttachment      sql.NullString          `json:"surat_attachment"`
		SuratAttachmentReply sql.NullString          `json:"surat_attachment_reply"`
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
		"trx_surat.surat_id",
		"trx_surat.surat_number",
		"trx_surat.arsip_number",
		"trx_surat.surat_date",
		"trx_surat.received_date",
		"trx_surat.classification",
		"c.category_name",
		"c.category_id",
		"p.priority_name",
		"p.priority_id",
		"st.status_alias",
		"st.status_id",
		"trx_surat.subject",
		"s.username as sender_name",
		"s.user_id as sender_id",
		"sd.division_name as sender_division",
		"sd.division_id as sender_division_id",
		"sp.position_name as sender_position",
		"sp.position_id as sender_position_id",
		"r.username as receiver_name",
		"r.user_id as receiver_id",
		"rd.division_name as receiver_division",
		"rd.division_id as receiver_division_id",
		"rp.position_name as receiver_position",
		"rp.position_id as receiver_position_id",
		"ty.type_name as surat_type",
		"trx_surat.type_id as surat_type_id",
		"tt.memo_name as template_name",
		"tt.template_id as template_id",
		"trx_surat.template_content",
		"tt.memo_header as template_header",
		"tt.memo_footer as template_footer",
		"trx_surat.surat_notes",
		"trx_surat.surat_attachment",
		"trx_surat.surat_attachment_reply",
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
	if surat.SuratAttachment.Valid {
		surat.SuratAttachment.String = "uploads/surat_attachment/" + surat.SuratAttachment.String
	}
	if surat.SuratAttachmentReply.Valid {
		surat.SuratAttachmentReply.String = "uploads/surat_attachment/" + surat.SuratAttachmentReply.String
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

func PrintPDF(c echo.Context) (err error) {
	request := new(model.FindId)
	type PrintData struct {
		TemplateContent      sql.NullString          `json:"template_content"`
		TemplateHeader 		 sql.NullString          `json:"template_header"`
		TemplateFooter       sql.NullString          `json:"template_footer"`
	}
	var (
		printdata PrintData
	)
	if err := c.Bind(request); err != nil {
		return err
	}
	con := model.InitDB()
	defer con.Close()

	if err := squirrel.Select(
		"trx_surat.template_content",
		"tt.memo_header as template_header",
		"tt.memo_footer as template_footer",
	).
		From("trx_surat").
		Where(squirrel.Eq{"trx_surat.surat_id": request.Id}).
		LeftJoin("mst_memo_template tt on tt.template_id = trx_surat.template_id").
		RunWith(con).
		QueryRow().
		Scan(
			&printdata.TemplateContent,
			&printdata.TemplateHeader,
			&printdata.TemplateFooter,
		); err != nil {
		return err
	}

	format := printdata.TemplateHeader.String + printdata.TemplateContent.String + printdata.TemplateFooter.String

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
    if err != nil {
        return err
    }

    wkhtmltopdf.SetPath(`C://Users/Josh/go/pkg/mod/github.com/!sebastiaan!klippert/go-wkhtmltopdf@v1.5.0`)

    pdfg.Dpi.Set(300)
	pdfg.Orientation.Set("Portrait")
	pdfg.Grayscale.Set(false)

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(format)))

	err = pdfg.Create()
	if err != nil {
		return err
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("./model/uploads/cetak_surat_"+strconv.FormatUint(request.Id, 1)+".pdf")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, printdata)
}
