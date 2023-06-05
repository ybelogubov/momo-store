package dependencies

import (
	"gitlab.praktikum-services.ru/yu.belogubov/momo-store/internal/store/dumplings"
	"gitlab.praktikum-services.ru/yu.belogubov/momo-store/internal/store/dumplings/fake"
)

// NewFakeDumplingsStore returns new fake store for app
func NewFakeDumplingsStore() (dumplings.Store, error) {
	packs := []dumplings.Product{
		{
			ID:          1,
			Name:        "Пельмени",
			Description: "С говядиной",
			Price:       5.00,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/1.jpeg",
		},
		{
			ID:          2,
			Name:        "Хинкали",
			Description: "Со свининой",
			Price:       3.50,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/2.jpeg",
		},
		{
			ID:          3,
			Name:        "Манты",
			Description: "С мясом молодых бычков",
			Price:       2.75,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/3.jpeg",
		},
		{
			ID:          4,
			Name:        "Буузы",
			Description: "С телятиной и луком",
			Price:       4.00,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/4.jpeg",
		},
		{
			ID:          5,
			Name:        "Цзяоцзы",
			Description: "С говядиной и свининой",
			Price:       7.25,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/5.jpeg",
		},
		{
			ID:          6,
			Name:        "Гедза",
			Description: "С соевым мясом",
			Price:       3.50,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/6.jpeg",
		},
		{
			ID:          7,
			Name:        "Дим-самы",
			Description: "С уткой",
			Price:       2.65,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/7.jpeg",
		},
		{
			ID:          8,
			Name:        "Момо",
			Description: "С бараниной",
			Price:       5.00,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/8.jpeg",
		},
		{
			ID:          9,
			Name:        "Вонтоны",
			Description: "С креветками",
			Price:       4.10,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/1.jpeg",
		},
		{
			ID:          10,
			Name:        "Баоцзы",
			Description: "С капустой",
			Price:       4.20,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/2.jpeg",
		},
		{
			ID:          11,
			Name:        "Кундюмы",
			Description: "С грибами",
			Price:       5.45,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/3.jpeg",
		},
		{
			ID:          12,
			Name:        "Курзе",
			Description: "С крабом",
			Price:       3.25,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/4.jpeg",
		},
		{
			ID:          13,
			Name:        "Бораки",
			Description: "С говядиной и бараниной",
			Price:       4.00,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/6.jpeg",
		},
		{
			ID:          14,
			Name:        "Равиоли",
			Description: "С рикоттой",
			Price:       2.90,
			Image:       "https://storage.yandexcloud.net/ibelogubov-praktikum-devops/momo-store-static/images/7.jpeg",
		},
	}

	store := fake.NewStore()
	store.SetAvailablePacks(packs...)

	return store, nil
}
