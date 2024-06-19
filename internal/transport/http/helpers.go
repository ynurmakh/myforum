package http

import "strconv"

type CheckedCategory struct {
	Category_id   int
	Category_name string
	IsChecked     bool
}

func (t *Transport) GetCategoriesForTemplate(categoriesList []string) (checkedList *[]CheckedCategory, idList []int, err error) {
	categories, err := t.service.GetCategiries()
	if err != nil {
		return nil, nil, err
	}

	idList = []int{}
	for _, c := range categoriesList {
		num, err := strconv.Atoi(c)
		if err != nil {
			return nil, nil, err
		}
		idList = append(idList, num)
	}

	checkedList = &[]CheckedCategory{}
	for _, c := range *categories {
		checked := func() bool {
			for _, num := range idList {
				if num == c.Category_id {
					return true
				}
			}
			return false
		}()
		*checkedList = append(*checkedList, CheckedCategory{
			Category_id:   c.Category_id,
			Category_name: c.Category_name,
			IsChecked:     checked,
		})
	}
	return checkedList, idList, nil
}
