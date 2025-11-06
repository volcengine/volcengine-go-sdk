package endpoints

import (
	"bytes"
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
	"regexp"
	"strings"
	"text/template"
)

const (
	defaultFormat = `{{.Service}}{{.Region}}.{{.SiteStack}}.com`
)

type StandardEndpointResolverVariable struct {
	Service   string
	Region    string
	SiteStack string
	Extension map[string]interface{}
}

type StandardEndpointResolver struct {
	Format         string
	Variables      StandardEndpointResolverVariable
	CustomServices map[string]ServiceInfo
}

func NewStandardEndpointResolver() *StandardEndpointResolver {
	return &StandardEndpointResolver{
		Format: defaultFormat,
		Variables: StandardEndpointResolverVariable{
			Service:   "",
			Region:    "",
			SiteStack: SiteStackVolcIPv4,
			Extension: nil,
		},
	}
}

type StandardEndpointOptions struct {
	Format         string
	SiteStack      string
	Extension      map[string]interface{}
	CustomServices map[string]ServiceInfo
}

func NewStandardEndpointResolverWithOptions(options *StandardEndpointOptions) *StandardEndpointResolver {
	if options == nil {
		return NewStandardEndpointResolver()
	}
	return &StandardEndpointResolver{
		Format: options.Format,
		Variables: StandardEndpointResolverVariable{
			Service:   "",
			Region:    "",
			SiteStack: options.SiteStack,
			Extension: options.Extension,
		},
		CustomServices: options.CustomServices,
	}
}

type ServiceInfo struct {
	Service  string
	IsGlobal bool
}

var ServiceInfos = map[string]ServiceInfo{
	"vpc":                         {Service: "vpc", IsGlobal: false},
	"ecs":                         {Service: "ecs", IsGlobal: false},
	"billing":                     {Service: "billing", IsGlobal: true},
	"ark":                         {Service: "ark", IsGlobal: false},
	"iam":                         {Service: "iam", IsGlobal: true},
	"mcs":                         {Service: "mcs", IsGlobal: false},
	"rocketmq":                    {Service: "rocketmq", IsGlobal: false},
	"bytehouse":                   {Service: "bytehouse", IsGlobal: false},
	"dns":                         {Service: "dns", IsGlobal: true},
	"autoscaling":                 {Service: "autoscaling", IsGlobal: false},
	"spark":                       {Service: "spark", IsGlobal: false},
	"cloud_detect":                {Service: "cloud_detect", IsGlobal: false},
	"filenas":                     {Service: "filenas", IsGlobal: false},
	"escloud":                     {Service: "escloud", IsGlobal: false},
	"flink":                       {Service: "flink", IsGlobal: false},
	"cp":                          {Service: "cp", IsGlobal: false},
	"vefaas":                      {Service: "vefaas", IsGlobal: false},
	"ml_platform":                 {Service: "ml_platform", IsGlobal: false},
	"edx":                         {Service: "edx", IsGlobal: true},
	"dcdn":                        {Service: "dcdn", IsGlobal: true},
	"cdn":                         {Service: "cdn", IsGlobal: true},
	"kafka":                       {Service: "kafka", IsGlobal: false},
	"certificate_service":         {Service: "certificate_service", IsGlobal: true},
	"waf":                         {Service: "waf", IsGlobal: true},
	"rds_mssql":                   {Service: "rds_mssql", IsGlobal: false},
	"cloudtrail":                  {Service: "cloudtrail", IsGlobal: false},
	"vei_api":                     {Service: "vei_api", IsGlobal: true},
	"cen":                         {Service: "cen", IsGlobal: true},
	"rabbitmq":                    {Service: "rabbitmq", IsGlobal: false},
	"vmp":                         {Service: "vmp", IsGlobal: false},
	"volc_observe":                {Service: "volc_observe", IsGlobal: false},
	"dataleap":                    {Service: "dataleap", IsGlobal: false},
	"fw_center":                   {Service: "fw_center", IsGlobal: true},
	"redis":                       {Service: "redis", IsGlobal: false},
	"mcdn":                        {Service: "mcdn", IsGlobal: true},
	"cloudidentity":               {Service: "cloudidentity", IsGlobal: false},
	"vedbm":                       {Service: "vedbm", IsGlobal: false},
	"cv":                          {Service: "cv", IsGlobal: true},
	"translate":                   {Service: "translate", IsGlobal: true},
	"cloud_trail":                 {Service: "cloud_trail", IsGlobal: false},
	"bio":                         {Service: "bio", IsGlobal: false},
	"nta":                         {Service: "nta", IsGlobal: true},
	"elasticmapreduce":            {Service: "elasticmapreduce", IsGlobal: false},
	"vepfs":                       {Service: "vepfs", IsGlobal: false},
	"seccenter":                   {Service: "seccenter", IsGlobal: true},
	"advdefence":                  {Service: "advdefence", IsGlobal: true},
	"tis":                         {Service: "tis", IsGlobal: true},
	"organization":                {Service: "organization", IsGlobal: true},
	"vke":                         {Service: "vke", IsGlobal: false},
	"Redis":                       {Service: "Redis", IsGlobal: false},
	"privatelink":                 {Service: "privatelink", IsGlobal: false},
	"RocketMQ":                    {Service: "RocketMQ", IsGlobal: false},
	"Kafka":                       {Service: "Kafka", IsGlobal: false},
	"rds_mysql":                   {Service: "rds_mysql", IsGlobal: false},
	"rds_postgresql":              {Service: "rds_postgresql", IsGlobal: false},
	"storage_ebs":                 {Service: "storage_ebs", IsGlobal: false},
	"clb":                         {Service: "clb", IsGlobal: false},
	"alb":                         {Service: "alb", IsGlobal: false},
	"FileNAS":                     {Service: "FileNAS", IsGlobal: false},
	"configcenter":                {Service: "configcenter", IsGlobal: false},
	"cr":                          {Service: "cr", IsGlobal: false},
	"sts":                         {Service: "sts", IsGlobal: false},
	"mongodb":                     {Service: "mongodb", IsGlobal: false},
	"transitrouter":               {Service: "transitrouter", IsGlobal: false},
	"Volc_Observe":                {Service: "Volc_Observe", IsGlobal: false},
	"dms":                         {Service: "dms", IsGlobal: false},
	"auto_scaling":                {Service: "auto_scaling", IsGlobal: false},
	"directconnect":               {Service: "directconnect", IsGlobal: false},
	"kms":                         {Service: "kms", IsGlobal: false},
	"dbw":                         {Service: "dbw", IsGlobal: false},
	"dts":                         {Service: "dts", IsGlobal: false},
	"natgateway":                  {Service: "natgateway", IsGlobal: false},
	"tos":                         {Service: "tos", IsGlobal: false},
	"TLS":                         {Service: "TLS", IsGlobal: false},
	"vpn":                         {Service: "vpn", IsGlobal: false},
	"vod":                         {Service: "vod", IsGlobal: false},
	"quota":                       {Service: "quota", IsGlobal: true},
	"ecs_ops":                     {Service: "ecs_ops", IsGlobal: true},
	"as_ops":                      {Service: "as_ops", IsGlobal: true},
	"account_management":          {Service: "account_management", IsGlobal: true},
	"account_management_byteplus": {Service: "account_management_byteplus", IsGlobal: true},
	"bandwidthquota":              {Service: "bandwidthquota", IsGlobal: true},
	"psa_manager":                 {Service: "psa_manager", IsGlobal: true},
	"dc_controller":               {Service: "dc_controller", IsGlobal: false},
	"eps_platform_trade":          {Service: "eps_platform_trade", IsGlobal: false},
	"eps_platform_fund":           {Service: "eps_platform_fund", IsGlobal: false},
	"commercialization":           {Service: "commercialization", IsGlobal: true},
	"veecp_openapi":               {Service: "veecp_openapi", IsGlobal: false},
	"orgnization":                 {Service: "orgnization", IsGlobal: true},
	"coze":                        {Service: "coze", IsGlobal: true},
	"sec_agent":                   {Service: "sec_agent", IsGlobal: true},
	"sec_intelligent_dev":         {Service: "sec_intelligent_dev", IsGlobal: true},
	"vegame":                      {Service: "vegame", IsGlobal: false},
	"acep":                        {Service: "acep", IsGlobal: true},
	"private_zone":                {Service: "private_zone", IsGlobal: true},
	"sqs":                         {Service: "sqs", IsGlobal: false},
	"resourcecenter":              {Service: "resourcecenter", IsGlobal: true},
	"aiotvideo":                   {Service: "aiotvideo", IsGlobal: true},
	"apig":                        {Service: "apig", IsGlobal: false},
	"bmq":                         {Service: "bmq", IsGlobal: false},
	"bytehouse_ce":                {Service: "bytehouse_ce", IsGlobal: false},
	"cloudmonitor":                {Service: "cloudmonitor", IsGlobal: false},
	"emr":                         {Service: "emr", IsGlobal: false},
	"ga":                          {Service: "ga", IsGlobal: true},
	"graph":                       {Service: "graph", IsGlobal: false},
	"gtm":                         {Service: "gtm", IsGlobal: true},
	"hbase":                       {Service: "hbase", IsGlobal: false},
	"metakms":                     {Service: "metakms", IsGlobal: false},
	"na":                          {Service: "na", IsGlobal: true},
	"resource_share":              {Service: "resource_share", IsGlobal: true},
	"speech_saas_prod":            {Service: "speech_saas_prod", IsGlobal: true},
	"tag":                         {Service: "tag", IsGlobal: true},
	"vefaas_dev":                  {Service: "vefaas_dev", IsGlobal: false},
	"vms":                         {Service: "vms", IsGlobal: false},
	"eco_partner":                 {Service: "eco_partner", IsGlobal: true},
	"smc":                         {Service: "smc", IsGlobal: true},
}

type RegionInfo struct{}

type regionChecker struct {
	WhiteRegions map[string]RegionInfo
	Regexp       *regexp.Regexp
}

func (r *regionChecker) Validate(region string) bool {
	if r.WhiteRegions != nil {
		if _, ok := r.WhiteRegions[region]; ok {
			return true
		}
	}
	if r.Regexp != nil {
		return r.Regexp.MatchString(region)
	}

	return false
}

var regionMatcher = regionChecker{
	WhiteRegions: map[string]RegionInfo{
		"ap-singapore-1":          {},
		"ap-southeast-1":          {},
		"ap-southeast-2":          {},
		"ap-southeast-3":          {},
		"byteplus-global":         {},
		"cn-beijing":              {},
		"cn-beijing-autodriving":  {},
		"cn-beijing-selfdrive":    {},
		"cn-beijing2":             {},
		"cn-beijing300":           {},
		"cn-changsha-sdv":         {},
		"cn-chengdu":              {},
		"cn-chengdu-sdv":          {},
		"cn-chongqing-sdv":        {},
		"cn-datong":               {},
		"cn-east-1-dedicated":     {},
		"cn-gaofang-bj":           {},
		"cn-gaofang-gz1":          {},
		"cn-gaofang-nt1":          {},
		"cn-gaofang-nt2":          {},
		"cn-gaofang-nt3":          {},
		"cn-gaofang-nt4":          {},
		"cn-gaofang-nt5":          {},
		"cn-guangzhou":            {},
		"cn-guilin-boe":           {},
		"cn-hangzhou":             {},
		"cn-hjxj":                 {},
		"cn-hjzg":                 {},
		"cn-hlbx":                 {},
		"cn-hlxj":                 {},
		"cn-hlzg":                 {},
		"cn-hongkong":             {},
		"cn-hongkong-pop":         {},
		"cn-lfbx":                 {},
		"cn-lfxj":                 {},
		"cn-lfzg":                 {},
		"cn-macau-pop-sdv":        {},
		"cn-mainland":             {},
		"cn-nanjing-bbit":         {},
		"cn-ningbo-sdv":           {},
		"cn-north-1":              {},
		"cn-north-1-dedicated":    {},
		"cn-north-boe":            {},
		"cn-shanghai":             {},
		"cn-shanghai-autodriving": {},
		"cn-taiwan-boe":           {},
		"cn-wuhan":                {},
		"cn-wulanchabu":           {},
		"cn-xian-boe-sdv":         {},
		"overseas-1":              {},
		"rec-cn":                  {},
		"rec-sg":                  {},
	},
	Regexp: regexp.MustCompile(
		`^(?:[a-z]{2}-[a-z]+(?:-[a-z]+)?|(?:cn|ap|eu|na|sa|me|af)-[a-z]+-\d+(?:-(?:finance|exclusive|local|inner))?)$`,
	),
}

func standardizeDomainServiceCode(serviceCode string) string {
	return strings.ReplaceAll(strings.ToLower(serviceCode), "_", "-")
}

const (
	// IPVersionIPv4Only will use IPv4 only for endpoint resolving.
	IPVersionIPv4Only = "IPv4Only"
	// IPVersionDualStack will use dual stack for endpoint resolving.
	IPVersionDualStack = "DualStack"
)

const (
	SiteStackVolcIPv4      = "volcengineapi"
	SiteStackVolcDualStack = "volcengine-api"
)

func (r *StandardEndpointResolver) EndpointFor(service, region string,
	opts ...func(*Options)) (resolvedEndpoint ResolvedEndpoint, err error) {
	var options Options
	options.Set(opts...)

	if !regionMatcher.Validate(region) {
		return resolvedEndpoint, volcengineerr.New("InvalidRegion",
			fmt.Sprintf("invalid region %s for standard endpoint resolver, "+
				"please upgrade the sdk endpoint resolver to the latest version", region), nil)
	}

	if len(r.Format) == 0 {
		r.Format = defaultFormat
	}
	r.Variables.Service = standardizeDomainServiceCode(service)

	serviceInfo, ok := ServiceInfos[service]
	if !ok {
		if r.CustomServices != nil {
			serviceInfo, ok = r.CustomServices[service]
		}
		if !ok {
			return resolvedEndpoint, volcengineerr.New("ServiceNotFound",
				fmt.Sprintf("service %s not found in ServiceInfos or CustomServices, "+
					"please upgrade the sdk endpoint resolver to the latest version", service), nil)
		}
	}

	isGlobal := serviceInfo.IsGlobal
	if !isGlobal {
		r.Variables.Region = "." + region
	}

	ipVersion := options.IPVersion

	if ipVersion == IPVersionDualStack {
		r.Variables.SiteStack = SiteStackVolcDualStack
	} else if ipVersion == IPVersionIPv4Only {
		r.Variables.SiteStack = SiteStackVolcIPv4
	}

	t, err := template.New("endpoint").Parse(r.Format)
	if err != nil {
		return resolvedEndpoint,
			volcengineerr.New("TemplateParseError",
				fmt.Sprintf("failed to parse endpoint template for format %s", r.Format), nil)
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, &r.Variables)
	if err != nil {
		return resolvedEndpoint,
			volcengineerr.New("TemplateExecuteError",
				fmt.Sprintf("failed to execute template for format %s, variable %v", r.Format, r.Variables), nil)
	}

	resolvedEndpoint.URL = buf.String()

	return resolvedEndpoint, nil
}
