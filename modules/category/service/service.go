package categoryservice

// type ICategoryQueryRepo interface {
// 	FindByID(ctx context.Context, id uuid.UUID) (*categorymodel.Category, error)
// 	List(
// 		ctx context.Context,
// 		pagingDTO *sharedmodel.PagingDTO,
// 		filterDTO *categorymodel.FilterCategoryDTO,
// 	) ([]categorymodel.Category, error)
// }

// type ICategoryCommandRepo interface {
// 	Insert(ctx context.Context, data *categorymodel.Category) error
// 	Update(ctx context.Context, id uuid.UUID, data *categorymodel.CategoryUpdateDTO) error
// 	Delete(ctx context.Context, id uuid.UUID, isHard bool) error
// }

// type ICategoryRepository interface {
// 	ICategoryQueryRepo
// 	ICategoryCommandRepo
// }

// Dependencies Injection by constructor/new function
// func NewCategoryService(catRepo ICategoryRepository) *CategoryService {
// 	return &CategoryService{catRepo: catRepo}
// }

// type CategoryService struct {
// 	catRepo ICategoryRepository // private
// }

// // Dependencies Injection by setter method
// func (s *CategoryService) SetCategoryRepository(catRepo ICategoryRepository) {
// 	s.catRepo = catRepo
// }
