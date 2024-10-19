package enum

// DebtType represents the type of debt
type DebtType string

// Debt types
const (
	DebtTypeMortgage DebtType = "mortgage"
	DebtTypeCredit   DebtType = "credit"
	DebtTypeLoan     DebtType = "loan"
	DebtTypeOther    DebtType = "other"
)
