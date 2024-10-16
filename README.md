# 快手磁力引擎MarketingAPI Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/kwai-marketing-api.svg)](https://pkg.go.dev/github.com/bububa/kwai-marketing-api)
[![Go](https://github.com/bububa/kwai-marketing-api/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/kwai-marketing-api/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/kwai-marketing-api/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/kwai-marketing-api/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/kwai-marketing-api.svg)](https://github.com/bububa/kwai-marketing-api)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/kwai-marketing-api)](https://goreportcard.com/report/github.com/bububa/kwai-marketing-api)
[![GitHub license](https://img.shields.io/github/license/bububa/kwai-marketing-api.svg)](https://github.com/bububa/kwai-marketing-api/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/kwai-marketing-api.svg)](https://GitHub.com/bububa/kwai-marketing-api/releases/)

- Oauth2 授权 (api/oauth)
  - 生成授权链接 [ Url(ctx context.Context, clt *core.SDKClient, req *oauth.UrlRequest) string ]
  - 获取AccessToken [ AccessToken(ctx context.Context, clt *core.SDKClient, authCode String) (*oauth.AccessTokenResponse, error) ]
  - 刷新Token [ RefreshToken(ctx context.Context, clt *core.SDKClient, refreshToken string) (*oauth.AccessTokenResponse, error)]
  - 拉取token下授权广告账户接口 [ ApprovalList(ctx context.Context, clt \*core.SDKClient, accessToken string, pageNo int, pageSize int) ([]uint64, error) ]
- 账号服务
  - 广告主 (api/advertiser)
    - 获取广告主信息 [ Info(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID int64) (*advertiser.Info, error) ]
    - 获取广告账户余额信息 [ FundGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID int64) (float64, error) ]
    - 获取广告主账户流水信息 [ FundDailyFlows(ctx context.Context, clt *core.SDK, accessToken string, req *advertiser.FundDailyFlowsRequest) (*advertiser.FundDailyFlowsResponse, error) ]
  - 账户罗盘(api/adcompass)
    - 获取罗盘绑定广告主列 [ Advertisers(ctx context.Context, clt \*core.SDKClient, accessToken string, advertiserID uint64) ([]adcompass.Advertiser, error) ]
    - 磁力罗盘对外 quota 腾挪接口 [ QuotaTending(ctx context.Context, clt *core.SDKClient, accessToken string, req *adcompass.QuotaTendingRequest) (string, error) ]
- 广告投放
  - 智能创编 (api/dsp)
    - 广告计划 (api/dsp/campaign)
      - 创建广告计划 [ Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.CreateRequest) (uint64, error) ]
      - 修改广告计划 [ Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.UpdateRequest) (uint64, error) ]
      - 获取广告计划信息 [ List(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.ListRequest) (*campaign.ListResponse, error) ]
      - 修改广告计划状态 [ UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.UpdateStatusRequest) ([]uint64, error) ]
    - 广告组 (api/dsp/unit)
      - 创建广告组 [ Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.CreateRequest) (uint64, error) ]
      - 修改广告组 [ Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateRequest) (uint64, error) ]
      - 查询广告组 [ List(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.ListRequest) (*unit.ListResponse, error) ]
      - 修改广告组预算 [ UpdateDayBudget(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateDayBudgetRequest) error ]
      - 修改广告组状态 [ UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateStatusRequest) ([]int64, error) ]
      - 修改广告组出价 [ UpdateBid(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateBidRequest) error ]
      - 批量获取监测链接接口 [ GetMonitorURLs(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.GetMonitorURLsRequest) ([]unit.UnitMonitorURL, error) ]
      - 监测链接批量更新接口 [ BatchUpdateMonitorURLs(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.BatchUpdateMonitorURLsRequest) ([]unit.UnitMonitorURL, error) ]
    - 广告创意 (api/dsp/creative)
      - 创建自定义创意 [ Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.CreateRequest) (uint64, error) ]
      - 创建程序化创意 [ AdvancedCreativeCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedCreativeCreateRequest) (uint64, error) ]
      - 修改自定义创意 [ Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.UpdateRequest) (uint64, error) ]
      - 修改程序化创意 [ AdvancedCreativeUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedCreativeUpdateRequest) (uint64, error) ]
      - 批量修改自定义创意 [ BatchUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.BatchUpdateRequest) (*creative.BatchUpdateResponse, error) ]
      - 查询自定义创意 [ List(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.ListRequest) (*creative.ListResponse, error) ]
      - 查询程序化创意 [ AdvancedCreativeList(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedCreativeListRequest) (*creative.AdvancedCreativeListResponse, error) ]
      - 修改创意状态 [ UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.UpdateStatusRequest) ([]uint64, error) ]
      - 创意体验 [ Preview(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.PreviewRequest) error ]
  - 获取各层级信息
    - 获取广告计划信息 [ campaign.List(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.ListRequest) (*campaign.ListResponse, error) ]
    - 获取广告组信息 [ unit.List(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.ListRequest) (*unit.ListResponse, error) ]
    - 获取广告创意信息 [ creative.List(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.ListRequest) (*creative.ListResponse, error) ]
    - 获取程序化创意2.0信息 [ creative.AdvancedProgramList(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramListRequest) (*creative.AdvancedProgramListResponse, error) ]
    - 获取程序化创意2.0审核信息 [ creative.AdvancedProgramReviewDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramReviewDetailRequest) (*creative.AdvancedProgramReviewDetail, error) ]
    - 账户操作记录信息查询 [ tool.OperationRecordList(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.OperationRecordListRequest) (*tool.OperationRecordListResponse, error) ]
    - 定向人群预估查询 [ tool.AudiencePredict(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.AudiencePredictRequest) (int64, error) ]
  - 账户层级
    - 账户日预算查询 [ advertiser.BudgetGet(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID int64) (*advertiser.Budget, error) ]
    - 修改账户预算 [ advertiser.UpdateBudget(ctx context.Context, clt *core.SDKClient, accessToken string, req *advertiser.UpdateBudgetRequest) error ]
  - 广告计划(api/campaign)
    - 创建广告计划 [ Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.CreateRequest) (int64, error) ]
    - 修改广告计划 [ Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.UpdateRequest) (int64, error) ]
    - 修改广告计划预算 [ UpdateBudget(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.UpdateBudgetRequest) error ]
    - 修改广告计划状态 [ UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *campaign.UpdateStatusRequest) ([]int64, error) ]
  - 广告组(api/unit)
    - 创建广告组 [ Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.CreateRequest) (int64, error) ]
    - 创建创建联盟定投广告组 [ CreateUnion(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.CreateUnionRequest) (int64, error) ]
    - 修改广告组 [ Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateRequest) (int64, error) ]
    - 修改联盟定投广告组 [ UpdateUnion(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateUnionRequest) (int64, error) ]
    - 修改广告组预算 [ UpdateDayBudget(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateDayBudgetRequest) error ]
    - 修改广告组状态 [ UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateStatusRequest) ([]int64, error) ]
    - 修改广告组出价 [ UpdateBid(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.UpdateBidRequest) error ]
  - 广告创意(api/creative)
    - 创建创意 [ Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.CreateRequest) (int64, error) ]
    - 创建程序化2.0创意 [ AdvancedProgramCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramCreateRequest) (int64, error) ]
    - 批量创建&修改创意 [ BatchUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.BatchUpdateRequest) (*creative.BatchUpdateResponse, error) ]
    - 修改创意 [ Update(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.UpdateRequest) (int64, error) ]
    - 修改程序化2.0创意 [ AdvancedProgramUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.AdvancedProgramUpdateRequest) error ]
    - 修改创意状态 [ UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.UpdateStatusRequest) ([]int64, error) ]
    - 创意体验 [ Preview(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.PreviewRequest) error ]
    - 创意标签填写建议 [ CreativeTagAdvise(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.CreativeTagAdviseRequest) (*creative.CreativeTagAdviseResponse, error) ]
  - 高级创意(api/asset)
    - 获取高级创意列表 [ AdvCardList(ctx context.Context, clt *core.SDKClient, accessToken string, req *asset.AdvCardListRequest) (*asset.AdvCardListResponse, error) ]
    - 创建高级创意接口 [ AdvCardCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *asset.AdvCardCreateRequest) ([]int64, error) ]
    - 删除高级创意接口 [ AdvCardRemove(ctx context.Context, clt *core.SDKClient, accessToken string, req *asset.AdvCardRemoveRequest) ([]int64, error) ]
  - 搜索广告工具
    - 关键词管理 (api/wordinfo)
      - 获取关键词列表 [ List(ctx context.Context, clt *core.SDKClient, accessToken string, req *wordinfo.ListRequest) (*wordinfo.ListResponse, error) ]
      - 创建关键词 [ Create(ctx context.Context, clt *core.SDKClient, accessToken string, req *wordinfo.CreateRequest) (*wordinfo.CreateResponse, error) ]
      - 修改关键词匹配方式 [ UpdateMatchType(ctx context.Context, clt *core.SDKClient, accessToken string, req *wordinfo.UpdateMatchTypeRequest) ([]uint64, error) ]
      - 修改关键词投放状态 [ UpdateStatus(ctx context.Context, clt *core.SDKClient, accessToken string, req *wordinfo.UpdateStatusRequest) ([]uint64, error) ]
- 数据报表
  - 广告数据报表 (api/report)
    - 代理商数据 [ AgentReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.AgentReportRequest) (*report.AgentReportResponse, error) ]
    - 广告主数据 [ AccountReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.AccountReportRequest) (*report.ReportResponse, error) ]
    - 广告计划数据 [ CampaignReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.CampaignReportRequest) (*report.ReportResponse, error) ]
    - 广告单元数据 [ UnitReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.UnitReportRequest) (*report.ReportResponse, error) ]
    - 广告创意数据 [ CreativeReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.CreativeReportRequest) (*report.ReportResponse, error) ]
    - 程序化创意数据 [ ProgramCreativeReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.ProgramCreativeReportRequest) (*report.ReportResponse, error) ]
    - 广告素材数据 [ CreativeReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.MaterialReportRequest) (*report.ReportResponse, error) ]
    - 人群分析数据 [ AudienceReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.AudienceReportRequest) (*report.ReportResponse, error) ]
    - 小店通转化数据 [ MerchantDeatailReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.MerchantDetailReportRequest) (*report.MerchantDetailReportResponse, error) ]
    - 关键词报表 [ WordInfoReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.WordInfoReportRequest) (*report.ReportResponse, error) ]
    - 搜索词报表 [ QueryWordReport(ctx context.Context, clt *core.SDKClient, accessToken string, req *report.QueryWordReportRequest) (*report.ReportResponse, error) ]
- 素材管理(api/file)
  - 图片素材
    - 上传图片v1接口 [ AdImageUploadV1(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdImageUploadRequestV1) (*file.Image, error) ]
    - 上传图片v2接口 [ AdImageUploadV2(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdImageUploadRequestV2) (*file.Image, error) ]
    - 查询图片信息get接口 [ AdImageGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdImageGetRequest) (*file.Image, error) ]
    - 查询图片信息list接口 [ AdImageList(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdImageListRequest) (*file.AdImageListResponse, error) ]
  - 视频素材
    - 上传视频接口v1 [ AdVideoUploadV1(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoUploadRequestV1) (string, error) ]
    - 上传视频接口v2 [ AdVideoUploadV2(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoUploadRequestV2) (*file.Video, error) ]
    - 获取视频信息get接口 [ AdVideoGet(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoGetRequest) ([]file.Video, error) ]
    - 查询视频信息list接口 [ AdVideoList(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoListRequest) (*file.AdVideoListResponse, error) ]
  - 视频库
    - 视频库-推送视频 [ AdVideoShare(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoShareRequest) ([]file.AdVideoShareDetail, error) ]
    - 视频库-批量更新视频功能 [ AdVideoUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoUpdateRequest) error ]
    - 视频库-删除视频标签 [ AdVideoTagDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoTagDeleteRequest) error ]
    - 视频关联创意数查询 [ AdVideoRelateCreatives(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdVideoRelateCreativesRequest) ([]file.AdVideoRelatedCreatives, error) ]
    - 查询账户共享视频库按钮是否开启 [ dsp.video.QueryAutoShareSwitch(ctx context.Context, clt *core.SDKClient, accessToken string, req *video.QueryAutoShareSwitchRequest) (*video.QueryAutoShareSwitchResponse, error) ]
- 建站管理(api/page)
  - 获取魔力建站落地页组信息列表 [ List(ctx context.Context, clt *core.SDKClient, accessToken string, req *page.ListRequest) (*page.ListResponse, error) ]
  - 批量转赠 [ BatchGive(ctx context.Context, clt *core.SDKClient, accessToken string, req *page.BatchGiveRequest) error ]
  - 魔力建站落地页更新CID信息 [ CidInfoUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *page.CidInfoUpdateRequest) (uint64, error) ]
- 工具
  - 查询工具
    - 获取可选的深度转化目标 [ unit.OcpcConversionInfos(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.OcpcConversionInfosRequest) (*unit.OcpcConversionInfosResponse, error) ]
    - 获取可选的定向标签 [ tool.TargetingTagsList(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.TargetingTagsListRequest) (*tool.TargetingTag, error) ]
    - 获取可选的应用定向 [ tool.AppSearch(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.AppSearchRequest) (*tool.TargetingApp, error) ]
    - 获取可选的推荐封面 [ tool.KeyFrame(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.KeyFrameRequest) ([]string, error) ]
    - 获取可选的动态词包 [ tool.CreativeWordList(ctx context.Context, clt \*core.SDKClient, accessToken string, advertiserID int64) ([]tool.CreativeWord, error) ]
    - 获取行动号召按钮 [ creative.ActionBarTextList(ctx context.Context, clt *core.SDKClient, accessToken string, req *creative.ActionBarTextListRequest) ([]string, error) ]
    - 获取可选的封面贴纸样式 [ tool.CreativeWordStyles(ctx context.Context, clt \*core.SDKClient, accessToken string, advertiserID int64) ([]tool.CreativeWordStyle, error) ]
    - 获取可用的转化目标 [ tool.ConvertList(ctx context.Context, clt *core.SDKClient, accessToken string, req *tool.ConvertListRequest) (*tool.ConvertListResponse, error) ]
    - 获取可选白名单接口 [ advertiser.WhiteList(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID int64) (*advertiser.WhiteListResponse, error) ]
    - 获取地域定向 [ region.List(ctx context.Context, clt \*core.SDKClient, accessToken string) (map[string]region.Region, error) ]
    - 获取商圈地域定向 [ region.DistrictList(ctx context.Context, clt \*core.SDKClient, accessToken string, advertiserID int64) (map[string]region.District, error) ]
    - 获取可用咨询组件列表 [ lp.ConsultList(ctx context.Context, clt *core.SDKClient, accessToken string, req *lp.ConsultListRequest) (*lp.ConsultListResponse, error) ]
  - 出价建议
    - 获取广告组出价建议 [ tool.unit.SuggestBidDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.SuggestBidDetailRequest) ([]unit.SuggestBidUnit, error) ]
    - 获取广告组曝光、转化预估 [ tool.unit.BidPredict(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.BidPredictRequest) (*unit.BidPredict, error) ]
    - 获取广告组投放预估曲线 [ tool.unit.BidTrendPredictDetail(ctx context.Context, clt *core.SDKClient, accessToken string, req *unit.BidTrendPredictDetailRequest) (*unit.BidTrendPredict, error) ]
  - 投前预估
    - 投前预估列表页接口 [ tools.preunit.PredicationTaskList(ctx context.Context, clt *core.SDKClient, accessToken string, req *preput.PredicationTaskListRequest) (*preput.PredicationTaskListResponse, error) ]
    - 投前预估详情 [ tools.preput.PredicationTaskDetails(ctx context.Context, clt *core.SDKClient, accessToken string, req *preput.PredicationTaskCreateRequest) (*preput.AdPredicationTaskDetail, error) ]
    - 投前预估任务管理接口 [ tool.preput.PredicationTaskManagement(ctx context.Context, clt *core.SDKClient, accessToken string, req *preput.PredicationTaskManagementRequest) (*preput.RealTaskResult, error) ]
    - 创建投前预估任务 [ tool.preput.PredicationTaskCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *preput.PredicationTaskCreateRequest) (*preput.RealTaskResult, error) ]
  - 功能名单
    - 获取创意分类标签白名单客户 [ advertiser.WhiteList(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID int64) (*advertiser.WhiteListResponse, error) ]
    - 获取联盟投放白名单 [ advertiser.WhiteList(ctx context.Context, clt *core.SDKClient, accessToken string, advertiserID int64) (*advertiser.WhiteListResponse, error) ]
  - 应用列表
    - 创建应用 [ file.AdAppCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdAppCreateRequest) (*file.App, error) ]
    - 修改应用 [ file.AdAppUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdAppUpdateRequest) (*file.App, error) ]
    - 获取应用列表 [ file.AdAppList(ctx context.Context, clt *core.SDKClient, accessToken string, req *file.AdAppListRequest) (*file.AdAppListResponse, error) ]
  - 定向模版(新) (api/dsp/target)
    - 查询定向模板 [ TemplateDetails(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateDetailsRequest) (*target.TemplateDetailsResponse, error) ]
    - 创建定向模板 [ TemplateCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateCreateRequest) (uint64, error) ]
    - 更新定向模板 [ TemplateUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateUpdateRequest) (uint64, error) ]
    - 删除定向模板 [ TemplateDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateDeleteRequest) (uint64, error) ]
    - 根据店铺名称查询商圈信息 [ OptionDistanceList(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.OptionDistanceListRequest) (*target.OptionDistanceListResponse, error) ]
    - 定向模板同步 [ TemplateUnitSync(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateUnitSyncRequest) (*target.TemplateUnitSyncResponse, error) ]
    - 模板同步失败查询 [ TemplateSyncHistory(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateSyncHistoryRequest) (*target.TemplateSyncHistoryResponse, error) ]
    - 查询模板关联的广告列表接口 [ TemplateRelatedUnitList(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateRelatedUnitListRequest) (*target.TemplateRelatedUnitListResponse, error) ]
    - 查询待升级模板列表 [ TemplateUpgradeList(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateUpgradeListRequest) ([]target.TemplateUpgradeItem, error) ]
    - 模板升级 [ TemplateUpgrade(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateUpgradeRequest) (int64, error) ]
  - 定向模版
    - 创建定向模板 [ target.TemplateCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateCreateRequest) (*target.Template, error) ]
    - 查询定向模板接口 [ target.TemplateList(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateListRequest) (*target.TemplateListResponse, error) ]
    - 修改定向模板 [ target.TemplateUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateUpdateRequest) (*target.Template, error) ]
    - 删除定向模板 [ target.TemplateDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *target.TemplateDeleteRequest) error ]
  - 直播推广 (api/dsp/jingbell)
    - 小铃铛推送 [ jingbel.Share(ctx context.Context, clt *core.SDKClient, accessToken string, req *jingbell.ShareRequest) error ]
  - 原生广告投放工具 (api/dsp/native)
    - 开启原生扩量开关接口 [ OpenAccountNative(ctx context.Context, clt *core.SDKClient, accessToken string, req *native.OpenAccountNativeRequest) error ]
  - 评论管理 (api/comment)
    - 评论列表数据查询 [ List(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ListRequest) (*comment.ListResponse, error) ]
    - 回复评论 [ Reply(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ReplyRequest) ([]comment.ReplyResult, error) ]
    - 评论树查询 [ Tree(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.TreeRequest) (*comment.TreeResponse, error) ]
    - 屏蔽评论 [ Shield(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ShieldRequest) error ]
    - 增加屏蔽评论信息 [ ShieldInfoCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ShieldInfoCreateRequest) ([]uint64, error) ]
    - 删除屏蔽评论信息 [ ShieldInfoDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ShieldInfoDeleteRequest) error ]
    - 获取屏蔽评论信息列表 [ ShieldInfoList(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.ShieldInfoListRequest) (*comment.ShieldInfoListResponse, error) ]
    - 评论置顶 [ SetTop(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.SetTopRequest) (uint64, error) ]
    - 取消评论置顶 [ CancelTop(ctx context.Context, clt *core.SDKClient, accessToken string, req *comment.CancelTopRequest) (uint64, error) ]
- DMP人群管理(api/dmp)
  - 人群包上传接口 [ PopulationUpload(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationUploadRequest) (*dmp.Population, error) ]
  - 人群包更新接口 [ PopulationUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationUpdateRequest) (*dmp.Population, error) ]
  - 人群列表查询接口 [ PopulationList(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationListRequest) ([]dmp.Population, error) ]
  - 人群包删除接口 [ PopulationDelete(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationDeleteRequest) error ]
  - 人群包跨账户推送 [ PopulationAccountsPush(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationAccountsPushRequest) (*dmp.PopulationAccountsPushResponse, error) ]
  - 人群包上线接口 [ PopulationPush(ctx context.Context, clt *core.SDKClient, accessToken string, req *dmp.PopulationPushRequest) error ]
- 应用管理 (api/appcenter)
  - 文件上传 (api/appcenter/upload)
    - 上传 APK 文件 [ Apk(ctx context.Context, clt *core.SDKClient, accessToken string, req *upload.ApkRequest) (string, error) ]
    - 上传图片 [ Pic(ctx context.Context, clt *core.SDKClient, accessToken string, req *upload.PicRequest) (string, error) ]
  - 应用创编
    - 创建 Android 应用 [ CreateAndroid(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.CreateAndroidRequest) (*app.App, error) ]
    - 创建 iOS 应用 [ CreateIos(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.CreateIosRequest) (*app.App, error) ]
    - 编辑 Android 应用 [ UpdateAndroid(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.UpdateAndroidRequest) (*app.App, error) ]
    - 更新 iOS 应用 [ UpdateIos(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.UpdateIosRequest) (*app.App, error) ]
  - 应用查询For单元投放
    - 获取新版应用发布列表【单元创编】[ app.ReleaseList(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.ReleaseListRequest) (*app.ListResponse, error) ]
    - 获取新版分包发布列表【单元创编】 [ subpkg.ReleaseList(ctx context.Context, clt *core.SDKClient, accessToken string, req *subpkg.ReleaseListRequest) (*subpkg.ListResponse, error) ]
  - 应用查询For应用中心
    - 获取应用列表 [ app.List(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.ListRequest) (*app.ListResponse, error) ]
    - 获取应用详情 [ app.Detail(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.DetailRequest) (*app.App, error) ]
  - 应用操作
    - iOS 应用上报更新 [ app.IosUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.IosUpdateRequest) error ]
    - 发布应用 [ app.Release(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.ReleaseRequest) (*app.App, error) ]
    - 应用上架 [ app.Online(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.OnlineRequest) error ]
    - 应用下架 [ app.Offline(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.OfflineRequest) error ]
    - 应用商店上下架 [ app.OfflineAppStores(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.OfflineAppStoresRequest) error ]
  - 应用分包
    - 新建应用分包 [ subpkg.Add(ctx context.Context, clt *core.SDKClient, accessToken string, req *subpkg.AddRequest) ([]subpkg.SubPackage, error) ]
    - 更新/恢复/删除应用分包 [ subpkg.Mod(ctx context.Context, clt *core.SDKClient, accessToken string, req *subpkg.ModRequest) error ]
    - 修改应用分包备注 [ subpkg.Description(ctx context.Context, clt *core.SDKClient, accessToken string, req *subpkg.DescriptionRequest) error ]
    - 获取分包管理/回收站列表 [ subpkg.List(ctx context.Context, clt *core.SDKClient, accessToken string, req *subpkg.ListRequest) (*subpkg.ListResponse, error) ]
    - 分包失败重新构建 [ app.RetryBuildSubPackage(ctx context.Context, clt *core.SDKClient, accessToken string, req *app.RetryBuildSubPackageRequest) (int, error) ]
- 商品库 (api/dsp/dpa)
  - 查询 DPA 模板信息 [ TemplateList(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.TemplateListRequest) (*dpa.TemplateListResponse, error) ]
  - 获取商品库类目树 [ CategoryList(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.CategoryListRequest) (*dpa.CategoryListResponse, error) ]
  - 获取商品库列表 [ LibraryList(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.LibraryListRequest) (*dpa.LibraryListResponse, error) ]
  - 创建商品 [ ProductBatchCreate(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ProductBatchCreateRequest) ([]dpa.ProductUpdateResult, error) ]
  - 更新商品 [ ProductBatchUpdate(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ProductBatchUpdateRequest) ([]dpa.ProductUpdateResult, error) ]
  - 获取商品列表 [ ProductBatchQuery(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ProductBatchQueryRequest) (*dpa.ProductBatchQueryResponse, error) ]
  - 获取商品列表(游标) [ ProductCursorQuery(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.ProductCursorQueryRequest) (*dpa.ProductCursorQueryResponse, error) ]
  - 获取SDPA创意视频模板 [ CreativeTemplateList(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.CreativeTemplateListRequest) (*dpa.CreativeTemplateListResponse, error) ]
  - 批量模板合成SDPA创意视频 [ CreativeVideoGenerate(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.CreativeVideoGenerateRequest) ([]dpa.GenerateVideoResult, error) ]
  - CID服务商投放SDPA接口 [ SecretCidLink(ctx context.Context, clt *core.SDKClient, accessToken string, req *dpa.SecretCidLinkRequest) error ]
- 数据上报管理 (api/track)
  - 转化回传 [ Activate(req *track.ActivateRequest) error ]
  - 点击检测链接 [ Click(baseUrl string, fields []string) string ]

# Reference

[API文档](https://developers.e.kuaishou.com/docs)
