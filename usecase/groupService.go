package usecase

import "../infrastructure"

type GroupService interface {
	GetGroups() []string
}

type groupServiceImpl struct {
	repository infrastructure.GroupRepository
}

func CreateGroupService(groupRepository infrastructure.GroupRepository) GroupService {
	return &groupServiceImpl{
		groupRepository,
	}
}

func (service *groupServiceImpl) GetGroups() []string {
	return service.repository.GetGroups()
}
