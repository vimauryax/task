package constants

import "time"

const (
	ADMINCOLLECTION              = "admin"
	POCKETSCOLLECTION            = "pockets"
	POCKETSCOLLECTIONV2          = "pockets-v2"
	LATESTPOCKETDETAILS          = "latest-pocket-details"
	POCKETSCOLLECTIONV3          = "pockets-v3"
	USERPOCKETSCOLLECTION        = "user-pockets"
	COLLECTIONS                  = "collections"
	Finvu                        = "finvu"
	WATCHLISTSCOLLECTION         = "watchLists"
	WATCHLISTSTOCKSCOLLECTION    = "dummy"
	WATCHLISTSTOCKSCOLLECTIONNEW = "watchlistsStocks"
	PINS                         = "pins"
	IDLOGPARAM                   = "id"
	EMAILUSERIDMAPPINGCOLLECTION = "emailUserIds"
	StocksContracts              = "stockContracts"
	POCKETSTRANSACTIONS          = "pocketTransactions"
	POCKETSTRANSACTIONSDetails   = "pocketTransactionsDetails"
	POCKETSTRANSACTIONSDetailsV2 = "pocketTransactionsDetailsV2"
	ALLBANKACCOUNTS              = "bankaccount-details"
	IPODATA                      = "ipo-data"
	IPOGMPDATA                   = "ipo-gmp-data"
	TRADINGUSERS                 = "trading-users"
	CLIENTCOLLECTION             = "users"
	BANKSTAEMENTCOLLECTION       = "bankstatement-details"
	FETCHFINANCIALS              = "fetch-financials"
	UPIPREFERENCE                = "upi-preference"
	CLIENTNOTIFICATION           = "user-notification"
	HOLIDAYCALENDAR              = "holiday-calendar"
	CLIENTDETAILS                = "client-details"
	EIPOData                     = "eipo-data"
	EIPOAllData                  = "eipo-currentDataDummy"
	EIPODataCollection           = "eipo-dataCollection"
	APPDETAILS                   = "App-Details"
	AccountFreezeStatus          = "account-freeze-status"
	INSTRUMENTS_COLLECTION       = "instruments"
	USERWATCHLISTS               = "user-watchlists"
	USERWATCHLISTMAPPING         = "user-watchlist-mapping"

	EQUITY    = "equity"
	CURRENCY  = "currency"
	COMMODITY = "commodity"
	FUTOPT    = "futopt"
	DELIVERY  = "delivery"
	INTRADAY  = "intraday"
	FUTURES   = "futures"
	OPTIONS   = "options"
	SELL      = "sell"
	BUY       = "buy"
	NSE       = "nse"
	BSE       = "bse"
	NFO       = "nfo"
	BFO       = "bfo"
	CDS       = "cds"
	MCX       = "mcx"
	LIVE      = "live"
	MARKET    = "MARKET"
	DAY       = "DAY"
	CNC       = "CNC"
	MIS       = "MIS"
	REGULAR   = "REGULAR"
	IOC       = "IOC"
	CURRENT   = "CURRENT"
	CLOSED    = "CLOSED"
	UPCOMING  = "UPCOMING"
	OCO       = "oco"
	MTF       = "MTF"

	NORMALCOMMODITY     = 1
	CASTORSEEDCOMMODITY = 2
	KAPASCOMMODITY      = 3
	PEPPERCOMMODITY     = 4
	RBDPMOLEINCOMMODITY = 5

	BASKETTYPE             = "NORMAL"
	BASKETPRODUCTTYPE      = "ALL"
	ORDETTYPE              = "MARKET"
	ORDERVALIDITY          = "DAY"
	ORDERDISCLOSEDQUANTITY = 0
	ORDERPRODUCT           = "CNC"
	ORDERSERIES            = "EQ"
	DEVICETYPE             = "WEB"
	ORDEREXECUTIONTYPE     = "REGULAR"
	USERORERID             = 10002

	BASKETEXECUTETYPE    = "NML"
	BASKETSQUAREOFF      = false
	BASKETEXECUTIONSTATE = true

	BEARER  = "Bearer"
	SPACEDB = "space-db"
	TRUE    = true

	TotalRevenue     = "'Total Revenue'"
	TotalExpenses    = "'Total Expenses'"
	NetCurrentAssets = "'Net Current Asset'"
	NetProfit        = "Net Profit"
	Sector           = "same"
	Mcap             = "mcap"
	ONEYEAR          = 31536000
	EIGHTYDAYS       = 6912000
	IndexExchange    = "NSE_INDICES"
	NiftyToken       = "26000"
	DayWise          = "3"
	DataDuration     = "1"
	EXCHANGE         = "BSE"
	GUESTUSERTYPE    = "guest"
	NoPocketFound    = "No pockets found"

	TRADINGGUESTUSER = 2
	COMPLETEKYCUSER  = 1
	UNIDENTIFIEDUSER = 0

	NSESMALLCAP   = "nsesmallcap"
	NSEMIDCAP     = "nsemidcap"
	NSELARGECAP   = "nselargecap"
	SMALLMCAPTYPE = "Small Cap"
	MIDMCAPTYPE   = "Mid Cap"
	LARGEMCAPTYPE = "Large Cap"

	ADMINMESSAGE = "admin_message"
	ALERT        = "alert"
	ADMIN        = "ADMIN"
)

const (
	INFO    = 1
	ERROR   = 2
	WARNING = 3
)
const (
	ZERO           = 0
	ONE            = 1
	TWO            = 2
	THREE          = 3
	FIVE           = 5
	NineDays       = 9
	TwelveDays     = 12
	THIRTEEN       = 13
	FOURTEEN       = 14
	FIFTEEN        = 15
	FIFTY          = 50
	Hundred        = 100
	TwentyDays     = 20
	TwentySixDays  = 26
	ThirtyFiveDays = 35
	ThirtyFourDays = 34

	CandleTypeThree             = "3"
	UnixOneDaySeconds           = 86400
	IlliquidStocksMinTradingVol = 2000
)

const (
	YYYYMMDD   = "2006-01-02" // date format
	DDMMYYYY   = "02-01-2006"
	TIMEFORMAT = "2006-01-02 15:04:05"
)

var PeersFetchBy = [5]string{"Mcap", "PeTtm", "PbTtm", "EpsTtm", "RoeTtm"}

const (
	DAILYANNOUNCEMENT    = "Announcement/nse"
	BOARDMEETING         = "Board-Meetings/bse"
	CHANGEOFNAME         = "Change-of-Name/all"
	SPLITS               = "Split/all/"
	MERGER               = "Merger-Demergers/all/"
	DIVIDENDANNOUNCEMENT = "Dividend/all/"
	BULKDEALS            = "BulkBlockDeal/bse/bulk/"
	BLOCKDEALS           = "BulkBlockDeal/bse/block/"
	BONUS                = "Bonus"
)

const (
	Daily    = "daily"
	Meeting  = "meeting"
	Name     = "name"
	Splits   = "splits"
	Merger   = "merger"
	Div      = "div"
	Bulk     = "bulk"
	Block    = "block"
	Merged   = "merged"
	Demerged = "demerged"
	AGM      = "agm"
	EGM      = "egm"
	Bonus    = "bonus"
	Pledge   = "PLEDGE"
	UnPledge = "UNPLEDGE"
)

const (
	FetchFundsTypeTL  = "all"
	FetchFundsUrlTL   = "/api/v1/funds/view"
	FetchProfileUrlTL = "/api/v1/user/profile"
	FundsUpdateUrl    = "/api/v1/backoffice/funds/transactions"
	Debit             = "DEBIT"
	Payout            = "PAYOUT"
	Success           = "SUCCESS"
	Started           = "STARTED"
	CalculateWB       = "calculateWB"
	PartialPayoutKey  = "payout:"
	PayoutTime        = 16
)

const (
	CoCode          = "cocode"
	Isin            = "isin"
	DateFormat      = "2-Jan-2006"
	TwoMonthsInMins = 86400
	IPOPrefix       = "IPOPrefix"
)

const (
	ReservesAndSurplus        = "   Reserves and Surplus"
	TotalEquityAndLiabilities = "TOTAL EQUITY AND LIABILITIES"
	ShareCapital              = "Share Capital "
	TotalCurrentLiabilities   = "Total Current Liabilities"
	TotalAssets               = "TOTAL ASSETS"
	OtherLiabiities           = "Other Liabilities"
	OtherAssets               = "Other Assets"
	CFI                       = "Cash Flow from Investing Activities"
	CFF                       = "Cash Flow from Financing Activities"
	OtherInvestingItems       = "Other investing items"
	RepaymentOfBorrowings     = "Repayment of borrowings"
	OfLongTermBorrowing       = "Of the Long Tem Borrowings"
	OfShortTermBorrowing      = "Of the Short Term Borrowings"
	TotalEquity               = "Total Equity"
)

// Health Components Consts
const (
	REDIS    = "redis/cache"
	MONGO    = "mongo"
	POSTGRES = "postgres"
)

const (
	NonUserValue     = 0
	TradingUserValue = 1
	GuestUserValue   = 2
)

const (
	SearchStockKey    = "stocks"
	SegmentWiseStocks = "stocks-segment"
	StockMetadata     = "stock_metadata"
)

const (
	CurrentQuarter  = 1
	PreviousQuarter = 0
	Capacity        = 20
)

const (
	ChartDataSizeError = "number of entries in chart data not sufficient"
	CCIFactor          = 0.015
)

const (
	MongoNoDocError = "mongo: no documents in result"
	RedisReadFailed = "readMessages, Failed to read from redis"
	UnmarshalFailed = "readMessages, Packet unable to unmarshal"
	InvalidBondIsin = "invalid bond isin"
	NoDataPresent   = "No data present"
)

const (
	ISIN      = "isin"
	NSESymbol = "Nsesymbol"
	BSECode   = "bsecode"
)

const (
	LogSpace         = "log_space"
	BankIndustryCode = "00000034"
)

const (
	CLIENTID = "clientid"
	EMAILID  = "emailid"
)

const (
	Order                   = "ORDER|Orderid:"
	Trade                   = "TRADE|Orderid:"
	Pending                 = "PENDING"
	Completed               = "COMPLETED"
	Partial                 = "PARTIAL"
	EXIT                    = "EXIT"
	Complete                = "COMPLETE"
	CancelConfirmed         = "CANCEL_CONFIRMED"
	TradeCompleted          = "TRADECOMPLETED"
	Rejected                = "REJECTED"
	AbsoluteTargetPriceType = "ABSOLUTE"
	BookType                = "0"
	BrokerId                = "11378"
	UniqueKey               = "uniqueKey"
	PROCESS                 = "PROCESS"
	PROCEED                 = "PROCEED"
	CANCELLED               = "CANCELLED"
)

// UnauthorizedMap Error code Map
var ValidExchangeMap = map[string]bool{
	"NSE": true,
	"NFO": true,
	"CDS": true,
	"BSE": true,
	"BFO": true,
	"MCX": true,
}

var CommodityMap = map[string]int{
	"CASTORSEED": 2,
	"KAPAS":      3,
	"PEPPER":     4,
	"RBDPALM":    5,
}

const (
	ProfitAfterTax                   = "Profit After Tax"
	TotalNonCurrentLiabilities       = "Total Non Current Liabilities"
	NetIncDecInCashAndCashEquivalent = "Net Inc./(Dec.) in Cash and Cash Equivalent"
	TotalLiabilities                 = "Total Liabilities"
	NonBankingBalanceSheetSearch     = "'TOTAL ASSETS', 'Total Non Current Liabilities','Total Current Liabilities'"
	BankingBalanceSheetSearch        = "'TOTAL ASSETS','TOTAL EQUITY AND LIABILITIES'"
)

var BANKSBSETOKEN = map[string]bool{
	"500180": true,
	"543942": true,
	"500247": true,
	"500469": true,
	"532477": true,
	"532218": true,
	"532209": true,
	"532652": true,
	"500112": true,
	"532461": true,
	"543596": true,
	"543279": true,
	"542904": true,
	"532187": true,
	"540065": true,
	"532180": true,
	"542867": true,
	"532215": true,
	"532505": true,
	"539437": true,
	"532648": true,
	"532885": true,
	"532210": true,
	"590003": true,
	"532134": true,
	"532772": true,
	"532814": true,
	"532174": true,
	"541153": true,
	"533295": true,
	"532388": true,
	"532149": true,
	"500116": true,
	"540611": true,
	"532525": true,
	"543386": true,
	"543243": true,
	"532483": true,
}

var BANKSNSESYMBOL = map[string]bool{
	"HDFCBANK":   true,
	"UTKARSHBNK": true,
	"KOTAKBANK":  true,
	"FEDERALBNK": true,
	"UNIONBANK":  true,
	"SOUTHBANK":  true,
	"J&KBANK":    true,
	"KTKBANK":    true,
	"SBIN":       true,
	"PNB":        true,
	"TMB":        true,
	"SURYODAY":   true,
	"UJJIVANSFB": true,
	"INDUSINDBK": true,
	"RBLBANK":    true,
	"DHANBANK":   true,
	"CSBBANK":    true,
	"AXISBANK":   true,
	"UCOBANK":    true,
	"IDFCFIRSTB": true,
	"YESBANK":    true,
	"CENTRALBK":  true,
	"CUB":        true,
	"KARURVYSYA": true,
	"BANKBARODA": true,
	"DCBBANK":    true,
	"INDIANB":    true,
	"ICICIBANK":  true,
	"BANDHANBNK": true,
	"PSB":        true,
	"IOB":        true,
	"BANKINDIA":  true,
	"IDBI":       true,
	"AUBANK":     true,
	"MAHABANK":   true,
	"FINOPB":     true,
	"EQUITASBNK": true,
	"CANBK":      true,
}

var BANKSISIN = map[string]bool{
	"INE040A01034": true,
	"INE735W01017": true,
	"INE237A01028": true,
	"INE171A01029": true,
	"INE692A01016": true,
	"INE683A01023": true,
	"INE168A01041": true,
	"INE614B01018": true,
	"INE062A01020": true,
	"INE160A01022": true,
	"INE668A01016": true,
	"INE428Q01011": true,
	"INE551W01018": true,
	"INE095A01012": true,
	"INE976G01028": true,
	"INE680A01011": true,
	"INE679A01013": true,
	"INE238A01034": true,
	"INE691A01018": true,
	"INE092T01019": true,
	"INE528G01035": true,
	"INE483A01010": true,
	"INE491A01021": true,
	"INE036D01028": true,
	"INE028A01039": true,
	"INE503A01015": true,
	"INE562A01011": true,
	"INE090A01021": true,
	"INE545U01014": true,
	"INE608A01012": true,
	"INE565A01014": true,
	"INE084A01016": true,
	"INE008A01015": true,
	"INE949L01017": true,
	"INE457A01014": true,
	"INE02NC01014": true,
	"INE063P01018": true,
	"INE476A01014": true,
}

const (
	PostgresPenFilePath    = "./resources/DigiCertGlobalRootCA.crt.pem"
	PostgresDevPenFilePath = "./resources/DevDigiCertGlobalRootCA.crt.pem"
)

const (
	Production = "prod"
)

const (
	NFO_           = "NFO_"
	Call_          = "_CE"
	Put_           = "_PE"
	DateTimeFormat = "02Jan06"
)

const (
	TradeConfirmationDateRange   = "TradeConfirmationDateRange"
	GetBillDetailsCdsl           = "GetBillDetailsCdsl"
	ScripWiseCosting             = "ScripWiseCosting"
	GetLongTermShortTerm         = "GetLongTermShortTerm"
	Profile                      = "Profile"
	TradeConfirmation            = "TradeConfirmation"
	OpenPosition                 = "OpenPosition"
	GetHolding                   = "GetHolding"
	GetMarginOnDate              = "GetMarginOnDate"
	FinancialLedgerBalanceOnDate = "FinancialLedgerBalanceOnDate"
	GetFinancial                 = "GetFinancial"
	GetFinancialDateRange        = "GetFinancialDateRange"
	NetPositionFO                = "NetPositionfo"
	HoldingCumFinancial          = "HoldingCumFinancial"
	GetCommodityTransaction      = "GetComTransaction"
	GetFNOTransaction            = "GetFoTransaction"
	GetEquityTransaction         = "GetCapTransaction"
)

const (
	ServiceVersion        = "1.0.0"
	ElasticApmServiceName = "ELASTIC_APM_SERVICE_NAME"
	ElasticApmEnvironment = "ELASTIC_APM_ENVIRONMENT"
	ElasticApmServerURL   = "ELASTIC_APM_SERVER_URL"
)

const (
	FundsTypeAll   = "all"
	OpeningBalance = "Opening Balance"
	MarginUsed     = "Margin Used"
	Historical     = "historical"
	// Completed      = "completed"
	Segment          = "equity"
	Payin            = "Pay In"
	EquityCreditSell = "Equity Credit Sell"
	AdhocDeposit     = "Adhoc Deposit"
	MtfMarginBlocked = "MTF Margin Blocked"
)

const (
	BondsDataCollection = "bonds-details"
)

const (
	FinvuConsentDescription     = "Wealth Management Service"
	FinvuTemplateName           = "FINVUDEMO_TESTING"
	FinvuUserSessionID          = "sessionid123"
	FinvuDataConsumerId         = "fiu@dhanaprayoga"
	FinvuPurposeCode            = "103"
	FinvuPurposeText            = "Aggregated statement"
	FinvuCategoryType           = "Financial Reporting"
	FinvuConsentMode            = "STORE"
	FinvuFetchType              = "PERIODIC"
	FinvuFrequencyValue         = 30
	FinvuFrequencyUnit          = "DAY"
	FinvuDataLifeValue          = 1
	FinvuDataLifeUnit           = "YEAR"
	FinvuConsentExpiry          = "2025-04-10T12:59:22+0000"
	FinvuAaID                   = "cookiejar-aa@finvu.in"
	FinvuAuthTokenKey           = "Finvu_Auth_Token"
	FinvuConsentRequestPlus     = "ConsentRequestPlus"
	FinvuConsentStatusRequested = "REQUESTED"
	FinvuConsentStatus          = "ConsentStatus"
	FinvuConsentStatusACCEPTED  = "ACCEPTED"
	FinvuFiRequestStatusREADY   = "READY"
	FinvuFIDataFetch            = "FIDataFetch"
	BankStatement               = "BANKSTATEMENT"
	FinvuFIStatus               = "FIStatus"
	FinvuFIRequest              = "FIRequest"
	FinvuConsent                = "Consent"
	New                         = "NEW"

	FinvuInvalidTokenMessage = "Invalid token. User does not have permission"
	FinvuInternalServerError = "Internal Server Error"
)

var FinvuFip = []string{"BARB0KIMXXX"}
var FinvuConsentTypes = []string{"PROFILE", "SUMMARY", "TRANSACTIONS"}
var FinvuFiTypes = []string{"DEPOSIT", "RECURRING_DEPOSIT", "TERM-DEPOSIT", "EQUITIES"}
var FinvuConsentStart = time.Date(2023, 4, 10, 12, 59, 22, 403, time.UTC)
var FinvuFIDataRangeFrom = time.Date(2022, 4, 10, 12, 59, 22, 403, time.UTC)
var FinvuFIDataRangeTo = time.Date(2025, 4, 10, 12, 59, 22, 403, time.UTC)

const (
	LogEdis           = "log_edis"
	PledgeTxnPageSize = 20
)

const (
	IPO_KEY    = "ipo-key"
	IPO_EXPIRY = 180
)

const (
	FAILED  = "failed"
	SUCCESS = "success"
)

const (
	TypeCapsStandalone   = "'S'"
	TypeSmallStandalone  = "'s'"
	TypeCapsConsolidate  = "'C'"
	TypeSmallConsolidate = "'c'"
)

const (
	ShilpiDateFormat         = "2-Jan-2006"
	ShilpiDateFormatWithTime = "2006-01-02 15:04:05.0"
	DateWithTimeFormat       = "02-01-2006 15:04:05"
)

const (
	FundPayment  = "Fund Payment"
	FundReceived = "Fund Received"
)

const (
	FF  = "FF"
	FX  = "FX"
	CE  = "CE"
	PE  = "PE"
	FNO = "FNO"
	CUR = "CUR"
	COM = "COM"
	FUT = "fut"
)

const (
	AwsS3CredConfig   = ".awsS3CredConfig"
	ReportsFolderName = ".reportsFolderName"
)

const (
	DpChargesReport          = "Dp_Charges_Report_"
	LedgerReport             = "Ledger_Report_"
	TradebookReport          = "Tradebook_Report_"
	OpenPositionReport       = "Open_Position_Report_"
	FnoPnlReport             = "Fno_Pnl_Report_"
	HoldingFinancialReport   = "Holding_Financial_Report_"
	CommodityTradebookReport = "Commodity_Tradebook_Report_"
	FnoTradebookReport       = "Fno_Tradebook_Report_"
	EquityTradebookReport    = "Equity_Tradebook_Report_"
)

const (
	LocalEnv = "local"
)

const (
	CALL = "CALL"
	PUT  = "PUT"
)

const (
	ReportData = "ReportData"
	ReportFile = "ReportFile"
)

const (
	ReportsCachingTTL = 360 //6 hours
	MainRedis         = "mainRedis"
	OrderRedis        = "orderRedis"
	Redis             = "redis"
)

const (
	NoRowPG = "sql: no rows in result set"
)

const (
	SegmentEquity          = "Equity"
	SegmentIndex           = "indices"
	SegmentCommodity       = "Commodities"
	SegmentIndices         = "INDICES"
	Commodity              = "Commodity"
	FutnOpt                = "F&O"
	EquityTradebook        = "Equity Tradebook"
	EquityTradebookSegment = "CAPITAL"
)

const (
	MongoQueryParam = "retryWrites=true&w=majority&maxIdleTimeMS=60000&connectTimeoutMS=150000&socketTimeoutMS=90000"
)

const (
	PinsSize = 6
)

var LatencyThresholdHigh int64
var LatencyThresholdLow int64

const (
	Topic                       = "topic"
	TopicExchange               = "topic_exchange"
	PKTFLKyc1                   = "PKTFLKyc1"
	KeyLedgerReport             = "PKTFLLegder"
	KeyCommodityTradebookReport = "PKTFLCommodityTradebook"
	KeyFnoTradebookReport       = "PKTFLFnoTradebook"
	KeyDpChargesReport          = "PKTFLDpCharges"
	KeyHoldingFinancialReport   = "PKTFLHoldingFinancial"
	KeyTradebookEquityReport    = "PKTFLTradebookEquityReport"
)

const (
	TokenSuccess   = "Access Token Found"
	TokenFailure   = "Access Token Not Found"
	SessionKey     = "space:session_id:"
	UserSessionKey = "space:user_session_id:"
)

var WatchlistIds = map[string]bool{
	"wl1": true,
	"wl2": true,
	"wl3": true,
	"wl4": true,
	"wl5": true,
}

const (
	WatchlistSize       = 20
	WatchlistStocksSize = 50
)

const (
	GetALLFromHashDelayTime  = 30
	PocketfulLoginSessionTTL = 24 * time.Hour
)

const (
	FreshdeskTicketURL = "tickets"
)

const (
	Global = "global"
)

const (
	CommodityTradebookReportName = "Tradebook & Charges for Commodity"
	FnoTradebookReportName       = "Tradebook & Charges for F&O"
	HoldingReportName            = "Holding Statement"
	DPChargesReportName          = "DP Charges Statement"
	EquityTradebookReportName    = "Tradebook & Charges for Equity"
)

var Env string

const (
	PocketOrderSleepSeconds = 3
)

const (
	OTPSENDER              = "PKTFUL"
	AccountFreeze          = "ACCOUNT_FREEZE_"
	OtpSetTimeRedis        = 15
	OtpLen                 = 4
	WaitBeforeReRequest    = 30
	LoginChangeOTP         = "LoginChangeOTP_"
	OTPVerifiedTill        = "OTPVerifiedTill_"
	LoginChangeOTPReqLimit = "LoginChangeOTPReqLimit|"
)

var LocationKolkata *time.Location

const (
	ASIAKOLKATA = "Asia/Kolkata"
)

// hotfix 1.0.157

const (
	Rebalance = "rebalance"
	Repair    = "repair"
	NONE      = "none"
)

const (
	Charts = "/api/v1/charts"
)

//secret update

const (
	SubscriptionKey = "subscription_ids"
)

const (
	ClientMembers = "client_ids:"
)

const (
	ACCOUNTDETAILSUPDATE = "account-details-update"
)

//1.0.112

const (
	NSETL = 1
	BSETL = 6
)

const (
	Active = "Active"
	Paused = "Paused"
)

var SectorMapping = map[string][]string{
	"Banks":               {"Banks"},
	"IT":                  {"IT - Hardware", "IT - Software"},
	"Finance":             {"Finance", "Financial Services", "Insurance", "Credit Rating Agencies"},
	"Automobile":          {"Automobile", "Auto Ancillaries"},
	"Healthcare":          {"Healthcare"},
	"FMCG":                {"FMCG"},
	"Chemicals":           {"Chemicals"},
	"Construction":        {"Construction"},
	"Consumer Durables":   {"Consumer Durables"},
	"Agriculture":         {"Agro Chemicals", "Fertilizers"},
	"Alcohol":             {"Alcoholic Beverages"},
	"Logistics":           {"Logistics"},
	"Entertainment":       {"Entertainment", "Media - Print/Television/Radio"},
	"Aerospace & Defence": {"Aerospace & Defence"},
	"Electronics":         {"Electronics"},
	"Engineering":         {"Engineering"},
	"Hospitality":         {"Hotels & Restaurants"},
	"Metals":              {"Ferro Alloys", "Mining & Mineral products", "Non Ferrous Metals"},
	"Oil & Gas":           {"Crude Oil & Natural Gas", "Gas Distribution"},
	"Capital Goods":       {"Capital Goods - Electrical Equipment", "Capital Goods-Non Electrical Equipment"},
	"Cement":              {"Cement", "Cement - Products"},
	"Air Transport":       {"Air Transport Service"},
	"Cables":              {"Cables"},
	"Cast & Forge":        {"Castings, Forgings & Fastners"},
	"Ceramic Products":    {"Ceramic Products"},
	"Computer Education":  {"Computer Education"},
	"Jewellery":           {"Diamond, Gems and Jewellery"},
	"Diversified":         {"Diversified"},
	"Dry cells":           {"Dry cells"},
	"Edible Oil":          {"Edible Oil"},
	"Glass & Products":    {"Glass & Glass Products"},
	"Infrastructure":      {"Infrastructure Developers & Operators"},
	"Leather":             {"Leather"},
	"Miscellaneous":       {"Miscellaneous"},
	"Bearings":            {"Bearings"},
	"Education":           {"Education"},
	"Platforms":           {"E-Commerce/App based Aggregator"},
	"InvITs":              {"Infrastructure Investment Trusts"},
	"Marine Infra":        {"Marine Port & Services"},
}

type Index struct {
	Name     string
	Token    string
	SectName string
}

var SectorToIndices = map[string]struct {
	NSEIndices []Index
	BSEIndices []Index
}{
	"00000006": {
		NSEIndices: []Index{
			{Name: "Nifty Bank", Token: "26009", SectName: "Banks"},
			{Name: "Nifty PSU Bank", Token: "26025", SectName: "Banks"},
			{Name: "Nifty Pvt Bank", Token: "26047", SectName: "Banks"},
		},
		BSEIndices: []Index{
			{Name: "BANKEX", Token: "12", SectName: "Banks"},
		},
	},
	"00000033": {
		NSEIndices: []Index{
			{Name: "Nifty IT", Token: "26008", SectName: "IT"},
		},
		BSEIndices: []Index{},
	},
	"00000034": {
		NSEIndices: []Index{},
		BSEIndices: []Index{
			{Name: "BSE Information Technology", Token: "34", SectName: "IT"},
		},
	},
	"00000026": {
		NSEIndices: []Index{
			{Name: "Nifty Financial Services Index", Token: "26037", SectName: "Finance"},
		},
		BSEIndices: []Index{},
	},
	"00000087": {
		NSEIndices: []Index{
			{Name: "Nifty Financial Services 25/50 Index", Token: "26065", SectName: "Finance"},
		},
		BSEIndices: []Index{},
	},
	"00000067": {
		NSEIndices: []Index{
			{
				Name: "Nifty Financial Services Ex-Bank Index", Token: "26118", SectName: "Finance"},
		},
		BSEIndices: []Index{},
	},
	"00000004": {
		NSEIndices: []Index{
			{Name: "Nifty Auto Index", Token: "26029", SectName: "Automobile"},
		},
		BSEIndices: []Index{
			{Name: "BSE AUTO", Token: "13", SectName: "Automobile"},
		},
	},
	"00000005": {
		NSEIndices: []Index{
			{Name: "Nifty EV", Token: "26101", SectName: "Automobile"},
		},
		BSEIndices: []Index{},
	},
	"00000030": {
		NSEIndices: []Index{
			{Name: "Nifty Healthcare Index", Token: "26069", SectName: "Healthcare"},
			{Name: "Nifty MidSmall Healthcare Index", Token: "26084", SectName: "Healthcare"},
		},
		BSEIndices: []Index{},
	},
	"00000019": {
		NSEIndices: []Index{
			{Name: "Nifty Oil and Gas Index", Token: "26071", SectName: "Oil & Gas"},
		},
		BSEIndices: []Index{
			{Name: "BSE OIL & GAS", Token: "15", SectName: "Oil & Gas"},
		},
	},

	"00000040": {
		NSEIndices: []Index{
			{Name: "Nifty Metal Index", Token: "26030", SectName: "Metals"},
		},
		BSEIndices: []Index{
			{Name: "BSE METAL", Token: "14", SectName: "Metals"},
		},
	},

	"00000027": {
		NSEIndices: []Index{
			{Name: "Nifty FMCG Index", Token: "26021", SectName: "FMCG"},
		},
		BSEIndices: []Index{
			{Name: "BSE Fast Moving Consumer Goods", Token: "25252", SectName: "FMCG"},
		},
	},
	"00000037": {
		NSEIndices: []Index{
			{Name: "Nifty Media Index", Token: "26031", SectName: "Entertainment"},
		},
		BSEIndices: []Index{},
	},
}

const (
	Period1D = "1D"
	Period1W = "1W"
	Period1M = "1M"
	Period6M = "6M"
	Period1Y = "1Y"
	Period3Y = "3Y"
	Period5Y = "5Y"
)

var ChartDataPeriod = []string{"1D", "1W", "1M", "6M", "1Y", "3Y", "5Y"}

// Search API Constants
const (
	ContractSearchPageLimit = 20 // Default number of results per page for contract search API
)

const (
	DefaultLimitToFetchShareholdingAggregate = 5 // default number of results for shareholding pattern api
)

const (
	GSecFBilSource   = "GSecBonds-FBIL"
	StripsFbilSource = "GSecStripsBonds-FBIL"
	SDLFbilSource    = "SDLBonds-Fbil"
	CPBondsSource    = "NSDL entire list of commercial papers"
	DebtBondSource   = "NSDL entire list of debt instruments"
	CcilBondSource   = "CCILBonds-Tbills"
	SecuritySource   = "List of Securities Report"
	Taxation         = "Taxfree"
)

const (
	StateGov  = "STATE-GOV"
	TBill     = "T-BILL"
	TaxFree   = "TAX-FREE"
	Gov       = "GOV"
	Corporate = "CORPORATE"
)

const (
	NewsLimit = 20
)

const (
	ActiveStatus   = "active"
	UpcomingStatus = "upcoming"
	ClosedStatus   = "closed"
	ListedStatus   = "listed"
)

const (
	ClosedIpo = "closedIPO"
	ActiveIpo = "activeIPO"
)

const (
	AlphaNumericCharacter              = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	IpoNseTransactions                 = "ipoNseTransactions"
	Demat                              = "demat"
	IPO_UPI_FLAG                       = "Y"
	ACTIVITY_NEW                       = "new"
	ACTIVITY_MODIFY                    = "modify"
	ACTIVITY_CANCEL                    = "cancel"
	IPO_State_APPLIED                  = "applied"
	IPO_Payment_Status_Mandate_Pending = "mandate-pending"
)

const (
	NseIpoToken          = "NseIpoToken"
	NseIpoTokenCacheTime = 55
)

var Sources = []string{
	"Dividend",
	"Bonus",
	"Rights",
	"AGM",
	"EGM",
	"Split",
	"Merger-Demergers",
	"Buyback",
}

const (
	InternalServerError = "500"
	BadRequest          = "400"
)

// Errors Code Map
var ErrorCodeMap = map[string]string{
	"P11001": "INTERNAL SERVER ERROR",
	"P11002": "INVALID PARAMETERS",
	"P11003": "METHOD NOT ALLOWED",
	"P11004": "INVALID IP",
	"P11006": "YOUR USER ID IS INVALID",
	"P11007": "USER IS BLOCKED",
	"P11008": "USER IS BLOCKED FOR TRADING",
	"P11010": "MOBILE OR EMAIL ALREADY EXISTS",
	"P11011": "INVALID USERID",
	"P11012": "INVALID TOKEN",
	"P11013": "TOKEN MISSING",
	"P11014": "ID DOES NOT EXISTS",
	"P11015": "TIP DOES NOT EXISTS",
	"P11016": "ID ALREADY EXISTS",
	"P11017": "INVALID REQUEST",
	"P11018": "INVALID FRONT PART",
	"P11019": "DETAILS DOES NOT EXISTS",
	"P11020": "ADMIN INVALID CREDS",
	"P11021": "USER VIDEO IS INVALID",
	"P11022": "USER SELFIE IS INVALID",
	"P11030": "INVALID MOBILE NO",
	"P11031": "INVALID EMAIL ID",
	"P11032": "INVALID EMAIL OTP",
	"P11033": "INVALID MOBILE OTP",
	"P11034": "INVALID DEVICE TYPE",
	"P11040": "INVALID TYPE",
	"P11041": "INVALID CLIENT",
	"P11042": "POCKET ALREADY EXISTS",
	"P11043": "POCKET DOES NOT EXISTS",
	"P11044": "INVALID USERID OR PASSWORD",
	"P11045": "COLLECTION ALREADY EXISTS",
	"P11046": "COLLECTION DOES NOT EXISTS",
	"P11047": "WATCHLISTS ALREADY EXISTS",
	"P11048": "WATCHLISTS DOES NOT EXISTS",
	"P11049": "PIN DOES NOT EXISTS",
	"P11050": "NO More Pins can be added! Capacity full!",
	"P11051": "Failed to fetch chart Data from TradeLab chart data api!",
	"P11052": "AUTHENTICATION FAILED",
	"P11053": "ISIN IS EMPTY",
	"P11054": "WatchLists Stock Already Exists",
	"P11055": "Invalid Exchange",
	"P11056": "Decoding Header Error",
	"P11057": "Empty Profile Response",
	"P11058": "Empty Ip Address",
	"P11059": "Invalid Device ID",
	"P11060": "Token Expired",
	"P11061": "Upi Don't Exist",
	"P11062": "Duplicate Upi",
	"P11063": "Invalid Date",
	"P11064": "Invalid Isin",
	"P11065": "Mismatch Auth Client",
	"P11066": "Invalid Display name",
	"P11067": "Your request is already proceed, try to request another payout. You can't cancel a payout request after it has started processing",
	"P11068": "Pins size exceeds the available capacity",
	"P11069": "Invalid Pin Index",
	"P11070": "Name of Basket must be less than 30 characters",
	"P11071": "You cannot place another payout request while a previous request is processing.",
	"P11072": "empty client id and email id",
	"P11073": "Invalid Otp",
	"P11074": "Account Freeze Invalid Request",
	"P11075": "Invalid Header",
	"P11076": "App Does Not Exists",
	"P11077": "Lot Size exceeds the available quantity",
	"P11078": "Invalid Page",
	"P11079": "No Data Found",
	"P11080": "Please wait atleast 30 seconds before requesting for otp again!",
	"P11081": "OTP for LoginV4  credentials like Password, Mpin, change Not Verified! Please Verify OTP first!",
	"P11082": "Action Required!, Update your pocket from your portfolio",
	"P11083": "Payout not permitted between 15:45 and 16:00 IST. Please try after 16:00.",
}
