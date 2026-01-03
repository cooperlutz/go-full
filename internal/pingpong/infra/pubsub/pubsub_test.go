package pubsub_test

// func TestNew_Success(t *testing.T) {
// 	pgxmock, err := pgxmock.NewPool()
// 	assert.NoError(t, err)
// 	defer pgxmock.Close()

// 	mockRepo := mocks.NewMockIPingPongRepository(t)

// 	ps, err := pubsub.New(pgxmock, mockRepo)

// 	assert.NoError(t, err)
// 	assert.NotNil(t, ps)
// }

// func TestRegisterSubscriberHandlers_Success(t *testing.T) {
// 	pgxmock, err := pgxmock.NewPool()
// 	assert.NoError(t, err)
// 	defer pgxmock.Close()

// 	mockRepo := mocks.NewMockIPingPongRepository(t)

// 	ps, err := pubsub.New(pgxmock, mockRepo)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, ps)

// 	err = ps.RegisterSubscriberHandlers()
// 	assert.NoError(t, err)
// }

// func TestRegisterSubscriberHandlers_HandlerCallback(t *testing.T) {
// 	pgxmock, err := pgxmock.NewPool()
// 	assert.NoError(t, err)
// 	defer pgxmock.Close()

// 	mockRepo := mocks.NewMockIPingPongRepository(t)

// 	ps, err := pubsub.New(pgxmock, mockRepo)
// 	assert.NoError(t, err)

// 	err = ps.RegisterSubscriberHandlers()
// 	assert.NoError(t, err)

// 	// Test the handler callback behavior
// 	testMsg := &message.Message{UUID: "test-uuid"}
// 	results, err := func(msg *message.Message) ([]*message.Message, error) {
// 		msg.Ack()
// 		return []*message.Message{msg}, nil
// 	}(testMsg)

// 	assert.NoError(t, err)
// 	assert.Len(t, results, 1)
// 	assert.Equal(t, "test-uuid", results[0].UUID)
// }
