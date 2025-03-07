// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	pg "github.com/go-pg/pg"
	mock "github.com/stretchr/testify/mock"

	pipelineConfig "github.com/devtron-labs/devtron/internal/sql/repository/pipelineConfig"
)

// PipelineRepository is an autogenerated mock type for the PipelineRepository type
type PipelineRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id, tx
func (_m *PipelineRepository) Delete(id int, tx *pg.Tx) error {
	ret := _m.Called(id, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, *pg.Tx) error); ok {
		r0 = rf(id, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterDeploymentDeleteRequestedPipelineIds provides a mock function with given fields: cdPipelineIds
func (_m *PipelineRepository) FilterDeploymentDeleteRequestedPipelineIds(cdPipelineIds []int) (map[int]bool, error) {
	ret := _m.Called(cdPipelineIds)

	var r0 map[int]bool
	var r1 error
	if rf, ok := ret.Get(0).(func([]int) (map[int]bool, error)); ok {
		return rf(cdPipelineIds)
	}
	if rf, ok := ret.Get(0).(func([]int) map[int]bool); ok {
		r0 = rf(cdPipelineIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int]bool)
		}
	}

	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(cdPipelineIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveByAppId provides a mock function with given fields: appId
func (_m *PipelineRepository) FindActiveByAppId(appId int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(appId)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(appId)
	}
	if rf, ok := ret.Get(0).(func(int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(appId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(appId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveByAppIdAndEnvironmentId provides a mock function with given fields: appId, environmentId
func (_m *PipelineRepository) FindActiveByAppIdAndEnvironmentId(appId int, environmentId int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(appId, environmentId)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(appId, environmentId)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(appId, environmentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(appId, environmentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveByAppIdAndEnvironmentIdV2 provides a mock function with given fields:
func (_m *PipelineRepository) FindActiveByAppIdAndEnvironmentIdV2() ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called()

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*pipelineConfig.Pipeline, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*pipelineConfig.Pipeline); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveByAppIdAndPipelineId provides a mock function with given fields: appId, pipelineId
func (_m *PipelineRepository) FindActiveByAppIdAndPipelineId(appId int, pipelineId int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(appId, pipelineId)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(appId, pipelineId)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(appId, pipelineId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(appId, pipelineId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveByAppIds provides a mock function with given fields: appIds
func (_m *PipelineRepository) FindActiveByAppIds(appIds []int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(appIds)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func([]int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(appIds)
	}
	if rf, ok := ret.Get(0).(func([]int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(appIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(appIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveByEnvId provides a mock function with given fields: envId
func (_m *PipelineRepository) FindActiveByEnvId(envId int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(envId)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(envId)
	}
	if rf, ok := ret.Get(0).(func(int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(envId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(envId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveByEnvIdAndDeploymentType provides a mock function with given fields: environmentId, deploymentAppType, exclusionList, includeApps
func (_m *PipelineRepository) FindActiveByEnvIdAndDeploymentType(environmentId int, deploymentAppType string, exclusionList []int, includeApps []int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(environmentId, deploymentAppType, exclusionList, includeApps)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int, string, []int, []int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(environmentId, deploymentAppType, exclusionList, includeApps)
	}
	if rf, ok := ret.Get(0).(func(int, string, []int, []int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(environmentId, deploymentAppType, exclusionList, includeApps)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string, []int, []int) error); ok {
		r1 = rf(environmentId, deploymentAppType, exclusionList, includeApps)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveByEnvIds provides a mock function with given fields: envId
func (_m *PipelineRepository) FindActiveByEnvIds(envId []int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(envId)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func([]int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(envId)
	}
	if rf, ok := ret.Get(0).(func([]int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(envId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(envId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveByInFilter provides a mock function with given fields: envId, appIdIncludes
func (_m *PipelineRepository) FindActiveByInFilter(envId int, appIdIncludes []int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(envId, appIdIncludes)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int, []int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(envId, appIdIncludes)
	}
	if rf, ok := ret.Get(0).(func(int, []int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(envId, appIdIncludes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int, []int) error); ok {
		r1 = rf(envId, appIdIncludes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveByNotFilter provides a mock function with given fields: envId, appIdExcludes
func (_m *PipelineRepository) FindActiveByNotFilter(envId int, appIdExcludes []int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(envId, appIdExcludes)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int, []int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(envId, appIdExcludes)
	}
	if rf, ok := ret.Get(0).(func(int, []int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(envId, appIdExcludes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int, []int) error); ok {
		r1 = rf(envId, appIdExcludes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllPipelineInLast24Hour provides a mock function with given fields:
func (_m *PipelineRepository) FindAllPipelineInLast24Hour() ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called()

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*pipelineConfig.Pipeline, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*pipelineConfig.Pipeline); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllPipelinesByChartsOverrideAndAppIdAndChartId provides a mock function with given fields: chartOverridden, appId, chartId
func (_m *PipelineRepository) FindAllPipelinesByChartsOverrideAndAppIdAndChartId(chartOverridden bool, appId int, chartId int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(chartOverridden, appId, chartId)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(bool, int, int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(chartOverridden, appId, chartId)
	}
	if rf, ok := ret.Get(0).(func(bool, int, int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(chartOverridden, appId, chartId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(bool, int, int) error); ok {
		r1 = rf(chartOverridden, appId, chartId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAppAndEnvDetailsByPipelineId provides a mock function with given fields: id
func (_m *PipelineRepository) FindAppAndEnvDetailsByPipelineId(id int) (*pipelineConfig.Pipeline, error) {
	ret := _m.Called(id)

	var r0 *pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*pipelineConfig.Pipeline, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *pipelineConfig.Pipeline); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAppAndEnvironmentAndProjectByPipelineIds provides a mock function with given fields: pipelineIds
func (_m *PipelineRepository) FindAppAndEnvironmentAndProjectByPipelineIds(pipelineIds []int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(pipelineIds)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func([]int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(pipelineIds)
	}
	if rf, ok := ret.Get(0).(func([]int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(pipelineIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(pipelineIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAutomaticByCiPipelineId provides a mock function with given fields: ciPipelineId
func (_m *PipelineRepository) FindAutomaticByCiPipelineId(ciPipelineId int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(ciPipelineId)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(ciPipelineId)
	}
	if rf, ok := ret.Get(0).(func(int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(ciPipelineId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ciPipelineId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByCiPipelineId provides a mock function with given fields: ciPipelineId
func (_m *PipelineRepository) FindByCiPipelineId(ciPipelineId int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(ciPipelineId)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(ciPipelineId)
	}
	if rf, ok := ret.Get(0).(func(int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(ciPipelineId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ciPipelineId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByCiPipelineIdsIn provides a mock function with given fields: ciPipelineIds
func (_m *PipelineRepository) FindByCiPipelineIdsIn(ciPipelineIds []int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(ciPipelineIds)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func([]int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(ciPipelineIds)
	}
	if rf, ok := ret.Get(0).(func([]int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(ciPipelineIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(ciPipelineIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: id
func (_m *PipelineRepository) FindById(id int) (*pipelineConfig.Pipeline, error) {
	ret := _m.Called(id)

	var r0 *pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*pipelineConfig.Pipeline, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *pipelineConfig.Pipeline); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByIdsIn provides a mock function with given fields: ids
func (_m *PipelineRepository) FindByIdsIn(ids []int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(ids)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func([]int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByIdsInAndEnvironment provides a mock function with given fields: ids, environmentId
func (_m *PipelineRepository) FindByIdsInAndEnvironment(ids []int, environmentId int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(ids, environmentId)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func([]int, int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(ids, environmentId)
	}
	if rf, ok := ret.Get(0).(func([]int, int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(ids, environmentId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func([]int, int) error); ok {
		r1 = rf(ids, environmentId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: pipelineName
func (_m *PipelineRepository) FindByName(pipelineName string) (*pipelineConfig.Pipeline, error) {
	ret := _m.Called(pipelineName)

	var r0 *pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*pipelineConfig.Pipeline, error)); ok {
		return rf(pipelineName)
	}
	if rf, ok := ret.Get(0).(func(string) *pipelineConfig.Pipeline); ok {
		r0 = rf(pipelineName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pipelineName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByParentCiPipelineId provides a mock function with given fields: ciPipelineId
func (_m *PipelineRepository) FindByParentCiPipelineId(ciPipelineId int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(ciPipelineId)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(ciPipelineId)
	}
	if rf, ok := ret.Get(0).(func(int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(ciPipelineId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(ciPipelineId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByPipelineTriggerGitHash provides a mock function with given fields: gitHash
func (_m *PipelineRepository) FindByPipelineTriggerGitHash(gitHash string) (*pipelineConfig.Pipeline, error) {
	ret := _m.Called(gitHash)

	var r0 *pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*pipelineConfig.Pipeline, error)); ok {
		return rf(gitHash)
	}
	if rf, ok := ret.Get(0).(func(string) *pipelineConfig.Pipeline); ok {
		r0 = rf(gitHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(gitHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIdsByAppIdsAndEnvironmentIds provides a mock function with given fields: appIds, environmentIds
func (_m *PipelineRepository) FindIdsByAppIdsAndEnvironmentIds(appIds []int, environmentIds []int) ([]int, error) {
	ret := _m.Called(appIds, environmentIds)

	var r0 []int
	var r1 error
	if rf, ok := ret.Get(0).(func([]int, []int) ([]int, error)); ok {
		return rf(appIds, environmentIds)
	}
	if rf, ok := ret.Get(0).(func([]int, []int) []int); ok {
		r0 = rf(appIds, environmentIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	if rf, ok := ret.Get(1).(func([]int, []int) error); ok {
		r1 = rf(appIds, environmentIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindIdsByProjectIdsAndEnvironmentIds provides a mock function with given fields: projectIds, environmentIds
func (_m *PipelineRepository) FindIdsByProjectIdsAndEnvironmentIds(projectIds []int, environmentIds []int) ([]int, error) {
	ret := _m.Called(projectIds, environmentIds)

	var r0 []int
	var r1 error
	if rf, ok := ret.Get(0).(func([]int, []int) ([]int, error)); ok {
		return rf(projectIds, environmentIds)
	}
	if rf, ok := ret.Get(0).(func([]int, []int) []int); ok {
		r0 = rf(projectIds, environmentIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	if rf, ok := ret.Get(1).(func([]int, []int) error); ok {
		r1 = rf(projectIds, environmentIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindNumberOfAppsWithCdPipeline provides a mock function with given fields: appIds
func (_m *PipelineRepository) FindNumberOfAppsWithCdPipeline(appIds []int) (int, error) {
	ret := _m.Called(appIds)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]int) (int, error)); ok {
		return rf(appIds)
	}
	if rf, ok := ret.Get(0).(func([]int) int); ok {
		r0 = rf(appIds)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]int) error); ok {
		r1 = rf(appIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAppAndEnvDetailsForDeploymentAppTypePipeline provides a mock function with given fields: deploymentAppType, clusterIds
func (_m *PipelineRepository) GetAppAndEnvDetailsForDeploymentAppTypePipeline(deploymentAppType string, clusterIds []int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(deploymentAppType, clusterIds)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(deploymentAppType, clusterIds)
	}
	if rf, ok := ret.Get(0).(func(string, []int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(deploymentAppType, clusterIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []int) error); ok {
		r1 = rf(deploymentAppType, clusterIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArgoPipelineByArgoAppName provides a mock function with given fields: argoAppName
func (_m *PipelineRepository) GetArgoPipelineByArgoAppName(argoAppName string) (pipelineConfig.Pipeline, error) {
	ret := _m.Called(argoAppName)

	var r0 pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (pipelineConfig.Pipeline, error)); ok {
		return rf(argoAppName)
	}
	if rf, ok := ret.Get(0).(func(string) pipelineConfig.Pipeline); ok {
		r0 = rf(argoAppName)
	} else {
		r0 = ret.Get(0).(pipelineConfig.Pipeline)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(argoAppName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArgoPipelinesHavingLatestTriggerStuckInNonTerminalStatuses provides a mock function with given fields: deployedBeforeMinutes, getPipelineDeployedWithinHours
func (_m *PipelineRepository) GetArgoPipelinesHavingLatestTriggerStuckInNonTerminalStatuses(deployedBeforeMinutes int, getPipelineDeployedWithinHours int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(deployedBeforeMinutes, getPipelineDeployedWithinHours)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(deployedBeforeMinutes, getPipelineDeployedWithinHours)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(deployedBeforeMinutes, getPipelineDeployedWithinHours)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(deployedBeforeMinutes, getPipelineDeployedWithinHours)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArgoPipelinesHavingTriggersStuckInLastPossibleNonTerminalTimelines provides a mock function with given fields: pendingSinceSeconds, timeForDegradation
func (_m *PipelineRepository) GetArgoPipelinesHavingTriggersStuckInLastPossibleNonTerminalTimelines(pendingSinceSeconds int, timeForDegradation int) ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called(pendingSinceSeconds, timeForDegradation)

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]*pipelineConfig.Pipeline, error)); ok {
		return rf(pendingSinceSeconds, timeForDegradation)
	}
	if rf, ok := ret.Get(0).(func(int, int) []*pipelineConfig.Pipeline); ok {
		r0 = rf(pendingSinceSeconds, timeForDegradation)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(pendingSinceSeconds, timeForDegradation)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEnvOverrideId provides a mock function with given fields: envOverrideId
func (_m *PipelineRepository) GetByEnvOverrideId(envOverrideId int) ([]pipelineConfig.Pipeline, error) {
	ret := _m.Called(envOverrideId)

	var r0 []pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int) ([]pipelineConfig.Pipeline, error)); ok {
		return rf(envOverrideId)
	}
	if rf, ok := ret.Get(0).(func(int) []pipelineConfig.Pipeline); ok {
		r0 = rf(envOverrideId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(envOverrideId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEnvOverrideIdAndEnvId provides a mock function with given fields: envOverrideId, envId
func (_m *PipelineRepository) GetByEnvOverrideIdAndEnvId(envOverrideId int, envId int) (pipelineConfig.Pipeline, error) {
	ret := _m.Called(envOverrideId, envId)

	var r0 pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (pipelineConfig.Pipeline, error)); ok {
		return rf(envOverrideId, envId)
	}
	if rf, ok := ret.Get(0).(func(int, int) pipelineConfig.Pipeline); ok {
		r0 = rf(envOverrideId, envId)
	} else {
		r0 = ret.Get(0).(pipelineConfig.Pipeline)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(envOverrideId, envId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConnection provides a mock function with given fields:
func (_m *PipelineRepository) GetConnection() *pg.DB {
	ret := _m.Called()

	var r0 *pg.DB
	if rf, ok := ret.Get(0).(func() *pg.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pg.DB)
		}
	}

	return r0
}

// GetPartiallyDeletedPipelineByStatus provides a mock function with given fields: appId, envId
func (_m *PipelineRepository) GetPartiallyDeletedPipelineByStatus(appId int, envId int) (pipelineConfig.Pipeline, error) {
	ret := _m.Called(appId, envId)

	var r0 pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (pipelineConfig.Pipeline, error)); ok {
		return rf(appId, envId)
	}
	if rf, ok := ret.Get(0).(func(int, int) pipelineConfig.Pipeline); ok {
		r0 = rf(appId, envId)
	} else {
		r0 = ret.Get(0).(pipelineConfig.Pipeline)
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(appId, envId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPostStageConfigById provides a mock function with given fields: id
func (_m *PipelineRepository) GetPostStageConfigById(id int) (*pipelineConfig.Pipeline, error) {
	ret := _m.Called(id)

	var r0 *pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (*pipelineConfig.Pipeline, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) *pipelineConfig.Pipeline); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PipelineExists provides a mock function with given fields: pipelineName
func (_m *PipelineRepository) PipelineExists(pipelineName string) (bool, error) {
	ret := _m.Called(pipelineName)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(pipelineName)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(pipelineName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(pipelineName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: pipeline, tx
func (_m *PipelineRepository) Save(pipeline []*pipelineConfig.Pipeline, tx *pg.Tx) error {
	ret := _m.Called(pipeline, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*pipelineConfig.Pipeline, *pg.Tx) error); ok {
		r0 = rf(pipeline, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDeploymentAppCreatedInPipeline provides a mock function with given fields: deploymentAppCreated, pipelineId, userId
func (_m *PipelineRepository) SetDeploymentAppCreatedInPipeline(deploymentAppCreated bool, pipelineId int, userId int32) error {
	ret := _m.Called(deploymentAppCreated, pipelineId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, int, int32) error); ok {
		r0 = rf(deploymentAppCreated, pipelineId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UndoDelete provides a mock function with given fields: id
func (_m *PipelineRepository) UndoDelete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UniqueAppEnvironmentPipelines provides a mock function with given fields:
func (_m *PipelineRepository) UniqueAppEnvironmentPipelines() ([]*pipelineConfig.Pipeline, error) {
	ret := _m.Called()

	var r0 []*pipelineConfig.Pipeline
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*pipelineConfig.Pipeline, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*pipelineConfig.Pipeline); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pipelineConfig.Pipeline)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: pipeline, tx
func (_m *PipelineRepository) Update(pipeline *pipelineConfig.Pipeline, tx *pg.Tx) error {
	ret := _m.Called(pipeline, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(*pipelineConfig.Pipeline, *pg.Tx) error); ok {
		r0 = rf(pipeline, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCdPipelineAfterDeployment provides a mock function with given fields: deploymentAppType, cdPipelineIdIncludes, userId, delete
func (_m *PipelineRepository) UpdateCdPipelineAfterDeployment(deploymentAppType string, cdPipelineIdIncludes []int, userId int32, delete bool) error {
	ret := _m.Called(deploymentAppType, cdPipelineIdIncludes, userId, delete)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []int, int32, bool) error); ok {
		r0 = rf(deploymentAppType, cdPipelineIdIncludes, userId, delete)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCdPipelineDeploymentAppInFilter provides a mock function with given fields: deploymentAppType, cdPipelineIdIncludes, userId, deploymentAppCreated, delete
func (_m *PipelineRepository) UpdateCdPipelineDeploymentAppInFilter(deploymentAppType string, cdPipelineIdIncludes []int, userId int32, deploymentAppCreated bool, delete bool) error {
	ret := _m.Called(deploymentAppType, cdPipelineIdIncludes, userId, deploymentAppCreated, delete)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []int, int32, bool, bool) error); ok {
		r0 = rf(deploymentAppType, cdPipelineIdIncludes, userId, deploymentAppCreated, delete)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPipelineRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewPipelineRepository creates a new instance of PipelineRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPipelineRepository(t mockConstructorTestingTNewPipelineRepository) *PipelineRepository {
	mock := &PipelineRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
