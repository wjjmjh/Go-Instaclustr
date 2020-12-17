package src

type Cluster struct {
	ID                         string       `json:"id"`
	ClusterName                string       `json:"clusterName"`
	ClusterStatus              string       `json:"clusterStatus"`
	CassandraVersion           string       `json:"cassandraVersion"`
	Username                   string       `json:"username"`
	InstaclustrUserPassword    string       `json:"instaclustrUserPassword"`
	SlaTier                    string       `json:"slaTier"`
	ClusterCertificateDownload string       `json:"clusterCertificateDownload"`
	PciCompliance              string       `json:"pciCompliance"`
	DataCentres                []DataCentre `json:"dataCentres"`
}

type DataCentre struct {
	ID                            string   `json:"id"`
	Name                          string   `json:"name"`
	Provider                      string   `json:"provider"`
	CdcNetwork                    string   `json:"cdcNetwork"`
	Bundles                       []string `json:"bundles"`
	ClientEncryption              bool     `json:"clientEncryption"`
	PasswordAuthentication        bool     `json:"passwordAuthentication"`
	UserAuthorization             bool     `json:"userAuthorization"`
	UsePrivateBroadcastRPCAddress bool     `json:"usePrivateBroadcastRPCAddress"`
	PrivateIPOnly                 bool     `json:"privateIPOnly"`
	Nodes                         []Node   `json:"nodes"`
	NodeCount                     int      `json:"nodeCount"`
	EncryptionKeyId               []string `json:"encryptionKeyId"`
	ResizeTargetNodeSize          string   `json:"resizeTargetNodeSize"`
}

type Node struct {
	ID             string   `json:"id"`
	Size           string   `json:"size"`
	Rack           string   `json:"rack"`
	PublicAddress  []string `json:"publicAddress"`
	PrivateAddress []string `json:"privateAddress"`
	NodeStatus     string   `json:"nodeStatus"`
	SparkMaster    bool     `json:"sparkMaster"`
	SparkJobserver bool     `json:"sparkJobserver"`
	Zeppelin       bool     `json:"zeppelin"`
}
