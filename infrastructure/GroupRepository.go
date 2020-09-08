package infrastructure

type GroupRepository interface {
	GetGroups() []string
}

type groupRepositoryImpl struct{}

func CreateGroupRepository() GroupRepository {
	return &groupRepositoryImpl{}
}

func (g *groupRepositoryImpl) GetGroups() []string {
	return []string{
		"group1",
		"group2",
	}
}
