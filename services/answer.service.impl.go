package services

import (
	"time"

	"gorm.io/gorm"
	"lawencon.com/imamfarisi/dao"
	"lawencon.com/imamfarisi/models"
)

var answerDao dao.AnswerDao = dao.AnswerDaoImpl{}

type AnswerServiceImpl struct{}

func (AnswerServiceImpl) CreateAnswer(hdr *models.AnswerHdr, dtl *[]models.AnswerDtl) error {
	return g.Transaction(func(tx *gorm.DB) error {
		hdr.CreatedDate = time.Now()
		if err := answerDao.CreateAnswerHdr(hdr, tx); err != nil {
			return err
		}

		for _, v := range *dtl {
			v.AnswerHdrId = *&hdr.Id
			v.CreatedDate = time.Now()
			if err := dtlDao.CreateAnswerDtl(&v, tx); err != nil {
				return err
			}
		}

		return nil
	})
}
