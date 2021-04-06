package design



// 适配器模式


type ICreateServer interface {
	CreateServer(cpu,mem float64)error
}


type QiNiuClient struct {}

func(c *QiNiuClient) RunInstance(cpu,mem float64) error{
	return nil
}

type QiNiuClientAdapter struct {
	Client QiNiuClient
}

func(q *QiNiuClientAdapter) CreateServer(cpu,mem float64) error{
	return q.Client.RunInstance(cpu,mem)
}

type AliYunClient struct {}

func(c *AliYunClient) CreateServer(cpu,mem float64) error{
	return nil
}

type AliYunClientAdapter struct {
	Client AliYunClient
}

func(a *AliYunClientAdapter) CreateServer(cpu,mem float64) error{
	return a.Client.CreateServer(cpu,mem)
}