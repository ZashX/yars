package engine

import (
	"testing"

	"gitlab.com/yars-ai/yars/ast"
	"gitlab.com/yars-ai/yars/repository"
)

func TestBuilderMovies(t *testing.T) {
	project, err := ast.Parse(
		`
		project movies;

		unit user {
			ID string;
			Age int32;
			BannedMovies []string;
			BannedAuthors []string;
		}
		
		unit movie {
			ID string;
			Author string;
			Flags []bool;
		}
		
		action rate user@movie;
		
		action watched user@movie;
		
		recommend bestFilms user@movie {
			rate: rate;
			filter:
				not movie.ID in user.BannedMovies and
				not movie.Author in user.BannedAuthors and
				not watched;
		}
		`)

	if err != nil {
		t.Fatalf("can't parse project: %v", err)
	}

	_, err = NewBuilder(repository.NewFabricMock("movies"), project).Build()
	if err != nil {
		t.Fatalf("can't build project: %v", err)
	}
}
