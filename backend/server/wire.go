// +build wireinject

package main

import (
	"backend/internal/controllers"
	"backend/internal/drivers"
	"backend/internal/repositories"
	"backend/internal/services"
	"database/sql"
	"fmt"
	"github.com/google/wire"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	WalletPath       string `yaml:"WalletPath"`
	HLConfigPath     string `yaml:"HLConfigPath"`
	HLWalletIdentity string `yaml:"HLWalletIdentity"`
	PublicKeyPath    string `yaml:"PublicKeyPath"`
	PrivateKeyPath   string `yaml:"PrivateKeyPath"`
	ChannelName      string `yaml:"ChannelName"`
	ContractName     string `yaml:"ContractName"`
	OrgMspId         string `yaml:"OrgMspId"`
	DbUser           string `yaml:"DbUser"`
	DbName           string `yaml:"DbName"`
	DbPassword       string `yaml:"DbPassword"`
	DbHost           string `yaml:"DbHost"`
	DbPort           string `yaml:"DbPort"`
	DbSslmode        string `yaml:"DbSslmode"`
}

func InitConfig() Config {
	configYaml, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(fmt.Sprintf("Could not read config.yaml file %s", err.Error()))
	}

	var config Config
	err = yaml.Unmarshal(configYaml, &config)
	if err != nil {
		panic(fmt.Sprintf("Could not parse yaml file %s", err.Error()))
	}

	return config
}

func InitializeWallet() *gateway.Wallet {
	config := InitConfig()
	wallet, err := gateway.NewFileSystemWallet(config.WalletPath)
	if err != nil {
		panic(fmt.Sprintf("Could not create wallet %s", err.Error()))
	}
	return wallet
}

func NewConfigFilePath() drivers.ConfigFilePath {
	config := InitConfig()
	return drivers.ConfigFilePath(config.HLConfigPath)
}

func NewWalletIdentity() drivers.WalletIdentity {
	config := InitConfig()
	return drivers.WalletIdentity(config.HLWalletIdentity)
}

func MakeHLClientProvider(
	iConfigFilePath drivers.ConfigFilePath,
) drivers.HLClientProvider {
	provider, err := drivers.MakeHLChannelProvider(iConfigFilePath)
	if err != nil {
		panic(fmt.Sprintf("Could not initialize HL client provider %s", err.Error()))
	}

	return provider
}

func InitializeHLClientProvider() drivers.HLClientProvider {
	wire.Build(MakeHLClientProvider, NewConfigFilePath)
	return drivers.HLClientProvider{}
}

func InitializeHLClient() *msp.Client {
	// provider := InitializeHLClientProvider()
	// client, err := provider.GetClient()
	// if err != nil {
	// 	panic(fmt.Sprintf("could not initialize client provider %s", err.Error()))
	// }
	// return client
	return nil
}

func InitializeHLIdentityService() drivers.HLIdentityService {
	wire.Build(drivers.MakeHLIdentityService, InitializeHLClient)
	return drivers.HLIdentityService{}
}

func NewMspId() drivers.MspId {
	config := InitConfig()
	return drivers.MspId(config.OrgMspId)
}

func InitializeHLGatewayInitializer() drivers.HLGatewayInitializer {
	wire.Build(drivers.MakeHLGatewayInitializer, NewMspId, InitializeWallet, NewConfigFilePath, NewWalletIdentity, InitializeHLIdentityService)
	return drivers.HLGatewayInitializer{}
}

func InitializePublicKey() services.PublicKey {
	config := InitConfig()
	data, err := os.ReadFile(config.PublicKeyPath)
	if err != nil {
		panic(err.Error())
	}

	return services.PublicKey(string(data))
}

func InitializePrivateKey() services.PrivateKey {
	config := InitConfig()
	data, err := os.ReadFile(config.PrivateKeyPath)
	if err != nil {
		panic(err.Error())
	}

	return services.PrivateKey(string(data))
}

func InitializeGraphContractSignature() services.GraphContractSignature {
	wire.Build(services.MakeGraphContractSignature, InitializePublicKey, InitializePrivateKey)
	return services.GraphContractSignature{}
}

func InitializeChannelName() drivers.ChannelName {
	config := InitConfig()
	return drivers.ChannelName(config.ChannelName)
}

func InitializeContractName() drivers.ContractName {
	config := InitConfig()
	return drivers.ContractName(config.ContractName)
}

var Set = wire.NewSet(
	drivers.MakeSmartContractDriverHL,
	wire.Bind(new(drivers.SmartContractDriverI), new(drivers.SmartContractDriverHL)),
	InitializeHLGatewayInitializer,
	InitializeChannelName,
	InitializeContractName,
)

/// Initialize repositories
var db *sql.DB = nil

func InitializeSqlDriver() *sql.DB {
	if db == nil {
		var err error
		config := InitConfig()
		connStr := ""
		connStr += " user=" + config.DbUser
		connStr += " dbname=" + config.DbName
		connStr += " password=" + config.DbPassword
		connStr += " host=" + config.DbHost
		connStr += " sslmode=" + config.DbSslmode
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			panic("could not connect to database")
		}
	}
	return db
}

var MaterialRepositorySet = wire.NewSet(
	InitializeMaterialRepositorySql,
	wire.Bind(new(repositories.MaterialRepositoryI), new(repositories.MaterialRepositorySql)),
)

func InitializeMaterialRepositorySql() repositories.MaterialRepositorySql {
	wire.Build(repositories.MakeMaterialRepositorySql, InitializeSqlDriver)
	return repositories.MaterialRepositorySql{}
}

func InitializePeerRepositorySql() repositories.PeerRepositorySql {
	wire.Build(repositories.MakePeerRepositorySql, InitializeSqlDriver)
	return repositories.PeerRepositorySql{}
}

var PeerSet = wire.NewSet(
	InitializePeerRepositorySql,
	wire.Bind(new(repositories.PeerRepositoryI), new(repositories.PeerRepositorySql)),
)

func InitializePeerController() controllers.PeersController {
	wire.Build(controllers.MakePeersController, PeerSet)
	return controllers.PeersController{}
}

func InitializeUserKeyRepositorySql() repositories.UserKeyRepositorySql {
	wire.Build(repositories.MakeUserKeyRepositorySql, InitializeSqlDriver)
	return repositories.UserKeyRepositorySql{}
}

var UserKeySet = wire.NewSet(
	InitializeUserKeyRepositorySql,
	wire.Bind(new(repositories.UserKeyRepositoryI), new(repositories.UserKeyRepositorySql)),
)

/// Finished initializing repositories

func InitializeMaterialContract() services.MaterialContract {
	wire.Build(services.MakeMaterialContract, Set, InitializeGraphContractSignature, InitializePublicKey)
	return services.MaterialContract{}
}

func InitializeMaterialRepositoryService() services.MaterialRepositoryService {
	wire.Build(services.MakeMaterialRepositoryService, MaterialRepositorySet, UserKeySet)
	return services.MaterialRepositoryService{}
}

func InitializeMaterialContractController() controllers.MaterialContractController {
	wire.Build(controllers.MakeMaterialContractController, InitializeMaterialContract, InitializeMaterialRepositoryService)
	return controllers.MaterialContractController{}
}
