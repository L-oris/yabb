package inject

import (
	"fmt"

	"github.com/L-oris/yabb/inject/types"
	"github.com/L-oris/yabb/repository/bucketrepository"
	"github.com/L-oris/yabb/repository/db"
	"github.com/L-oris/yabb/repository/postrepository"
	"github.com/sarulabs/di"
)

func repositories() []di.Def {
	postRepository := di.Def{
		Name: types.PostRepository.String(),
		Build: func(ctn di.Container) (interface{}, error) {
			return postrepository.New(
				&postrepository.Config{
					DB: db.BlogDB,
				},
			), nil
		},
	}

	bucketRepository := di.Def{
		Name: types.BucketRepository.String(),
		Build: func(ctn di.Container) (interface{}, error) {
			repo, err := bucketrepository.New(
				bucketrepository.Config{
					BucketName: "yabb",
				},
			)
			if err != nil {
				return nil, fmt.Errorf("could not create bucket: %s", err.Error())
			}

			return repo, nil
		},
	}

	return []di.Def{
		postRepository, bucketRepository,
	}
}
