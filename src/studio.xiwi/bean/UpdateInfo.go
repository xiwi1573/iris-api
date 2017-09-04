package bean

/**
 * app更新信息
 */
type UpdateAppInfo struct {
	CurrVerCode    int16
	CurrVerCodeStr string

	LastVerCode    int16
	LastVerCodeStr string

	IsForceUpdate    bool
	DownloadLink     string
	UpdateNotifyInfo string
}
