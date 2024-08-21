package campaign

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignByCampaignId(input GetDetailCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.repository.FindCampaignByUserId(userId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindAllCampaign()

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaignByCampaignId(input GetDetailCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindCampaignByCampaignId(input.Id)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
