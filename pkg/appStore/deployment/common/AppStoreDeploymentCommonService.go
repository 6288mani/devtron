/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package appStoreDeploymentCommon

import (
	"github.com/devtron-labs/devtron/internal/util"
	appStoreBean "github.com/devtron-labs/devtron/pkg/appStore/bean"
	"github.com/devtron-labs/devtron/pkg/appStore/deployment/repository"
	"github.com/go-pg/pg"
	"github.com/google/go-github/github"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/xanzy/go-gitlab"
	"go.uber.org/zap"
	"net/http"
)

type AppStoreDeploymentCommonService interface {
	GetInstalledAppByClusterNamespaceAndName(clusterId int, namespace string, appName string) (*appStoreBean.InstallAppVersionDTO, error)
	GetInstalledAppByInstalledAppId(installedAppId int) (*appStoreBean.InstallAppVersionDTO, error)
	ParseGitRepoErrorResponse(err error) (bool, error)
}

type AppStoreDeploymentCommonServiceImpl struct {
	logger                 *zap.SugaredLogger
	installedAppRepository repository.InstalledAppRepository
}

func NewAppStoreDeploymentCommonServiceImpl(logger *zap.SugaredLogger, installedAppRepository repository.InstalledAppRepository) *AppStoreDeploymentCommonServiceImpl {
	return &AppStoreDeploymentCommonServiceImpl{
		logger:                 logger,
		installedAppRepository: installedAppRepository,
	}
}

func (impl AppStoreDeploymentCommonServiceImpl) GetInstalledAppByClusterNamespaceAndName(clusterId int, namespace string, appName string) (*appStoreBean.InstallAppVersionDTO, error) {
	installedApp, err := impl.installedAppRepository.GetInstalledApplicationByClusterIdAndNamespaceAndAppName(clusterId, namespace, appName)
	if err != nil {
		if err == pg.ErrNoRows {
			impl.logger.Warnw("no installed apps found", "clusterId", clusterId)
			return nil, nil
		} else {
			impl.logger.Errorw("error while fetching installed apps", "clusterId", clusterId, "error", err)
			return nil, err
		}
	}

	if installedApp.Id > 0 {
		installedAppVersion, err := impl.installedAppRepository.GetInstalledAppVersionByInstalledAppIdAndEnvId(installedApp.Id, installedApp.EnvironmentId)
		if err != nil {
			return nil, err
		}
		return impl.convert(installedApp, installedAppVersion), nil
	}

	return nil, nil
}

func (impl AppStoreDeploymentCommonServiceImpl) GetInstalledAppByInstalledAppId(installedAppId int) (*appStoreBean.InstallAppVersionDTO, error) {
	installedAppVersion, err := impl.installedAppRepository.GetActiveInstalledAppVersionByInstalledAppId(installedAppId)
	if err != nil {
		return nil, err
	}
	installedApp := &installedAppVersion.InstalledApp
	return impl.convert(installedApp, installedAppVersion), nil
}

// converts db object to bean
func (impl AppStoreDeploymentCommonServiceImpl) convert(chart *repository.InstalledApps, installedAppVersion *repository.InstalledAppVersions) *appStoreBean.InstallAppVersionDTO {
	chartVersionApp := installedAppVersion.AppStoreApplicationVersion
	chartRepo := installedAppVersion.AppStoreApplicationVersion.AppStore.ChartRepo
	return &appStoreBean.InstallAppVersionDTO{
		EnvironmentId:         chart.EnvironmentId,
		Id:                    chart.Id,
		AppId:                 chart.AppId,
		TeamId:                chart.App.TeamId,
		TeamName:              chart.App.Team.Name,
		AppOfferingMode:       chart.App.AppOfferingMode,
		ClusterId:             chart.Environment.ClusterId,
		Namespace:             chart.Environment.Namespace,
		AppName:               chart.App.AppName,
		EnvironmentName:       chart.Environment.Name,
		InstalledAppId:        chart.Id,
		InstalledAppVersionId: installedAppVersion.Id,
		InstallAppVersionChartDTO: &appStoreBean.InstallAppVersionChartDTO{
			AppStoreChartId: chartVersionApp.AppStore.Id,
			ChartName:       chartVersionApp.Name,
			ChartVersion:    chartVersionApp.Version,
			InstallAppVersionChartRepoDTO: &appStoreBean.InstallAppVersionChartRepoDTO{
				RepoName: chartRepo.Name,
				RepoUrl:  chartRepo.Url,
				UserName: chartRepo.UserName,
				Password: chartRepo.Password,
			},
		},
		DeploymentAppType:            chart.DeploymentAppType,
		AppStoreApplicationVersionId: installedAppVersion.AppStoreApplicationVersionId,
	}
}

func (impl AppStoreDeploymentCommonServiceImpl) ParseGitRepoErrorResponse(err error) (bool, error) {
	//update values yaml in chart
	noTargetFound := false
	if err != nil {
		if errorResponse, ok := err.(*github.ErrorResponse); ok && errorResponse.Response.StatusCode == http.StatusNotFound {
			impl.logger.Errorw("no content found while updating git repo on github, do auto fix", "error", err)
			noTargetFound = true
		}
		if errorResponse, ok := err.(azuredevops.WrappedError); ok && *errorResponse.StatusCode == http.StatusNotFound {
			impl.logger.Errorw("no content found while updating git repo on azure, do auto fix", "error", err)
			noTargetFound = true
		}
		if errorResponse, ok := err.(*azuredevops.WrappedError); ok && *errorResponse.StatusCode == http.StatusNotFound {
			impl.logger.Errorw("no content found while updating git repo on azure, do auto fix", "error", err)
			noTargetFound = true
		}
		if errorResponse, ok := err.(*gitlab.ErrorResponse); ok && errorResponse.Response.StatusCode == http.StatusNotFound {
			impl.logger.Errorw("no content found while updating git repo gitlab, do auto fix", "error", err)
			noTargetFound = true
		}
		if err.Error() == util.BITBUCKET_REPO_NOT_FOUND_ERROR {
			impl.logger.Errorw("no content found while updating git repo bitbucket, do auto fix", "error", err)
			noTargetFound = true
		}
	}
	return noTargetFound, err
}
