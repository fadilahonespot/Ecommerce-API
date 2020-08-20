package usecase

import "ecommerce/model"

func (e *ProductUsecaseImpl) AddWachlist(wachlist *model.Watchlist)(*model.Watchlist, error) {
	return e.productRepo.AddWachlist(wachlist)
}

func (e *ProductUsecaseImpl) ViewWachlistByUserId(id int)(*[]model.Watchlist, error) {
	return e.productRepo.ViewWachlistByUserId(id)
}

func (e *ProductUsecaseImpl) ViewWachlistByProductId(id int)(*model.Watchlist, error) {
	return e.productRepo.ViewWachlistByProductId(id)
}

func (e *ProductUsecaseImpl) DeleteWatchlistById(id int) error {
	return e.productRepo.DeleteWatchlistById(id)
}

func (e *ProductUsecaseImpl) DeleteWachlistByProductId(id int) error {
	tx := e.ongkirRepo.BeginTrans()
	err := e.productRepo.DeleteWachlistByProductId(id, tx)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}
