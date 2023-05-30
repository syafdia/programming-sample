package types

type AivenClient = interface{}

type GCloudClient = interface{}

type BashExecutor interface {
	Run(args []string) (string, error)
}
