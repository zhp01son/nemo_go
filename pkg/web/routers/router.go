package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/hanc00l/nemo_go/pkg/web/controllers"
)

func init() {
	//beego v2.0.2后手工注册路由风格：
	web.CtrlGet("/", (*controllers.LoginController).IndexAction)
	web.CtrlPost("/", (*controllers.LoginController).LoginAction)
	web.CtrlGet("/logout", (*controllers.LoginController).LogoutAction)

	web.CtrlGet("/config-list", (*controllers.ConfigController).IndexAction)
	web.CtrlPost("/config-list", (*controllers.ConfigController).LoadDefaultConfigAction)
	web.CtrlPost("/config-change-password", (*controllers.ConfigController).ChangePasswordAction)
	web.CtrlGet("/custom-list", (*controllers.ConfigController).CustomAction)
	web.CtrlPost("/custom-load", (*controllers.ConfigController).LoadCustomConfigAction)
	web.CtrlPost("/custom-save", (*controllers.ConfigController).SaveCustomConfigAction)
	web.CtrlPost("/config-save-taskslice", (*controllers.ConfigController).SaveTaskSliceNumberAction)
	web.CtrlPost("/config-save-portscan", (*controllers.ConfigController).SavePortscanAction)
	web.CtrlPost("/config-save-fingerprint", (*controllers.ConfigController).SaveFingerprintAction)
	web.CtrlPost("/config-upload-xraypoc", (*controllers.ConfigController).UploadXrayPocAction)
	web.CtrlPost("/config-save-notify", (*controllers.ConfigController).SaveTaskNotifyAction)

	web.CtrlGet("/dashboard", (*controllers.DashboardController).IndexAction)
	web.CtrlPost("/dashboard", (*controllers.DashboardController).GetStatisticDataAction)
	web.CtrlPost("/dashboard-task-info", (*controllers.DashboardController).GetTaskInfoAction)
	web.CtrlPost("/worker-list", (*controllers.DashboardController).WorkerAliveListAction)
	web.CtrlPost("/onlineuser-list", (*controllers.DashboardController).OnlineUserListAction)
	web.CtrlPost("/dashboard-task-started-info", (*controllers.DashboardController).GetStartedTaskInfoAction)
	web.CtrlPost("/worker-reload", (*controllers.DashboardController).ManualReloadWorkerAction)
	web.CtrlPost("/worker-filesync", (*controllers.DashboardController).ManualWorkerFileSyncAction)

	web.CtrlGet("/ip-list", (*controllers.IPController).IndexAction)
	web.CtrlPost("/ip-list", (*controllers.IPController).ListAction)
	web.CtrlGet("/ip-info", (*controllers.IPController).InfoAction)
	web.CtrlPost("/ip-delete", (*controllers.IPController).DeleteIPAction)
	web.CtrlPost("/port-attr-delete", (*controllers.IPController).DeletePortAttrAction)
	web.CtrlGet("/ip-statistics", (*controllers.IPController).StatisticsAction)
	web.CtrlGet("/ip-memo-get", (*controllers.IPController).GetMemoAction)
	web.CtrlPost("/ip-memo-update", (*controllers.IPController).UpdateMemoAction)
	web.CtrlGet("/ip-memo-export", (*controllers.IPController).ExportMemoAction)
	web.CtrlPost("/ip-color-tag", (*controllers.IPController).MarkColorTagAction)
	web.CtrlPost("/ip-import-portscan", (*controllers.IPController).ImportPortscanResultAction)

	web.CtrlGet("/domain-list", (*controllers.DomainController).IndexAction)
	web.CtrlPost("/domain-list", (*controllers.DomainController).ListAction)
	web.CtrlGet("/domain-info", (*controllers.DomainController).InfoAction)
	web.CtrlPost("/domain-delete", (*controllers.DomainController).DeleteDomainAction)
	web.CtrlPost("/domain-attr-delete", (*controllers.DomainController).DeleteDomainAttrAction)
	web.CtrlPost("/domain-onlineapi-attr-delete", (*controllers.DomainController).DeleteDomainOnlineAPIAttrAction)
	web.CtrlGet("/domain-statistics", (*controllers.DomainController).StatisticsAction)
	web.CtrlGet("/domain-memo-get", (*controllers.DomainController).GetMemoAction)
	web.CtrlPost("/domain-memo-update", (*controllers.DomainController).UpdateMemoAction)
	web.CtrlGet("/domain-memo-export", (*controllers.DomainController).ExportMemoAction)
	web.CtrlPost("/domain-color-tag", (*controllers.DomainController).MarkColorTagAction)

	web.CtrlGet("/vulnerability-list", (*controllers.VulController).IndexAction)
	web.CtrlPost("/vulnerability-list", (*controllers.VulController).ListAction)
	web.CtrlGet("/vulnerability-info", (*controllers.VulController).InfoAction)
	web.CtrlPost("/vulnerability-delete", (*controllers.VulController).DeleteAction)
	web.CtrlPost("/vulnerability-load-xray-pocfile", (*controllers.VulController).LoadXrayPocFileAction)
	web.CtrlPost("/vulnerability-load-nuclei-pocfile", (*controllers.VulController).LoadNucleiPocFileAction)

	web.CtrlGet("/org-list", (*controllers.OrganizationController).IndexAction)
	web.CtrlPost("/org-list", (*controllers.OrganizationController).ListAction)
	web.CtrlPost("/org-get", (*controllers.OrganizationController).GetAction)
	web.CtrlPost("/org-getall", (*controllers.OrganizationController).GetAllAction)
	web.CtrlGet("/org-add", (*controllers.OrganizationController).AddIndexAction)
	web.CtrlPost("/org-add", (*controllers.OrganizationController).AddSaveAction)
	web.CtrlPost("/org-update", (*controllers.OrganizationController).UpdateAction)
	web.CtrlPost("/org-delete", (*controllers.OrganizationController).DeleteAction)

	web.CtrlGet("/task-list", (*controllers.TaskController).IndexAction)
	web.CtrlPost("/task-list", (*controllers.TaskController).ListAction)
	web.CtrlGet("/task-info-run", (*controllers.TaskController).InfoAction)
	web.CtrlPost("/task-delete-run", (*controllers.TaskController).DeleteAction)
	web.CtrlPost("/task-stop-run", (*controllers.TaskController).StopAction)
	web.CtrlPost("/task-start-portscan", (*controllers.TaskController).StartPortScanTaskAction)
	web.CtrlPost("/task-start-batchscan", (*controllers.TaskController).StartBatchScanTaskAction)
	web.CtrlPost("/task-start-domainscan", (*controllers.TaskController).StartDomainScanTaskAction)
	web.CtrlPost("/task-start-vulnerability", (*controllers.TaskController).StartPocScanTaskAction)
	web.CtrlPost("/task-batch-delete", (*controllers.TaskController).DeleteBatchAction)
	web.CtrlPost("/task-start-xscan", (*controllers.TaskController).StartXScanTaskAction)
	web.CtrlGet("/task-info-main", (*controllers.TaskController).InfoMainAction)
	web.CtrlPost("/task-delete-main", (*controllers.TaskController).DeleteMainAction)

	web.CtrlGet("/task-cron-list", (*controllers.TaskController).IndexCronAction)
	web.CtrlPost("/task-cron-list", (*controllers.TaskController).ListCronAction)
	web.CtrlGet("/task-cron-info", (*controllers.TaskController).InfoCronAction)
	web.CtrlPost("/task-cron-delete", (*controllers.TaskController).DeleteCronAction)
	web.CtrlPost("/task-cron-disable", (*controllers.TaskController).DisableCronTaskAction)
	web.CtrlPost("/task-cron-enable", (*controllers.TaskController).EnableCronTaskAction)
	web.CtrlPost("/task-cron-run", (*controllers.TaskController).RunCronTaskAction)

	web.CtrlGet("/key-word-list", (*controllers.KeySearchController).IndexAction)
	web.CtrlPost("/key-word-list", (*controllers.KeySearchController).ListAction)
	web.CtrlPost("/key-word-add", (*controllers.KeySearchController).AddSaveAction)
	web.CtrlPost("/key-word-del", (*controllers.KeySearchController).DeleteKeyWordAction)
}
