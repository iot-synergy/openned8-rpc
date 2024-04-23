
¼
openned8.protoopenned8"
Empty"
BeanMsg
msg (	Rmsg"
IdString
id (	Rid":
PageInfo
Page (RPage
PageSize (RPageSize"H
SdkUsage
UserId (	RUserId
All (RAll
Used (RUsed"ï
AppInfo
Id (	RId
	CreatedAt (R	CreatedAt
	UpdatedAt (R	UpdatedAt
UserId (	RUserId
AppName (	RAppName
Summary (	RSummary 
AppCategory (RAppCategory 
UseIndustry (RUseIndustry(
AppCategoryName	 (	RAppCategoryName(
UseIndustryName
 (	RUseIndustryName
AppKey (	RAppKey
	AppSecret (	R	AppSecret"¢
AppInfoCreateReq
UserId (	RUserId
AppName (	RAppName
Summary (	RSummary 
AppCategory (RAppCategory 
UseIndustry (RUseIndustry"²
AppInfoUpdateReq
Id (	RId
UserId (	RUserId
AppName (	RAppName
Summary (	RSummary 
AppCategory (RAppCategory 
UseIndustry (RUseIndustry":
AppInfoDeleteReq
Id (	RId
UserId (	RUserId"Â

AppListReq.
PageInfo (2.openned8.PageInfoRPageInfo
AppName (	RAppName(
AppCategoryName (	RAppCategoryName(
UseIndustryName (	RUseIndustryName
UserId (	RUserId"J
ApplistInfo
Total (RTotal%
Data (2.openned8.AppInfoRData">
CategorylistResp*
Data (2.openned8.CategoryInfoRData"2
CategoryInfo
Id (RId
Name (	RName">
IndustrylistResp*
Data (2.openned8.IndustryInfoRData"2
IndustryInfo
Id (RId
Name (	RName"‘
ActiveCodeListReq.
PageInfo (2.openned8.PageInfoRPageInfo
AppId (	RAppId
Status (RStatus

ExpireDate (R
ExpireDate"X
ActiveCodeListInfo
Total (RTotal,
Data (2.openned8.ActiveCodeInfoRData"d
ActiveCodeResp
code (Rcode
msg (	Rmsg,
data (2.openned8.ActiveCodeInfoRdata"ö
ActiveCodeInfo
Id (	RId
	CreatedAt (R	CreatedAt
	UpdatedAt (R	UpdatedAt
	ActiveKey (	R	ActiveKey
UserId (	RUserId
AppId (	RAppId
ActiveIP (	RActiveIP
DeviceSN (	RDeviceSN
	DeviceMac	 (	R	DeviceMac&
DeviceIdentity
 (	RDeviceIdentity

ActiveDate (R
ActiveDate

ActiveType (R
ActiveType

ActiveFile (	R
ActiveFile
Version (	RVersion
	StartDate (R	StartDate

ExpireDate (R
ExpireDate
Status (RStatus"^
ActiveCodeCreatReq
UserId (	RUserId
Quantity (RQuantity
AppId (	RAppId".
UserSdkUsageQueryReq
UserId (	RUserId"A
UserSdkUsageUpdateReq
UserId (	RUserId
All (RAll2 
openned8:
	appCreate.openned8.AppInfoCreateReq.openned8.AppInfo:
	appUpdate.openned8.AppInfoUpdateReq.openned8.AppInfo:
	appDelete.openned8.AppInfoDeleteReq.openned8.BeanMsg7
appQuery.openned8.AppListReq.openned8.ApplistInfo<
categoryQuery.openned8.Empty.openned8.CategorylistResp<
industryQuery.openned8.Empty.openned8.IndustrylistRespL
activeCodeQuery.openned8.ActiveCodeListReq.openned8.ActiveCodeListInfoI
activeCodeCreat.openned8.ActiveCodeCreatReq.openned8.ActiveCodeRespG
queryUserSdkUsage.openned8.UserSdkUsageQueryReq.openned8.SdkUsageI
updateUserSdkUsage.openned8.UserSdkUsageUpdateReq.openned8.SdkUsageBZ
./openned8bproto3