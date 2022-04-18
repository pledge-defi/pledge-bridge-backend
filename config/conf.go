package config

var Config *Conf

type Conf struct {
	Mysql     MysqlConfig
	Redis     RedisConfig
	Jwt       JwtConfig
	Env       EnvConfig
	Task      TaskConfig
	Email     EmailConfig
	BscNet    BscNetConfig
	EthNet    EthNetConfig
	Threshold ThresholdConfig
	Contract  ContractConfig
}

type BscNetConfig struct {
	BridgeToken string `toml:"bridge_token"`
	NetUrl      string `toml:"net_url"`
	ChainId     int64  `toml:"chain_id"`
}

type EthNetConfig struct {
	BridgeToken string `toml:"bridge_token"`
	NetUrl      string `toml:"net_url"`
	ChainId     int64  `toml:"chain_id"`
}

type TaskConfig struct {
	Duration uint64 `toml:"duration"`
}

type EnvConfig struct {
	Port          string `toml:"port"`
	IsDev         bool   `toml:"is_dev"`
	ReleaseDurDev int64  `toml:"release_dur_dev"`
}

type JwtConfig struct {
	SecretKey  string `toml:"secret_key"`
	ExpireTime int    `toml:"expire_time"` // duration, s
}

type MysqlConfig struct {
	Address      string `toml:"address"`
	Port         string `toml:"port"`
	DbName       string `toml:"db_name"`
	UserName     string `toml:"user_name"`
	Password     string `toml:"password"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	MaxLifeTime  int    `toml:"max_life_time"`
}

type RedisConfig struct {
	Address     string `toml:"address"`
	Port        string `toml:"port"`
	Db          string `toml:"db"`
	Password    string `toml:"password"`
	MaxIdle     int    `toml:"max_idle"`
	MaxActive   int    `toml:"max_active"`
	IdleTimeout int    `toml:"idle_timeout"`
}

type EmailConfig struct {
	Username string   `toml:"username"`
	Pwd      string   `toml:"pwd"`
	Host     string   `toml:"host"`
	Port     string   `toml:"port"`
	From     string   `toml:"from"`
	Subject  string   `toml:"subject"`
	To       []string `toml:"to"`
	Cc       []string `toml:"cc"`
}

type ContractConfig struct {
	PlgrToken                    string `toml:"plgr_token"`
	MplgrToken                   string `toml:"mplgr_token"`
	ContractAdminTokenPrivateKey string `toml:"contract_admin_token_private_key"`
	BridgeTokenOne               string `toml:"bridge_token_one"`
	BridgeTokenTwo               string `toml:"bridge_token_two"`
	BridgeTokenThree             string `toml:"bridge_token_three"`
}

type ThresholdConfig struct {
	BridgeBsc           string `toml:"bridge_bsc"`
	BridgeEth           string `toml:"bridge_eth"`
	BridgeTokenOneBnb   string `toml:"bridge_token_one_bnb"`
	BridgeTokenOneEth   string `toml:"bridge_token_one_eth"`
	BridgeTokenTwoBnb   string `toml:"bridge_token_two_bnb"`
	BridgeTokenTwoEth   string `toml:"bridge_token_two_eth"`
	BridgeTokenThreeBnb string `toml:"bridge_token_three_bnb"`
	BridgeTokenThreeEth string `toml:"bridge_token_three_eth"`
}
