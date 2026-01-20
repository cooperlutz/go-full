package outbound_test

// func TestNewPostgresAdapter(t *testing.T) {
// 	testDB := mocks.NewMockDBTX(t)
// 	repo := outbound.NewPostgresAdapter(testDB)
// 	assert.NotNil(t, repo)
// }

// func TestPostgres_FindAll(t *testing.T) {
// 	ctx := context.Background()
// 	testDB := mocks.NewMockIQuerierExamination(t)
// 	repo := outbound.PostgresAdapter{
// 		Handler: testDB,
// 	}
// 	testDB.On("FindAllExams", mock.Anything, mock.Anything).Return(
// 		outbound.FixtureExams, nil,
// 	)
// 	exams, err := repo.FindAll(ctx)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, exams)
// }

// func TestPostgres_FindAll_Error(t *testing.T) {
// 	ctx := context.Background()
// 	testDB := mocks.NewMockIQuerierExamination(t)
// 	repo := outbound.PostgresAdapter{
// 		Handler: testDB,
// 	}
// 	testDB.On("FindAllExams", mock.Anything, mock.Anything).Return(
// 		nil, assert.AnError,
// 	)
// 	exams, err := repo.FindAll(ctx)
// 	assert.Error(t, err)
// 	assert.Nil(t, exams)
// }
