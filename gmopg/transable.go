package gmopg

// ErrorResult is
type ErrorResult struct {
	ErrCode string `json:"errCode"`
	ErrInfo string `json:"errInfo"`
}

// ErrorResults is
type ErrorResults []ErrorResult

// Constant type definition
type (
	// Method is
	Method string
	// PayType is
	PayType string
	// Status is
	Status string
	// JobCd is
	JobCd string
	// SeqMode is
	SeqMode string
	// DefaultFlag is
	DefaultFlag string
	// CvsCode is
	CvsCode string
)

//
const (
	Lump             = Method("1")
	Installment      = Method("2")
	BonusLump        = Method("3")
	Revolving        = Method("4")
	BonusInstallment = Method("5")

	Cash        = PayType("Z")
	Credit      = PayType("0")
	Suica       = PayType("1")
	Edy         = PayType("2")
	Cvs         = PayType("3")
	DirectDebit = PayType("28")

	Unprocessed   = Status("UNPROCESSED")
	Authenticated = Status("AUTHENTICATED")
	Check         = Status("CHECK")
	Capture       = Status("CAPTURE")
	Auth          = Status("AUTH")
	Sales         = Status("SALES")
	Void          = Status("VOID")
	Return        = Status("RETURN")
	Returnx       = Status("RETURNX")
	Sauth         = Status("SAUTH")
	Reqsuccess    = Status("REQSUCCESS")
	Paysuccess    = Status("PAYSUCCESS")
	Payfail       = Status("PAYFAIL")
	Expired       = Status("EXPIRED")
	Cancel        = Status("CANCEL")

	JCheck   = JobCd("CHECK")
	JCapture = JobCd("CAPTURE")
	JAuth    = JobCd("AUTH")
	JSales   = JobCd("SALES")
	JVoid    = JobCd("VOID")
	JReturn  = JobCd("RETURN")
	JReturnx = JobCd("RETURNX")
	JSauth   = JobCd("SAUTH")
)
