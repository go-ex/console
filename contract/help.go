package contract

type Help interface {
	Configure() Configure
	Execute(input Input)
	HelpExecute(con Configure)
}
