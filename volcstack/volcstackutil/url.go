package volcstackutil

type UrlInfo struct {
	UseSSL bool
	//Locate bool
	//UseInternal                 bool
	CustomerDomain string
	//CustomerDomainIgnoreService bool
}

type ServiceInfo struct {
	Service string
	Region  string
}

const (
	url = "open.volcengineapi.com"
	//internalUrl = "open.volcengineapi.com"
	http  = "http"
	https = "https"
)

func UrlWithServiceInfo(urlInfo *UrlInfo, info ServiceInfo) string {
	p := Protocol(urlInfo.UseSSL)
	if urlInfo.CustomerDomain != "" {
		//if !urlInfo.CustomerDomainIgnoreService {
		//	return p + "://" + info.Service + "." + urlInfo.CustomerDomain
		//} else {
		//	return p + "://" + urlInfo.CustomerDomain
		//}
		return p + "://" + urlInfo.CustomerDomain

	} else {
		//if urlInfo.UseInternal {
		//	return p + "://" + internalUrl
		//}
		//if urlInfo.Locate && &info.Region != nil {
		//	return p + "://" + info.Service + "." + info.Region + url
		//}
		return p + "://" + url
	}
}
func Url(urlInfo *UrlInfo) string {
	p := Protocol(urlInfo.UseSSL)
	if urlInfo.CustomerDomain != "" {
		return p + "://" + urlInfo.CustomerDomain
	} else {
		return p + "://" + url
	}
}

func Protocol(useSSL bool) string {
	if useSSL {
		return https
	}
	return http
}
