package service

import (
	"time"

	"lawencon.com/bootest/dao"
	"lawencon.com/bootest/model"
)

var candidatDao dao.CandidatProfileDao = dao.CandidatProfileDaoImpl{}

type CandidatProfileServiceImpl struct{}

func (CandidatProfileServiceImpl) CreateCandidat(data *model.CandidateProfiles) (e error) {
	defer catchError(&e)
	data.CreatedDate = time.Now()
	return candidatDao.CreateCandidat(data)
}